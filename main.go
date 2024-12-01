package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rossheat/openai-tune/create"
	"github.com/rossheat/openai-tune/option"
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

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createFileID := createCmd.String("file-id", "", "File ID of JSONL data file uploaded to OpenAI")
	configFile := createCmd.String("config", "", "YAML file containing your custom fine-tune settings")

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

		options := option.Upload{
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
	case "create":
		createCmd.Parse(os.Args[2:])
		if *createFileID == "" {
			fmt.Println("please specify -file-id")
			createCmd.PrintDefaults()
			os.Exit(1)
		}

		options := option.Create{
			FileID:       *createFileID,
			ConfigFile:   *configFile,
			OpenAIAPIKey: openAIAPIKey,
		}

		err := create.Create(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
