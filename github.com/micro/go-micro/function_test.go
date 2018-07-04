package micro

import (
	"context"
	"sync"
	"testing"

	"github.com/micro/go-micro/registry/mock"
	proto "github.com/micro/go-micro/server/debug/proto"
)

func TestFunction(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	// create service
	fn := NewFunction(
		Registry(mock.NewRegistry()),
		Name("test.function"),
		AfterStart(func() error {
			wg.Done()
			return nil
		}),
	)

	// we can't test fn.Init as it parses the command line
	// fn.Init()

	ch := make(chan error, 2)

	go func() {
		// run service
		ch <- fn.Run()
	}()

	// wait for start
	wg.Wait()

	// test call debug
	req := fn.Client().NewRequest(
		"test.function",
		"Debug.Health",
		new(proto.HealthRequest),
	)

	rsp := new(proto.HealthResponse)

	err := fn.Client().Call(context.TODO(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.Status != "ok" {
		t.Fatalf("function response: %s", rsp.Status)
	}

	if err := <-ch; err != nil {
		t.Fatal(err)
	}
}
