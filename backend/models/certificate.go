package models

import "time"

type Certificate struct {
	ID           string    `firestore:"id"`           // Will be Generated
	StudentName  string    `firestore:"studentName"`  // Name of the student
	StudentEmail string    `firestore:"studentEmail"` // Student email
	Course       string    `firestore:"course"`       // Course
	DateIssued   time.Time `firestore:"dateIssued"`   // Date of issue
	IssuerName   string    `firestore:"issuerName"`   // Organization
	Revoked      string    `firestore:"revoked"`      // Revoke status
	NoOfIssue    string    `firestore:"noOfIssue"`    // Issued certificate number
}
