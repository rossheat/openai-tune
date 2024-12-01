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
	uploadFile := uploadCmd.String("file", "", "JSONL data file to upload")
	uploadList := uploadCmd.Bool("list", false, "list all uploaded files with purpose 'fine-tune'")

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

		if (*uploadFile == "") == (!*uploadList) {
			fmt.Println("please specify either -file or -list")
			uploadCmd.PrintDefaults()
			os.Exit(1)
		}

		options := options.Upload{
			File:         *uploadFile,
			OpenAIAPIKey: openAIAPIKey,
		}

		if *uploadList {
			err := upload.List(options)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			err := upload.Upload(options)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

	}
}
