package bigquery

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
)

func GetProjectID() (projectBQID string, projectDSID string, projectNotFound error) {
	projectIDEnvBQ := "GCP_BIGQUERY_PROJECT_ID"
	projectIDEnvBQDS := "GCP_BQ_DATASET_PROJECT_ID"

	projectBQID = os.Getenv(projectIDEnvBQ)
	if projectBQID == "" {
		return "", "", errors.New("Unable to get GCP_BIGQUERY_PROJECT_ID from environment variables")
	}
	projectDSID = os.Getenv(projectIDEnvBQDS)
	if projectDSID == "" {
		return "", "", errors.New("Unable to get GCP_BQ_DATASET_PROJECT_ID from environment variables")
	}
	return projectBQID, projectDSID, nil
}

//connection in big query and return context and client dataset
func ConnectBQ(projectID string) (context.Context, *bigquery.Client, error) {
	ctx := context.Background()

	if projectID == "" {
		fmt.Println("GCP_STORAGE_PROJECT_ID environment variable must be set.")
		os.Exit(1)
	}

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	// defer client.Close()

	return ctx, client, err
}

// query returns a row iterator suitable for reading query results.
func QueryBQ(ctx context.Context, client *bigquery.Client, query string) (*bigquery.RowIterator, error) {

	statement := client.Query(query)
	return statement.Read(ctx)
}
