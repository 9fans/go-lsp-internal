// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

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

func reply(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request, result any, err error) error {
	if r.Notif {
		// no response is sent on notification
		return err
	}
	if err != nil {
		rpcerr := &jsonrpc2.Error{
			Code:    jsonrpc2.CodeInternalError,
			Message: err.Error(),
		}
		return conn.ReplyWithError(ctx, r.ID, rpcerr)
	}
	return conn.Reply(ctx, r.ID, result)
}

func sendParseError(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request, err error) error {
	if r.Notif {
		// no response is sent on notification
		return err
	}
	rpcerr := &jsonrpc2.Error{
		Code:    jsonrpc2.CodeParseError,
		Message: err.Error(),
	}
	return conn.ReplyWithError(ctx, r.ID, rpcerr)
}

type serverHandler struct {
	server     Server
	errHandler func(error, *jsonrpc2.Request)
}

func NewServerHandler(server Server, errHandler func(error, *jsonrpc2.Request)) jsonrpc2.Handler {
	return &serverHandler{
		server:     server,
		errHandler: errHandler,
	}
}

func (h *serverHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
	ok, err := ServerDispatch(ctx, h.server, conn, r)
	if !ok {
		rpcerr := &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: "method not implemented",
		}
		err = conn.Reply(ctx, r.ID, rpcerr)
	}
	if h.errHandler != nil {
		h.errHandler(err, r)
	}
}

type clientHandler struct {
	client     Client
	errHandler func(error, *jsonrpc2.Request)
}

func NewClientHandler(client Client, errHandler func(error, *jsonrpc2.Request)) jsonrpc2.Handler {
	return &clientHandler{
		client:     client,
		errHandler: errHandler,
	}
}

func (h *clientHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
	ok, err := ClientDispatch(ctx, h.client, conn, r)
	if !ok {
		rpcerr := &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: "method not implemented",
		}
		err = conn.Reply(ctx, r.ID, rpcerr)
	}
	if h.errHandler != nil {
		h.errHandler(err, r)
	}
}
