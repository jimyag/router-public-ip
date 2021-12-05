package ip

import "net"

type myIp struct {
	Ip
}

func (this *myIp) GetExternalIp() (net.IP, error) {
	return this.Ip.GetExternalIp()
}
