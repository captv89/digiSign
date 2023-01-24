package cmd

import (
	"digiSign/sign"
	"fmt"

	"github.com/spf13/cobra"
)

var privateKeyFilePath string

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a pdf file",
	Long:  "Sign a pdf file using the provided private key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pdfSigner := &sign.PdfSigner{}

		if privateKeyFilePath != "" {
			_ = pdfSigner.ReadPrivateKey(privateKeyFilePath)
		} else {
			pdfSigner.GenerateKeys()
		}
		pdfSigner.ReadPdfFile(args[0])
		pdfSigner.SignPdfFile()
		fmt.Println("PDF file signed successfully!")
		pdfSigner.SaveSignedPdfFile("output/result.pdf")
		fmt.Println("Signed PDF & signature file saved locally!")
		if privateKeyFilePath == "" {
			pdfSigner.SaveKeysToFile("output/private.key", "output/public.key")
			fmt.Println("Keys saved to file. Keep them safe!")
		}
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	signCmd.Flags().StringVarP(&privateKeyFilePath, "private-key", "k", "", "Path to the private key file")
}
