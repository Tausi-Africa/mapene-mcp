// mapene-mcp — a Black Swan product. © 2026 BSAi Global Limited.
// Built on the Mifos Initiative mcp-mifosx base, which is licensed under the
// Mozilla Public License, v. 2.0 (http://mozilla.org/MPL/2.0/); that license is
// retained for the base. Black Swan modifications and the mapene-mcp / Black
// Swan / Mapene names and marks are reserved. See NOTICE.
package server

import (
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/openMF/mcp-mifosx/go/adapter"
	"github.com/openMF/mcp-mifosx/go/tools"
)

type MifosMcpServer struct {
	MCPServer *mcpserver.MCPServer
	Registry  *tools.Registry
}

func NewMifosMcpServer(fineractClient *adapter.FineractClient) *MifosMcpServer {
	appServer := mcpserver.NewMCPServer("mapene-mcp", "1.0.0-blackswan")

	registry := &tools.Registry{
		Server:   appServer,
		Fineract: fineractClient,
	}

	registry.RegisterAllTools()

	return &MifosMcpServer{
		MCPServer: appServer,
		Registry:  registry,
	}
}

func (s *MifosMcpServer) Serve() error {
	return mcpserver.ServeStdio(s.MCPServer)
}
