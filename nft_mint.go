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
    "math/big"
)

func main() {
	client := liteclient.NewConnectionPool()

	// connect to mainnet lite server
	err := client.AddConnection(context.Background(), "135.181.140.212:13206", "K0t3+IWLOXHYMvMcrGZDPs+pn58a17LFbnXoQkKc2xw=")
	if err != nil {
		panic(err)
	}

	// initialize ton api lite connection wrapper
	api := ton.NewAPIClient(client)
	w := getWallet(api)

	log.Println("Deploy wallet:", w.WalletAddress().String())

	collectionAddr := address.MustParseAddr("")
	collection := nft.NewCollectionClient(api, collectionAddr)

	collectionData, err := collection.GetCollectionData(context.Background())
	if err != nil {
		panic(err)
	}

	nftAddr, err := collection.GetNFTAddressByIndex(context.Background(), collectionData.NextItemIndex)
	if err != nil {
		panic(err)
	}


    for i:=0; i < 500; i++ {
        nextIndex := new(big.Int).Add(collectionData.NextItemIndex, big.NewInt(int64((i+1))))
        mintData, err := collection.BuildMintPayload(collectionData.NextItemIndex, w.WalletAddress(), tlb.MustFromTON("0.01"), &nft.ContentOffchain{URI: fmt.Sprint(nextIndex) + ".json",})


        fmt.Println("Minting NFT...")
        mint := wallet.SimpleMessage(collectionAddr, tlb.MustFromTON("0.025"), mintData)
    
        err = w.Send(context.Background(), mint, true)
        if err != nil {
            panic(err)
        }
    
        fmt.Println("Minted NFT:", nftAddr.String())

        if err != nil {
            panic(err)
        }
    }

	newData, err := nft.NewItemClient(api, nftAddr).GetNFTData(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Minted NFT addr: ", nftAddr.String())
	fmt.Println("NFT Owner:", newData.OwnerAddress.String())
}

func getWallet(api *ton.APIClient) *wallet.Wallet {
	words := strings.Split("", " ")
	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		panic(err)
	}
	return w
}