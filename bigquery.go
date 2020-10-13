package bigquery

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/bigquery"
)

//export GOOGLE_APPLICATION_CREDENTIALS=/home/jean-ferreira/go/src/github.com/jeansferreira/bigquery/projeto-bigquery-291719-accfbedbbde6.json

//Get project ID
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

	projectBQID, projectDSID, err := GetProjectID()
	if err != nil {
		log.Fatalf("bigquery.QueryBQ: %v", err)
	}

	query = strings.Replace(query, "projectBQID", projectBQID, -1)
	query = strings.Replace(query, "projectDSID", projectDSID, -1)

	ctx, client, err = ConnectBQ(projectBQID)
	if err != nil {
		log.Fatalf("bigquery.ConnectBQ: %v", err)
	}

	q := client.Query(query)
	it, err := q.Read(ctx)
	if err != nil {
		log.Fatalf("bigquery.Read: %v", err)
	}
	return it, err
}

// func main() {
// 	//variavel
// 	projectID := os.Getenv("GCP_STORAGE_PROJECT_ID")

// 	_query := "SELECT full_name, age FROM projectBQID.projectDSID.pessoa"

// 	ctx, client, err := ConnectBQ(projectID)
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	it, err := QueryBQ(ctx, client, _query)
// 	for {
// 		var values []bigquery.Value
// 		err := it.Next(&values)
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("bigquery.iterator: %v", err)
// 		}
// 		fmt.Println(values)
// 	}
// }
