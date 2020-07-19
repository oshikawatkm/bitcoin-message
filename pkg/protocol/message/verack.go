package message

type Verack struct{}

func NewVerack() protocol.Message {
	return &Verack{}
}

func NewVerack()
