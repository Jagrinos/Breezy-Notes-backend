package notes

import (
	"BreeZy_Backend_vol_0/internal/net/mongonet"
	"BreeZy_Backend_vol_0/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCRUD(t *testing.T) {
	var c mongonet.Client
	err := c.Connect()
	assert.Nil(t, err)
	defer func(c *mongonet.Client) {
		err := c.Disconnect()
		assert.Nil(t, err)
	}(&c)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := uuid.NewString()

	n := models.Note{
		Id:           id,
		Title:        "test",
		CreationTime: time.Now().Local().Unix(),
		LastChange:   time.Now().Local().Unix(),
		Users: []string{
			"test",
			"test2",
		},
	}
	err = Create(c.Client, n, ctx)
	assert.Nil(t, err)

	bs, err := GetAllByUser(c.Client, "test", ctx)
	assert.Nil(t, err)

	if len(bs) == 0 {
		assert.Errorf(t, errors.New("error in GetAllInNote"), "nothing get")
	}

	err = UpdateLastChange(c.Client, id, ctx)
	assert.Nil(t, err)

	err = Delete(c.Client, id, ctx)
	assert.Nil(t, err)
}

//func TestTime(t *testing.T) {
//	tm := time.Now().Local()
//	log.Println(tm)
//	u := tm.Unix()
//	log.Println(u)
//	tmfu := time.Unix(u, 0).Local()
//	log.Println(tmfu)
//}
