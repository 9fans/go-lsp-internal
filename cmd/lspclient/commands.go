package main

import (
	"context"
	"log"
	"os"

	"github.com/9fans/go-lsp-internal/lsp/protocol"
)

func initialize(server protocol.Server) (*protocol.ServerCapabilities, error) {
	log.Printf("calling initialize")
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	traceValues := protocol.TraceValues(protocol.Verbose)
	result, err := server.Initialize(context.Background(), &protocol.ParamInitialize{
		XInitializeParams: protocol.XInitializeParams{
			Trace: &traceValues,
			Capabilities: protocol.ClientCapabilities{
				Workspace: protocol.WorkspaceClientCapabilities{
					WorkspaceFolders: true,
				},
			},
		},
		WorkspaceFoldersInitializeParams: protocol.WorkspaceFoldersInitializeParams{
			WorkspaceFolders: []protocol.WorkspaceFolder{
				{
					URI:  "file://" + cwd,
					Name: "default",
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	log.Printf("got initilize result: %#v", result.Capabilities)
	err = server.Initialized(context.Background(), &protocol.InitializedParams{})
	if err != nil {
		return nil, err
	}
	log.Printf("initialized notification sent")
	return &result.Capabilities, nil
}

func definition(server protocol.Server, serverCap *protocol.ServerCapabilities) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("definition cap: %#v", serverCap.DefinitionProvider)
	switch dcap := serverCap.DefinitionProvider.Value.(type) {
	case bool:
		if !dcap {
			log.Print("server not capable")
			return
		}
	}
	def, err := server.Definition(context.Background(), &protocol.DefinitionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.DocumentURI("file://" + cwd + "/main.go"),
			},
			Position: protocol.Position{
				Line:      24,
				Character: 2,
			},
		},
	})
	if err != nil {
		log.Printf("definition failed: %v", err)
		return
	}
	log.Printf("got definition: %v\n", def)
}
