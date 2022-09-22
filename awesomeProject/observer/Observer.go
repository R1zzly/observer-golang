package observer

type Observer interface {
	HandleEvent([]string)
}
