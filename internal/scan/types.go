package scan

import (
	"fmt"
	"net"
	"time"
	"github.com/FalconOpsLLC/udpz/internal/data"

	"github.com/rs/zerolog"
	"github.com/txthinking/socks5"
)

type UdpProbeScanner struct {
	HostConcurrency uint
	PortConcurrency uint
	ProbeCount      uint
	Retransmissions uint
	ReadTimeout     time.Duration

	Logger *zerolog.Logger
	Broker data.UdpDataBroker

	proxy            *socks5.Client
	useProxy         bool
	scanAllAddresses bool

	resultsLive chan PortResult
	results     []PortResult
	resultsMap  map[string]map[uint16][]PortResult
}

type PortResult struct {
	Host      Host                  `yaml:"host" json:"host"`
	Port      uint16                `yaml:"port" json:"port"`
	Transport string                `yaml:"transport" json:"transport"`
	Service   data.UdpServiceOutput `yaml:"service" json:"service"`
	Probe     data.UdpProbeOutput   `yaml:"probe" json:"probe"`
	Response  string                `yaml:"response" json:"response"`
}

type Target struct {
	Type   string `yaml:"type" json:"type"`
	Target string `yaml:"source" json:"source"`

	targetString string
}

type Host struct {
	Host   net.IP `yaml:"host" json:"host"`
	Target Target `yaml:"target" json:"target"`

	hostString string
}

func (host *Host) String() string {
	if host.hostString == "" {
		host.hostString = host.Host.String()

		if host.Host.To16() != nil {
			host.hostString = fmt.Sprintf("[%s]", host.hostString)
		}
	}
	return host.hostString
}

func (target *Target) String() string {
	if target.targetString == "" {
		target.targetString = fmt.Sprintf("%s:%s", target.Type, target.Target)
	}
	return target.targetString
}
