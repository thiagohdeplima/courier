package courier

import "errors"

type Connections struct {
	Subscriptions []Subscription
}

func (connections *Connections) AddSubscription(subscription Subscription) {
	connections.Subscriptions = append(connections.Subscriptions, subscription)
}

func (connections *Connections) Start() error {
	if len(connections.Subscriptions) == 0 {
		return errors.New("must have at least one subscription to start")
	}

	return nil
}
