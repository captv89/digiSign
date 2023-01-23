package sign

import (
	"crypto/ed25519"
	"encoding/pem"
	"errors"
	"os"
)

// GenerateKeys to generate the ed25519 private and public keys
func (p *PdfSigner) GenerateKeys() {
	var err error
	p.PublicKey, p.PrivateKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
}

// SaveKeysToFile to save the ed25519 private and public keys to file
func (p *PdfSigner) SaveKeysToFile(privateKeyFile, publicKeyFile string) error {
	// Encode the public key to PEM format
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: p.PublicKey})
	// Write the public key to a file
	err := os.WriteFile(publicKeyFile, publicKeyPEM, 0644)
	if err != nil {
		return err
	}

	// Encode the private key to PEM format
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: p.PrivateKey})
	// Write the private key to a file
	err = os.WriteFile(privateKeyFile, privateKeyPEM, 0600)
	if err != nil {
		return err
	}
	return nil
}

// ReadPrivateKey to read the ed25519 private key from file
func (p *PdfSigner) ReadPrivateKey(filePath string) error {
	// Read the private key file
	privateKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	p.PrivateKey, err = p.DecodePEMKeys(privateKeyPEM, "PRIVATE KEY")
	return nil
}

// ReadPublicKey to read the ed25519 public key from file
func (p *PdfSigner) ReadPublicKey(filePath string) error {
	// Read the public key file
	publicKeyPEM, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	p.PublicKey, err = p.DecodePEMKeys(publicKeyPEM, "PUBLIC KEY")
	if err != nil {
		return err
	}
	return nil
}

// DecodePEMKeys to decode the PEM encoded private and public keys
func (p *PdfSigner) DecodePEMKeys(aKeyPEM []byte, keyType string) ([]byte, error) {
	// Decode the PEM encoded public key
	block, _ := pem.Decode(aKeyPEM)
	if block == nil {
		return nil, errors.New("failed to decode PEM encoded public key")
	}
	if block.Type != keyType {
		return nil, errors.New("invalid PEM block type")
	} else {
		return block.Bytes, nil
	}
}