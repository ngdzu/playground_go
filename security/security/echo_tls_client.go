package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
)

func EchoTLS_client_skip_server_verification_demo(messageToSend,
	clientCertFile,
	clientPrivateKeyFilename,
	serverCertFilename,
	caCertFilename,
	host string) {

	// client certificate to present to the server. The server
	// will verify this certificate
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientPrivateKeyFilename)
	if err != nil {
		log.Fatal("Fail to load certificate from file:", clientCertFile)
	}

	// caCert, err := LoadCertificateFromPemFile(caCertFilename)
	// if err != nil {
	// 	log.Fatal("Fail to load CA certificate", err)
	// }

	// serverCert, err := LoadCertificateFromPemFile(serverCertFilename)
	// if err != nil {
	// 	log.Fatal("Fail to load CA certificate", err)
	// }

	// Both CA and server certificates are needeed
	// certPool := x509.NewCertPool()
	// certPool.AddCert(caCert)
	// certPool.AddCert(serverCert)

	config := &tls.Config{
		// Required to accept self-signed certs
		InsecureSkipVerify: true,

		// Provide your client certificate if necessary
		// This client certificate will be presented to the other side
		// to be verified
		Certificates: []tls.Certificate{
			clientCert,
		},

		// ServerName is used to verify the hostname (unless you are
		// skipping verification)
		// It is also included in the handshake in case the server uses
		// virtual hosts Can also just be an IP address
		// instead of a hostname.
		ServerName: "localhost",

		// NOTE if InsecureSkipVerify is true, we don't need to provide this
		// since we won't verify the server
		// RootCAs that you are willing to accept
		// If RootCAs is nil, the host's default root CAs are used
		// RootCAs: certPool,
	}

	connection, err := tls.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("Fail to dial to %s. %v", host, err)
	}
	defer connection.Close()

	// Write data to socket
	byteWritten, err := connection.Write([]byte(messageToSend))
	if err != nil {
		log.Fatal("Fail to write to socket", err)
	}
	fmt.Printf("Wrote %d bytes to socket", byteWritten)

	// Read data back from the socket
	buffer := make([]byte, 100)
	bytesRead, err := connection.Read(buffer)
	if err != nil {
		log.Fatal("Fail to read from socket", err)
	}

	fmt.Println("bytes received:", bytesRead)
	fmt.Println("Message received:", string(buffer))
}

func EchoTLS_client_with_server_verification_demo(messageToSend,
	clientCertFile,
	clientPrivateKeyFilename,
	serverCertFilename,
	caCertFilename,
	host string) {

	// client certificate to present to the server. The server
	// will verify this certificate
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientPrivateKeyFilename)
	if err != nil {
		log.Fatal("Fail to load certificate from file:", clientCertFile)
	}

	caCert, err := LoadCertificateFromPemFile(caCertFilename)
	if err != nil {
		log.Fatal("Fail to load CA certificate", err)
	}

	serverCert, err := LoadCertificateFromPemFile(serverCertFilename)
	if err != nil {
		log.Fatal("Fail to load CA certificate", err)
	}

	// Both CA and server certificates are needeed
	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)
	certPool.AddCert(serverCert)

	config := &tls.Config{
		InsecureSkipVerify: false, // will verify server

		// Provide your client certificate if necessary
		// This client certificate will be presented to the other side
		// to be verified
		Certificates: []tls.Certificate{
			clientCert,
		},

		// ServerName is used to verify the hostname (unless you are
		// skipping verification)
		// It is also included in the handshake in case the server uses
		// virtual hosts Can also just be an IP address
		// instead of a hostname.
		ServerName: "localhost",

		// NOTE if InsecureSkipVerify is true, we don't need to provide this
		// since we won't verify the server
		// RootCAs that you are willing to accept
		// If RootCAs is nil, the host's default root CAs are used
		RootCAs: certPool,
	}

	connection, err := tls.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("Fail to dial to %s. %v", host, err)
	}
	defer connection.Close()

	// Write data to socket
	byteWritten, err := connection.Write([]byte(messageToSend))
	if err != nil {
		log.Fatal("Fail to write to socket", err)
	}
	fmt.Printf("Wrote %d bytes to socket", byteWritten)

	// Read data back from the socket
	buffer := make([]byte, 100)
	bytesRead, err := connection.Read(buffer)
	if err != nil {
		log.Fatal("Fail to read from socket", err)
	}

	fmt.Println("bytes received:", bytesRead)
	fmt.Println("Message received:", string(buffer))
}
