package stream

import (
	"context"
	"fmt"
	"sync"

	"connectrpc.com/connect"
)

type StreamInterface[T any] interface {
	Send(data *T)
	Run() error
	Close()
}

// Stream wraps a connect.ServerStream.
type Stream[T any] struct {
	mu sync.Mutex
	// stream is the underlying connect stream
	// that does the actual transfer of data
	// between the server and a client
	stream *connect.ServerStream[T]
	// context is the context of the stream
	ctx context.Context
	// The channel that we listen to for any
	// new data that we need to send to the client.
	ch chan *T
	// closed is a flag that indicates whether
	// the stream has been closed.
	closed bool
}

// newStream creates a new stream.
func NewStream[T any](ctx context.Context, st *connect.ServerStream[T]) *Stream[T] {
	return &Stream[T]{
		stream: st,
		ctx:    ctx,
		ch:     make(chan *T),
	}
}

// Close closes the stream.
func (s *Stream[T]) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.closed {
		close(s.ch)
	}
	s.closed = true
}

// Run runs the stream.
// Run will block until the stream is closed.
func (s *Stream[T]) Run() error {
	defer s.Close()
	for {
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		case data, ok := <-s.ch:
			if !ok {
				return connect.NewError(connect.CodeCanceled, fmt.Errorf("stream closed"))
			}
			if err := s.stream.Send(data); err != nil {
				return err
			}
		}
	}
}

// Send sends data to this stream's connected client.
func (s *Stream[T]) Send(data *T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.closed {
		s.ch <- data
	}
}
