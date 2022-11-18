package main

import "sample/security"

func main() {
	caPrivateKeyFilename := "out/private.pem"
	caPublicKeyFilename := "out/public.pem"
	caCertFilename := "out/cacert.pem"
	security.GenerateRSAKeyPair(caPrivateKeyFilename, caPublicKeyFilename)
	security.GenerateCertificate(caPrivateKeyFilename, true, caCertFilename)

	// create certificate for server
	serverPrivateKeyFilename := "out/server_private_key.pem"
	serverPublicKeyFilename := "out/server_public_key.pem"
	security.GenerateRSAKeyPair(serverPrivateKeyFilename, serverPublicKeyFilename)

	// create CSR for server
	serverCsrFilename := "out/server_csr.pem"
	security.CreateCertificateSigningRequest(serverPrivateKeyFilename, serverCsrFilename)
	// Sign CSR with with CA
	signedServerCertificateFilename := "out/signedServerCert.pem"
	security.SignCertificateWithCA(caCertFilename, caPrivateKeyFilename, serverCsrFilename, signedServerCertificateFilename)

	// create private key for certificate signing request (CSR)
	clientPrivateKeyFilename := "out/client_private_key.pem"
	clientPublicKeyFilename := "out/client_public_key.pem"
	security.GenerateRSAKeyPair(clientPrivateKeyFilename, clientPublicKeyFilename)

	// create CSR for client
	clientCsrFilename := "out/client_csr.pem"
	security.CreateCertificateSigningRequest(clientPrivateKeyFilename, clientCsrFilename)
	// Sign CSR with with CA
	signedClientCertificateFilename := "out/signedClientCert.pem"
	security.SignCertificateWithCA(caCertFilename, caPrivateKeyFilename, clientCsrFilename, signedClientCertificateFilename)

	host := "localhost:9034"
	security.EchoTLS_server_demo(signedServerCertificateFilename, serverPrivateKeyFilename, signedClientCertificateFilename, caCertFilename, host)
}
