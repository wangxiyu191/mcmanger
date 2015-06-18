package mc

//Minecraft操作
import (
	"github.com/thinkofdeath/steven/protocol"
)

var (
	address string
	conn    protocol.Conn
)

func Init(addr string) {
	address = addr

}

func Status() (status protocol.StatusReply, err error) {
	conn, err := protocol.Dial(address)
	if err != nil {
		return protocol.StatusReply{}, err
	}
	defer conn.Close()

	response, _, err := conn.RequestStatus()
	if err != nil {
		return protocol.StatusReply{}, err
	}

	return response, nil
}
