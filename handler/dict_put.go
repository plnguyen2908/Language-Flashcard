package handler

import (
	"context"
	"fmt"
	"net/http"
	"service/mongoDB"
	"service/platform/dict"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DictPut(uri string, db string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		document := dict.Word{}

		c.Bind(&document)
		client, err := mongoDB.Connect(uri)
		defer client.Disconnect(context.TODO())

		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		filter := map[string]interface{}{
			"word": document.Word,
		}

		_, err = mongoDB.FindOne(client, db, id, filter)

		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusInternalServerError)
		} else if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		} else {
			update := map[string]interface{}{
				"$push": map[string]interface{}{
					"definition": document.Definition[0],
					"type":       document.Type[0],
				},
			}

			err = mongoDB.UpsertOne(client, db, id, filter, update)

			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}

			c.Status(http.StatusAccepted)
		}
	}
}
