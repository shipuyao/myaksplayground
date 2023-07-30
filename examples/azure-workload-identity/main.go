package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func main() {
	// create a secret client with the default credential
	// DefaultAzureCredential will use the environment variables injected by the Azure Workload Identity
	// mutating webhook to authenticate with Azure Key Vault.

	//cred, err := azidentity.NewDefaultAzureCredential(nil)
	//if err != nil {
	//      log.Fatal(err)
	//}

	cred, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{ClientID: "d26641b9-074b-4e46-8c1f-cb3a513b2502"})

	token, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: []string{"https://graph.microsoft.com/.default"}})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(token.Token)

	client, err := armresources.NewResourceGroupsClient("412d1f37-bc7c-422c-bf7e-93099a2feab0", cred, nil)
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

	log.Print("done")
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
