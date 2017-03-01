package conn_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"gitlab.com/yaotsu/core/conn"
)

func TestConn(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Request System")
}

type MockConnection struct {
	Connected map[conn.Connectable]bool
	ReqSent   []conn.Request
}

func NewMockConnection() *MockConnection {
	return &MockConnection{
		make(map[conn.Connectable]bool),
		make([]conn.Request, 0)}
}

func (c *MockConnection) Attach(connectable conn.Connectable) error {
	c.Connected[connectable] = true
	return nil
}

func (c *MockConnection) Detach(connectable conn.Connectable) error {
	c.Connected[connectable] = false
	return nil
}

func (c *MockConnection) Send(req conn.Request) *conn.Error {
	c.ReqSent = append(c.ReqSent, req)
	return nil
}

type MockRequest struct {
	*conn.BasicRequest
}

func NewMockRequest() *MockRequest {
	return &MockRequest{conn.NewBasicRequest()}
}
