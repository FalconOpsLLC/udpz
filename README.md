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
git clone https://github.com/FalconOps-Cybersecurity/udpz.git
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

1. Clone the repository:
```sh
git clone https://github.com/FalconOps-Cybersecurity/udpz.git
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

| Service name                                                 | Port(s)                                   | Probe name(s)                                                 |
| ------------------------------------------------------------ | ----------------------------------------- | ------------------------------------------------------------- |
| Apple Remote Desktop (ARD)                                   | 3283                                      | ARD generic                                                   |
| BitTorrent DHT                                               | 6881                                      | BitTorrent DHT ping                                           |
| Building Automation & Control Networks (BACNet)              | 47808                                     | BACNet ReadPropertyMultiple request                           |
| Character Generator Protocol (CharGen)                       | 19                                        | CharGen generic                                               |
| Citrix WinFrame Remote Desktop Server                        | 1604                                      | WinFrame generic                                              |
| Connectionless Lightweight Directory Access Protocol (CLDAP) | 389                                       | CLDAP root DSE query                                          |
| Constrained Application Protocol (CoAP)                      | 5683, 5684                                | CoAP generic                                                  |
| Datagram Transport Layer Security (DTLS)                     | 443, 2221, 4433, 5061, 5349, 10161        | DTLS client hello, DTLS application data                      |
| Distributed Network Protocol 3 (DNP3)                        | 20000                                     | DNP3 Request Link Status                                      |
| Domain Name System (DNS)                                     | 53                                        | DNS NS query, DNS A query (localhost), DNS version.bind query |
| EtherNet/IP                                                  | 44818, 2222                               | EtherNet/IP list identity request                             |
| Factory Interface Network Service (FINS)                     | 9600                                      | FINS DATA READ request                                        |
| HID Discovery Protocol                                       | 4070                                      | HID Discovery generic                                         |
| Highway Addressable Remote Transducer Industrial Protocol    | 5094                                      | HART-IP generic                                               |
| IBM-DB2                                                      | 523                                       | DB2 GETADDR Request                                           | 
| Intelligent Platform Management Interface (IPMI)             | 623                                       | IPMI RMCP                                                     |
| Internet Key Exchange (IKE)                                  | 500, 4500                                 | IKE generic                                                   |
| Kerberos                                                     | 88                                        | Kerberos AS-REQ                                               |
| Lantronix Discovery                                          | 30718                                     | Lantronix Discovery search                                    |
| Layer 2 Tunneling Protocol (L2TP)                            | 1702                                      | L2TP generic                                                  |
| Memcache                                                     | 11211                                     | Memcache Version, Memcache Stats                              |
| Microsoft Structured Query Language (SQL) Server             | 1434                                      | MSSQL ping                                                    |
| Microsoft Windows Remote Procedure Call (MSRPC)              | 135                                       | MSRPC ncadg_ip_udp bind                                       |
| Mitsubishi MELSEC-Q                                          | 5006                                      | MELSEC-Q Get CPU Info                                         |
| Moxa NPort                                                   | 4800, 4001                                | Moxa NPort Enum                                               |
| Multicast Domain Name System (mDNS)                          | 5353                                      | mDNS reverse lookup                                           |
| Network Address Translation Port Mapping Protocol (NAT-PMP)  | 5351                                      | NAT-PMP address request                                       |
| Network Basic Input/Output System (NetBIOS)                  | 137                                       | NetBIOS stat                                                  |
| Network File System (NFS)                                    | 2049                                      | NFS generic                                                   |
| Network Time Protocol (NTP)                                  | 123                                       | NTPv4 request, NTPv2 request                                  |
| OpenVPN (Virtual Private Networking)                         | 1194                                      | OpenVPN HARD RESET CLIENT                                     |
| PCWorx                                                       | 1962                                      | PCWorx generic                                                |
| PROFInet Context Manager                                     | 34964                                     | PROFInet Read Implicit request                                |
| Quote of the Day (QOTD)                                      | 17                                        | QOTD Ping                                                     |
| Remote Authentication Dial-In User Service (RADIUS)          | 1812, 1645, 1813                          | RADIUS generic                                                |
| Remote Desktop Protocol (RDP) over UDP                       | 3389                                      | RDPUDP SYN request                                            |
| Remote Procedure Call (RPC)                                  | 111                                       | Portmap RPC dump                                              |
| Routing Information Protocol next generation (RIPng)         | 521                                       | RIPng request                                                 |
| Routing Information Protocol (RIP)                           | 520                                       | RIPv2 request                                                 |
| Service Location Protocol (SLP)                              | 427                                       | SLP generic                                                   |
| Session Initiation Protocol (SIP)                            | 5060, 5061, 2543                          | SIP INVITE request                                            |
| Session Traversal Utilities for NAT (STUN)                   | 3478, 3470, 19302, 1990                   | STUN binding request                                          |
| Simple Network Management Protocol (SNMP) - v1, v2c, v3      | 161, 162, 6161, 8161, 10161, 10162, 11161 | SNMPv1 get-request, SNMPv2c get-request, SNMPv3 get-request   |
| Symantec PCAnywhere                                          | 5632                                      | PCAnywhere info                                               |
| Trivial File Transfer Protocol (TFTP)                        | 69, 247, 6969                             | TFTP read request                                             |
| Ubiquiti Networks AirControl Management Discovery Protocol   | 10001                                     | Ubiquiti discover V1, Ubiquiti discover V2                    |
| Universal Plug and Play (UPnP)                               | 1900, 5000                                | UPnP search                                                   |
| VxWorks Wind Debug Agent ONCRPC                              | 17185                                     | WDBRPC info                                                   |
| Web Services Discovery (WSD)                                 | 3702                                      | WSD discovery, WSD blank SOAP                                 |
| X Display Manager Control Protocol (XDMCP)                   | 177                                       | XDMCP query                                                   |

## Inspiration / Credits
- [Nmap](https://nmap.org/)
- [UDPx](https://github.com/nullt3r/udpx)
