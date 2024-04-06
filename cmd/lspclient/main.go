package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/9fans/go-lsp-internal/lsp/protocol"
	"github.com/sourcegraph/jsonrpc2"
)

func run() error {
	cmd := exec.Command("gopls", "-vv", "serve", "-rpc.trace", "-debug=localhost:6600", "-logfile=gopls.log")
	writer, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = os.Stderr

	serverStarted := make(chan struct{})
	go func() {
		err := cmd.Start()
		if err != nil {
			log.Fatalf("LSP server did not start: %v", err)
		}
		log.Printf("server started")
		close(serverStarted)
		err = cmd.Wait()
		if err != nil {
			log.Fatalf("LSP server exited: %v", err)
		}
	}()
	<-serverStarted

	rwc := &readWriteCloser{
		reader,
		writer,
	}
	steam := jsonrpc2.NewBufferedStream(rwc, jsonrpc2.VSCodeObjectCodec{})
	handler := protocol.NewClientHandler(&clientHandler{}, nil)
	logger := &connLogger{}
	conn := jsonrpc2.NewConn(context.Background(), steam, handler, jsonrpc2.LogMessages(logger))
	server := protocol.NewServer(conn)
	serverCap, err := initialize(server)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		switch line {
		case "def":
			definition(server, serverCap)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("reading standard input: %v", err)
	}
	return nil
}

type readWriteCloser struct {
	io.ReadCloser
	io.WriteCloser
}

func (rwc *readWriteCloser) Close() error {
	err := rwc.WriteCloser.Close()
	err1 := rwc.ReadCloser.Close()
	if err == nil {
		err = err1
	}
	return err
}

type connLogger struct{}

func (logger *connLogger) Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}
