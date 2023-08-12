package service

import (
	"backend/contracts"
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func DeployCourseCall(sender string, priceInETH float64, courseId uint) (string, error) {

// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/e62196cd12084efc85fca2472c42b04f") // Replace with your Ethereum node URL
// 	if err != nil {
// 		fmt.Println("Failed to connect to the Ethereum client:", err)
// 		return "", err
// 	}
// 	priceInWei, err := util.ConvertEthToWei(priceInETH)
// 	if err != nil {
// 		fmt.Println("Failed to convert to wei:", err)
// 		return "", err
// 	}

// 	contractInstance, err := BindCourseManagerContract(client)
// 	if err != nil {
// 		fmt.Println("Failed to find instance:", err)
// 		return "", err
// 	}
// 	senderAddress := common.HexToAddress(sender)
// 	auth := createTransactOpts(senderAddress, client, priceInWei)

// 	tx, err := contractInstance.DeployCourse(auth, big.NewInt(int64(courseId)), priceInWei)
// 	if err != nil {
// 		fmt.Println("Failed to deploy course:", err)
// 		return "", err
// 	}
// 	txHash := tx.Hash().Hex()

// 	return txHash, nil
// }

func BindCourseManagerContract(client *ethclient.Client) (*contracts.Contracts, error) {

	course_manager_address := os.Getenv("COURSE_MANAGER_ADDRESS")
	contractAddress := common.HexToAddress(course_manager_address)
	contractInstance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return contractInstance, nil
}

func createTransactOpts(senderAddress common.Address, client *ethclient.Client, valueInWei *big.Int) *bind.TransactOpts {

	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil
	}

	return &bind.TransactOpts{
		From:     senderAddress,
		Nonce:    big.NewInt(int64(nonce)),
		Value:    valueInWei,
		GasLimit: uint64(600000),
		GasPrice: gasPrice,
	}
}
