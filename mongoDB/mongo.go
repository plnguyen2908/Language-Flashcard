package mongoDB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, error) {
	// Set connection options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// Connection established
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func InsertOne(client *mongo.Client, db string, collectionName string, document interface{}) error {

	// Uncomment these one if you want to add unduplicated document
	// _, err := findOne(client, db, collectionName, document)

	// if err == nil {
	// 	return fmt.Errorf("The document has already exists")
	// } else if err != mongo.ErrNoDocuments {
	// 	return err
	// }

	collection := client.Database(db).Collection(collectionName)

	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}

func InsertMany(client *mongo.Client, db string, collectionName string, document []interface{}) error {
	// Access the specified collection
	collection := client.Database(db).Collection(collectionName)

	// Insert the document
	_, err := collection.InsertMany(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}

func FindOne(client *mongo.Client, db string, collectionName string, filter interface{}) (*mongo.SingleResult, error) {
	collection := client.Database(db).Collection(collectionName)

	res := collection.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func Find(client *mongo.Client, db string, collectionName string, filter interface{}) (*mongo.Cursor, error) {
	collection := client.Database(db).Collection(collectionName)

	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateOne(client *mongo.Client, db string, collectionName string, filter interface{}, update interface{}) error {
	collection := client.Database(db).Collection(collectionName)

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func UpsertOne(client *mongo.Client, db string, collectionName string, filter interface{}, update interface{}) error {
	collection := client.Database(db).Collection(collectionName)

	options := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, options)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMany(client *mongo.Client, db string, collectionName string, filter interface{}, update interface{}) error {
	collection := client.Database(db).Collection(collectionName)

	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(client *mongo.Client, db string, collectionName string, filter interface{}) error {
	collection := client.Database(db).Collection(collectionName)

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMany(client *mongo.Client, db string, collectionName string, filter interface{}) error {
	collection := client.Database(db).Collection(collectionName)

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
