package main

import (
	"context"
	"fmt"

	"github.com/melonrush13/keyvault_helper/config"
	"github.com/melonrush13/keyvault_helper/keyvault"
)

func main() {
	fmt.Println("In main")
	config.LoadSettings()
	ctx := context.Background()

	//func CreateKey(ctx context.Context, vaultName, keyName string) (key keyvault.KeyBundle, err error) {

	_, err := keyvault.CreateKey(ctx, "melVault", "meldev1316")
	if err != nil {
		fmt.Println("Error is ", err)
	}
}
