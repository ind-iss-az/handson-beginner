package main

import (
	"os"
	"log"
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
)

type ParameterStore struct{
	param Parameter
}

type Parameter struct{
	db_host		string
	db_user		string
	db_password	string
}

func newParameterStore() (*ParameterStore, error) {
	ps := new(ParameterStore)
	err := ps.LoadParameter()
	if err != nil {
        return nil, err
    }
	return ps, nil
}

func (ps *ParameterStore) LoadParameter() (error) {
	var err error
	tag := os.Getenv("TAG")
	keyvaultName := "handson-" + tag + "-kv"
	log.Println(tag)
	ps.param.db_host, err = ps.Get(tag + "-dbhost", keyvaultName)
	if err != nil {
        return err
    }
	ps.param.db_user, err = ps.Get(tag + "-dbuser", keyvaultName)
	if err != nil {
        return err
    }
	ps.param.db_password, err = ps.Get(tag + "-dbpassword", keyvaultName)
	if err != nil {
        return err
	}
	return nil
}

func (ps *ParameterStore) Get(name string, keyvaultName string) (string, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Printf("unable to get your authorizer object: %v", err)
		return "", err
	}
	keyClient := keyvault.New()
	keyClient.Authorizer = authorizer
	param, err := keyClient.GetSecret(context.Background(), fmt.Sprintf("https://%s.vault.azure.net", keyvaultName), name, "")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return *param.Value, nil
}