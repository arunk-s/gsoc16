package headers

// File: include/uapi/asm-generic/socket.h
var SockOptNameLookup = map[int]string{
	1:  "SO_DEBUG",
	2:  "SO_REUSEADDR",
	3:  "SO_TYPE",
	4:  "SO_ERROR",
	5:  "SO_DONTROUTE",
	6:  "SO_BROADCAST",
	7:  "SO_SNDBUF",
	8:  "SO_RCVBUF",
	9:  "SO_KEEPALIVE",
	10: "SO_OOBINLINE",
	11: "SO_NO_CHECK",
	12: "SO_PRIORITY",
	13: "SO_LINGER",
	14: "SO_BSDCOMPAT",
	15: "SO_REUSEPORT",
	16: "SO_PASSCRED",
	17: "SO_PEERCRED",
	18: "SO_RCVLOWAT",
	19: "SO_SNDLOWAT",
	20: "SO_RCVTIMEO",
	21: "SO_SNDTIMEO",
	22: "SO_SECURITY_AUTHENTICATION",
	23: "SO_SECURITY_ENCRYPTION_TRANSPORT",
	24: "SO_SECURITY_ENCRYPTION_NETWORK",
	25: "SO_BINDTODEVICE",
	26: "SO_ATTACH_FILTER",
	27: "SO_DETACH_FILTER",
	28: "SO_PEERNAME",
	29: "SO_TIMESTAMP",
	30: "SO_ACCEPTCONN",
	31: "SO_PEERSEC",
	32: "SO_SNDBUFFORCE",
	33: "SO_RCVBUFFORCE",
	34: "SO_PASSSEC",
	35: "SO_TIMESTAMPNS",
	36: "SO_MARK",
	37: "SO_TIMESTAMPING",
	38: "SO_PROTOCOL",
	39: "SO_DOMAIN",
	40: "SO_RXQ_OVFL",
	41: "SO_WIFI_STATUS",
	42: "SO_PEEK_OFF",
	43: "SO_NOFCS",
	44: "SO_LOCK_FILTER",
	45: "SO_SELECT_ERR_QUEUE",
	46: "SO_BUSY_POLL",
	47: "SO_MAX_PACING_RATE",
	48: "SO_BPF_EXTENSIONS",
	49: "SO_INCOMING_CPU",
	50: "SO_ATTACH_BPF",

	// PPC has these different
	116: "SO_RCVLOWAT",
	117: "SO_SNDLOWAT",
	118: "SO_RCVTIMEO",
	119: "SO_SNDTIMEO",
	120: "SO_PASSCRED",
	121: "SO_PEERCRED",
}
