package websocket

type ConnectionStore interface {
	// Assoicate this userId and connectionId pair
	Create(userId string, connectionId string) error
	// Note this connection is no longer use
	Delete(connectionId string) error
	// Return the list of connections for a given user
	ForUser(userId string) ([]string, error)
	// Refresh any aging timers
	Heartbeat(connectionId string) error
}
