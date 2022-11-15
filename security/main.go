package main

import "sample/security/mypackage"

func runEchoTLSserver_demo() {
	certFilename := "testdata/cert.pem" // self-signed
	privFilename := "testdata/private.pem"
	host := "localhost:9034"
	clientCertFilename := "testdata/signed_cert.pem"

	mypackage.EchoTLS_server_demo(certFilename, privFilename, clientCertFilename, host)
}

func runGenerateCACertificate_demo() {
	outputCertificateFilename := "testdata/cacert.pem"
	isCA := true
	mypackage.GenerateCertificate("testdata/private.pem", isCA, outputCertificateFilename)
}

func main() {
	// mypackage.Readfull()
	// readatleast()
	// bufferedreader()
	// scanner()
	// zipcreate()
	// mypackage.Zipread()
	// gzipcompress()
	// mypackage.Gzipuncompress()
	// usingflag()
	// mypackage.Which()
	// mypackage.SymbolicLink()
	// mypackage.HashFile()
	// mypackage.HashLargeFile()
	// mypackage.HashingPassword()
	// mypackage.CryptoRandom()
	// mypackage.SignMessage_demo()
	// mypackage.GenerateCertificate_demo("testdata/private.pem", true)
	// mypackage.GenerateCertificate_demo("testdata/private.pem", false)
	// mypackage.CertificateSigningRequest_demo()
	// mypackage.CreateCertificateWithCA_demo()

	caPrivateKeyFilename := "out/private.pem"
	caPublicKeyFilename := "out/public.pem"
	caCertCertificateFilename := "out/cacert.pem"
	mypackage.GenerateRSAKeyPair(caPrivateKeyFilename, caPublicKeyFilename)
	mypackage.GenerateCertificate(caPrivateKeyFilename, true, caCertCertificateFilename)

	// create private key for certificate signing request (CSR)
	privateKeyPathForCSR := "out/private_key_for_csr.pem"
	publicKeyPathForCSR := "out/public_key_for_csr.pem"
	mypackage.GenerateRSAKeyPair(privateKeyPathForCSR, publicKeyPathForCSR)

	// create CSR
	csrFilename := "out/csr.pem"
	mypackage.CreateCertificateSigningRequest(privateKeyPathForCSR, csrFilename)
	// Sign CSR with with CA
	signedCertificateFilename := "out/signedCert.pem"
	mypackage.SignCertificateWithCA(caCertCertificateFilename, caPrivateKeyFilename, csrFilename, signedCertificateFilename)

	// runEchoTLSserver()
}
