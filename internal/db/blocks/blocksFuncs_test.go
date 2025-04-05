package blocks

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

// TODO запустить кластер в докере, для работы транзакций

//func TestOperationsWithTransaction(t *testing.T) {
//	var c mongonet.Client
//	err := c.Connect()
//	assert.Nil(t, err)
//	defer func(c *mongonet.Client) {
//		err := c.Disconnect()
//		assert.Nil(t, err)
//	}(&c)
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	wc := writeconcern.Majority()
//	txnOptions := options.Transaction().SetWriteConcern(wc)
//
//	session, err := c.Client.StartSession()
//	assert.Nil(t, err)
//	defer session.EndSession(ctx)
//	_, err = session.WithTransaction(ctx, func(sesCtx context.Context) (interface{}, error) {
//		id := uuid.NewString()
//
//		b := models.Block{
//			Id:     id,
//			Type:   "test",
//			NoteId: "test",
//			Content: []models.TextBlock{
//				{
//					Style: "default",
//					Text:  "Обычный",
//				},
//				{
//					Style: "italic",
//					Text:  "Это курсив",
//				},
//			},
//		}
//		if err = Create(c.Client, b, sesCtx); !assert.Nil(t, err) {
//			return nil, err
//		}
//
//		if _, err = GetAllInNote(c.Client, "test", sesCtx); !assert.Nil(t, err) {
//			return nil, err
//		}
//
//		a := []models.TextBlock{
//			{
//				Style: "bald",
//				Text:  "new жирный",
//			},
//			{
//				Style: "italic",
//				Text:  "новый курсив",
//			},
//		}
//		if err = UpdateContent(c.Client, id, a, sesCtx); !assert.Nil(t, err) {
//			return nil, err
//		}
//
//		if err = Delete(c.Client, id, sesCtx); !assert.Nil(t, err) {
//			return nil, err
//		}
//
//		return nil, errors.New("no need to commit")
//	}, txnOptions)
//}

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

	b := models.Block{
		Id:     id,
		Type:   "test",
		NoteId: "test",
		Content: []models.TextBlock{
			{
				Style: "default",
				Text:  "Обычный",
			},
			{
				Style: "italic",
				Text:  "Это курсив",
			},
		},
	}
	err = Create(c.Client, b, ctx)
	assert.Nil(t, err)

	bs, err := GetAllInNote(c.Client, "test", ctx)
	assert.Nil(t, err)

	if len(bs) == 0 {
		assert.Errorf(t, errors.New("error in GetAllInNote"), "nothing get")
	}

	a := []models.TextBlock{
		{
			Style: "bald",
			Text:  "new жирный",
		},
		{
			Style: "italic",
			Text:  "новый курсив",
		},
	}
	err = UpdateContent(c.Client, id, a, ctx)
	assert.Nil(t, err)

	err = Delete(c.Client, id, ctx)
	assert.Nil(t, err)
}
