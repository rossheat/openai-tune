package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rossheat/openai-tune/options"
	"github.com/rossheat/openai-tune/upload"
	"github.com/rossheat/openai-tune/utils"
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

	openAIAPIKey, err := utils.GetOpenAIAPIKeyFromEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "upload":
		uploadCmd.Parse(os.Args[2:])
		if *uploadFile == "" {
			fmt.Println("please specify a file using -file")
			uploadCmd.PrintDefaults()
			os.Exit(1)
		}
		options := options.Upload{
			File:         *uploadFile,
			Purpose:      *uploadPurpose,
			OpenAIAPIKey: openAIAPIKey,
		}
		upload.Upload(options)
	}
}
