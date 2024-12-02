# OpenAI Tune

A pleasant CLI tool for fine-tuning OpenAI models

## Quickstart

1. Confirm you have Go (version>=1.16) installed:
```bash
go version
```

2. Install openai-tune:
```bash
go install github.com/rossheat/openai-tune@latest
```

3. Confirm installation: 
```bash
openai-tune
```

If you see "command not found" you may need to add `$GOPATH/bin` to your PATH: 

For bash:
```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc && source ~/.bashrc
```

For zsh: 
```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc
```

4. Add your OpenAI API key to your environment:
```bash
export OPENAI_API_KEY=<OPENAI_API_KEY>
```

## Usage

Upload a JSONL file for fine-tuning:
```bash
openai-tune upload -file <path-to-jsonl-file>
```

List all uploaded files with the purpose 'fine-tune':
```bash
openai-tune upload -list
```

Create a fine-tuning job with default settings:
```bash
openai-tune create -file-id <file-id> -model <model-name>
```

See this [list of models](https://platform.openai.com/docs/guides/fine-tuning/#which-models-can-be-fine-tuned) that can be fine-tuned.

Create a fine-tuning job with custom settings from a YAML config:
```bash
openai-tune create -config <path-to-yaml>
```

See [config.example.yml](config.example.yml) for all available options.

List fine-tuning jobs:
```bash
openai-tune list [-limit <n>] [-after <job-id>]
```

Get information about a specific fine-tuning job:
```bash
openai-tune get <job-id>
```

Cancel a fine-tuning job:
```bash
openai-tune cancel <job-id>
```

## Weights & Biases Integration

In order to enable logging with W&B, first please follow these [instructions](https://platform.openai.com/docs/guides/fine-tuning#weights-and-biases-integration) to enable the W&B integration for your OpenAI account.

Once you've enabled the W&B integration on your OpenAI account, specified your W&B integration in your config file, and created your fine-tuning job, you can view the job in W&B by navigating to the W&B project you specified in the job creation request. Your run should be located at the URL:
`https://wandb.ai/<WANDB-ENTITY>/<WANDB-PROJECT>/runs/ftjob-ABCDEF`

