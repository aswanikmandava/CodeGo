A goroutine is a function that runs concurrently with other code in the same program

Goroutines are:
    Lightweight compared to OS threads
    Managed by the Go runtime
    Cheap to create and destroy
    Designed to scale to many concurrent tasks

You don't create or manage threads directly. Go manages it for you

When should you use a goroutine?
    you want task to run concurrently
    the task can run independently
    you don't need the result immediately
    the task might block or wait