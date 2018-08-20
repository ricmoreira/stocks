package repositories

import (
	"context"
	"log"

	"stocks/config"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/aggregateopt"
	"github.com/mongodb/mongo-go-driver/mongo/changestreamopt"
	"github.com/mongodb/mongo-go-driver/mongo/collectionopt"
	"github.com/mongodb/mongo-go-driver/mongo/countopt"
	"github.com/mongodb/mongo-go-driver/mongo/deleteopt"
	"github.com/mongodb/mongo-go-driver/mongo/distinctopt"
	"github.com/mongodb/mongo-go-driver/mongo/dropcollopt"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo/insertopt"
	"github.com/mongodb/mongo-go-driver/mongo/replaceopt"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
)

// MongoCollection is an interface to abstract the Collection for mongo
type MongoCollection interface {
	Aggregate(ctx context.Context, pipeline interface{}, opts ...aggregateopt.Aggregate) (mongo.Cursor, error)
	Clone(opts ...collectionopt.Option) (*mongo.Collection, error)
	Count(ctx context.Context, filter interface{}, opts ...countopt.Count) (int64, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...deleteopt.Delete) (*mongo.DeleteResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...deleteopt.Delete) (*mongo.DeleteResult, error)
	Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...distinctopt.Distinct) ([]interface{}, error)
	Drop(ctx context.Context, opts ...dropcollopt.DropColl) error
	Find(ctx context.Context, filter interface{}, opts ...findopt.Find) (mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...findopt.One) *mongo.DocumentResult
	FindOneAndDelete(ctx context.Context, filter interface{}, opts ...findopt.DeleteOne) *mongo.DocumentResult
	FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...findopt.ReplaceOne) *mongo.DocumentResult
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...findopt.UpdateOne) *mongo.DocumentResult
	Indexes() mongo.IndexView
	InsertMany(ctx context.Context, documents []interface{}, opts ...insertopt.Many) (*mongo.InsertManyResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...insertopt.One) (*mongo.InsertOneResult, error)
	Name() string
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...replaceopt.Replace) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...updateopt.Update) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, options ...updateopt.Update) (*mongo.UpdateResult, error)
	Watch(ctx context.Context, pipeline interface{}, opts ...changestreamopt.ChangeStream) (mongo.Cursor, error)
}

type DBCollections struct {
	StockMov MongoCollection
}

// Returns a mongo database with collections indexes set
func NewDBCollections(config *config.Config) *DBCollections {

	client, err := mongo.NewClient(config.MongoHost)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.MongoDatabaseName)

	stockMovCollection := db.Collection("stock_movements")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongo database successfully")

	return &DBCollections{
		StockMov: stockMovCollection,
	}
}
