package MongoDB

import (
	"banking_ledger/internal/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

type MongoTxRepository struct {
	TransactionCollection *mongo.Collection
}

func NewMongoDBAccountRepository() (*MongoTxRepository, error) {
	Port := os.Getenv("Mongodb_port")
	dbName := os.Getenv("Db_name")
	collectionName := os.Getenv("collection")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%s", Port))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err.Error())
	}
	MongoClient = client
	TransactionCollection := client.Database(dbName).Collection(collectionName)

	return &MongoTxRepository{TransactionCollection: TransactionCollection}, nil
}

func (TxCollection *MongoTxRepository) CreateTransaction(ctx context.Context, tx *models.Transaction) (any, error) {
	doc := bson.M{
		"account_id":  tx.AccountID,
		"amount":      tx.Amount,
		"type":        tx.Type,
		"description": tx.Description,
		"timestamp":   time.Now(),
	}
	result, err := TxCollection.TransactionCollection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	log.Println("Tranaction is updated to DB", doc)

	return result.InsertedID, nil
}

func (TxCollection *MongoTxRepository) GetTansaction(ctx context.Context, id string) (*models.Transaction, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	var account models.Transaction

	if err := TxCollection.TransactionCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&account); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no tranactions found")
		}
		return nil, err
	}
	return &account, nil

}

func (TxCollection *MongoTxRepository) GetTransactions(ctx context.Context, fromDate, toDate *time.Time,
	account_id string) (*[]models.Transaction, error) {
	filter := bson.M{
		"account_id": account_id,
		"timestamp": bson.M{
			"$gte": fromDate,
			"$lte": toDate.Add(24 * time.Hour),
		},
	}

	log.Default().Print(filter)

	curser, err := TxCollection.TransactionCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer curser.Close(ctx)

	var res []models.Transaction
	if err := curser.All(ctx, &res); err != nil {
		return nil, err
	}

	fmt.Println(res)

	return &res, nil

}
