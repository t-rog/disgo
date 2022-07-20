package wrapper

import (
	"sync"
)

// spool represents a synchronized Session pool.
var spool sync.Pool

// NewSession gets a Session from a pool.
func NewSession() *Session {
	if s := spool.Get(); s != nil {
		return s.(*Session) //nolint:forcetypeassert
	}

	return new(Session)
}

// putSession puts a Session into the pool.
func putSession(s *Session) {
	s.Lock()
	defer s.Unlock()

	// reset the Session.
	s.ID = ""
	s.Seq = 0
	s.Endpoint = ""
	s.Context = nil
	s.Conn = nil
	s.heartbeat = nil
	s.manager = nil

	spool.Put(s)
}

// gpool represents a synchronized Gateway Payload pool.
var gpool sync.Pool

// getPayload gets a Gateway Payload from the pool.
func getPayload() *GatewayPayload {
	if g := gpool.Get(); g != nil {
		return g.(*GatewayPayload) //nolint:forcetypeassert
	}

	return new(GatewayPayload)
}

// putPayload puts a Gateway Payload into the pool.
func putPayload(g *GatewayPayload) {
	// reset the Gateway Payload.
	g.Op = 0
	g.Data = nil
	g.SequenceNumber = nil
	g.EventName = nil

	gpool.Put(g)
}
