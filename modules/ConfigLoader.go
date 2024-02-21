package modules

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/joho/godotenv"
)

func LoadConfig() map[string]string {
	ctx := context.Background()
	mapConfig := make(map[string]string)

	errEnv := godotenv.Load("./../../.env")
	// errEnv := godotenv.Load("/canopy/app/.env")
	if errEnv != nil {
		panic(errEnv)
	}

	cl := initVaultClient()

	err := cl.SetToken(os.Getenv("VAULT_TOKEN"))
	if err != nil {
		panic(err)
	}
	
	resp, err := cl.Read(ctx, "v1/topsecret/data/canopyScrapper")
	if err != nil {
		panic(err)
	}

	data, ok := resp.Data["data"].(map[string]interface{})
	if !ok {
		panic("not map interface")
	}

	for key, value := range data {
        strKey := fmt.Sprintf("%v", key)
        strValue := fmt.Sprintf("%v", value)

        mapConfig[strKey] = strValue
    }

	return mapConfig
}

func initVaultClient() *vault.Client {
	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://195.85.19.218:8200"),
		vault.WithRequestTimeout(10*time.Second),
	)

	if err != nil {
		panic(err)
	}
	
	return client
}
