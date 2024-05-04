package memory

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mo3et/itv-learn-24/kafka/kratos-event/event"
)

func TestSendAndReceive(t *testing.T) {
	send, receive := NewMemory("test")
	err := receive.Receive(context.Background(), func(ctx context.Context, event event.Event) error {
		t.Logf(fmt.Sprintf("key:%s, value:%s\n", event.Key(), event.Value()))
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 5; i++ {
		err := send.Send(context.Background(), &Message{
			key:   "kratos",
			value: []byte("hello world"),
		})
		if err != nil {
			t.Error(err)
		}
	}

	time.Sleep(5 * time.Second)
}
