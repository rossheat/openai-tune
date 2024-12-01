package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rossheat/openai-tune/upload"
)

func PrintUsage() {
	panic("unimplemented")
}

func main() {
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	uploadFile := uploadCmd.String("file", "", "training data file (JSONL format) to upload")
	uploadPurpose := uploadCmd.String("purpose", "fine-tune", "purpose of the file")

	if len(os.Args) < 2 {
		PrintUsage()
	}

	switch os.Args[1] {
	case "upload":
		uploadCmd.Parse(os.Args[2:])
		if *uploadFile == "" {
			fmt.Println("please specify a file using -file")
			uploadCmd.PrintDefaults()
			os.Exit(1)
		}
		upload.Upload(uploadFile, uploadPurpose)
	}
}
