package courier

type Subscription struct {
	Name      string
	BrokerURL string

	Process func(msg *Message) (interface{}, error)
}
