package model

import (
	"fmt"
	"net/smtp"
)

type Sender struct {
	Host     string
	Email    string
	Password string
	NickName string
	Port     int
	SSL      bool
}

func (this *Sender) Auth() (auth smtp.Auth) {
	return smtp.PlainAuth("", this.Email, this.Password, this.Host)
}

func (this *Sender) HostAddress() string {
	return fmt.Sprintf("%s:%d", this.Host, this.Port)
}
