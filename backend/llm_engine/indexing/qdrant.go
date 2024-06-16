package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr                  = flag.String("addr", "localhost:6334", "the address to connect to")
	collectionName        = "TOPIC_DETAILS"
	vectorSize     uint64 = 1024
	distance              = pb.Distance_Cosine
)

func CreateQdrantCollection() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.DialContext(context.Background(), *addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// create grpc collection client
	collections_client := pb.NewCollectionsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Check Qdrant version
	qdrantClient := pb.NewQdrantClient(conn)
	healthCheckResult, err := qdrantClient.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("Could not get health: %v", err)
	} else {
		log.Printf("Qdrant version: %s", healthCheckResult.GetVersion())
	}

	// Create new collection
	_, err = collections_client.Create(ctx, &pb.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     vectorSize,
				Distance: distance,
			},
		}},
	})
	if err != nil {
		log.Fatalf("Could not create collection: %v", err)
	} else {
		log.Println("Collection", collectionName, "created")
	}

}
