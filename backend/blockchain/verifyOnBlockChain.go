package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// VerifyCertificateHash checks if the hash matches what's on-chain
func VerifyCertificateHash(certID string, hash [32]byte) (bool, error) {
	onChainHash, err := contract.GetCertificateHash(&bind.CallOpts{Context: context.Background()}, certID)
	if err != nil {
		return false, err
	}
	return onChainHash == hash, nil
}
