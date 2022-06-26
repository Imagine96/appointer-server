package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING = "mongodb+srv://dbuser960601:bW8LGlpDktP1flP3@cluster0.2tpoq.mongodb.net/?retryWrites=true&w=majority"
const DB_NAME = "Cluster0"
const DRAFTMAN_COLLECTION_NAME = "draftman"
const SCHEDULE_COLLECTION_NAME = "schedule"

/*

	TODO

	*DB
		01 connect to mongo atlas cluster
		02 disconnect from mongo atlas cluster

	*Functionalities
		01 d.Draftman CRUD
		02 Schedule CRUD


	TO CHECK
	*DB 01 02
	*Functionalities
*/

//DB 01
func ConnectToDB() (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	return client, ctx, cancel, err
}

//DB 02
func CloseDBClient(c *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := c.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()
}

//*Functionalities 01
//readOne
func GetOneDraftmanById(c *mongo.Client, ctx context.Context, id string) (*Draftman, error) {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)
	var res bson.M
	err := collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&res)
	if err != nil {
		return nil, err
	}
	if bsonBytes, err := bson.Marshal(res); err != nil {
		return nil, err
	} else {
		var res Draftman
		if err := bson.Unmarshal(bsonBytes, &res); err != nil {
			return nil, err
		}
		return &res, nil
	}
}

//readAll
func GetDraftmanList(c *mongo.Client, ctx context.Context, limit int) ([]Draftman, error) {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	match := []Draftman{}
	res, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for res.Next(ctx) {
		var item Draftman
		if err := res.Decode(&item); err != nil {
			return nil, err
		}
		match = append(match, item)
	}

	res.Close(ctx)
	return match, nil
}

//insertOne
func InsertNewDraftMan(c *mongo.Client, ctx context.Context, d Draftman) (string, error) {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)
	if res, err := collection.InsertOne(ctx, d); err != nil {
		return "", err
	} else {
		id := res.InsertedID.(string)
		return id, nil
	}
}

//updateOne
func UpdateDraftmanById(c *mongo.Client, ctx context.Context, d *Draftman) error {

	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)

	replacement := bson.M{"$set": d}

	if _, err := collection.UpdateByID(ctx, d.Id, replacement); err != nil {
		return err
	}

	return nil
}

//deleteOne
func DeleteDraftmanById(c *mongo.Client, ctx context.Context, id string) error {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)

	if _, err := collection.DeleteOne(ctx, primitive.M{"_id": id}); err != nil {
		return err
	}

	return nil
}

//*Functionalities 02

//insertOne
func InsertNewSchedule(c *mongo.Client, ctx context.Context, s Schedule) (string, error) {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)
	if res, err := collection.InsertOne(ctx, s); err != nil {
		return "", err
	} else {
		id := res.InsertedID.(string)
		return id, nil
	}
}

//updateOne
func UpdateScheduleById(c *mongo.Client, ctx context.Context, s *Schedule) error {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)
	replacement := bson.M{"$set": s}

	if _, err := collection.UpdateByID(ctx, s.Id, replacement); err != nil {
		return err
	}

	return nil
}

// deleteOne
func DeleteScheduleById(c *mongo.Client, ctx context.Context, id string) error {
	collection := c.Database(DB_NAME).Collection(DRAFTMAN_COLLECTION_NAME)
	if _, err := collection.DeleteOne(ctx, primitive.M{"_id": id}); err != nil {
		return err
	}

	return nil
}
