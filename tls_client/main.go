package main

import "sample/security/mypackage"

// NOTE run the server at tls_server/main.go first to generate all
// the certificate files. Come back here and run the client
func main() {
	messageToSend := "Hello\n"
	host := "localhost:9034"
	clientCertFilename := "../tls_server/out/signedClientCert.pem"
	caCertFilename := "../tls_server/out/cacert.pem"
	clientPrivateKeyFilename := "../tls_server/out/client_private_key.pem"
	serverCertificateFilename := "../tls_server/out/signedServerCert.pem"
	mypackage.EchoTLS_client_demo(messageToSend,
		clientCertFilename,
		clientPrivateKeyFilename,
		serverCertificateFilename,
		caCertFilename,
		host)
}
