package iam

import (
	"context"
	"fmt"
	"log"

	v1 "google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

func Run(projectID, SAFilePath string) {
	ctx := context.Background()
	s := getService(ctx, SAFilePath)

	//listSAKeys(s, "email")
	createSA(ctx, s, projectID, "golang-sa", "Golang SA")
	listSA(s, projectID)
}

func getService(ctx context.Context, SAFilePath string) *v1.Service {
	s, err := v1.NewService(ctx, option.WithCredentialsFile(SAFilePath))
	if err != nil {
		log.Fatalf("iam.NewService: %v", err)
	}

	return s
}

func createSA(ctx context.Context, s *v1.Service, projectID, name, displayName string) {
	request := &v1.CreateServiceAccountRequest{
		AccountId:      name,
		ServiceAccount: &v1.ServiceAccount{DisplayName: displayName},
	}
	account, err := s.Projects.ServiceAccounts.Create("projects/"+projectID, request).Do()
	if err != nil {
		log.Fatalf("Projects.ServiceAccounts.Create: %v", err)
	}

	fmt.Printf("Created Service Account: %+v", account)
}

func listSA(s *v1.Service, projectID string) {
	response, err := s.Projects.ServiceAccounts.List("projects/" + projectID).Do()
	if err != nil {
		log.Fatalf("Projects.ServiceAccounts.List: %v", err)
	}

	fmt.Printf("Service Accounts:\n")
	for _, account := range response.Accounts {
		fmt.Printf("%v\n", account.Name)
	}
}

func listSAKeys(s *v1.Service, serviceAccountEmail string) {
	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	response, err := s.Projects.ServiceAccounts.Keys.List(resource).Do()
	if err != nil {
		log.Fatalf("Projects.ServiceAccounts.Keys.List: %v", err)
	}

	for _, key := range response.Keys {
		fmt.Printf("Listing key: %v", key.Name)
	}
	fmt.Printf("Keys: %+v \n", response.Keys)
}
