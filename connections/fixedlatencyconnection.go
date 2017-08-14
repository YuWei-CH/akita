package connections

import (
	"log"

	"gitlab.com/yaotsu/core"
	"gitlab.com/yaotsu/core/util"
)

// FixedLatencyConnection provides a way to connect two component directly so that
// fixed latency would happen.
type FixedLatencyConnection struct {
	EndPoints map[core.Connectable]bool
	latency   int
	freq      util.Freq
	engine    core.Engine
}

// NewFixedLatencyConnection creates a new FixedLatencyConnection object
func NewFixedLatencyConnection(engine core.Engine, latency int, freq util.Freq) *FixedLatencyConnection {
	c := FixedLatencyConnection{}
	c.EndPoints = make(map[core.Connectable]bool)
	c.latency = latency
	c.freq = freq
	c.engine = engine
	return &c
}

// Attach adds a Connectable object into the end point list of the
// FixedLatencyConnection.
func (c *FixedLatencyConnection) Attach(connectable core.Connectable) {
	c.EndPoints[connectable] = true
}

// Detach removes a Connectable from the end point list of the
// FixedLatencyConnection
func (c *FixedLatencyConnection) Detach(connectable core.Connectable) {
	if _, ok := c.EndPoints[connectable]; !ok {
		log.Panicf("connectable if not attached")
	}

	delete(c.EndPoints, connectable)
}

// Send of a FixedLatencyConnection invokes receiver's Recv method
func (c *FixedLatencyConnection) Send(req core.Req) *core.Error {
	t := c.freq.NCyclesLater(c.latency, req.SendTime())
	evt := NewDeliverEvent(t, c, req)
	c.engine.Schedule(evt)
	return nil
}

// Handle defines how the FixedLatencyConnection handles events
func (c *FixedLatencyConnection) Handle(evt core.Event) error {
	switch evt := evt.(type) {
	case *DeliverEvent:
		return c.handleDeliverEvent(evt)
	}
	return nil
}

func (c *FixedLatencyConnection) handleDeliverEvent(evt *DeliverEvent) error {
	req := evt.Req
	req.SetRecvTime(evt.Time())
	err := req.Dst().Recv(req)
	if err != nil {
		if !err.Recoverable {
			log.Fatal(err)
		} else {
			evt.SetTime(err.EarliestRetry)
			c.engine.Schedule(evt)
		}
	}
	return nil
}