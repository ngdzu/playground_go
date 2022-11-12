package main

import "sample/security/mypackage"

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
	// mypackage.GenRSA_demo()
	// mypackage.SignMessage_demo()
	// mypackage.GenerateCertificate_demo("testdata/private.pem", true)
	// mypackage.GenerateCertificate_demo("testdata/private.pem", false)
	mypackage.CertificateSigningRequest_demo()
}
