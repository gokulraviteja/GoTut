# Go Notes

- Concurrency
- Parallelism
- GoRoutines
- Wait Groups
- Mutex
- Channels

# Concurrency vs Parallelism

![Screenshot 2022-01-12 at 12.50.11 PM.png](Go%20Notes%20b85fa572395d48e383aadfaa6ba082a0/Screenshot_2022-01-12_at_12.50.11_PM.png)

## Concurrency

Doing multiple tasks by splitting them into chunks. But not at the same time. ex: Eat food, watch FB, eat food, watch tv, watch FB, eat food.

## Parallelism

Doing all tasks together at the same time. ex: Calling multiple friends and assigning each task and they do the tasks parallelly

## Go Routines

Go Routines is the way how we achieve parallelism.

`Threads` - Managed by OS, Fixed Stack - 1Mb 

`Go Routines` - Managed by Go Runtime, Flexible Stack - 1Kb

`Do not communicate by sharing memory; instead, share memory by communicating`

## Wait Groups

Wait groups are like a modified version of time.sleep() 

1) Add method tells the go runtime that a new goroutine has been initialized and it has to maintain it. 

2) Done method tells that the task of that particular goroutine is finished.

3) Wait method is responsible for not exiting the program until all the go routines are done .

## Mutex

1) When a memory is shared between goroutines, locks can be applied on the memory to manage the access of memory between goroutines.

2) There is also a RWMutex which says that it allow all the reads to happen simultaneously ... but when a write is happening, it won't allow any other operation by applying a lock on the memory

## Race Condition

1) When two threads are simultaneously trying to access the same memory 

2) Can be avoided by using mutex

To check if there could be any race condition: go run   --race .

## Channels

- A channel is a way in which two goroutines can be communicated.
- Channel only allows a value to be passed in if somebody is listening to it. (be it when passing 2 values, 2 listeners should be there)

Buffered Channel can have N no of extra values.

`myChan := make(chan int,1)`

`mychan ←5`

`mychan ←4`

`←mychan`

In the above example, its having 1 extra value as a buffer, it is holding 4 in buffer though not getting consumed.

- Receive only channel :`go func(ch <-chan int, wg *sync.WaitGroup)`
- Channel cannot be closed via receive only channel
- Generally closed channels can be listened to.
- If there are no values in the channel and the channel is closed, 0 will be returned when receiving.
- If the channel is open and no values in it, and still a listener is present, it panic
- If the channel is open and there are more listeners than the senders, it panic