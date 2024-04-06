package main

import (
	"context"
	"log"

	"github.com/9fans/go-lsp-internal/lsp/protocol"
)

type clientHandler struct{}

// $/logTrace
func (c *clientHandler) LogTrace(context.Context, *protocol.LogTraceParams) error {
	log.Println("unimplemented")
	return nil
}

// $/progress
func (c *clientHandler) Progress(context.Context, *protocol.ProgressParams) error {
	log.Println("unimplmented")
	return nil
}

// client/registerCapability
func (c *clientHandler) RegisterCapability(context.Context, *protocol.RegistrationParams) error {
	log.Println("unimplmented")
	return nil
}

// client/unregisterCapability
func (c *clientHandler) UnregisterCapability(context.Context, *protocol.UnregistrationParams) error {
	log.Println("unimplmented")
	return nil
}

// telemetry/event
func (c *clientHandler) Event(context.Context, *interface{}) error {
	log.Println("unimplmented")
	return nil
}

// textDocument/publishDiagnostics
func (c *clientHandler) PublishDiagnostics(context.Context, *protocol.PublishDiagnosticsParams) error {
	log.Println("unimplmented")
	return nil
}

// window/logMessage
func (c *clientHandler) LogMessage(context.Context, *protocol.LogMessageParams) error {
	log.Println("unimplmented")
	return nil
}

// window/showDocument
func (c *clientHandler) ShowDocument(context.Context, *protocol.ShowDocumentParams) (*protocol.ShowDocumentResult, error) {
	log.Println("unimplmented")
	return nil, nil
}

// window/showMessage
func (c *clientHandler) ShowMessage(context.Context, *protocol.ShowMessageParams) error {
	log.Println("unimplmented")
	return nil
}

// window/showMessageRequest
func (c *clientHandler) ShowMessageRequest(context.Context, *protocol.ShowMessageRequestParams) (*protocol.MessageActionItem, error) {
	log.Println("unimplmented")
	return nil, nil
}

// window/workDoneProgress/create
func (c *clientHandler) WorkDoneProgressCreate(context.Context, *protocol.WorkDoneProgressCreateParams) error {
	log.Println("unimplmented")
	return nil
}

// workspace/applyEdit
func (c *clientHandler) ApplyEdit(context.Context, *protocol.ApplyWorkspaceEditParams) (*protocol.ApplyWorkspaceEditResult, error) {
	log.Println("unimplmented")
	return nil, nil
}

// workspace/codeLens/refresh
func (c *clientHandler) CodeLensRefresh(context.Context) error {
	log.Println("unimplmented")
	return nil
}

// workspace/configuration
func (c *clientHandler) Configuration(context.Context, *protocol.ParamConfiguration) ([]protocol.LSPAny, error) {
	log.Println("unimplmented")
	return nil, nil
}

// workspace/diagnostic/refresh
func (c *clientHandler) DiagnosticRefresh(context.Context) error {
	log.Println("unimplmented")
	return nil
}

// workspace/inlayHint/refresh
func (c *clientHandler) InlayHintRefresh(context.Context) error {
	log.Println("unimplmented")
	return nil
}

// workspace/inlineValue/refresh
func (c *clientHandler) InlineValueRefresh(context.Context) error {
	log.Println("unimplmented")
	return nil
}

// workspace/semanticTokens/refresh
func (c *clientHandler) SemanticTokensRefresh(context.Context) error {
	log.Println("unimplmented")
	return nil
}

// workspace/workspaceFolders
func (c *clientHandler) WorkspaceFolders(context.Context) ([]protocol.WorkspaceFolder, error) {
	log.Println("unimplmented")
	return nil, nil
}
