package mypackage

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"os"
)

// # Create CSR
// openssl req -new -key priv.pem -out csr.pem
// # View details to verify request was created properly
// openssl req -verify -in csr.pem -text -noout

func WritePemBlockToFile(block *pem.Block, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	return pem.Encode(file, block)
}

func WriteCSRToPemFile(csr []byte, filePath string) error {
	block := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr,
	}

	return WritePemBlockToFile(block, filePath)
}

func CreateCertificateSigningRequest(inputPrivateKeyPath, outputCsrFilename string) {

	privateKey := LoadPrivateKeyFromPemFile(inputPrivateKeyPath)
	nameInfo := pkix.Name{
		Organization:       []string{"My Organization"},
		CommonName:         "localhost",
		OrganizationalUnit: []string{"Business Unit Name"},
		Country:            []string{"US"},
		Province:           []string{"MA"},
		Locality:           []string{"Boston"},
	}

	csrTemplate := &x509.CertificateRequest{
		Version:            2,
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          privateKey.PublicKey,
		Subject:            nameInfo,

		// subject alternate name
		// Subject Alternate Name values.
		DNSNames:       []string{"Business Unit Name"},
		EmailAddresses: []string{"test@localhost"},
		IPAddresses:    []net.IP{},
	}

	// create CSR based on the template
	csr, err := x509.CreateCertificateRequest(rand.Reader, csrTemplate, privateKey)
	if err != nil {
		log.Fatal("Error creating certificate signing request. ", err)
	}

	if WriteCSRToPemFile(csr, outputCsrFilename) != nil {
		fmt.Println("Fail to create CSR")
	} else {
		fmt.Println("Successfully create CSR.")
	}
}
