package database

import (
	"context"
	"fmt"
	"i-am-here/app/internal/models"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetUsers() ([]models.User, error)
	GetTestData() ([]models.Test, error)
	CreateUser(user models.User) error
}

type service struct {
	db *mongo.Client
}

var (
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
)

func New() Service {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.zdr9b.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", username, password)).SetServerAPIOptions(serverAPI))

	if err != nil {
		log.Fatal(err)

	}
	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetTestData() ([]models.Test, error) {
	collection := s.db.Database("iah").Collection("test")
	log.Println("Connected to database and collection")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var data []models.Test
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to find Data: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var testData models.Test
		if err = cursor.Decode(&testData); err != nil {
			log.Printf("Failed to decode user: %v", err)
			return nil, err
		}
		data = append(data, testData)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return data, nil
}

func (s *service) GetUsers() ([]models.User, error) {
	collection := s.db.Database("iah").Collection("users")
	log.Println("Connected to database and collection")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to find users: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			log.Printf("Failed to decode user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return users, nil
}

func (s *service) CreateUser(user models.User) error {
	collection := s.db.Database("iah").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	return err
}
