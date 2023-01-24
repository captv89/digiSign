package cmd

import (
	"digiSign/sign"
	"fmt"
	"github.com/spf13/cobra"
)

var publicKeyFilePath string
var signKeyFilePath string

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the sign of pdf file",
	Long:  `Verify the sign of pdf file using the provided public key.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pdfSigner := &sign.PdfSigner{}
		if publicKeyFilePath == "" {
			fmt.Print("Please provide the public key file path")
			_, err := fmt.Scanf("%s", &publicKeyFilePath)
			if err != nil {
				fmt.Println(err)
			}
		}
		if signKeyFilePath == "" {
			fmt.Print("Please provide the public key file path")
			_, err := fmt.Scanf("%s", &signKeyFilePath)
			if err != nil {
				fmt.Println(err)
			}
		}

		err := pdfSigner.ReadPublicKey(publicKeyFilePath)
		if err != nil {
			return
		}
		pdfSigner.ReadSignature(signKeyFilePath)
		pdfSigner.ReadPdfFile(args[0])
		if pdfSigner.VerifyPdfFile() {
			fmt.Println("PDF file verified successfully!")
		} else {
			fmt.Println("PDF file not verified!")
		}

	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.Flags().StringVarP(&publicKeyFilePath, "public-key", "p", "", "Path to the public key file")

	verifyCmd.Flags().StringVarP(&signKeyFilePath, "sign-file", "s", "", "Path to the public key file")
}
