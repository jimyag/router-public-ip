package ip

import (
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

type Ip struct {
}

func (this *Ip) GetExternalIp() (net.IP, error) {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	content, _ := ioutil.ReadAll(resp.Body)
	ip := net.ParseIP(string(content))
	if ip == nil {
		return nil, errors.New("ip:" + string(content) + "format error")
	}
	return ip, nil
}
