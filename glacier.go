package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
	"io/ioutil"
	"log"
)

func main() {

	//Get the names of all the directories in our current root
	// these will be the vault names
	vnames, valerr := ioutil.ReadDir("./")
	if valerr != nil {
		log.Fatal(valerr)
	}
	//vaultName := "gotest1"

	//Setup the policy for data retrieval - HT Andrew
	params := &glacier.SetDataRetrievalPolicyInput{
		Policy: &glacier.DataRetrievalPolicy{
			Rules: []*glacier.DataRetrievalRule{
				&glacier.DataRetrievalRule{
					Strategy: aws.String("FreeTier"),
				},
			},
		},
	}

	svc := glacier.New(session.New(&aws.Config{Region: aws.String("eu-west-1")}))
	_, polerr := svc.SetDataRetrievalPolicy(params)
	if polerr != nil {
		fmt.Println(polerr)
		return
	}
	//And now setup the vaults
	for _, vaultName := range vnames {
		_, err := svc.CreateVault(&glacier.CreateVaultInput{
			VaultName: aws.String(vaultName.Name()),
		})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Created vault", vaultName.Name())
	}
}
