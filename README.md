# UDPz

UDPz is a speedy, portable, cross-platform UDP port scanner written in Go.

![UDPz session](https://github.com/user-attachments/assets/91caf4ce-4710-4304-afe4-eec81af956b4)

## Purpose

UDPz was created to address the need for a fast and efficient tool to scan UDP services across multiple hosts. 
Traditional network scanning tools like `nmap` often provide slower UDP scanning capabilities, which can be a bottleneck for network administrators and security professionals.

UDPz aims to fill this gap by providing a robust solution that can be easily integrated into existing workflows due to it's logging capabilities and flexible output options.

## Features

- **Root-less**: In direct contrast to other UDP scanning solutions like [nmap](https://github.com/nmap/nmap), UDPz **does not** require root or privileged access.
- **Concurrent Scanning**: Utilizes goroutines and channels to perform flexible concurrent scans, significantly speeding up the scanning process.
- **Structured Logging**: Uses [zerolog](https://github.com/rs/zerolog) for detailed and structured logging, making it easier to analyze scan results.
- **Flexible Target Resolution**: Supports loading IP addresses, CIDR ranges, and hostnames from arguments, or from a file.
- **Proxy Support**: Offers SOCKS5 proxy support for UDP tunneling.
- **Error Handling**: Gracefully handles errors during scanning, ensuring the process continues even if some targets fail.

> [!WARNING]
> The SOCKS5 client will only work with SOCKS5 servers that implement UDP support.

## Installation

> [!TIP]
> Docker installation is recommended for higher performance.

### Docker

1. Clone the repository:
```sh
git clone https://github.com/FalconOpsLLC/udpz.git
cd udpz
```

2. Build the Docker image:
```sh
docker build --tag="udpz" --network="host" .
```

3. Run the Docker container:
```sh
alias udpz='sudo docker run --rm --network="host" --name="udpz" --volume="$PWD:/output" -it udpz'
udpz [flags] [targets ...]
```

> [!TIP]
> When using the docker image with file I/O, make sure to utilize Docker volumes to properly interact with files on the host filesystem.

### Standard Installation

#### `go install`

```sh
go install -ldflags="-s -w" github.com/FalconOpsLLC/udpz@latest
```

#### Manual Installation

1. Clone the repository:
```sh
git clone https://github.com/FalconOpsLLC/udpz.git
cd udpz
```

2. Build the project:
```sh
go build
```

3. Run the binary:
```sh
./udpz [flags] [targets ...]
```

## Usage

```
Usage:
  udpz [flags] [IP|hostname|CIDR|file ...]

Flags:
  -v, --version              version for udpz
  -o, --output string        Save output to file
  -O, --log string           Output log messages to file
      --append               Append results to output file (default true)
  -f, --format string        Output format [text, pretty, csv, tsv, json, yaml, auto] (default "auto")
  -F, --log-format string    Output log format [pretty, json, auto] (default "auto")
  -l, --list                 List available services / probes
  -p, --probes string        comma-delimited list of service probes
      --tags string          comma-delimited list of target service tags
  -c, --host-tasks uint      Maximum Number of hosts to scan concurrently (default 10)
  -P, --port-tasks uint      Number of Concurrent scan tasks per host (default 100)
  -r, --retries uint         Number of probe retransmissions per probe (default 2)
  -t, --timeout uint         UDP Probe timeout in milliseconds (default 3000)
  -A, --all                  Scan all resolved addresses instead of just the first (default true)
  -S, --socks string         SOCKS5 proxy address as HOST:PORT
      --socks-user string    SOCKS5 proxy username
      --socks-pass string    SOCKS5 proxy password
      --socks-timeout uint   SOCKS5 proxy timeout in milliseconds (default 3000)
  -D, --debug                Enable debug logging (Very noisy!)
  -T, --trace                Enable trace logging (Very noisy!)
  -q, --quiet                Disable info logging
  -s, --silent               Disable ALL logging
  -h, --help                 help for udpz
```

The target argument(s) can be an IP address, hostname, CIDR, or file(s) containing targets.


### Examples

- Simple scan of a host and CIDR limited to one packet retransmission and without informational logging:
```
./udpz --retries 1 --quiet 10.10.197.101 10.10.197.102/31

╭───────────────┬───────────┬──────┬───────┬──────────┬──────────────────────╮
│ HOST          │ TRANSPORT │ PORT │ STATE │ SERVICE  │ PROBES               │
├───────────────┼───────────┼──────┼───────┼──────────┼──────────────────────┤
│ 10.10.197.101 │ UDP       │   53 │ OPEN  │ DNS      │ DNS A query          │
│ 10.10.197.101 │ UDP       │   88 │ OPEN  │ Kerberos │ Kerberos AS-REQ      │
│ 10.10.197.101 │ UDP       │  123 │ OPEN  │ NTP      │ NTPv4 request        │
│ 10.10.197.101 │ UDP       │  389 │ OPEN  │ CLDAP    │ CLDAP root DSE query │
│ 10.10.197.103 │ UDP       │   53 │ OPEN  │ DNS      │ DNS A query          │
│ 10.10.197.103 │ UDP       │   88 │ OPEN  │ Kerberos │ Kerberos AS-REQ      │
│ 10.10.197.103 │ UDP       │  123 │ OPEN  │ NTP      │ NTPv4 request        │
│ 10.10.197.103 │ UDP       │  389 │ OPEN  │ CLDAP    │ CLDAP root DSE query │
╰───────────────┴───────────┴──────┴───────┴──────────┴──────────────────────╯
```

- UDP scan with custom number of workers and timeout with debug logging:
```
./udpz -f pretty -p 100 -t 2000 localhost --debug

2:31PM INF cmd/root.go:187 > Starting scanner
2:31PM DBG pkg/scan/scan.go:297 > Calculating unique probe count service_count=49
2:31PM DBG pkg/scan/scan.go:304 > Calculated unique probe count probe_count=93
2:31PM DBG pkg/scan/scan.go:310 > Calculated total probe count total_probes=279
2:31PM DBG pkg/scan/scan.go:316 > Resolving targets target_count=1
2:31PM DBG pkg/scan/scan.go:139 > Resolved target hostname addresses=1 target=localhost
2:31PM DBG pkg/scan/scan.go:389 > Port closed host=127.0.0.1 port=427 target=localhost
2:31PM DBG pkg/scan/scan.go:377 > Skipping closed port host=127.0.0.1 port=427 target=localhost

...
```

- Scan multiple hosts using a CIDR range:
```
./udpz -f pretty 10.10.14.0/24
```


## Supported Services

| Service name | Port(s) | Probes(s) |
| -------------| ------- | ------------- |
| Apple Remote Desktop (ARD) | 3283 | ard:generic |
| Building Automation & Control Networks (BACNet) | 47808 | bacnet:readpropertymultiple |
| BitTorrent Distributed Hash Table (DHT) | 6881 | bittorrent:dht-ping |
| Character Generator Protocol | 19 | chargen:generic |
| Connectionless Lightweight Directory Access Protocol (CLDAP) | 389 | cldap:rootdse |
| Constrained Application Protocol (CoAP) | 5683, 5684 | coap:generic |
| IBM-DB2 | 523 | db2:getaddr |
| Distributed Network Protocol 3 (DNP3) | 20000 | dnp3:requestlinkstatus |
| Domain Name System (DNS) | 53 | dns:ns, dns:a, dns:version |
| Datagram Transport Layer Security (DTLS) | 443, 2221, 3391, 4433, 5061, 5349, 10161 | dtls:client-hello |
| EtherNet/IP | 44818, 2222 | enip:list-identity |
| Factory Interface Network Service (FINS) | 9600 | fins:data-read |
| Highway Addressable Remote Transducer Industrial Protocol | 5094 | hart-ip:generic |
| HID Discovery Protocol | 4070 | hid-discovery:generic |
| Internet Key Exchange (IKE) | 500, 4500 | ike:generic, ike:malformed |
| Intelligent Platform Management Interface (IPMI) | 623 | ipmi:rmcp |
| Kerberos Key Distribution Center (KDC) | 88 | kerberos:asreq |
| KNXNet/IP (Konnex) | 3671 | knx:tunnel |
| Layer 2 Tunneling Protocol (L2TP) | 1701, 1702 | l2tp:generic |
| Lantronix Discovery | 30718 | lantronix:search |
| Multicast Domain Name System (mDNS) | 5353 | mdns:reverse |
| Mitsubishi MELSEC-Q | 5006, 5001, 5007 | melsec-q:getcpuinfo |
| Memcache | 11211 | memcache:version, memcache:stats |
| Modbus over UDP | 502 | modbus:read-coils |
| Moxa NPort | 4800, 4001 | moxa-nport:enum |
| MQTT-SN Discovery | 1884, 1885 | mqtt-sn:searchgw |
| Microsoft Windows Remote Procedure Call (MSRPC) | 135 | msrpc:ncadg-ip-udp |
| Microsoft Structured Query Language (SQL) Server | 1434 | mssql:ping |
| Network Address Translation Port Mapping Protocol (NAT-PMP) | 5351 | nat-pmp:address |
| Network Basic Input/Output System (NetBIOS) | 137 | netbios:stat |
| Network File System (NFS) | 2049 | nfs:generic |
| Network Time Protocol (NTP) | 123 | ntp:v4, ntp:v2 |
| OpenVPN (Virtual Private Networking) | 1194 | openvpn:hardresetclient |
| Symantec PCAnywhere | 5632 | pca:info |
| PCWorx | 1962 | pcworx:generic |
| Sun Remote Procedure Call (RPC) | 111 | portmap:rpc-dump |
| PROFInet Context Manager | 34964 | profinet-cm:lookup |
| Quote of the Day (QOTD) | 17 | qotd:ping |
| Remote Authentication Dial-In User Service (RADIUS) | 1812, 1645, 1813 | radius:generic |
| Remote Desktop Protocol (RDP) over UDP | 3389 | rdp:syn |
| Routing Information Protocol (RIP) | 520 | rip:v2 |
| Routing Information Protocol next generation (RIPng) | 521 | ripng:request |
| Session Initiation Protocol (SIP) | 5060, 5061, 2543 | sip:invite |
| Service Location Protocol (SLP) | 427 | slp:generic |
| Simple Network Management Protocol (SNMP) | 161, 162, 6161, 8161, 10161, 10162, 11161 | snmp:v1-get-request, snmp:v2c-get-request, snmp:v3-get-request |
| Session Traversal Utilities for NAT (STUN) | 3478, 3470, 19302, 1990 | stun:bind |
| Trivial File Transfer Protocol (TFTP) | 69, 247, 6969 | tftp:read |
| Ubiquiti Networks AirControl Management Discovery Protocol | 10001 | ubiquiti:discover-v1, ubiquiti:discover-v2 |
| Universal Plug and Play (UPnP) | 1900, 5000, 62078 | upnp:search |
| VxWorks Wind Debug Agent ONCRPC | 17185 | wdbrpc:info |
| Citrix WinFrame Remote Desktop Server | 1604 | winframe:generic |
| Web Services Discovery (WSD) | 3702 | wsd:discovery, wsd:blank |
| X Display Manager Control Protocol (XDMCP) | 177 | xdmcp:query |

## Inspiration / Credits
- [Nmap](https://nmap.org/)
- [UDPx](https://github.com/nullt3r/udpx)
