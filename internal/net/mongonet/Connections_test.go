package mongonet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Run("connection", func(t *testing.T) {
		t.Parallel()
		var c Client

		err := c.Connect()
		assert.Nil(t, err)

		err = c.Disconnect()
		assert.Nil(t, err)
	})
}
