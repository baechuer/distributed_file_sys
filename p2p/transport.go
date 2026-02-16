package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}

// Transport is anything that handles communication between the nodes in the networks
// It can be TCP, UDP, websocket, etc.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
