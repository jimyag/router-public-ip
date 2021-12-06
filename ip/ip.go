package ip

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetExternalIp() (string, error) {
	//resp, err := http.Get("https://myexternalip.com/raw")
	resp, err := http.Get("https://www.taobao.com/help/getip.php")
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	content := string(c)
	content = content[16 : len(content)-3]
	return content, nil
}

func GetIpMsg(ip string) (string, error) {
	//https://ip.taobao.com/outGetIpInfo?ip=&accessKey=alibaba-inc
	resp, err := http.Get(fmt.Sprintf("https://ip.taobao.com/outGetIpInfo?ip=%s&accessKey=alibaba-inc", ip))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(c), nil

}
