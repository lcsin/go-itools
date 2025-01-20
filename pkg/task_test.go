package pkg

import (
	"context"
	"testing"
	"time"
)

func TestTimeoutTask(t *testing.T) {
	job := func(ctx context.Context) error {
		for {
			time.Sleep(time.Second * 3)
			break
		}
		for {
			time.Sleep(time.Second * 5)
			break
		}
		return nil
	}

	now := time.Now()
	if err := TimeoutTask(context.Background(), time.Second*10, job); err != nil {
		t.Log(err)
	}
	t.Log(time.Now().Sub(now))
}
