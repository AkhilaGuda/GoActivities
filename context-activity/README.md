# Context aware - HTTP server
- To register the process route with process handler
    - http.HandleFunc("/route",function that handle the route)
    - http.HandleFunc("/process",processHandler)
- Starting HTTP server on port 8080
    - http.ListenAndServe(":8080",nil) it return err 
    - nil indicates to use default mux that already has handler (process)
- processHandler function :
    - This function handles requests on /process endpoint
- Created new context with timeout of 10seconds
    - Passed r.Context() as parent context because each http request have its own context and we need to check if parent context is disconnected
    - The new context will automatically cancelled after 10seconds or if client closes
- To ensure resources are closed : defer cancel()
- created done channel to signal when the background task is done
- Started a go routine to simulate background work 
    - Inside loop for each iteration ensured to check for ctx is cancelled or timeout and default working status
    - After successfull task completion signalling that work is completed with close(done)
- Outside the loop waiting for either work is done or context to be cancelled 
    - case done : work completed before timeout/cancellation
    - case context done: work was cancelled due to closed connection or timeout

## How to run locally
- 