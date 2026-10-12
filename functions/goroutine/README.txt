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

A program(main goroutine) that runs goroutine by default uses a single core. There is a component called scheduler that spawns a new goroutine (child goroutine)
when one goroutine finishes execution or hits a blocking call (eg: http.Get()) where it waits for a response from the web server.
scheduler will pause the current and gives cpu time to another goroutine.

scheduler has to monitor which goroutines the cpu time has to be allocated and constantly shuffle among the child goroutines.
This is referred to as concurrent execution of goroutines

Go can configured to leverage multiple cores, in that case schedule will assign one core to one child goroutine. This is referred
to as parallel execution of goroutines.

In the Go programming language, GOMAXPROCS is a runtime setting that controls the maximum number of operating system (OS) threads that can execute user-level Go code simultaneously.
Essentially, it determines the "available parallelism" of your application by defining how many CPU cores Go can utilize at any given split second.

The Go scheduler relies on three internal abstractions to manage concurrency:
    • G (Goroutine): The lightweight user-space thread containing your actual code.
    • M (Machine): A heavy, standard operating system thread.
    • P (Processor): A logical context or processing unit required to run Go code.

GOMAXPROCS directly sets the number of P contexts. Because an M (OS thread) needs a P to run a G (Goroutine), 
GOMAXPROCS acts as a hard cap on how many goroutines can execute in true parallel across your CPU cores at any single moment

Channel
-------
A channel can be initialized using make() that tells to create a channel that passes data of a specific type (eg:- string)
    c := make(chan string)

Channel allows to communicate from a main goroutine to a child goroutine and vice-versa
------------------------------------------------------------------
                                                Child goroutine-1
Main Goroutine -----------> Channel <---------- Child goroutine-2
                                                Child goroutine-3
------------------------------------------------------------------

Sending data with channels
--------------------------
c <- "hello"      // this tells to send "hello" into the channel variable
mydata <- c       // wait for the data to be sent into channel variable. when we get one, assign the data to mydata variable
fmt.Println(<- c) // wait for the data to be sent into channel variable. when we get one, print it out immediately