package sign

import "crypto/ed25519"

// PdfSigner to hold the data ed25519 private key, public key and the pdf file
type PdfSigner struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
	PdfFile    []byte
	Signature  []byte
}
