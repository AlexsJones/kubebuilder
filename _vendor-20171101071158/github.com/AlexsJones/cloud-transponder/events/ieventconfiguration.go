package event

//IEventConfiguration for specific connection strings etc
type IEventConfiguration interface {
	GetTopic() (string, error)
	GetConnectionString() (string, error)
	GetSubscriptionString() (string, error)
}

//GetTopic set topic...
func GetTopic(i IEventConfiguration) (string, error) {
	return i.GetTopic()
}

//GetConnectionString to connect too
func GetConnectionString(i IEventConfiguration) (string, error) {
	return i.GetConnectionString()
}

//GetSubscriptionString ...
func GetSubscriptionString(i IEventConfiguration) (string, error) {
	return i.GetSubscriptionString()
}
