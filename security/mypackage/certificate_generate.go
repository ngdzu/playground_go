package mypackage

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

func setupCertificateTemplate(isCA bool) x509.Certificate {
	notBefore := time.Now()
	notAfter := notBefore.Add(time.Hour * 24 * 365) // plus one year

	// generate random serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	randomNumber, err := rand.Int(rand.Reader, serialNumberLimit)

	if err != nil {
		log.Fatal("error generating randome serial number", err)
	}
	nameInfo := pkix.Name{
		Organization:       []string{"My Organization"},
		CommonName:         "localhost",
		OrganizationalUnit: []string{"My org unit"},
		Country:            []string{"US"},
		Province:           []string{"MA"},
		Locality:           []string{"Boston"},
	}

	// create certificate template
	certTemplate := x509.Certificate{
		SerialNumber: randomNumber,
		Subject:      nameInfo,
		NotBefore:    notBefore,
		NotAfter:     notAfter,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

		// For ExtKeyUsage, default to any, but can specify to use
		// only as server or client authentication, code signing, etc
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	// to create CA,
	if isCA {
		certTemplate.IsCA = true
		certTemplate.KeyUsage = certTemplate.KeyUsage | x509.KeyUsageCertSign
	}

	// Add any IP addresses and hostnames covered by this cert
	// This example only covers localhost
	certTemplate.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
	certTemplate.DNSNames = []string{"localhost", "localhost.local"}

	return certTemplate
}

func writeCertToPemFile(outputFilename string, derBytes []byte) {
	parentDir := filepath.Dir(outputFilename)
	if _, err := os.Lstat(parentDir); os.IsNotExist(err) {
		if err = os.MkdirAll(parentDir, os.ModePerm); err != nil {
			log.Fatal("can't create directory", parentDir)
		}
	}

	certOutFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer certOutFile.Close()

	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	err = pem.Encode(certOutFile, block)
	if err != nil {
		log.Fatal(err)
	}

}

// Create a self-sign certificate
func GenerateCertificate_demo(inputPrivateKeyPath string, isCA bool) {
	// Private key of signer - self signed means signer==signee
	privateKey := LoadPrivateKeyFromPemFile(inputPrivateKeyPath)

	// Public key of signee. Self signing means we are the signer and
	// the signee so we can just pull our public key from our private key
	publicKey := privateKey.PublicKey

	certTemplate := setupCertificateTemplate(isCA)

	// create and sign certificate with private key
	certPem, err := x509.CreateCertificate(
		rand.Reader,
		&certTemplate,
		&certTemplate,
		&publicKey, // MUST pass pointer
		privateKey, // privateKey is pointer
	)
	if err != nil {
		log.Fatal(err)
	}

	if isCA {
		writeCertToPemFile("out/cacert.pem", certPem)
	} else {
		writeCertToPemFile("out/cert.pem", certPem)
	}

	fmt.Print("Successfully create a certificate.")
}
