package model

import (
	"github.com/jordan-wright/email"
)

type LoadSender interface {
	LoadSender() (*Sender, error)
}

type LoadReceiver interface {
	LoadReceiver() ([]Receiver, error)
}

type LoadEmail interface {
	LoadEmail() (*email.Email, error)
	LoadEmailFile() (*email.Email, error)
}
