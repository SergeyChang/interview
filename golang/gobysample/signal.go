package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig, ok := <-sigs
		if ok {
			fmt.Println(sig)
		}
		done <- struct{}{}
	}()

	fmt.Println("waiting signal..")
	<-done
	fmt.Println("exit")
}
