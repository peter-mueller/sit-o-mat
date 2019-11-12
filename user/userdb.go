package user

import (
	"context"
	"gocloud.dev/docstore"
	"log"
	"os"
)

func userCollection() *docstore.Collection {
	ctx := context.Background()

	url := lookupEnv("SITOMAT_COLLECTION_USER", "mem://user/name")
	coll, err := docstore.OpenCollection(ctx, url)
	if err != nil {
		panic(err)
	}
	return coll
}

func lookupEnv(env string, alternative string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		log.Printf("Using default for %v: %v", env, alternative)
		return alternative
	}
	return value
}
