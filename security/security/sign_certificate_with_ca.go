package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func LoadCertificateFromPemFile(certificateFilePath string) (*x509.Certificate, error) {
	certText, err := os.ReadFile(certificateFilePath)
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(certText)

	return x509.ParseCertificate(block.Bytes)
}

func LoadCertificateRequestFromPemFile(csrFilePath string) (*x509.CertificateRequest, error) {
	certText, err := os.ReadFile(csrFilePath)
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(certText)

	return x509.ParseCertificateRequest(block.Bytes)
}

func WriteSignedRequestToPemFile(csr []byte, filePath string) {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: csr,
	}

	WritePemBlockToFile(block, filePath)
}

func SignCertificateWithCA(caCertFilename, caPrivateKeyFilename, csrFilename, outputSingedCertificateFilename string) {
	caCert, err := LoadCertificateFromPemFile(caCertFilename)
	if err != nil {
		log.Fatal(err)
	}
	caPrivateKey := LoadPrivateKeyFromPemFile(caPrivateKeyFilename)

	csrCert, err := LoadCertificateRequestFromPemFile(csrFilename)
	if err != nil {
		log.Fatal("fail to load CSR", err)
	}
	csrPublicKey, ok := csrCert.PublicKey.(*rsa.PublicKey)
	if !ok {
		log.Fatal("can't extract public key from CSR certificate")
	}

	// create client certificate template
	certTemplate := &x509.Certificate{
		Signature:          csrCert.Signature,
		SignatureAlgorithm: csrCert.SignatureAlgorithm,

		PublicKeyAlgorithm: csrCert.PublicKeyAlgorithm,
		PublicKey:          csrPublicKey,

		SerialNumber: big.NewInt(2),
		Issuer:       caCert.Subject,
		Subject:      csrCert.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	signed, err := x509.CreateCertificate(
		rand.Reader,
		certTemplate,
		caCert,
		csrPublicKey,
		caPrivateKey,
	)
	if err != nil {
		log.Fatal("fail to sign certificate", err)
	}

	WriteCertToPemFile(outputSingedCertificateFilename, signed)
}
