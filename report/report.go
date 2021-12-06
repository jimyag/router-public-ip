package report

import (
	"crypto/tls"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"os"
	ip2 "router-public-ip/ip"
	"router-public-ip/model"
	"strconv"
)

type MyReport struct {
	Sender    *model.Sender
	Receivers []model.Receiver
	Email     *email.Email
}

func (this *MyReport) LoadSender() (*model.Sender, error) {
	opencast, err := os.Open("source/sender.csv")
	if err != nil {
		return nil, err
	}
	defer func(opencast *os.File) {
		err := opencast.Close()
		if err != nil {
		}
	}(opencast)
	read := csv.NewReader(opencast)
	csvSender, err := read.Read()
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(csvSender[4])
	if err != nil {
		return nil, err
	}
	var ssl bool
	if csvSender[5] == "TRUE" || csvSender[5] == "true" {
		ssl = true
	} else {
		ssl = false
	}
	sender := &model.Sender{
		Host:     csvSender[0],
		Email:    csvSender[1],
		Password: csvSender[2],
		NickName: csvSender[3],
		Port:     port,
		SSL:      ssl}

	this.Sender = sender
	return sender, nil
}

func (this *MyReport) LoadReceiver() ([]model.Receiver, error) {
	opencast, err := os.Open("source/receiver.csv")
	if err != nil {
		return nil, err
	}
	defer func(opencast *os.File) {
		err := opencast.Close()
		if err != nil {
		}
	}(opencast)
	read := csv.NewReader(opencast)
	csvReceiver, err := read.ReadAll()

	if err != nil {
		return nil, err
	}
	if len(csvReceiver) == 0 {
		return nil, errors.New("vacant data")
	}
	receivers := []model.Receiver{}

	for _, receiver := range csvReceiver {
		receivers = append(receivers, model.Receiver{Email: receiver[0], NickName: receiver[1]})
	}
	this.Receivers = receivers
	return receivers, nil
}
func (this *MyReport) LoadEmail() (*email.Email, error) {
	mail := &email.Email{}
	send, err := this.LoadSender()
	if err != nil {
		return nil, err
	}
	receivers, err := this.LoadReceiver()
	if err != nil {
		return nil, err
	}
	receiversE := model.ReceiverEmail(receivers)
	mail.To = receiversE
	mail.From = fmt.Sprintf("%s<%s>", send.NickName, send.Email)
	mail.Subject = "公网ip"
	ip, err := ip2.GetExternalIp()
	if err != nil {
		return nil, err
	}
	ipMsg, err := ip2.GetIpMsg(ip)
	if err != nil {
		return nil, err
	}
	mail.HTML = []byte(ipMsg)
	this.Email = mail
	return mail, nil
}

func (this *MyReport) Send() error {
	_, err := this.LoadEmail()
	if err != nil {
		return err
	}
	if this.Sender.SSL {
		err := this.Email.SendWithTLS(
			this.Sender.HostAddress(),
			this.Sender.Auth(),
			&tls.Config{
				ServerName: this.Sender.Host,
			},
		)
		if err != nil {
			return err
		}
		return nil
	}
	err = this.Email.Send(this.Sender.HostAddress(), this.Sender.Auth())
	if err != nil {
		return err
	}
	return nil
}
func (this *MyReport) LoadEmailFile() (*email.Email, error) {
	return nil, nil
}
