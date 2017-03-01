package pipe

import (
	"bufio"
	"context"
	"errors"
	"os"

	"bitbucket.org/tsetsova/decode-prototype/hub/expando"
)

// FromStdIn pipes data from stdin
func FromStdIn(ctx context.Context) *stdin {

	errors := make(chan error)

	return &stdin{
		ctx:    ctx,
		errors: errors,
	}
}

type stdin struct {
	ctx    context.Context
	errors chan error
}

// Channel returns a new channel to start processing messages
func (s *stdin) Channel() stdinChannel {
	channel := make(chan expando.Input)
	return stdinChannel{Out: channel, errors: s.errors}
}

type stdinChannel struct {
	errors chan error
	Out    chan expando.Input
}

// Errors returns a channel of errors
func (s stdinChannel) Errors() chan error {
	return s.errors
}

// Next starts the process of getting the next message
func (s stdinChannel) Next() {

	contents, err := getInputFromStdIn()

	if err != nil {
		s.errors <- err
	} else {
		s.Out <- expando.Input{Payload: contents}
	}
}

// if we are being piped some input return it else error
func getInputFromStdIn() ([]byte, error) {

	fi, err := os.Stdin.Stat()

	if err != nil {
		return []byte{}, err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return []byte{}, errors.New("input expected from stdin")
	}

	reader := bufio.NewReader(os.Stdin)

	return reader.ReadBytes('\n')
}
