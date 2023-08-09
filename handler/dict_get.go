package handler

import (
	"context"
	"fmt"
	"net/http"
	"service/mongoDB"
	"service/platform/dict"

	"github.com/gin-gonic/gin"
)

func DictGet(uri string, db string) gin.HandlerFunc {
	return func(c *gin.Context) {

		client, err := mongoDB.Connect(uri)

		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		defer client.Disconnect(context.TODO())

		database := client.Database(db)

		emptyFilter := map[string]interface{}{}
		collectionNames, err := database.ListCollectionNames(context.TODO(), emptyFilter)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, collectionNames)
	}
}

func DictGetOne(uri string, db string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		client, err := mongoDB.Connect(uri)

		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		defer client.Disconnect(context.TODO())

		emptyFilter := map[string]interface{}{}
		res, err := mongoDB.Find(client, db, id, emptyFilter)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		words := make([]dict.Word, 0)
		for res.Next(context.TODO()) {
			var word dict.Word
			res.Decode(&word)
			words = append(words, word)
		}
		c.JSON(http.StatusOK, words)
	}
}
