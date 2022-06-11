package repository

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dimaunx/go-clean-example/pkg/entity"
)

var (
	mongoUri    = os.Getenv("MONGO_URI")
	mongoUser   = os.Getenv("MONGO_USER")
	mongoPasswd = os.Getenv("MONGO_PASSWD")
)

type MongoRepo struct{}

func NewMongoRepository() *MongoRepo {
	return &MongoRepo{}
}

func NewMongoClient() (*mongo.Client, error) {
	credential := options.Credential{
		Username: mongoUser,
		Password: mongoPasswd,
	}
	opts := options.Client().ApplyURI(mongoUri).SetAuth(credential)
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (MongoRepo) Save(ctx context.Context, d *entity.Device) (string, error) {
	client, err := NewMongoClient()
	if err != nil {
		return "", err
	}

	err = client.Connect(ctx)
	if err != nil {
		return "", err
	}
	defer client.Disconnect(ctx)

	col := client.Database("devices").Collection("devices")
	_, err = col.InsertOne(ctx, d)
	if err != nil {
		return "", err
	}
	return d.Id, nil
}

func (MongoRepo) FindAll(ctx context.Context) ([]entity.Device, error) {
	client, err := NewMongoClient()
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	var results []entity.Device
	col := client.Database("devices").Collection("devices")
	cur, err := col.Find(ctx, bson.D{})
	for cur.Next(ctx) {
		var result entity.Device
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (MongoRepo) FindById(ctx context.Context, id string) (*entity.Device, error) {
	client, err := NewMongoClient()
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	col := client.Database("devices").Collection("devices")
	filter := bson.D{{"id", id}}
	var result entity.Device
	err = col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
