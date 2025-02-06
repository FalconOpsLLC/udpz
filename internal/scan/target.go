package scan

import (
	"bufio"
	"net"
	"os"

	"github.com/rs/zerolog"
)

// Credit to @udhos -> https://gist.github.com/udhos/b468fbfd376aa0b655b6b0c539a88c03
func nextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}

func (sc *UdpProbeScanner) ResolveTargetLine(targetSource string, hosts chan Host) (ok bool) {

	var target Target
	//var host Host

	target.Target = targetSource

	if ip := net.ParseIP(targetSource); ip != nil {
		// If target can be parsed as IP address

		target.Type = "IP"

		sc.Logger.Debug().
			Str("target", target.String()).
			IPAddr("ip", ip).
			Msg("Target resolved")

		hosts <- Host{
			Target: target,
			Host:   ip,
		} // Target is resolved to IP address
		ok = true

	} else if ip, ipNet, err := net.ParseCIDR(targetSource); err == nil {
		// If target can be parsed as CIDR

		target.Type = "CIDR"

		sc.Logger.Debug().
			Str("target", target.String()).
			IPAddr("ip", ip).
			Str("cidr", ipNet.String()).
			Msg("Target CIDR resolved")

		for ipNet.Contains(ip) {
			hosts <- Host{
				Target: target,
				Host:   ip,
			}
			ip = nextIP(ip, 1)
		}
		ok = true

	} else if REGEX_HOSTNAME.MatchString(targetSource) {

		if ips, err := net.LookupIP(targetSource); err == nil {

			target.Type = "hostname"

			sc.Logger.Debug().
				Str("target", targetSource).
				Int("addresses", len(ips)).
				Msg("Resolved target hostname")

			for _, ip := range ips {

				sc.Logger.Debug().
					Str("target", targetSource).
					IPAddr("ip", ip).
					Msg("Queueing host")

				hosts <- Host{
					Target: target,
					Host:   ip,
				}

				if !sc.scanAllAddresses {
					break
				}
			}
			ok = true
		}
	}
	return
}

func (sc *UdpProbeScanner) ResolveTarget(targetSource string, hosts chan Host) (err error) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).ResolveTarget").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Msg("(*UdpProbeScanner).ResolveTarget(...)")

	if sc.ResolveTargetLine(targetSource, hosts) {

	} else if _, err := os.Stat(targetSource); err == nil {

		var file *os.File

		if file, err = os.OpenFile(targetSource, os.O_RDONLY, 0); err != nil {
			sc.Logger.Error().
				Err(err).
				Str("path", targetSource).
				Msg("Failed to open target input file")
		}
		defer func() {
			if err = file.Close(); err != nil {
				sc.Logger.Error().
					Err(err).
					Str("path", targetSource).
					Msg("Failed to close target input file")
			}
		}()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			sc.ResolveTargetLine(scanner.Text(), hosts)
		}
	} else {
		sc.Logger.Error().
			Str("target", targetSource).
			Msg("Could not resolve target")
	}

	sc.Logger.Trace().
		Str("type", "return").
		Str("function", "(*UdpProbeScanner).ResolveTarget").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Dict("return", zerolog.Dict().
			AnErr("err", err)).
		Msg("return <- (*UdpProbeScanner).ResolveTarget()")

	return nil
}
