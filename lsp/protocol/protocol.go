// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"
	"log"

	"github.com/sourcegraph/jsonrpc2"
)

type clientDispatcher struct {
	sender *jsonrpc2.Conn
}

func NewClient(conn *jsonrpc2.Conn) Client {
	return &clientDispatcher{
		sender: conn,
	}
}

type serverDispatcher struct {
	sender *jsonrpc2.Conn
}

func NewServer(conn *jsonrpc2.Conn) Server {
	return &serverDispatcher{
		sender: conn,
	}
}

func reply(ctx context.Context, conn *jsonrpc2.Conn, id jsonrpc2.ID, result any, err error) error {
	if err != nil {
		rpcerr := &jsonrpc2.Error{
			Code:    jsonrpc2.CodeInternalError,
			Message: err.Error(),
		}
		return conn.ReplyWithError(ctx, id, rpcerr)
	}
	return conn.Reply(ctx, id, result)
}

func sendParseError(ctx context.Context, conn *jsonrpc2.Conn, id jsonrpc2.ID, err error) error {
	rpcerr := &jsonrpc2.Error{
		Code:    jsonrpc2.CodeParseError,
		Message: err.Error(),
	}
	return conn.ReplyWithError(ctx, id, rpcerr)
}

type serverHandler struct {
	server Server
}

func NewServerHandler(server Server) jsonrpc2.Handler {
	return &serverHandler{
		server: server,
	}
}

func (h *serverHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
	ok, err := serverDispatch(ctx, h.server, conn, r)
	if !ok {
		rpcerr := &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: "method not implemented",
		}
		err = conn.Reply(ctx, r.ID, rpcerr)
	}
	if err != nil {
		log.Printf("rpc reply failed: %v", err)
	}
}

type clientHandler struct {
	client Client
}

func NewClientHandler(client Client) jsonrpc2.Handler {
	return &clientHandler{
		client: client,
	}
}

func (h *clientHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
	ok, err := clientDispatch(ctx, h.client, conn, r)
	if !ok {
		rpcerr := &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: "method not implemented",
		}
		err = conn.Reply(ctx, r.ID, rpcerr)
	}
	if err != nil {
		log.Printf("rpc reply failed: %v", err)
	}
}
