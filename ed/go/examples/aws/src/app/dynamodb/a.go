package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

func Run(cfg *aws.Config) {
	svc := dynamodb.New(session.New(), cfg)

	{
		tablesList, err := svc.ListTables(&dynamodb.ListTablesInput{})
		if err != nil {
			fmt.Println("Got error %w:", err.Error())
		}
		if tablesList == nil || tablesList.TableNames == nil {
			fmt.Println("Got nil output")
			return
		}

		moviesExsists := false
		for _, tableName := range tablesList.TableNames {
			if *tableName == "Movies" {
				moviesExsists = true
				break
			}
		}

		if !moviesExsists {
			f0(svc)
		}
	}

	{
		info, err := svc.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String("Movies")})
		if err != nil {
			fmt.Println("Got error %w:", err.Error())
		}
		if info != nil && info.Table != nil {
			v := info.Table.ItemCount
			fmt.Printf("Found count in table: %v \n", *v)
		}
	}

	f1(svc)
	f2(svc)
}

func f0(svc *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		// see: https://docs.aws.amazon.com/sdkforruby/api/Aws/DynamoDB/Types/AttributeValue.html
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("Year"), KeyType: aws.String("HASH")},   // partition key
			{AttributeName: aws.String("Title"), KeyType: aws.String("RANGE")}, // sort key
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Movies"),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:", err.Error())
		os.Exit(1)
	}
}

func f1(svc *dynamodb.DynamoDB) {
	item := Item{
		Year:   2015,
		Title:  "The Big New Movie",
		Plot:   "Nothing happens at all.",
		Rating: 0.0,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:", err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Movies"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:", err.Error())
		os.Exit(1)
	}
}

func f2(svc *dynamodb.DynamoDB) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String("2015"),
			},
			"Title": {
				S: aws.String("The Big New Movie"),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Found movie: ", result)
}
