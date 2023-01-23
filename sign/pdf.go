package sign

import (
	"crypto/ed25519"
	"os"
)

// ReadPdfFile to read the pdf file
func (p *PdfSigner) ReadPdfFile(filePath string) {
	var err error
	p.PdfFile, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	//fmt.Println("PDF File:", p.PdfFile)
}

// SignPdfFile to sign the pdf file
func (p *PdfSigner) SignPdfFile() {
	p.Signature = ed25519.Sign(p.PrivateKey, p.PdfFile)
	//fmt.Println("Length of signature", len(p.Signature))
}

// VerifyPdfFile to verify the pdf file
func (p *PdfSigner) VerifyPdfFile() bool {
	//fmt.Println("PDF Byte", p.PdfFile)
	// Get the original pdf file without the signature
	originalPdf := p.PdfFile[:len(p.PdfFile)-len(p.Signature)]
	//fmt.Printf("Original PDF: %x, Signature: %x, Signed PDF: %x", len(originalPdf), len(p.Signature), len(p.PdfFile))
	//fmt.Println("Original PDF Byte", originalPdf)
	return ed25519.Verify(p.PublicKey, originalPdf, p.Signature)
}

// SaveSignedPdfFile to save the signed pdf file
func (p *PdfSigner) SaveSignedPdfFile(filePath string) {
	var err error
	// append the signature to the pdf file
	p.PdfFile = append(p.PdfFile, p.Signature...)

	err = os.WriteFile(filePath, p.PdfFile, 0644)
	if err != nil {
		panic(err)
	}

	//Save the signature to a file
	// join filepath and signature
	signPath := filePath + ".sign"
	err = os.WriteFile(signPath, p.Signature, 0644)
	if err != nil {
		panic(err)
	}
}
