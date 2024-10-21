package operator

// // OUTDATED
// // This file contains cli functions for registering an operator with the AVS and printing status
// // However, all of this functionality has been moved to the plugin/ package
// // we are just waiting for eigenlayer-cli to be open sourced so we can completely get rid of this registration functionality in the operator

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/spf13/viper"

// 	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
// 	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
// 	"github.com/ethereum/go-ethereum/ethclient"

// 	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
// )

// func registerOperatorOnStartup(
// 	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
// 	mockTokenStrategyAddr common.Address,
// ) {
// 	err := RegisterOperatorWithAvs(operatorEcdsaPrivateKey)
// 	if err != nil {
// 		fmt.Println("Error registering operator with avs", "err", err)
// 		panic(err)
// 	}
// 	fmt.Println("Registered operator with avs")
// }

// // Registration specific functions
// func RegisterOperatorWithAvs(
// 	operatorEcdsaKeyPair *ecdsa.PrivateKey,
// ) error {
// 	// hardcode these things for now
// 	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
// 	socket := "Not Needed"
// 	operatorToAvsRegistrationSigSalt := [32]byte{123}

// 	ethRpcUrl := viper.GetString("eth_rpc_url")
// 	ethClient, err := ethclient.Dial(ethRpcUrl)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
// 	}

// 	curBlockNum, err := ethClient.BlockNumber(context.Background())
// 	if err != nil {
// 		fmt.Println("Unable to get current block number")
// 		return err
// 	}
// 	curBlock, err := ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
// 	if err != nil {
// 		fmt.Println("Unable to get current block")
// 		return err
// 	}
// 	sigValidForSeconds := int64(1_000_000)
// 	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time) + sigValidForSeconds)
// 	_, err = o.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
// 		context.Background(),
// 		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
// 		o.blsKeypair, quorumNumbers, socket,
// 	)
// 	if err != nil {
// 		fmt.Println("Unable to register operator with avs registry coordinator")
// 		return err
// 	}
// 	o.logger.Infof("Registered operator with avs registry coordinator.")

// 	return nil
// }

// func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
// 	return regcoord.BN254G1Point{
// 		X: p.X.BigInt(new(big.Int)),
// 		Y: p.Y.BigInt(new(big.Int)),
// 	}
// }
