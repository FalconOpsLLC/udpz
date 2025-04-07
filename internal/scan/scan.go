package scan

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
	"github.com/FalconOpsLLC/udpz/internal/data"

	"github.com/rs/zerolog"
	"github.com/txthinking/socks5"
)

func (sc *UdpProbeScanner) handleResult(pr PortResult) {

	hostString := pr.Host.Host.String()

	resultLog := sc.Logger.With().
		Str("target", pr.Host.Target.String()).
		IPAddr("host", pr.Host.Host).
		Uint16("port", pr.Port).
		Str("service", pr.Service.Slug).
		Str("probe", pr.Probe.Slug).
		Logger()

	resultLog.Debug().
		Str("response", pr.Response).
		Msg("Received response")

	if _, ok := sc.resultsMap[string(hostString)]; !ok {
		sc.resultsMap[hostString] = make(map[uint16][]PortResult)
	}
	if _, ok := sc.resultsMap[hostString][pr.Port]; !ok {
		resultLog.Info().Msg("Discovered UDP service")

		sc.results = append(sc.results, pr)
		sc.resultsMap[hostString][pr.Port] = []PortResult{pr}
	} else {
		sc.resultsMap[hostString][pr.Port] = append(sc.resultsMap[hostString][pr.Port], pr)
	}
}

func (sc *UdpProbeScanner) Length() int {
	return len(sc.results)
}

// TODO: ctx
func (sc *UdpProbeScanner) scanTask(host Host, port uint16, payload []byte) (result PortResult, err error) {

	/*
		sc.Logger.Trace().
			Str("function", "(*UdpProbeScanner).scanTask(Host, uint16, []byte) (PortResult, error)").
			Dict("arguments", zerolog.Dict().
				Interface("host", host).
				Uint16("port", port).
				Bytes("payload", payload)).
			Msg("(*UdpProbeScanner).scanTask")
	*/

	var conn net.Conn

	for {
		transport := "udp"
		address := fmt.Sprintf("%s:%d", host.String(), port)

		if sc.useProxy {
			/*
				sc.Logger.Trace().
					Str("function", "(*socks5.Client).Dial").
					Str("proxy", sc.proxy.Server). //sc.proxy.Server).
					Str("transport", transport).
					Str("address", address).
					Msg("(*socks5.Client).Dial(...)")
			*/

			conn, err = sc.proxy.Dial(transport, address)
		} else {
			/*
				sc.Logger.Trace().
					Str("function", "net.Dial").
					Str("transport", transport).
					Str("address", address).
					Msg("net.Dial(...)")
			*/

			conn, err = net.Dial(transport, address)
		}
		if err == nil {
			if err = conn.SetReadDeadline(time.Now().Add(sc.ReadTimeout)); err == nil {

				var response []byte
				var readLen int

				/*
					sc.Logger.Trace().
						Str("type", "(net.Conn).write").
						Str("transport", transport).
						Str("address", address).
						Bytes("data", payload).
						Msg("(net.Conn).Write(data)")
				*/

				conn.Write(payload)

				responseBuffer := make([]byte, RESPONSE_MAX_LEN)

				if readLen, err = conn.Read(responseBuffer); err != nil {
					if strings.Contains(err.Error(), "read: no route to host") {
						sc.Logger.Debug().
							Err(err).
							Msg("Resource limit reached")
						time.Sleep(10 * time.Millisecond)
						continue
					}

				} else {
					response = responseBuffer[:readLen]
					responseBuffer = nil
				}
				/*
					sc.Logger.Trace().
						Str("type", "connection.close").
						Str("transport", transport).
						Str("address", address).
						Bytes("data", response).
						Msg("(net.Conn).Close()")
				*/

				if terr := conn.Close(); terr != nil {
					sc.Logger.Debug().
						Stack().
						Err(terr).
						Str("transport", transport).
						Str("address", address).
						Msg("Failed to close connection")
				}

				if err == nil {
					result = PortResult{
						Port:      port,
						Transport: transport,
						Host:      host,
					}
					if readLen > 0 {
						result.Response = base64.StdEncoding.EncodeToString(response)
					}
				}
			}
			break
		} else {
			if strings.Contains(err.Error(), "connect: resource temporarily unavailable") {
				sc.Logger.Debug().
					Err(err).
					Msg("Resource limit reached")
				time.Sleep(10 * time.Millisecond)
				continue
			}
			break
		}
	}

	/*
		sc.Logger.Trace().
			Str("function", "(*UdpProbeScanner).scanTask(Host, uint16, []byte) (PortResult, error)").
			Dict("arguments", zerolog.Dict().
				Interface("host", host).
				Uint16("port", port).
				Bytes("payload", payload)).
			Dict("return", zerolog.Dict().
				AnErr("err", err).
				Interface("result", result)).
			Msg("return (*UdpProbeScanner).scanTask")
	*/
	return
}

