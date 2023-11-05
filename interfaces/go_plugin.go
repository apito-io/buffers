package interfaces

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type CodeGenerator interface {
	Greet() string
	GenCodes(projectId string, rebuild bool) (string, error)
}

type CodeGeneratorRPC struct{ client *rpc.Client }

func (g *CodeGeneratorRPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}
	return resp
}

func (g *CodeGeneratorRPC) GenCodes(projectId string, rebuild bool) (string, error) {
	var resp string
	err := g.client.Call("Plugin.GenCodes", map[string]interface{}{
		"project_id": projectId,
		"rebuild":    rebuild,
	}, &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

type CodeGeneratorRPCServer struct {
	// This is the real implementation
	Impl CodeGenerator
}

func (s *CodeGeneratorRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

func (s *CodeGeneratorRPCServer) GenCodes(args map[string]interface{}, resp *string) error {
	var err error
	*resp, err = s.Impl.GenCodes(args["project_id"].(string), args["rebuild"].(bool))
	return err
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
type CodeGeneratorPlugin struct {
	// Impl Injection
	Impl CodeGenerator
}

func (p *CodeGeneratorPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &CodeGeneratorRPCServer{Impl: p.Impl}, nil
}

func (CodeGeneratorPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &CodeGeneratorRPC{client: c}, nil
}
