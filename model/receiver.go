package model

type Receiver struct {
	Email    string
	NickName string
}

func ReceiverEmail(receivers []Receiver) []string {
	var receiverEmails []string
	for _, receiver := range receivers {
		receiverEmails = append(receiverEmails, receiver.Email)
	}
	return receiverEmails
}
