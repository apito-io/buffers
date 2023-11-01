package interfaces

import (
	"github.com/apito-cms/buffers/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/labstack/echo/v4"
	"net/rpc"
)

// Greeter is the interface that we're exposing as a plugin.
type Greeter interface {
	Greet() string
	ExeSchema(routerCtx echo.Context, cache *shared.ApplicationCache, projectDriver ApitoProjectDB) ([]byte, error)
}

// Here is an implementation that talks over RPC
type GreeterRPC struct{ client *rpc.Client }

func (g *GreeterRPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

func (g *GreeterRPC) ExeSchema(routerCtx echo.Context, cache *shared.ApplicationCache, projectDB ApitoProjectDB) ([]byte, error) {
	var resp []byte
	err := g.client.Call("Plugin.ExeSchema", map[string]interface{}{
		"routerCtx": routerCtx,
		"cache":     cache,
		"projectDB": projectDB,
	}, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return nil, err
	}
	return resp, nil
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

func (s *GreeterRPCServer) ExeSchema(args map[string]interface{}, resp []byte) error {
	resp, err := s.Impl.ExeSchema(
		args["routerCtx"].(echo.Context),
		args["cache"].(*shared.ApplicationCache),
		args["projectDB"].(ApitoProjectDB),
	)
	if err != nil {
		return err
	}
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type GreeterPlugin struct {
	// Impl Injection
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
