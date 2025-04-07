package scan

import (
	"encoding/base64"
	"strings"
	"sync"
	"github.com/FalconOpsLLC/udpz/internal/data"
)

func (sc *UdpProbeScanner) DefaultScan(hostWg *sync.WaitGroup, hosts chan Host, services map[string]data.UdpService) (err error) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).DefaultScan").
		Msg("(*UdpProbeScanner).DefaultScan(*sync.WaitGroup, chan Host, map[string]data.UdpService)")

	hostSem := make(chan struct{}, sc.HostConcurrency)

	for host := range hosts {

		host := host // Shadow variable
		portWg := sync.WaitGroup{}
		portSem := make(chan struct{}, sc.PortConcurrency)

		hostSem <- struct{}{}
		hostWg.Add(1)

		go func() {

			defer func() {
				hostWg.Done()
				<-hostSem
			}()

			for _, service := range services {
				for _, port := range service.Ports {

					// Port is unresponsive unless a response is received (duh)
					var portStatus uint8 = STATE_UNRESPONSIVE

					for _, probe := range service.Probes {

						probe := probe

						portSem <- struct{}{}
						portWg.Add(1)

						go func(wg *sync.WaitGroup,
							h Host, port uint16, probe data.UdpProbe,
							service data.UdpService, portStatus *uint8) {

							defer func() {
								wg.Done()
								<-portSem
							}()

							if probeBytes, err := base64.StdEncoding.DecodeString(probe.EncodedData); err == nil {

								portLog := sc.Logger.With().
									Str("target", h.Target.Target).
									IPAddr("host", h.Host).
									Uint16("port", port).
									Str("probe", probe.Slug).
									Logger()

								for i := 0; i <= int(sc.Retransmissions); i++ {

									// Don't send any more packets if port is closed
									if *portStatus == STATE_CLOSED {
										portLog.Trace().
											Msg("Skipping closed port")
									}

									if result, err := sc.scanTask(h, port, probeBytes); err != nil {

										// Check if error is a refused connection (port closed)
										if strings.Contains(err.Error(), "connection refused") {
											*portStatus = STATE_CLOSED
											portLog.Trace().
												Msg("Port closed")

											// Check if error is an io timeout (port unresponsive)
										} else if strings.Contains(err.Error(), "i/o timeout") {
											portLog.Debug().
												Msg("Port unresponsive")

										} else {
											// Scan task returned unexpected error
											portLog.Error().
												Stack().Err(err).
												Msg("Error in scan task")
										}
									} else {
										// Convert results to relevant output data
										result.Service = service.ToOutput()
										result.Probe = probe.ToOutput()

										sc.resultsLive <- result

										// Mark port as responsive to prevent further scanning
										*portStatus = STATE_RESPONSIVE
									}
								}
							} else {
								// If base64.StdEncoding.DecodeString fails to decode probe data
								sc.Logger.Error().
									Stack().Err(err).
									Interface("probe", probe).
									Msg("Failed to decode probe data")
							}
						}(&portWg, host, port, probe, service, &portStatus)
					}
				}
			}
			// Wait for port results
			portWg.Wait()
		}()
	}

	// Wait for host results
	hostWg.Wait()
	return
}
