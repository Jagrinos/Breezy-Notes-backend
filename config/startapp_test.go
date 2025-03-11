package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStartApp(t *testing.T) {
	err := os.Chdir("..")
	assert.Nil(t, err)

	t.Run("getPrivateKey", func(t *testing.T) {
		t.Parallel()

		_, err := loadPrivateKeyFromFile()
		assert.Nil(t, err)
	})
	t.Run("getPublicKey", func(t *testing.T) {
		t.Parallel()

		_, err := loadPublicKeyFromFile()
		assert.Nil(t, err)
	})
}
