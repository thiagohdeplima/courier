package courier_test

import (
	"testing"

	"github.com/thiagohdeplima/courier"
)

var (
	conn courier.Connections = courier.Connections{}

	subs courier.Subscription = courier.Subscription{
		Name:      "SUB",
		BrokerURL: "rabbitmq://rabbitmq",
		Process: func(msg *courier.Message) (string, error) {
			return "", nil
		},
	}
)

func TestCourierConnections(t *testing.T) {
	t.Run("when instantiate should have empty subscriptions", func(t *testing.T) {
		if len(conn.Subscriptions) > 0 {
			t.Errorf("excepted a empty list, got %#v", conn)
		}
	})

	t.Run("when AddSubscription is called should append Subscriptions", func(t *testing.T) {
		conn.AddSubscription(subs)

		if len(conn.Subscriptions) != 1 {
			t.Errorf("expected to append subscriptions, got %#v", conn)
		}
	})
}
