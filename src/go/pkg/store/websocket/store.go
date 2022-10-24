package websocket

import "context"

type ConnectionStore interface {
	// Assoicate this userId and connectionId pair
	Create(ctx context.Context, userId string, connectionId string) error
	// Note this connection is no longer use
	Delete(ctx context.Context, connectionId string) error
	// Return the list of connections for a given user
	ForUser(ctx context.Context, userId string) ([]string, error)
	// Refresh any aging timers
	Heartbeat(ctx context.Context, connectionId string) error
}
