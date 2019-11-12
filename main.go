package main

import (
	"fmt"
	"io"
	"os"
)

type Debugger bool

func (d Debugger) Debug(format string, args ...interface{}) {
	if d {
		msg := fmt.Sprintf(format, args...)
		fmt.Fprintf(os.Stderr, "[DEBUG] %s\n", msg)
	}
}

var dbg Debugger = false

type ServerConn struct{}

type iamremote struct{}

// Read reads up to len(data) bytes from the channel.
func (i *iamremote) Read(data []byte) (int, error) {
	return os.Stdin.Read(data)
}

// Write writes len(data) bytes to the channel.
func (i *iamremote) Write(data []byte) (int, error) {
	return os.Stdout.Write(data)
}

// Close signals end of channel use. No data may be sent after this
// call.
func (i *iamremote) Close() error {
	os.Stdin.Close()
	os.Stdout.Close()
	return nil
}

// CloseWrite signals the end of sending in-band
// data. Requests may still be sent, and the other side may
// still send data
func (i *iamremote) CloseWrite() error {
	return os.Stdout.Close()
}

// SendRequest sends a channel request.  If wantReply is true,
// it will wait for a reply and return the result as a
// boolean, otherwise the return value will be false. Channel
// requests are out-of-band messages so they may be sent even
// if the data stream is closed or blocked by flow control.
// If the channel is closed before a reply is returned, io.EOF
// is returned.
func (i *iamremote) SendRequest(name string, wantReply bool, payload []byte) (bool, error) {
	return false, nil
}

// Stderr returns an io.ReadWriter that writes to this channel
// with the extended data type set to stderr. Stderr may
// safely be read and written from a different goroutine than
// Read and Write respectively.
func (i *iamremote) Stderr() io.ReadWriter {
	return os.Stderr
}

func usage() {
	fmt.Fprintf(os.Stderr, "This program only can be used by Remote mode(winscp)!!!\n")
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(2)
}

func main() {
	for _, argv := range os.Args {
		if argv == "-h" || argv == "--help" {
			usage()
		}
	}
	scp := new(ServerConn)
	ch := new(iamremote)
	scp.SCPHandler(os.Args, ch)
}
