package app_init

import (
	"golang-basics/pkg/dynamodb"
	"context"
	"time"
)

func Init() context.Context {
	ctx := context.TODO()
	time.Local = time.FixedZone("IST", 19800)

	initialiseServices(ctx)
	initialiseDynamoDbClient(ctx)
	return ctx
}

func initialiseServices(ctx context.Context) {
	// Add your service initialisation code here
	// i.e. initialise your upstream http / grpc clients
}

func initialiseDynamoDbClient(ctx context.Context) {
	// hard coding the dynamodb client initialisation
	dynamodb.InitDynamoClient(ctx)
}

func CloseConnections(ctx context.Context) {
	// Add your service close code here
	// i.e. close your upstream http & grpc clients
	// database connections, etc.
}
