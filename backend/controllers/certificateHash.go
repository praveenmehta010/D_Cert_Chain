package controllers

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/models"
)

func GenerateCertificateHash(cert *models.Certificate) (issueCertificateHash string) {

	// Certificate Hash -> For the chain
	rawData := cert.StudentName + cert.StudentEmail + cert.Revoked + cert.NoOfIssue + "Magical Salt :- dcomep%%*%59-e48jcn cjd cidge98754*YjbVV**$^gbx##"
	hash := sha256.Sum256([]byte(rawData))
	issueCertificateHash = hex.EncodeToString(hash[:])

	return issueCertificateHash
}

func GenerateCertificateID(cert *models.Certificate) (certID string) {
	// Certificate ID -> For the DB
	certIDData := cert.StudentEmail + "Magical Salt :- jnj864hyf%$7VthI" + cert.Course + cert.NoOfIssue + cert.StudentName
	certIDHash := sha256.Sum256([]byte(certIDData))
	certID = hex.EncodeToString(certIDHash[:])
	return certID
}
