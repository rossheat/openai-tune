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

TODO 

Please see [config.example.yml](config.example.yml) to view all of the options you can set for custom fine-tuning.

## Weights & Biases Integration

You can configuration a W&B integration in your config file. Please follow these [instructions](https://platform.openai.com/docs/guides/fine-tuning#weights-and-biases-integration) to enable the W&B integration for your OpenAI account.