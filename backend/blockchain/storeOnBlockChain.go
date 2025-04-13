package blockchain

import (
	"errors"
	"log"
)

// StoreCertificateHash saves the hash to the blockchain
func StoreCertificateHash(certID string, hash [32]byte) error {
	tx, err := contract.StoreCertificateHash(auth, certID, hash)
	if err != nil {
		return errors.New("failed to store cert hash: " + err.Error())
	}
	log.Printf("Stored cert on-chain. Tx Hash: %s", tx.Hash().Hex())
	return nil
}
