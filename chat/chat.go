package chat

import "fmt"

type Server struct {
}

func (s *Server) SendAndReceiveMessages(obs ChatService_SendAndReceiveMessagesServer) error {
	a := make(chan *MessageRequest)
	b := make(chan error)

	go func() {
		for {
			a, b := <-obs.Recv(), <-obs.Recv()
			fmt.Printf("%v\n%v\n", a, b)
		}
	}()
}
