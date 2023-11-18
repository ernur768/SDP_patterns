package main

type Observer interface {
	Update()
}

type Observable interface {
	Register(Observer)
	NotifyAll()
}

type Follower struct{}

func (f *Follower) Update() {}

type Channel struct {
	observers []Observer
}

func (c *Channel) PostNews() {
	// post news
	c.NotifyAll()
}

func (c *Channel) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *Channel) NotifyAll() {
	for _, observer := range c.observers {
		observer.Update()
	}
}

func main() {
	groupChannel := Channel{}
	groupChannel.Register(&Follower{})
	groupChannel.Register(&Follower{})

	groupChannel.PostNews()
}
