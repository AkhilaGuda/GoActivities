package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
	// creating new context with timeout 10seconds
	// r.Context() is parent context it will be cancelled if the client disconncects
	// returned ctx context will automatically cancelled after 10seconds or if client closes
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	// to indicate user writing response processing has started
	fmt.Fprintln(w, "Processing started....")
	// creating channel to signal when background task is done
	done := make(chan struct{})

	// start go routine to simulate background task
	go func() {
		for i := 1; i <= 10; i++ {
			select {
			// if context is cancelled or time out
			case <-ctx.Done():
				err := ctx.Err()
				if err == context.Canceled {
					fmt.Println("Stopped due to Context is cancelled")
				} else {
					fmt.Println("Stopped due to Context is timeout")
				}
				return

			case <-time.After(1 * time.Second):
				fmt.Println("Working step: ", i)
			}
		}
		//signalling work is done
		close(done)
	}()
	select {
	case <-done:
		//work completed before timeout or cancellation
		fmt.Fprintln(w, "Processing completed successfully")
	case <-ctx.Done():
		//work was cancelled or timeout
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Fprintln(w, " Timeout occurred on server!")
		} else {
			fmt.Fprintln(w, " Client disconnected!")
		}
	}

}
func main() {
	//process route handler
	http.HandleFunc("/process", processHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
