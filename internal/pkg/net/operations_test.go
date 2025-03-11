package net

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"uasbreezy/config/views"
	"uasbreezy/internal/pkg/db/users"
	users2 "uasbreezy/pkg/users"
)

func TestUsersOperations(t *testing.T) {
	if db, err := Connect(); assert.Nil(t, err) {
		tx, err := db.Driver.Begin()
		assert.Nil(t, err)
		defer tx.Rollback()

		idtest := uuid.NewString()
		t.Run("getall", func(t *testing.T) {
			_, err = users.GetAll(tx)
			assert.Nil(t, err)
		})

		t.Run("create", func(t *testing.T) {
			err = users.Create(tx, views.User{
				Id:       idtest,
				Login:    "testlogin",
				Email:    "testemail@testemail.te",
				About:    "testabout",
				Password: "testpassword",
			})
			assert.Nil(t, err)
		})

		t.Run("auth", func(t *testing.T) {
			err = users2.Auth(tx, views.UserAuth{
				Login:    "testlogin",
				Password: "testpassword",
			})
			assert.Nil(t, err)
		})

		t.Run("getinfo", func(t *testing.T) {
			_, err = users2.GetInfo(tx, "testlogin")
			assert.Nil(t, err)
		})

		t.Run("update", func(t *testing.T) {
			err = users.Update(tx, views.UserNoId{
				Login:    "testlogin2",
				Email:    "testemail2@testemail2.te",
				About:    "testabout2",
				Password: "testpassword2",
			}, idtest)
			assert.Nil(t, err)
		})

		t.Run("delete", func(t *testing.T) {
			err = users.Delete(tx, idtest)
			assert.Nil(t, err)
		})

		err = Disconnect(db)
		assert.Nil(t, err)
	}
}
