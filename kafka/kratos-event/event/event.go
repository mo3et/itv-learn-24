package event

import "context"

// event 为接口声明文件

type Event interface {
	Key() string
	Value() []byte
}

type Handler func(context.Context, Event) error

type Sender interface {
	Send(ctx context.Context, msg Event) error
	Close() error
}

type Receiver interface {
	Receive(ctx context.Context, handler Handler) error
	Close() error
}
