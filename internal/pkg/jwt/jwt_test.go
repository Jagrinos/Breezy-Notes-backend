package jwt

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"uasbreezy/config"
)

func TestJWT(t *testing.T) {
	err := os.Chdir(filepath.Join("..", "..", ".."))
	assert.Nil(t, err)

	if err := config.SetupKeys(); assert.Nil(t, err) {
		var at, rt string
		t.Run("generate_tokens", func(t *testing.T) {
			t.Run("access_token", func(t *testing.T) {
				t.Parallel()

				var err error
				at, err = GenerateToken("testLogin", "ACCESS")
				assert.Nil(t, err)
			})
			t.Run("refresh_token", func(t *testing.T) {
				t.Parallel()

				var err error
				rt, err = GenerateToken("testLogin", "REFRESH")
				assert.Nil(t, err)
			})
		})
		t.Run("work_with_tokens", func(t *testing.T) {
			t.Run("get_login_from_tokens", func(t *testing.T) {
				t.Run("get_login_from_access_token", func(t *testing.T) {
					t.Parallel()

					var err error
					_, err = GetLoginFromToken(at)
					assert.Nil(t, err)
				})
				t.Run("get_login_from_refresh_token", func(t *testing.T) {
					t.Parallel()

					var err error
					_, err = GetLoginFromToken(rt)
					assert.Nil(t, err)
				})
			})

			t.Run("verify_tokens", func(t *testing.T) {
				t.Run("verify_access_token", func(t *testing.T) {
					t.Parallel()

					var err error
					_, err = VerifyToken(at)
					assert.Nil(t, err)
				})
				t.Run("verify_refresh_token", func(t *testing.T) {
					t.Parallel()

					var err error
					_, err = VerifyToken(rt)
					assert.Nil(t, err)
				})
			})

			t.Run("refresh_token", func(t *testing.T) {
				t.Parallel()

				newat, err := Refresh(rt, "testLogin")
				assert.Nil(t, err)

				_, err = VerifyToken(newat)
				assert.Nil(t, err)
			})
		})
	}
}
