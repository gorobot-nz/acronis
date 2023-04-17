package main

import (
	"fmt"
	acronis "github.com/gorobot-nz/acronis/pkg/client"
	"github.com/gorobot-nz/acronis/pkg/client/apimodels"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}
}

func main() {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	datacenterUrl := os.Getenv("DATACENTER_URL")

	acronisClient, err := acronis.NewAcronisClient(clientId, clientSecret, datacenterUrl)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	/*client, err := acronisClient.GetClient()
	if err != nil {
		return
	}

	tenant, err := acronisClient.CreateTenant(&apimodels.Tenant{
		Name:     "Customer Inc",
		Kind:     apimodels.TenantCustomerKind,
		ParentId: client.TenantId,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := acronisClient.CreateUser(&apimodels.User{
		TenantId: tenant.Id,
		Login:    "go_robot",
		Contact: apimodels.UserContacts{
			Email:     "nikita.zn11102001@gmail.com",
			FirstName: "Nikita",
			LastName:  "Zhamoidzik",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/

	services, err := acronisClient.FetchServices()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var serviceId string

	for _, val := range services {
		if val.Type == apimodels.BackupType {
			serviceId = val.Id
			break
		}
	}

	items, err := acronisClient.FetchOfferingItemsForChild(true)
	if err != nil {
		return
	}

	var filtered []apimodels.OfferingItem

	for index := range items {
		if items[index].ApplicationId == serviceId && items[index].Edition == apimodels.ProtectionPerGigabyte {
			filtered = append(filtered, items[index])
		}
	}

	for _, item := range filtered {
		_ = acronisClient.EnableOfferingItem("5370059d-dee1-42d9-a509-a18843c83ee9", &item)
		if err != nil {
			return
		}
	}
}
