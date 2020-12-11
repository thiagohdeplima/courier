package courier_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestConnectionsAddSubscription(t *testing.T) {
	t.Run("when instantiate should have empty subscriptions", func(t *testing.T) {
		assert.Empty(t, conn.Subscriptions)
	})

	t.Run("when AddSubscription is called should append Subscriptions", func(t *testing.T) {
		conn.AddSubscription(subs)

		assert.Len(t, conn.Subscriptions, 1)
	})
}
	})
}
