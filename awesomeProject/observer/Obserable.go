package observer

type Obserable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	notifyAll()
}
