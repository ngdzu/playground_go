package security

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
)

// Create a TLS listener and echo back data received by clients.
func EchoTLS_server_demo(serverCertFilename, serverPrivateKeyFilename, clientCertFilename, caCertFilename, host string) {
	serverCertificate, err := tls.LoadX509KeyPair(serverCertFilename, serverPrivateKeyFilename)
	if err != nil {
		log.Fatal("Fail to load certificate and private key")
	}

	clientCert, err := LoadCertificateFromPemFile(clientCertFilename)
	if err != nil {
		log.Fatal("fail to load client certificate", err)
	}

	caCert, err := LoadCertificateFromPemFile(caCertFilename)
	if err != nil {
		log.Fatal("fail to load client certificate", err)
	}

	// Create and add client certificate to cert pool
	certPool := x509.NewCertPool()
	certPool.AddCert(clientCert)
	certPool.AddCert(caCert)

	// set up certificate
	config := &tls.Config{

		// Pass the certificate that will be verified by the
		// other side
		Certificates: []tls.Certificate{
			serverCertificate,
		},
		// By default no client certificate is required.
		// To require and validate client certificates, specify the
		// ClientAuthType to be one of:
		// NoClientCert, RequestClientCert, RequireAnyClientCert,
		// VerifyClientCertIfGiven, RequireAndVerifyClientCert)
		ClientAuth: tls.RequireAndVerifyClientCert,
		// ClientAuth: tls.RequireAndVerifyClientCert
		// Define the list of certificates you will accept as
		// trusted certificate authorities with ClientCAs.
		// ClientCAs: *x509.CertPool
		ClientCAs: certPool,
	}

	// create TLS socket client
	listener, err := tls.Listen("tcp", host, config)
	if err != nil {
		log.Fatal("Fail to create listener", err)
	}
	defer listener.Close()

	// Listen forever
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("fail to accept connection", err)
		}

		// handle in a separate thread
		go handleConnection(connection)
	}

}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	socketReader := bufio.NewReader(connection)
	for {
		message, err := socketReader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading socket connection", err)
			return
		}

		fmt.Println(message)

		// echo
		n, err := connection.Write([]byte(message))
		if err != nil {
			log.Println("error writing to client socket", err)
			return
		}
		fmt.Printf("Wrote %d bytes back to the client.\n", n)
	}
}
