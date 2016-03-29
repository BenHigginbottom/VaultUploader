package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/glacier"
)

func main() {

    //Get the names of all the directories in our current root
    // these will be the vault names
    vnames, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
    //vaultName := "gotest1"

    svc := glacier.New(session.New(&aws.Config{Region: aws.String("eu-west-1")}))
    for _, vaultName := range vnames {
        _, err := svc.CreateVault(&glacier.CreateVaultInput{
        VaultName: (vaultName.Name()),
        })
    if err != nil {
    log.Println(err)
    return
}
fmt.Println("Created vault", vaultName)
}
}