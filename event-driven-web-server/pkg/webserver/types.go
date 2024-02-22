package webserver

type Protocol int8

const (
	TCP Protocol = iota
	UDP
)

func (p Protocol) String() string {
	switch p {
	case TCP:
		return "tcp"
	case UDP:
		return "udp"
	default:
		return ""
	}
}
