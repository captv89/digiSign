package cmd

import (
	"digiSign/sign"
	"fmt"

	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a pdf file",
	Long:  "Sign a pdf file using the provided private key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pdfSigner := &sign.PdfSigner{}
		pdfSigner.GenerateKeys()
		pdfSigner.ReadPdfFile(args[0])
		pdfSigner.SignPdfFile()
		fmt.Println("pdf file signed")
		pdfSigner.SaveSignedPdfFile("output/result.pdf")
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
