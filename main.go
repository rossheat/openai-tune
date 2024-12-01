// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rossheat/openai-tune/create"
	"github.com/rossheat/openai-tune/list"
	"github.com/rossheat/openai-tune/option"
	"github.com/rossheat/openai-tune/upload"
	"github.com/rossheat/openai-tune/utils"
)

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  openai-tune upload -file <path-to-jsonl-file>  Upload a JSONL file for fine-tuning")
	fmt.Println("  openai-tune upload -list                       List all uploaded files")
	fmt.Println("  openai-tune create -file-id <file-id> -model <model-name>  Create a fine-tuning job with default settings")
	fmt.Println("  openai-tune create -config <path-to-yaml>      Create a fine-tuning job with custom settings")
	fmt.Println("  openai-tune list [-limit <n>] [-after <job-id>] List fine-tuning jobs")
	os.Exit(1)
}

func main() {
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	uploadFile := uploadCmd.String("file", "", "JSONL data file to upload")
	uploadList := uploadCmd.Bool("list", false, "list all uploaded files with purpose 'fine-tune'")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createFileID := createCmd.String("file-id", "", "File ID of JSONL data file uploaded to OpenAI")
	createModel := createCmd.String("model", "gpt-4o-mini-2024-07-18", "Model to fine-tune (only used with -file-id)")
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
		if (*createFileID == "") == (*configFile == "") {
			fmt.Println("please specify either -file-id or -config")
			createCmd.PrintDefaults()
			os.Exit(1)
		}
		options := option.Create{
			FileID:       *createFileID,
			Model:        *createModel,
			ConfigFile:   *configFile,
			OpenAIAPIKey: openAIAPIKey,
		}
		err := create.Create(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		limit := listCmd.Int("limit", 0, "Number of fine-tuning jobs to retrieve")
		after := listCmd.String("after", "", "Retrieve jobs after this job ID")
		listCmd.Parse(os.Args[2:])

		options := option.List{
			OpenAIAPIKey: openAIAPIKey,
			Limit:        *limit,
			After:        *after,
		}

		err := list.List(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		PrintUsage()
	}
}
