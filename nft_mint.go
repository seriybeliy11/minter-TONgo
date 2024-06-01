package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/nft"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func main() {
	client := liteclient.NewConnectionPool()

	err := client.AddConnection(context.Background(), "5.9.10.15:48014", "3XO67K/qi+gu3T9v8G2hx1yNmWZhccL3O7SoosFo8G0=")
	if err != nil {
		panic(err)
	}

	api := ton.NewAPIClient(client)
	w := getWallet(api)

	log.Println("Deploy wallet:", w.WalletAddress().String())

	collectionAddr := address.MustParseAddr("адрес_деплоя_коллекции")
	collection := nft.NewCollectionClient(api, collectionAddr)

	collectionData, err := collection.GetCollectionData(context.Background())
	if err != nil {
		panic(err)
	}

	nftAddr, err := collection.GetNFTAddressByIndex(context.Background(), collectionData.NextItemIndex)
	if err != nil {
		panic(err)
	}

	mintData, err := collection.BuildMintPayload(collectionData.NextItemIndex, w.WalletAddress(), tlb.MustFromTON("0.01"), &nft.ContentOffchain{
		URI: fmt.Sprint(collectionData.NextItemIndex) + ".json",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Minting NFT...")
	mint := wallet.SimpleMessage(collectionAddr, tlb.MustFromTON("0.025"), mintData)

	err = w.Send(context.Background(), mint, true)
	if err != nil {
		panic(err)
	}

	fmt.Println("Minted NFT:", nftAddr.String())
}

func getWallet(api *ton.APIClient) *wallet.Wallet {
	words := strings.Split("сиды", " ")
	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		panic(err)
	}
	return w
}