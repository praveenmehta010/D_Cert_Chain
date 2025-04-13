package blockchain

import (
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/blockchain/certificate" // path to abigen package
)

var (
	client   *ethclient.Client
	contract *certificate.Certificate
	auth     *bind.TransactOpts

	contractAddr common.Address
)

func Init() {
	var err error

	// Load .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	privateKeyHex := os.Getenv("BLOCKCHAIN_PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("PRIVATE_KEY not found in environment")
	}

	contractAddressStr := os.Getenv("BlockChainAccountAddress")
	if contractAddressStr == "" {
		log.Fatal("BlockChainAccountAddress not found in environment")
	}
	contractAddr = common.HexToAddress(contractAddressStr)

	client, err = ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337)) // Hardhat default chain ID
	if err != nil {
		log.Fatalf("Failed to create auth: %v", err)
	}

	contract, err = certificate.NewCertificate(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}
}
