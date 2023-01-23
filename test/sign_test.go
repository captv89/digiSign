package test

import (
	"crypto/ed25519"
	"digiSign/sign"
	"os"
	"reflect"
	"testing"
)

// Test the PdfSigning
func TestSignPdf(t *testing.T) {
	pdfSigner := &sign.PdfSigner{}
	pdfSigner.GenerateKeys()
	pdfSigner.ReadPdfFile("test.pdf")

	// Signature creation test
	t.Run("Signature creation test", func(t *testing.T) {
		pdfSigner.SignPdfFile()
		if len(pdfSigner.Signature) == 0 {
			t.Errorf("Signature creation failed")
		}
	})

	// Saving the signed pdf file test
	t.Run("Saving the signed pdf file test", func(t *testing.T) {
		pdfSigner.SaveSignedPdfFile("test-signed.pdf")
		if _, err := os.Stat("test-signed.pdf"); os.IsNotExist(err) {
			t.Errorf("Saving the signed pdf file failed")
		}
	})

	// Signature verification test
	t.Run("Signature verification test", func(t *testing.T) {
		pdfSigner.ReadPdfFile("test-signed.pdf")
		check := pdfSigner.VerifyPdfFile()
		//t.Log("Check result:", check)
		if !check {
			t.Errorf("Signature verification failed")
		}
	})
}

// Test the keys generation
func TestGenerateKeys(t *testing.T) {
	pdfSigner := &sign.PdfSigner{}

	t.Run("Keys generation test", func(t *testing.T) {
		pdfSigner.GenerateKeys()
		if reflect.TypeOf(pdfSigner.PrivateKey) != reflect.TypeOf(ed25519.PrivateKey{}) {
			t.Error("Expected private key of type ed25519.PrivateKey")
		}

		if reflect.TypeOf(pdfSigner.PublicKey) != reflect.TypeOf(ed25519.PublicKey{}) {
			t.Error("Expected public key of type ed25519.PublicKey")
		}
	})

	//	Test the keys saving to file
	t.Run("Keys saving to file test", func(t *testing.T) {
		err := pdfSigner.SaveKeysToFile("private.key", "public.key")
		if err != nil {
			t.Error("Keys saving to file failed")
		}
	})

	//	Test the private keys reading from file
	t.Run("Keys reading from file test", func(t *testing.T) {
		err := pdfSigner.ReadPrivateKey("private.key")
		if err != nil {
			t.Error("Private Key reading from file failed")
		}
	})

	//	Test the public keys reading from file
	t.Run("Keys reading from file test", func(t *testing.T) {
		err := pdfSigner.ReadPublicKey("public.key")
		if err != nil {
			t.Error("Public Key reading from file failed")
		}
	})
}
