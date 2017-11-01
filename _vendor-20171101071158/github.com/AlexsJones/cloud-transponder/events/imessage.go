package event

//IMessage interface for generic messaage type
type IMessage interface {
	Ack()
	Nack()
	GetRaw() []byte
}

//Ack the incoming message
func Ack(m IMessage) {
	m.Ack()
}

//Nack the incoming message
func Nack(m IMessage) {
	m.Ack()
}

//GetRaw message
func GetRaw(m IMessage) []byte {
	return m.GetRaw()
}
