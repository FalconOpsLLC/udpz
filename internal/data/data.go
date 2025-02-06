/*
Credits:
- https://weberblog.net/the-ultimate-pcap/
- https://github.com/nullt3r/udpx
*/

package data

var (
	UDP_SERVICES_DEFAULT = map[string]UdpService{
		"ard": {
			Slug:        "ard",
			NameShort:   "ARD",
			Name:        "Apple Remote Desktop (ARD)",
			Description: `Apple Remote Desktop allows Mac OS users to remotely control or monitor other Mac OS desktop sessions`,
			Ports: []uint16{
				3283,
			},
			Probes: []UdpProbe{
				{
					Slug:        "ard:generic",
					Name:        "ARD generic",
					Service:     "ard",
					EncodedData: "ABQAAQM=",
				},
			},
			Tags: []string{
				"remote-desktop",
				"macos",
				"vendor",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=3283",
				"https://wikipedia.org/wiki/Apple_Remote_Desktop",
			},
		},
		"bacnet": {
			Slug:        "bacnet",
			NameShort:   "BACNet",
			Name:        "Building Automation & Control Networks (BACNet)",
			Description: `BACNet controls communication of building automation and control systems for applications such as heating, ventilating, air-conditioning control (HVAC), lighting control, access control, and fire detection systems and their associated equipment.`,
			Ports: []uint16{
				47808,
			},
			Probes: []UdpProbe{
				{
					Slug:        "bacnet:readpropertymultiple",
					Name:        "BACNet ReadPropertyMultiple request",
					Service:     "bacnet",
					EncodedData: "gQoAJQEEAgUBDgwCAAAAHgkMCRwJLAk4CTkJOglGCU0JeAl5Hw==",
				},
			},
			Tags: []string{
				"ics",
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=47808",
				"https://wikipedia.org/wiki/BACnet",
			},
		},
		"chargen": {
			Slug:        "chargen",
			NameShort:   "CharGen",
			Name:        "Character Generator Protocol",
			Description: `The CharGen protocol generates and replies with a packet containing arbitrary characters. Should be disabled if there is no specific need for it, source for potential attacks.`,
			Ports: []uint16{
				19,
			},
			Probes: []UdpProbe{
				{
					Slug:        "chargen:generic",
					Name:        "CharGen generic",
					Service:     "chargen",
					EncodedData: "AQ==",
				},
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=19",
				"https://wikipedia.org/wiki/Character_Generator_Protocol",
				"https://datatracker.ietf.org/doc/html/rfc864",
			},
		},
		"winframe": {
			Slug:        "winframe",
			NameShort:   "Citrix WinFrame",
			Name:        "Citrix WinFrame Remote Desktop Server",
			Description: `Citrix WinFrame is an adapted version of Windows NT that allows multiple clients running MS-DOS or Microsoft Windows to connect to a centralized server and access applications over a network. The Remote Desktop component for WinFrame eventually evolved into the Remote Desktop Services component in Windows`,
			Ports: []uint16{
				1604,
			},
			Probes: []UdpProbe{
				{
					Slug:        "winframe:generic",
					Name:        "WinFrame generic",
					Service:     "winframe",
					EncodedData: "HgABMAL9qOMAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
				},
			},
			Tags: []string{
				"remote-desktop",
				"windows",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=1604",
				"https://betawiki.net/wiki/Citrix_WinFrame",
			},
		},
		"coap": {
			Slug:        "coap",
			NameShort:   "CoAP",
			Name:        "Constrained Application Protocol (CoAP)",
			Description: `CoAP is a specialized web transfer protocol for use with constrained nodes and constrained (low-power,lossy) networks often comprised of 8-bit microcontrollers with limited memory and ROM.`,
			Ports: []uint16{
				5683,
				5684, // alt
			},
			Probes: []UdpProbe{
				{
					Slug:        "coap:generic",
					Name:        "CoAP generic",
					Service:     "coap",
					EncodedData: "QAF9cLsud2VsbC1rbm93bgRjb3Jl",
				},
			},
			Tags: []string{
				"iot",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=5683",
				"https://wikipedia.org/wiki/Constrained_Application_Protocol",
				"https://datatracker.ietf.org/doc/html/rfc7252",
			},
		},
		"db2": {
			Slug:        "db2",
			NameShort:   "DB2",
			Name:        "IBM-DB2",
			Description: `DB2 is a family of data management products including database servers, developed by IBM.`,
			Ports: []uint16{
				523,
			},
			Probes: []UdpProbe{
				{
					Slug:        "db2:getaddr",
					Name:        "DB2 GETADDR Request",
					Service:     "db2",
					EncodedData: "REIyR0VUQUREUgBTUUwwODAxMAA=",
				},
			},
			Tags: []string{
				"database",
				"vendor",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=523",
				"https://wikipedia.org/wiki/IBM_Db2",
			},
		},
		"dnp3": {
			Slug:        "dnp3",
			NameShort:   "DNP3",
			Name:        "Distributed Network Protocol 3 (DNP3)",
			Description: `Communication protocol used between components in process automation systems. Its main use is in utilities such as electric and water companies.`,
			Ports: []uint16{
				20000,
			},
			Probes: []UdpProbe{
				{
					Slug:        "dnp3:requestlinkstatus",
					Name:        "DNP3 Request Link Status",
					Service:     "dnp3",
					EncodedData: "BWQFyQAAAAA2TAVkBckBAAAA3o4FZAXJAgAAAJ+EBWQFyQMAAAB3RgVkBckEAAAAHZAFZAXJBQAAAPVSBWQFyQYAAAC0WAVkBckHAAAAXJoFZAXJCAAAABm5BWQFyQkAAADxewVkBckKAAAAsHEFZAXJCwAAAFizBWQFyQwAAAAyZQVkBckNAAAA2qcFZAXJDgAAAJutBWQFyQ8AAABzbwVkBckQAAAAEesFZAXJEQAAAPkpBWQFyRIAAAC4IwVkBckTAAAAUOEFZAXJFAAAADo3BWQFyRUAAADS9QVkBckWAAAAk/8FZAXJFwAAAHs9BWQFyRgAAAA+HgVkBckZAAAA1twFZAXJGgAAAJfWBWQFyRsAAAB/FAVkBckcAAAAFcIFZAXJHQAAAP0ABWQFyR4AAAC8CgVkBckfAAAAVMgFZAXJIAAAAAFPBWQFySEAAADpjQVkBckiAAAAqIcFZAXJIwAAAEBFBWQFySQAAAAqkwVkBcklAAAAwlEFZAXJJgAAAINbBWQFyScAAABrmQVkBckoAAAALroFZAXJKQAAAMZ4BWQFySoAAACHcgVkBckrAAAAb7AFZAXJLAAAAAVmBWQFyS0AAADtpAVkBckuAAAArK4FZAXJLwAAAERsBWQFyTAAAAAm6AVkBckxAAAAzioFZAXJMgAAAI8gBWQFyTMAAABn4gVkBck0AAAADTQFZAXJNQAAAOX2BWQFyTYAAACk/AVkBck3AAAATD4FZAXJOAAAAAkdBWQFyTkAAADh3wVkBck6AAAAoNUFZAXJOwAAAEgXBWQFyTwAAAAiwQ==",
				},
			},
			Tags: []string{
				"ics",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=20000",
				"https://en.wikipedia.org/wiki/DNP3",
			},
		},
		"dns": {
			Slug:        "dns",
			NameShort:   "DNS",
			Name:        "Domain Name System (DNS)",
			Description: `DNS is a hierarchical and distributed name service that provides a naming system for computers, services, and other resources in the Internet or other Internet Protocol (IP) networks`,
			Ports: []uint16{
				53,
			},
			Probes: []UdpProbe{
				{
					Slug:        "dns:ns",
					Name:        "DNS NS query",
					Service:     "dns",
					EncodedData: "/I4BIAABAAAAAAABAAACAAEAACkE0AAAAAAADAAKAAg9I8AK+dL7Og==",
				},
				{
					Slug:        "dns:a",
					Name:        "DNS A query",
					Service:     "dns",
					EncodedData: "AAABIAABAAAAAAABCWxvY2FsaG9zdAAAAQABAAApBNAAAAAAAAwACgAIl+OWjXjQ82o=",
				},
				{
					Slug:        "dns:version",
					Name:        "DNS version.bind query",
					Service:     "dns",
					EncodedData: "nkABIAABAAAAAAABB3ZlcnNpb24EYmluZAAAEAADAAApBNAAAAAAAAwACgAI0How5NhZLqA=",
				},
			},
			Tags: []string{
				"common",
				"active-directory",
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=53",
				"https://wikipedia.org/wiki/Domain_Name_System",
			},
		},
		"ipmi": {
			Slug:        "ipmi",
			NameShort:   "IPMI",
			Name:        "Intelligent Platform Management Interface (IPMI)",
			Description: `IPMI is a set of computer interface specifications for an autonomous computer subsystem that provides management and monitoring capabilities independent of the host system's CPU, firmware (BIOS or UEFI) and operating system`,
			Ports: []uint16{
				623,
			},
			Probes: []UdpProbe{
				{
					Slug:        "ipmi:rmcp",
					Name:        "IPMI RMCP",
					Service:     "ipmi",
					EncodedData: "BgD/BwAAAAAAAAAAAAkgGMiBADiOBLU=",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=623",
				"https://wikipedia.org/wiki/Intelligent_Platform_Management_Interface",
				"https://wiki.wireshark.org/IPMI",
			},
		},
		"cldap": {
			Slug:        "cldap",
			NameShort:   "CLDAP",
			Name:        "Connectionless Lightweight Directory Access Protocol (CLDAP)",
			Description: `Connectionless Lightweight Directory Access Protocol (CLDAP) is the connectionless variant of LDAP often used to query the RootDSE of a Microsoft Windows domain controller`,
			Ports: []uint16{
				389,
			},
			Probes: []UdpProbe{
				{
					Slug:        "cldap:rootdse",
					Name:        "CLDAP root DSE query",
					Service:     "cldap",
					EncodedData: "MIQAAAAtAgEBY4QAAAAkBAAKAQAKAQACAQACAQABAQCHC29iamVjdGNsYXNzMIQAAAAAAAo=",
				},
			},
			Tags: []string{
				"common",
				"windows",
				"active-directory",
			},
			References: []string{
				"https://wikipedia.org/wiki/Lightweight_Directory_Access_Protocol",
				"https://www.speedguide.net/port.php?port=389",
				"https://wiki.wireshark.org/MS-CLDAP",
			},
		},
		"mdns": {
			Slug:        "mdns",
			NameShort:   "mDNS",
			Name:        "Multicast Domain Name System (mDNS)",
			Description: `Multicast Domain Name System (mDNS) resolves hostnames to IP addresses within small networks that do not include a local name server. It is a zero-configuration service, using essentially the same programming interfaces, packet formats and operating semantics as unicast Domain Name System (DNS)`,
			Ports: []uint16{
				5353,
			},
			Probes: []UdpProbe{
				{
					Slug:        "mdns:reverse",
					Name:        "mDNS reverse lookup",
					Service:     "mdns",
					EncodedData: "G2wBIAABAAAAAAABATEBMAEwAzEyNwdpbi1hZGRyBGFycGEAAAwAAQAAKQTQAAAAAAAMAAoACH8B+nAODI6w",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=5353",
				"https://wikipedia.org/wiki/Multicast_DNS",
			},
		},
		"memcache": {
			Slug:        "memcache",
			Name:        "Memcache",
			NameShort:   "Memcache",
			Description: `Memcache is a general-purpose distributed memory-caching system often used to speed up dynamic database-driven websites by caching data and objects in RAM to reduce the number of times an external data source must be read`,
			Ports: []uint16{
				11211,
			},
			Probes: []UdpProbe{
				{
					Slug:        "memcache:version",
					Name:        "Memcache Version",
					Service:     "memcache",
					EncodedData: "AAEAAAABAAB2ZXJzaW9uDQo=",
				},
				{
					Slug:        "memcache:stats",
					Name:        "Memcache Stats",
					Service:     "memcache",
					EncodedData: "Wk0AAAABAABzdGF0cyBpdGVtcw0K",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=11211",
				"https://wikipedia.org/wiki/Memcached",
				"https://www.wireshark.org/docs/dfref/m/memcache.html",
			},
		},
		"melsec-q": {
			Slug:        "melsec-q",
			NameShort:   "MELSEC-Q",
			Name:        "Mitsubishi MELSEC-Q",
			Description: `Mitsubishi Electric MELSEC-Q Series PLCs use a proprietary network protocol for communication. The devices are used by equipment and manufacturing facilities to provide high-speed, large volume data processing and machine control.`,
			Ports: []uint16{
				5006,
				5001,
				5007,
			},
			Probes: []UdpProbe{
				{
					Slug:        "melsec-q:getcpuinfo",
					Name:        "MELSEC-Q Get CPU Info",
					Service:     "melsec-q",
					EncodedData: "VwAAAAAREQcAAP//AwAA/gMAABQAHAgKCAAAAAAAAAAEAQEBAAAAAAE=",
				},
			},
			Tags: []string{
				"ics",
				"vendor",
			},
			References: []string{
				"https://dl.mitsubishielectric.com/dl/fa/document/manual/plc/sh080008/sh080008ab.pdf",
				"https://github.com/xl7dev/ICSecurity/blob/master/icse-nse/melsecq-discover-udp.nse",
			},
		},
		"moxa-nport": {
			Slug:        "moxa-nport",
			NameShort:   "Moxa NPort",
			Name:        "Moxa NPort",
			Description: `Moxa NPort is a line of serial device servers often used to connect card readers and payment terminals to IP-based networks`,
			Ports: []uint16{
				4800,
				4001,
			},
			Probes: []UdpProbe{
				{
					Slug:        "moxa-nport:enum",
					Name:        "Moxa NPort Enum",
					Service:     "moxa-nport",
					EncodedData: "AQAACAAAAAA=",
				},
			},
			Tags: []string{
				"ics",
				"vendor",
			},
			References: []string{
				"https://github.com/xl7dev/ICSecurity/blob/45693d87b4cc2818d0ddf4a3e8d110eb41ffeec1/icse-nse/moxa-enum.nse",
				"https://www.moxa.com/getmedia/b4396fe8-eca9-4231-a9ea-2c2ea66bf61d/moxa-udp-mode-for-nport-tech-note-v2.0.pdf",
			},
		},
		"mssql": {
			Slug:        "mssql",
			NameShort:   "MSSQL",
			Name:        "Microsoft Structured Query Language (SQL) Server",
			Description: `MSSQL is a proprietary relational database management system developed by Microsoft. As a database server, it is a software product with the primary function of storing and retrieving data`,
			Ports: []uint16{
				1434,
			},
			Probes: []UdpProbe{
				{
					Slug:        "mssql:ping",
					Name:        "MSSQL ping",
					Service:     "mssql",
					EncodedData: "Ag==",
				},
			},
			Tags: []string{
				"windows",
				"database",
				"common",
				"vendor",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=1434",
				"https://wikipedia.org/wiki/Microsoft_SQL_Server",
			},
		},
		"nat-pmp": {
			Slug:        "nat-pmp",
			NameShort:   "NAT-PMP",
			Name:        "Network Address Translation Port Mapping Protocol (NAT-PMP)",
			Description: `NAT Port Mapping Protocol (NAT-PMP) is a network protocol for automatically establishing network address translation (NAT) settings and port forwarding configurations`,
			Ports: []uint16{
				5351,
			},
			Probes: []UdpProbe{
				{
					Slug:        "nat-pmp:address",
					Name:        "NAT-PMP address request",
					Service:     "nat-pmp",
					EncodedData: "AAA=",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=5351",
				"https://wikipedia.org/wiki/NAT_Port_Mapping_Protocol",
				"https://nmap.org/nsedoc/scripts/nat-pmp-info.html",
			},
		},
		"netbios": {
			Slug:        "netbios",
			NameShort:   "NetBIOS",
			Name:        "Network Basic Input/Output System (NetBIOS)",
			Description: `NetBIOS is a protocol often associated with Microsoft Windows that allows communication of data for files and printers over a network connection.`,
			Ports: []uint16{
				137,
			},
			Probes: []UdpProbe{
				{
					Slug:        "netbios:stat",
					Name:        "NetBIOS stat",
					Service:     "netbios",
					EncodedData: "5dgAAAABAAAAAAAAIENLQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBAAAhAAE=",
				},
			},
			Tags: []string{
				"windows",
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=137",
				"https://wikipedia.org/wiki/NetBIOS",
			},
		},
		"ntp": {
			Slug:        "ntp",
			NameShort:   "NTP",
			Name:        "Network Time Protocol (NTP)",
			Description: `NTP is a networking protocol for clock synchronization between computer systems over packet-switched, variable-latency data networks.`,
			Ports: []uint16{
				123,
			},
			Probes: []UdpProbe{
				{
					Slug:        "ntp:v4",
					Name:        "NTPv4 request",
					Service:     "ntp",
					EncodedData: "4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOqVvTC5xaAA",
				},
				// TODO: NTPv3?
				{
					Slug:        "ntp:v2",
					Name:        "NTPv2 request",
					Service:     "ntp",
					EncodedData: "FwADKgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
				},
			},
			Tags: []string{
				"common",
				"internet",
				"active-directory",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=123",
				"https://wikipedia.org/wiki/Network_Time_Protocol",
			},
		},
		"fins": {
			Slug:        "fins",
			NameShort:   "FINS",
			Name:        "Factory Interface Network Service (FINS)",
			Description: `FINS is a network protocol used by Omron PLCs, over different physical networks like Ethernet, Controller Link, DeviceNet and RS-232C`,
			Ports: []uint16{
				9600,
			},
			Probes: []UdpProbe{
				{
					Slug:        "fins:data-read",
					Name:        "FINS DATA READ request",
					Service:     "fins",
					EncodedData: "gAACAAAAAGMA7wUBAA==",
				},
			},
			Tags: []string{
				"ics",
			},
			References: []string{
				"https://svn.nmap.org/nmap/scripts/omron-info.nse",
				"https://github.com/xl7dev/ICSecurity/blob/45693d87b4cc2818d0ddf4a3e8d110eb41ffeec1/icse-nse/omronudp-info.nse",
			},
		},
		"openvpn": {
			Slug:        "openvpn",
			NameShort:   "OpenVPN",
			Name:        "OpenVPN (Virtual Private Networking)",
			Description: `The OpenVPN network protocol is a virtual private network (VPN) system that creates secure point-to-point connections using SSL/TLS for key exchange.`,
			Ports: []uint16{
				1194,
			},
			Probes: []UdpProbe{
				{
					Slug:        "openvpn:hardresetclient",
					Name:        "OpenVPN HARD RESET CLIENT",
					Service:     "openvpn",
					EncodedData: "OBISEhISEhISAAAAAAA4sSbe",
				},
			},
			Tags: []string{
				"common",
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=1194",
				"https://wikipedia.org/wiki/OpenVPN",
			},
		},
		"pca": {
			Slug:        "pca",
			NameShort:   "PCAnywhere",
			Name:        "Symantec PCAnywhere",
			Description: `pcAnywhere is a discontinued suite of remote desktop software by Symantec which allows authenticated clients to connect to computers running the pcAnywhere host software.`,
			Ports: []uint16{
				5632,
			},
			Probes: []UdpProbe{
				{
					Slug:        "pca:info",
					Name:        "PCAnywhere info",
					Service:     "pca",
					EncodedData: "TlE=",
				},
			},
			Tags: []string{
				"remote-desktop",
				"vendor",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=5632",
				"https://wikipedia.org/wiki/PcAnywhere",
			},
		},
		"portmap": {
			Slug:        "portmap",
			NameShort:   "Portmap/RPC",
			Name:        "Sun Remote Procedure Call (RPC)",
			Description: `Sun RPC is a remote procedure call system originally developed by Sun Microsystems as part of the Network File System (NFS) protocol.`,
			Ports: []uint16{
				111,
			},
			Probes: []UdpProbe{
				{
					Slug:        "portmap:rpc-dump",
					Name:        "Portmap RPC dump",
					Service:     "portmap",
					EncodedData: "Gqn/4QAAAAAAAAACAAGGoAAAAAIAAAAEAAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			Tags: []string{
				"unix",
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=111",
				"https://wikipedia.org/wiki/Sun_RPC",
			},
		},
		"qotd": {
			Slug:        "qotd",
			NameShort:   "QOTD",
			Name:        "Quote of the Day (QOTD)",
			Description: `A QOTD server simply sends a short daily message/quote without regard to the input`,
			Ports: []uint16{
				17,
			},
			Probes: []UdpProbe{
				{
					Slug:        "qotd:ping",
					Name:        "QOTD Ping",
					Service:     "qotd",
					EncodedData: "AQ==",
				},
			},
			Tags: []string{
				"unix",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=17",
				"https://wikipedia.org/wiki/QOTD",
			},
		},
		"rdp": {
			Slug:        "rdp",
			NameShort:   "RDPUDP",
			Name:        "Remote Desktop Protocol (RDP) over UDP",
			Description: `Microsoft's RDP UDP transport extension enables users to connect to a remote desktop over UDP. This is often used to improve connection speed for clients.`,
			Ports: []uint16{
				3389,
			},
			Probes: []UdpProbe{
				{
					Slug:        "rdp:syn",
					Name:        "MS-RDPEUDP SYN request",
					Service:     "rdp",
					EncodedData: "/////wBAGAFZWCQHBNAE0GDa1CQRyEqhqcP+Zz+0AAAAAAAAAAAAAAAAAAAAAAAAAAEBAY+z1q2O3EHZ8BZDPHeR+OlNvIbZuF2GTI0NqoTLlJigAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
				},
			},
			Tags: []string{
				"remote-desktop",
				"windows",
				"active-directory",
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=3389",
				"https://wikipedia.org/wiki/Remote_Desktop_Protocol",
				"https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpeudp",
			},
		},
		"rip": {
			Slug:        "rip",
			NameShort:   "RIP",
			Name:        "Routing Information Protocol (RIP)",
			Description: `Routing Information Protocol (RIP) is a dynamic distance-vector routing protocol that uses hop count as a metric to find the best path between the source and destination network. RIP is often limited to smaller networks due to its maximum hop count of 15 and lack of advanced features like CIDR or built-in authentication.`,
			Ports: []uint16{
				520,
			},
			Probes: []UdpProbe{
				{
					Slug:        "rip:v2",
					Name:        "RIPv2 request",
					Service:     "rip",
					EncodedData: "AQIAAAAAAAAAAAAAAAAAAAAAAAAAAAAQ",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=520",
				"https://wikipedia.org/wiki/Routing_Information_Protocol",
			},
		},
		"ripng": {
			Slug:        "ripng",
			NameShort:   "RIPng",
			Name:        "Routing Information Protocol next generation (RIPng)",
			Description: "Routing Information Protocol next generation (RIPng) is a dynamic distance-vector routing protocol designed for IPv6 that uses hop count as a metric to find the best path between the source and destination network. RIPng is often limited to smaller networks due to its maximum hop count of 15 and lack of advanced features like CIDR or built-in authentication.",
			Ports: []uint16{
				521,
			},
			Probes: []UdpProbe{
				{
					Slug:        "ripng:request",
					Name:        "RIPng request",
					Service:     "ripng",
					EncodedData: "AQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAQ",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=521",
				"https://en.wikipedia.org/wiki/Routing_Information_Protocol",
			},
		},
		"sip": {
			Slug:        "sip",
			NameShort:   "SIP",
			Name:        "Session Initiation Protocol (SIP)",
			Description: `Session Initiation Protocol (SIP) is a signaling protocol used for initiating, maintaining, and terminating communication sessions that include voice, video and messaging applications`,
			Ports: []uint16{
				5060,
				5061,
				2543,
			},
			Probes: []UdpProbe{
				{
					Slug:        "sip:invite",
					Name:        "SIP INVITE request",
					Service:     "sip",
					EncodedData: "T1BUSU9OUyBzaXA6bm0gU0lQLzIuMA0KVmlhOiBTSVAvMi4wL1VEUCBubTticmFuY2g9Zm9vO3Jwb3J0DQpGcm9tOiA8c2lwOm5tQG5tPjt0YWc9cm9vdA0KVG86IDxzaXA6bm0yQG5tMj4NCkNhbGwtSUQ6IDUwMDAwDQpDU2VxOiA0MiBPUFRJT05TDQpNYXgtRm9yd2FyZHM6IDcwDQpDb250ZW50LUxlbmd0aDogMA0KQ29udGFjdDogPHNpcDpubUBubT4NCkFjY2VwdDogYXBwbGljYXRpb24vc2RwDQoNCg==",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=5060",
				"https://wikipedia.org/wiki/Session_Initiation_Protocol",
			},
		},
		"snmp": {
			Slug:        "snmp",
			NameShort:   "SNMP",
			Name:        "Simple Network Management Protocol (SNMP)",
			Description: `Simple Network Management Protocol (SNMP) is an Internet Standard protocol for collecting and organizing information about managed network devices, and modifying that information to change device behavior. Devices that typically support SNMP include cable modems, routers, switches, servers, workstations, printers, and more.`,
			Ports: []uint16{
				161,
				162,
				6161,
				8161,
				10161,
				10162,
				11161,
			},
			Probes: []UdpProbe{
				{
					Slug:        "snmp:v1-get-request",
					Name:        "SNMPv1 get-request",
					Service:     "snmp",
					EncodedData: "MCkCAQAEBnB1YmxpY6AcAgRWWtxdAgEAAgEAMA4wDAYIKwYBAgEBAQAFAA==",
				},
				{
					Slug:        "snmp:v2c-get-request",
					Name:        "SNMPv2c get-request",
					Service:     "snmp",
					EncodedData: "MCYCAQEEBnB1YmxpY6EZAgTcY8KaAgEAAgEAMAswCQYFKwYBAgEFAA==",
				},
				{
					Slug:        "snmp:v3-get-request",
					Name:        "SNMPv3 get-request",
					Service:     "snmp",
					EncodedData: "MDoCAQMwDwICSmkCAwD/4wQBBAIBAwQQMA4EAAIBAAIBAAQABAAEADASBAAEAKAMAgI38AIBAAIBADAA",
				},
			},
			Tags: []string{
				"common",
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=161",
				"https://wikipedia.org/wiki/Simple_Network_Management_Protocol",
			},
		},
		"stun": {
			Slug:        "stun",
			NameShort:   "STUN",
			Name:        "Session Traversal Utilities for NAT (STUN)",
			Description: `Session Traversal Utilities for NAT (STUN) is a standardized set of methods, including a network protocol, for traversal of network address translator (NAT) gateways in applications of real-time voice, video, messaging, and other interactive communications`,
			Ports: []uint16{
				3478,
				3470,
				19302,
				1990,
			},
			Probes: []UdpProbe{
				{
					Slug:        "stun:bind",
					Name:        "STUN binding request",
					Service:     "stun",
					EncodedData: "AAEAACESpEIAAAAAAAAAAAAAAAA=",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=3478",
				"https://www.speedguide.net/port.php?port=19302",
				"https://www.shadowserver.org/what-we-do/network-reporting/accessible-stun-service-report/",
				"https://www.rfc-editor.org/rfc/rfc5389",
			},
		},
		"tftp": {
			Slug:        "tftp",
			NameShort:   "TFTP",
			Name:        "Trivial File Transfer Protocol (TFTP)",
			Description: `Trivial File Transfer Protocol (TFTP) is a simple lockstep File Transfer Protocol which allows a client to get a file from or put a file onto a remote host. One of its primary uses is in the early stages of nodes booting from a local area network. TFTP has been used for this application because it is very simple to implement`,
			Ports: []uint16{
				69,
				247,
				6969,
			},
			Probes: []UdpProbe{
				{
					Slug:        "tftp:read",
					Name:        "TFTP read request",
					Service:     "tftp",
					EncodedData: "AAEvYQBuZXRhc2NpaQA=",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=69",
				"wikipedia.org/wiki/Trivial_File_Transfer_Protocol",
			},
		},
		"ubiquiti": {
			Slug:      "ubiquiti",
			NameShort: "Ubiquiti AirControl",
			Name:      "Ubiquiti Networks AirControl Management Discovery Protocol",
			Ports: []uint16{
				10001,
			},
			Probes: []UdpProbe{
				{
					Slug:        "ubiquiti:discover-v1",
					Name:        "Ubiquiti discover V1",
					Service:     "ubiquiti",
					EncodedData: "AQAAAA==",
				},
				{
					Slug:        "ubiquiti:discover-v2",
					Name:        "Ubiquiti discover V2",
					Service:     "ubiquiti",
					EncodedData: "AggAAA==",
				},
			},
			Tags: []string{
				"vendor",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=10001",
			},
		},
		"upnp": {
			Slug:        "upnp",
			NameShort:   "UPnP",
			Name:        "Universal Plug and Play (UPnP)",
			Description: `Universal Plug and Play (UPnP) is a networking protocol that allows devices to automatically discover and communicate with each other on a network, facilitating seamless device integration and service usage. It is commonly used for tasks like media streaming and port forwarding but has been criticized for potential security vulnerabilities when exposed to the internet.`,
			Ports: []uint16{
				1900,
				5000,
				62078,
			},
			Probes: []UdpProbe{
				{
					Slug:        "upnp:search",
					Name:        "UPnP search",
					Service:     "upnp",
					EncodedData: "TS1TRUFSQ0ggKiBIVFRQLzEuMQ0KSE9TVDoyMzkuMjU1LjI1NS4yNTA6MTkwMA0KU1Q6c3NkcDphbGwNCk1BTjoic3NkcDpkaXNjb3ZlciINCg0K",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=1900",
				"https://www.speedguide.net/port.php?port=5000",
				"https://wikipedia.org/wiki/Universal_Plug_and_Play",
			},
		},
		"wdbrpc": {
			Slug:        "wdbrpc",
			NameShort:   "WDBRPC",
			Name:        "VxWorks Wind Debug Agent ONCRPC",
			Description: `WDBRPC is associated with the target agent, a debugging and monitoring component of the VxWorks operating system. The target agent allows for debugging and remote control of a VxWorks-based system.`,
			Ports: []uint16{
				17185,
			},
			Probes: []UdpProbe{
				{
					Slug:        "wdbrpc:info",
					Name:        "WDBRPC info",
					Service:     "wdbrpc",
					EncodedData: "Ggn6ugAAAAAAAAACVVVVVQAAAAEAAAABAAAAAAAAAAAAAAAAAAAAAP//VRIAAAA8AAAAAQAAAAIAAAAAAAAAAA==",
				},
			},
			Tags: []string{
				"iot",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=17185",
			},
		},
		"wsd": {
			Slug:        "wsd",
			NameShort:   "WSD",
			Name:        "Web Services Discovery (WSD)",
			Description: `Web Services Discovery (WSD) is a protocol designed by Microsoft that allows computers and networked devices to discover and connect to web services on a local network. This discovery protocol is commonly used in Windows environments for automatic detection of devices like printers, scanners, and other network services without requiring manual configuration.`,
			Ports: []uint16{
				3702,
			},
			Probes: []UdpProbe{
				{
					Slug:        "wsd:discovery",
					Name:        "WSD discovery",
					Service:     "wsd",
					EncodedData: "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPHNvYXA6RW52ZWxvcGUgeG1sbnM6c29hcD0iaHR0cDovL3d3dy53My5vcmcvMjAwMy8wNS9zb2FwLWVudmVsb3BlIiB4bWxuczp3c2E9Imh0dHA6Ly9zY2hlbWFzLnhtbHNvYXAub3JnL3dzLzIwMDQvMDgvYWRkcmVzc2luZyIgeG1sbnM6d3NkPSJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA0L2Rpc2NvdmVyeSIgeG1sbnM6d3NkcD0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3MvMjAwNi8wMi9kZXZwcm9mIj4KPHNvYXA6SGVhZGVyPjx3c2E6VG8+dXJuOnNjaGVtYXMteG1sc29hcC1vcmc6d3M6MjAwNTowNDpkaXNjb3Zlcnk8L3dzYTpUbz48d3NhOkFjdGlvbj5odHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA0L2Rpc2NvdmVyeS9Qcm9iZTwvd3NhOkFjdGlvbj48d3NhOk1lc3NhZ2VJRD51cm46dXVpZDpjZTA0ZGFkMC01ZDJjLTQwMjYtOTE0Ni0xYWFiZmMxZTQxMTE8L3dzYTpNZXNzYWdlSUQ+PC9zb2FwOkhlYWRlcj48c29hcDpCb2R5Pjx3c2Q6UHJvYmU+PHdzZDpUeXBlcz53c2RwOkRldmljZTwvd3NkOlR5cGVzPjwvd3NkOlByb2JlPjwvc29hcDpCb2R5Pjwvc29hcDpFbnZlbG9wZT4K",
				},
				{
					Slug:        "wsd:blank",
					Name:        "WSD blank SOAP",
					Service:     "wsd",
					EncodedData: "PDovPgo=",
				},
			},
			Tags: []string{
				"windows",
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=3702",
				"https://wikipedia.org/wiki/Web_Services_Discovery",
			},
		},
		"xdmcp": {
			Slug:        "xdmcp",
			NameShort:   "XDMCP",
			Name:        "X Display Manager Control Protocol (XDMCP)",
			Description: `The X Display Manager Control Protocol (XDMCP) is a network protocol in the X Window System that allows client devices to remotely initiate and manage graphical sessions on a server running an X display manager. It is often used in trusted, local network environments to enable remote desktop access, though its lack of encryption makes it unsuitable for unsecured networks.`,
			Ports: []uint16{
				177,
			},
			Probes: []UdpProbe{
				{
					Slug:        "xdmcp:query",
					Name:        "XDMCP query",
					Service:     "xdmcp",
					EncodedData: "AAEAAgABAA==",
				},
			},
			Tags: []string{
				"linux",
				"remote-desktop",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=177",
				"https://en.wikipedia.org/wiki/X_display_manager",
			},
		},
		"kerberos": {
			Slug:        "kerberos",
			NameShort:   "Kerberos",
			Name:        "Kerberos Key Distribution Center (KDC)",
			Description: "Kerberos is a network authentication protocol that uses tickets to allow devices to prove their identity to others in a secure manner.",
			Ports: []uint16{
				88,
			},
			Probes: []UdpProbe{
				{
					Slug:        "kerberos:asreq",
					Name:        "Kerberos AS-REQ",
					Service:     "kerberos",
					EncodedData: "anoweKEDAgEFogMCAQqkbDBqoAcDBQBAAAAAoREwD6ADAgEBoQgwBhsEbm1hcKIGGwR0ZXN0oxkwF6ADAgECoRAwDhsGa3JidGd0GwR0ZXN0pREYDzIwMjIxMTEzMjE0NTAyWqcGAgQJSnaBqA4wDAIBEgIBEQIBEAIBFw==",
				},
			},
			Tags: []string{
				"common",
				"windows",
				"active-directory",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=88",
				"https://wikipedia.org/wiki/Kerberos_(protocol)",
			},
		},
		"lantronix": {
			Slug:        "lantronix",
			NameShort:   "Lantronix",
			Name:        "Lantronix Discovery",
			Description: "The Lantronix Discovery protocol is used to discover Lantronix devices on a given subnet.",
			Ports: []uint16{
				30718,
			},
			Probes: []UdpProbe{
				{
					Slug:        "lantronix:search",
					Name:        "Lantronix Discovery search",
					Service:     "lantronix",
					EncodedData: "AAAA9g==",
				},
			},
			Tags: []string{
				"ics",
				"iot",
			},
			References: []string{
				"http://wiki.lantronix.com/wiki/Lantronix_Discovery_Protocol",
			},
		},
		"ike": {
			Slug:        "ike",
			NameShort:   "IKE",
			Name:        "Internet Key Exchange (IKE)",
			Description: `Internet Key Exchange (IKE) is a protocol used in IPsec to establish secure, authenticated communication channels between devices over the internet. It manages the exchange of cryptographic keys and negotiates security associations, ensuring secure data transmission in VPNs and other encrypted network communications.`,
			Ports: []uint16{
				500,
				4500, // Alt
			},
			Probes: []UdpProbe{
				{
					Slug:        "ike:generic",
					Name:        "IKE generic",
					Service:     "ike",
					EncodedData: "W15kwD6ZtREAAAAAAAAAAAEQAgAAAAAAAAABUAAAATQAAAABAAAAAQAAASgBAQAIAwAAJAEB",
				},
			},
			Tags: []string{
				"common",
				"internet",
			},
			References: []string{
				"https://en.wikipedia.org/wiki/Internet_Key_Exchange",
			},
		},
		"radius": {
			Slug:        "radius",
			NameShort:   "RADIUS",
			Name:        "Remote Authentication Dial-In User Service (RADIUS)",
			Description: `Remote Authentication Dial-In User Service (RADIUS) is a network protocol that provides centralized authentication, authorization, and accounting (AAA) for users connecting to a network. It is widely used in enterprise environments to manage secure network access for remote users and devices, particularly for VPNs and wireless networks.`, Ports: []uint16{
				1812,
				1645, // Alt
				1813, // Alt
			},
			Probes: []UdpProbe{
				{
					Slug:        "radius:generic",
					Name:        "RADIUS generic",
					Service:     "radius",
					EncodedData: "AWcAV0C2ZNv11oGyrb0XaVFRGMgBB3N0ZXZlAhLbxsS3WL4U8AWzh3yeL7YBBAbAqAAcBQYAAAB7UBJfD4ZH6Mib2IE2Qmj80EUyTwwCZgAKAXN0ZXZl",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://en.wikipedia.org/wiki/RADIUS",
			},
		},
		"dtls": {
			Slug:        "dtls",
			NameShort:   "DTLS",
			Name:        "Datagram Transport Layer Security (DTLS)",
			Description: `Datagram Transport Layer Security (DTLS) is a protocol that provides encryption, data integrity, and authentication for applications using UDP, such as real-time voice or video. It is commonly used in secure, latency-sensitive applications like VPNs, VoIP, and online gaming, ensuring data protection without the reliability overhead of TCP.`,
			Ports: []uint16{
				443,  // HTTPS
				2221, // ENIP over DTLS
				3391, // Remote Desktop Gateway
				4433,
				5061,
				5349,
				10161,
			},
			Probes: []UdpProbe{
				{
					Slug:        "dtls:client-hello",
					Name:        "DTLS client hello",
					Service:     "dtls",
					EncodedData: "Fv7/AAAAAAAAAAAAwAEAALQAAAAAAAAAtP79+Cy0aiJvA+Bs0BWn891pJkw5IbKS1EQIvqHUgDv3g/UAAAA4wCzAMACfzKnMqMyqwCvALwCewCTAKABrwCPAJwBnwArAFAA5wAnAEwAzAJ0AnAA9ADwANQAvAP8BAABSAAsABAMAAQIACgAMAAoAHQAXAB4AGQAYACMAAAAWAAAAFwAAAA0AKgAoBAMFAwYDCAcICAgJCAoICwgECAUIBgQBBQEGAQMDAwEDAgQCBQIGAg==",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://wikipedia.org/wiki/Datagram_Transport_Layer_Security",
			},
		},
		"enip": {
			Slug:        "enip",
			NameShort:   "EtherNet/IP",
			Name:        "EtherNet/IP",
			Description: `Ethernet/IP is an industrial network protocol that adapts the Common Industrial Protocol (CIP) to standard Ethernet. EtherNet/IP is one of the leading industrial protocols in the United States and is widely used in a range of industries including factory, hybrid and process.`,
			Ports: []uint16{
				44818,
				2222, // Alt
			},
			Probes: []UdpProbe{
				{
					Slug:        "enip:list-identity",
					Name:        "EtherNet/IP list identity request",
					Service:     "enip",
					EncodedData: "YwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
				},
			},
			Tags: []string{
				"ics",
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=44818",
				"https://www.speedguide.net/port.php?port=2222",
				"https://wikipedia.org/wiki/EtherNet/IP",
			},
		},
		"bittorrent": {
			Slug:        "bittorrent",
			NameShort:   "BitTorrent DHT",
			Name:        "BitTorrent Distributed Hash Table (DHT)",
			Description: `The BitTorrent Distributed Hash Table (DHT) is a decentralized system that allows BitTorrent clients to locate peers for file sharing without relying on a central tracker. It distributes information about which peers have specific files across a network of nodes, enabling efficient peer discovery for sharing content.`,
			Ports: []uint16{
				6881,
			},
			Probes: []UdpProbe{
				{
					Slug:        "bittorrent:dht-ping",
					Name:        "BitTorrent DHT ping",
					Service:     "bittorrent",
					EncodedData: "ZDE6YWQyOmlkMjA6YWJjZGVmZ2hpajAxMjM0NTY3ODllMTpxNDpwaW5nMTp0MjphYTE6eTE6cWU=",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=6881",
				"http://www.bittorrent.org/beps/bep_0005.html",
			},
		},
		"msrpc": {
			Slug:        "msrpc",
			NameShort:   "MSRPC",
			Name:        "Microsoft Windows Remote Procedure Call (MSRPC)",
			Description: `Microsoft Windows Remote Procedure Call (MSRPC) is a protocol that enables inter-process communication over a network or within the same system, allowing applications to request services from remote systems as if they were local. It is widely used for Windows network services like file sharing, authentication, and managing networked resources in Windows environments.`,
			Ports: []uint16{
				135,
			},
			Probes: []UdpProbe{
				{
					Slug:        "msrpc:ncadg-ip-udp",
					Name:        "MSRPC ncadg_ip_udp bind",
					Service:     "msrpc",
					EncodedData: "BQALAxAAAABIAAAAAQAAALgQuBAAAAAAAQAAAAAAAQDE/vyZYFIbELvLAKoAITR6AAAAAARdiIrrHMkRn+gIACsQSGACAAAA",
				},
			},
			Tags: []string{
				"windows",
				"active-directory",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=135",
				"https://wikipedia.org/wiki/Microsoft_RPC",
			},
		},
		"l2tp": {
			Slug:        "l2tp",
			NameShort:   "L2TP",
			Name:        "Layer 2 Tunneling Protocol (L2TP)",
			Description: `Layer 2 Tunneling Protocol (L2TP) is a protocol used to create secure tunnels for data transmission over the internet, often in VPN setups. Operating at Layer 2, it encapsulates data to enable secure, private connections between networks, though it typically relies on IPsec to provide encryption.`,
			Ports: []uint16{
				1701,
				1702,
			},
			Probes: []UdpProbe{
				{
					Slug:        "l2tp:generic",
					Name:        "L2TP generic",
					Service:     "l2tp",
					EncodedData: "yAIAawAAAAAAAAAAgAgAAAAAAAGACAAAAAIBAIAKAAAAAwAAAAOACgAAAAQAAAAAgAgAAAAGBpCAEAAAAAcxMDkuNi4xLjcygBMAAAAIeGVsZXJhbmNlLmNvbYAIAAAACUoygAgAAAAKAAQ=",
				},
			},
			Tags: []string{
				"common",
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=1701",
				"https://wikipedia.org/wiki/Layer_2_Tunneling_Protocol",
			},
		},
		"nfs": {
			Slug:        "nfs",
			NameShort:   "NFS",
			Name:        "Network File System (NFS)",
			Description: `Network File System (NFS) is a protocol that allows file access over a network in a manner similar to local file access, enabling users to read and write files on remote servers as if they were on their local machines. Commonly used in UNIX and Linux environments, NFS facilitates file sharing and management across different systems, promoting collaboration and centralized data storage.`,
			Ports: []uint16{
				2049,
			},
			Probes: []UdpProbe{
				{
					Slug:        "nfs:generic",
					Name:        "NFS generic",
					Service:     "nfs",
					EncodedData: "ZcpyKAAAAAAAAAACAAGGowAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			Tags: []string{
				"common",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=2049",
				"https://wikipedia.org/wiki/Network_File_System",
			},
		},
		"slp": {
			Slug:        "slp",
			NameShort:   "SLP",
			Name:        "Service Location Protocol (SLP)",
			Description: `Service Location Protocol (SLP) is a network protocol used for the discovery of network services in local area networks. It enables clients to find and connect to services without needing prior knowledge of their locations, facilitating dynamic service discovery and improving network resource management.`,
			Ports: []uint16{
				427,
			},
			Probes: []UdpProbe{
				{
					Slug:        "slp:generic",
					Name:        "SLP generic",
					Service:     "slp",
					EncodedData: "AgEAADYgAAAAAAABAAJlbgAAABVzZXJ2aWNlOnNlcnZpY2UtYWdlbnQAB2RlZmF1bHQAAAAA",
				},
			},
			Tags: []string{
				"internet",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=427",
				"https://wikipedia.org/wiki/Service_Location_Protocol",
			},
		},
		"hart-ip": {
			Slug:        "hart-ip",
			NameShort:   "HART-IP",
			Name:        "Highway Addressable Remote Transducer Industrial Protocol",
			Description: `The HART Communications Protocol (Highway Addressable Remote Transducer Protocol) is an early implementation of Fieldbus, a digital industrial automation protocol. Its most notable advantage is that it can communicate over legacy wiring.`,
			Ports: []uint16{
				5094,
			},
			Probes: []UdpProbe{
				{
					Slug:        "hart-ip:generic",
					Name:        "HART-IP generic",
					Service:     "hart-ip",
					EncodedData: "AQAAAAACAA0BAAB1MA==",
				},
			},
			Tags: []string{
				"ics",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=5094",
				"https://wikipedia.org/wiki/Highway_Addressable_Remote_Transducer_Protocol",
				"https://www.shodan.io/explore/category/industrial-control-systems",
				"https://wiki.wireshark.org/HART-IP",
			},
		},
		"hid-discovery": {
			Slug:        "hid-discovery",
			NameShort:   "HID Discovery",
			Name:        "HID Discovery Protocol",
			Description: ``, // TODO
			Ports: []uint16{
				4070,
			},
			Probes: []UdpProbe{
				{
					Slug:        "hid-discovery:generic",
					Name:        "HID Discovery generic",
					Service:     "hid-discovery",
					EncodedData: "ZGlzY292ZXI7MDEzOw==",
				},
			},
			Tags: []string{
				"ics",
			},
			References: []string{
				"https://www.exploit-db.com/exploits/44992",
				"https://www.pcworld.com/article/420378/flaw-in-popular-door-controllers-allow-hackers-to-easily-unlock-secure-doors.html",
				"https://www.shodan.io/search?query=%22HID+VertX%22+port%3A4070",
			},
		},
		"pcworx": {
			Slug:        "pcworx",
			NameShort:   "PCWorx",
			Name:        "PCWorx",
			Description: ``, // TODO
			Ports: []uint16{
				1962,
			},
			Probes: []UdpProbe{
				{
					Slug:        "pcworx:generic",
					Name:        "PCWorx generic",
					Service:     "pcworx",
					EncodedData: "AQEAGgAAAAB4gAADAAxJQkVUSDAxTjBfTQA=",
				},
			},
			Tags: []string{
				"ics",
			},
			References: []string{
				"https://www.speedguide.net/port.php?port=1962",
				"https://sergiusechel.medium.com/misconfiguration-in-ilc-gsm-gprs-devices-leaves-over-1-200-ics-devices-vulnerable-to-attacks-over-82c2d4a91561",
			},
		},
		"profinet-cm": {
			Slug:        "profinet-cm",
			NameShort:   "PROFInet CM",
			Name:        "PROFInet Context Manager",
			Description: ``, // TODO
			Ports: []uint16{
				34964,
			},
			Probes: []UdpProbe{
				{
					Slug:        "profinet-cm:lookup",
					Name:        "PROFInet lookup request",
					Service:     "profinet-cm",
					EncodedData: "BAAgABAAAAAAAAAAAAAAAAAAAAAAAAAACIOv4R9dyRGRpAgAKxSg+m4xmvUxmqNZ9ZoIACf1xjYAAAAAAwAAAMgrAAACAP////9MAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAQCg3pds0RGCcQCgJELffQEAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAA",
				},
				/*{
					Slug:        "profinet-cm:read-implicit",
					Name:        "PROFInet Read Implicit request",
					Service:     "profinet-cm",
					EncodedData: "BAAIABAAAAAAAKDel2zREYJxAAEAAwFaAQCg3pds0RGCcQCgJELffduruuwdAFRDslALAWMKuv0AAAAAAQAAAAAAAAAFAP////9UAAAAAABAgAAAQAAAAECAAAAAAAAAQAAAAAAJADwBAAAKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAD4QAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
				},*/
			},
			Tags: []string{
				"ics",
			},
			References: []string{
				"https://github.com/ITI/ICS-Security-Tools/blob/master/pcaps/profinet/profinet.pcap",
			},
		},
		/*
			"epl": {
				Slug:        "epl",
				NameShort:   "EPL",
				Name:        "Ethernet Powerlink",
				Description: `Real-time protocol for standard Ethernet. It is an open protocol managed by the Ethernet POWERLINK Standardization Group (EPSG). It was introduced by Austrian automation company B&R in 2001.`,

				// [TODO]

				Tags: []string{
					"ics",
				},
				References: []string{
					"https://wiki.wireshark.org/SampleCaptures#ethernet-powerlink-v1",
					"https://wiki.wireshark.org/SampleCaptures#ethernet-powerlink-v2",
				},
			},
				"steam-hltv": {
					Slug:        "steam-hltv",
					NameShort:   "Steam HLTV",
					Name:        "Steam (Valve gaming platform) Half-Life TV",
					Description: `Half-Life TV offers the ability to have an unlimited number of spectators watching online games. They can follow the game just like they would as a spectator on the game server.`,
					References: []string{
						"https://www.speedguide.net/port.php?port=27015",
						"https://help.steampowered.com/en/faqs/view/558D-FD60-531D-98BC",
					},
					Ports: []uint16{
						27015,
					},
					Probes: []UdpProbe{},
				},
		*/
		/*
			"sentinel": {
				Slug:        "sentinel",
				NameShort:   "Sentinel RMS",
				Name:        "Sentinel RMS License Manager",
				Description: `Sentinel RMS License Manager is an on-premises network service that enforces and manages licensing in a multi-user environment.`,
				Ports: []uint16{
					5093,
				},
				Probes: []UdpProbe{
					{
						Slug: "sentinel",
					},
				},
				Tags: []string{
					"vendor",
				},
				References: []string{
					"https://www.speedguide.net/port.php?port=5093",
					"https://docs.sentinel.thalesgroup.com/ldk/LDKdocs/SPNL/LDK_SLnP_Guide/Distributing/License_Manager/020-License_Manager_Type.htm",
				},
			},
		*/
	}
)
