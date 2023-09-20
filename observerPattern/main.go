package main

// interfaces -------------------------------

type Observer interface {
	update()
}

type Observable interface {
	register(Observer)
	notifyAll()
}

// structs ----------------------------------

type YoutubeUser struct {
}

type YoutubeChannel struct {
	observers []Observer
}

// methods -----------------------------------

func (y YoutubeUser) update() {
	println("hooray, a new video has been released")
}

func (y *YoutubeChannel) register(observer Observer) {
	y.observers = append(y.observers, observer)
}

func (y YoutubeChannel) notifyAll() {
	for _, observer := range y.observers {
		observer.update()
	}
}

func main() {
	toples := YoutubeChannel{observers: make([]Observer, 0, 3)}

	user_nt9e := YoutubeUser{}
	user_ngf9 := YoutubeUser{}
	user_h56j := YoutubeUser{}

	toples.register(user_nt9e)
	toples.register(user_ngf9)
	toples.register(user_h56j)

	toples.notifyAll()
}
