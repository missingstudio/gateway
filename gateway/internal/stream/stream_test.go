package stream

import (
	"fmt"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/missingstudio/ai/gateway/internal/mock"
)

type Data struct {
	Msg string
}

var messages = []*Data{
	{Msg: "Hello"},
	{Msg: "World"},
	{Msg: "Foo"},
	{Msg: "Bar"},
	{Msg: "Gandalf"},
	{Msg: "Frodo"},
	{Msg: "Bilbo"},
	{Msg: "Radagast"},
	{Msg: "Sauron"},
	{Msg: "Gollum"},
}

func TestStream(t *testing.T) {
	var counter uint32
	stream := mock.NewMockStream[Data](t)
	stream.SetCounter(&counter)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := stream.Run()
		t.Log(err)
	}()

	for _, data := range messages {
		stream.Send(data)
	}

	stream.Close()
	wg.Wait()

	// A total of 10 messages should have been sent.
	if counter != 10 {
		fmt.Println(counter)
		t.Errorf("expected 10, got %d", counter)
	}

	msgMsp := make(map[string]int)
	for _, data := range stream.Messages {
		msgMsp[data.Msg]++
	}

	if len(stream.Messages) != 10 {
		t.Errorf("expected 10 messages, got %d", len(stream.Messages))
	}
	if diff := cmp.Diff(messages, stream.Messages); diff != "" {
		t.Errorf("expected %v, got %v: %s", messages, stream.Messages, diff)
	}
}
