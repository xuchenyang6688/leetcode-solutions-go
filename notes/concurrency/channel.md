chan bool
chan string
chan int
chan struct{}

e.g.
done := chan bool
done <- true
<- done

close() means notify all, and the channel not receive any message again
<- true means notify one, and the channel can still receive message
