chan bool
chan string
chan int
chan struct{}

e.g.
`done := chan bool` //decalre and set value
`done <- true` //send bool message to channel
`<- done` // receive message from channel

`close(done)` means notify all, and the channel not receive any message again
`<- true` means notify one, and the channel can still receive message

Declare:
`fooCh chan struct{}`
Set Value:
`fooCh : make(chan struct{})` // if no buffer size, the sending operation `fooCh <- struct{}{}` to the channel will block forever
`fooCh : make(chan struct{}, 1)` // set buffer size

Define the goroutine(light thread) anonymous function and call
```go
go func() {

}()
```
or 
```go
func foo(){

}
go foo()
```
