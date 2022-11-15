package mypackage

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
)

func EchoTLS_client_demo(messageToSend, clientCertFile, clientPrivateKeyFilename, serverCertFilename, caCertFilename, host string) {
	// clientCert, err := LoadCertificateFromPemFile(clientCertFile)

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

	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)
	certPool.AddCert(serverCert)

	config := &tls.Config{
		// // Required to accept self-signed certs
		InsecureSkipVerify: true,

		// Provide your client certificate if necessary
		// This client certificate will be presented to the other side
		// to be verified
		Certificates: []tls.Certificate{
			clientCert,
			// {
			// 	Certificate: [][]byte{
			// 		clientCert.Raw,
			// 		caCert.Raw,
			// 	},
			// },
		},

		// ServerName is used to verify the hostname (unless you are
		// skipping verification)
		// It is also included in the handshake in case the server uses
		// virtual hosts Can also just be an IP address
		// instead of a hostname.
		ServerName: "localhost",

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
	fmt.Println("Message received:", buffer)
}
