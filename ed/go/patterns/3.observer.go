package main

type MassMedia interface { // Observer
	Notify(topic string)
}

type News struct { // Subject
	Topic       string
	subscribers []MassMedia
}

func (n *News) AttachObserver(mm MassMedia) {
	n.subscribers = append(n.subscribers, mm)
}

func (n News) Publish() {
	for i := 0; i < len(n.subscribers); i++ {
		n.subscribers[i].Notify(n.Topic)
	}
}

type TV struct {
}

func (tv TV) Notify(topic string) {
	println("Breaking news:", topic)
}

type Press struct {
}

func (p Press) Notify(topic string) {
	println("Fresh press:", topic)
}

func main() {
	n := News{Topic: "Real Madrid wins at Bayern Munich."}
	n.AttachObserver(TV{})
	n.AttachObserver(Press{})
	n.Publish()
}
