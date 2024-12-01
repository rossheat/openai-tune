package model

type Integration struct {
    Type  string      `yaml:"type" json:"type"`     
    Wandb WandbConfig `yaml:"wandb" json:"wandb"`   
}