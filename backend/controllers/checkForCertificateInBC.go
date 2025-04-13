package controllers

import (
	"log"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/blockchain"
)

func checkForCertificateInBC(certificateHash string, certificateHashBytes32 [32]byte) bool {

	checkForCertificateInChain, errInBC := blockchain.VerifyCertificateHash(certificateHash, certificateHashBytes32)

	if errInBC != nil {
		log.Printf("🔴 Blockchain lookup failed: %v", errInBC)
		// json.NewEncoder(w).Encode(map[string]bool{"valid": false})
		return false
	}

	if checkForCertificateInChain {
		// log.Println(" Certificate already exists in Blockchain chain")
		return true
	}
	return false
}
