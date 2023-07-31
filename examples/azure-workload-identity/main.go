package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func main() {
	// Using single Managed Identity
	//
	// DefaultAzureCredential will use the environment variables injected by the Azure Workload Identity
	// mutating webhook to authenticate with Azure Resource.
	//
	// Defaults to the value of the environment variable AZURE_CLIENT_ID.
	/*
		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
		      log.Fatal(err)
		}
	*/

	// Using multiple managed identities
	// injected managed idenitity client id by environment variable

	wiClientID := os.Getenv("WI_CLIENT_ID")
	subID := os.Getenv("SUB_ID")

	cred, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{ClientID: wiClientID})
	if err != nil {
		log.Fatal(err)
	}

	// Make sure your managed identity has premission to access your subscription
	client, err := armresources.NewResourceGroupsClient(subID, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	resourceGroups, err := listResourceGroup(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, resource := range resourceGroups {
		log.Printf("Resource Group Name: %s", *resource.Name)
	}
}

func printToken(cred *azidentity.DefaultAzureCredential) {
	token, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{"https://graph.microsoft.com/.default"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(token.Token)
}

func listResourceGroup(resourceGroupClient *armresources.ResourceGroupsClient, ctx context.Context) ([]*armresources.ResourceGroup, error) {

	resultPager := resourceGroupClient.NewListPager(nil)

	resourceGroups := make([]*armresources.ResourceGroup, 0)
	for resultPager.More() {
		pageResp, err := resultPager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		resourceGroups = append(resourceGroups, pageResp.ResourceGroupListResult.Value...)
	}
	return resourceGroups, nil
}
