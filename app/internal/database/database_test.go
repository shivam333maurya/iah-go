package database

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func mustStartMongoContainer() (func(context.Context) error, error) {
	dbContainer, err := mongodb.Run(context.Background(), "mongo:latest")
	if err != nil {
		return nil, err
	}

	if err != nil {
		return dbContainer.Terminate, err
	}

	if err != nil {
		return dbContainer.Terminate, err
	}

	return dbContainer.Terminate, err
}

func TestMain(m *testing.M) {
	teardown, err := mustStartMongoContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	m.Run()

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}
}

func TestNew(t *testing.T) {
	srv := New()
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New()

	stats := srv.Health()

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}
}
