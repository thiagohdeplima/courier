package courier

type Connections struct {
	Subscriptions []Subscription
}

func (connections *Connections) AddSubscription(subscription Subscription) {
	connections.Subscriptions = append(connections.Subscriptions, subscription)
}
