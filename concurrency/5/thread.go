package main

type Thread struct {
	joiner chan bool
}

func (t *Thread) Join() {
	<-t.joiner
}

func StartThread(f func()) *Thread {
	t := &Thread{
		joiner: make(chan bool),
	}
	go func() {
		f()
		t.joiner <- true
	}()
	return t
}
