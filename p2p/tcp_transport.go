package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	// Conn is the underlying connection
	conn net.Conn
	// if we dial a connection => outbound = true
	// if we accept a connection => outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	ListenerAddress string
	Listener        net.Listener

	shakeHands HandshakeFunc
	decoder    Decoder
	mu         sync.RWMutex
	peers      map[net.Addr]Peer
}

func NewTCPTransport(addr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:      NOPHandshakeFunc,
		ListenerAddress: addr,
	}
}
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenerAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}

		fmt.Printf("New incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, false)

	if err := t.shakeHands(conn); err != nil {

	}
	// Read loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {

			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}

}
