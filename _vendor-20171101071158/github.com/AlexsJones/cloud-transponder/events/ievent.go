package event

//IEvent interface
type IEvent interface {
	Connect(config IEventConfiguration) error
	Publish(m []byte) error
	Subscribe(func(IMessage)) error
	CancelSubscribe()
}

//Connect to the pubsub source
func Connect(i IEvent, c IEventConfiguration) error {
	return i.Connect(c)
}

//Publish a message
func Publish(i IEvent, m []byte) error {
	return i.Publish(m)
}

//Subscribe to a topic
func Subscribe(i IEvent, fn func(IMessage)) error {
	return i.Subscribe(fn)
}

//CancelSubscribe on any active subscription
func CancelSubscribe(i IEvent) {
	i.CancelSubscribe()
}
