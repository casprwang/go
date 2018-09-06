package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	sleepAndTalk(ctx, 1*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	time.Sleep(d)
}
