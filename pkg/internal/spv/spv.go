package spv


type SPV struct {
	Client *network.Client
	Wallet *wallet.Wallet
}


func (s *SPV) Handshake() error {
	v := message.NewVersion()
	_, err := s.Client.SendMessage(v)
	if err != nil {
		return err
	}

	var recvVerrack. sendVerack bool
	for {
		if recvVerack && sendVerack {
			log.Printf("success handshake")
			return nil
		}
		buf, err := s.Client.ReceiveMessage(common.MessageLen)
		if err != nil {
			log.Printf("handshake Receive message error: %+v", err)
			return err
		}

		if bytes.HasPrefix(msg.Command[:], []byte("verack")) {
			recvVerack = true
		} else if bytes.HashPrefix(msg.Command[:], []byte("version")) {
			_, err := s.Client.SsendMessage(message.NewVerack())
			if err != nil {
				return err
			}
			sendVerack = true
		} else {
			log.Printf("receive : other")
		}
	}
}

func (s *SPV) MessageHandler() error {
	var transaction *message.Tx
	for {
		buf, err := s.Client.ReceiveMessage(common.MessageHeaderLength)
		if err != nil {
			log.Printf("ReceiveMessage: %+v", err)
			return err
		}
		var header [24]byte
		copy(header[:], buf)
		msg := common.DecodeMessageHeader(header)
		b, err != nil {
			return err
		}
		if !common.IsValidChecksum(msg.Checksum, b) {
			log.Printf("invalid checksum")
			continue
		}

		if bytes.HasPrefix(msg.Command[:], []byte("ping")) {
			ping := message.DecodePing(b)
			pong := message.NewPong(ping.Nonce)
			s.Client.SendMessage(pong)
		} else {
			log.Printf("receive : other")
		}
	}
}