func (sc *UdpProbeScanner) Scan(targetSourceList []string, slugs []string, tags []string) (err error) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).Scan([]string, []string, []string) (error)").
		Dict("arguments", zerolog.Dict().
			Strs("targetSourceList", targetSourceList).
			Strs("slugs", slugs).
			Strs("tags", tags)).
		Msg("(*UdpProbeScanner).Scan(...)")

	var hostWg sync.WaitGroup

	services, probes := sc.Broker.Filter(slugs, tags)

	probeCount := uint(len(probes))
	serviceCount := uint(len(services))

	if probeCount == 0 {
		sc.Logger.Debug().
			Strs("slugs", slugs).
			Strs("tags", tags).
			Msg("No probes to run")
		return errors.New("no matching probes")
	}

	hosts := make(chan Host)

	go func(wg *sync.WaitGroup, c chan Host) {

		sc.Logger.Debug().
			Int("target_count", len(targetSourceList)).
			Msg("Resolving targets")

		for _, ts := range targetSourceList {
			sc.ResolveTarget(ts, c)
		}
		wg.Wait()
		close(c)

	}(&hostWg, hosts)

	go func() {
		// Handle live results
		for r := range sc.resultsLive {
			sc.handleResult(r)
		}
	}()

	sc.Logger.Info().
		Uint("service_count", serviceCount).
		Uint("probe_count", probeCount).
		Uint("packet_count", probeCount*(sc.Retransmissions+1)).
		Msg("Starting scan")

	sc.DefaultScan(&hostWg, hosts, services)
	return
}

func NewUdpProbeScanner(logger *zerolog.Logger, broker data.UdpDataBroker, scanAllAddresses bool,
	hostConcurrency uint, portConcurrency uint, retransmissions uint, readTimeout time.Duration,
	socks5Address string, socks5User string, socks5Password string, socks5Timeout int) (sc UdpProbeScanner, err error) {

	// Assemble a new UdpProbeScanner using supplied arguments

	sc.resultsLive = make(chan PortResult)
	sc.resultsMap = make(map[string]map[uint16][]PortResult)

	sc.scanAllAddresses = scanAllAddresses
	sc.HostConcurrency = hostConcurrency
	sc.PortConcurrency = portConcurrency
	sc.Retransmissions = retransmissions
	sc.ReadTimeout = readTimeout

	sc.Broker = broker
	sc.Logger = logger

	if socks5Address != "" {

		socksLog := sc.Logger.With().
			Str("address", socks5Address).
			Int("timeout", socks5Timeout).
			Logger()

		if socks5User != "" || socks5Password != "" {
			socksLog = socksLog.With().
				Str("user", socks5User).
				Logger()
		}

		sc.useProxy = true

		if socks5User != "" || socks5Password != "" {
			socksLog.Debug().Msg("Using SOCKS5 authentication")
		}

		if sc.proxy, err = socks5.NewClient(

			socks5Address, socks5User, socks5Password,
			socks5Timeout, socks5Timeout); err != nil {

			socksLog.Error().
				Stack().Err(err).
				Msg("Failed to initialize SOCKS5 proxy dialer")

		} else if _, err := sc.proxy.Dial("udp", "127.0.0.1:1"); err != nil {
			// This shouldn't send any packets through the proxy, justs tests if its available.

			socksLog.Error().
				Stack().Err(err).
				Msg("SOCKS5 connection failed")

		} else {
			socksLog.Info().Msg("Using SOCKS5 proxy")
		}

	}

	return
}
