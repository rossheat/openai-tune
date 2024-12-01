package model

type WandbConfig struct {
    Project string   `yaml:"project" json:"project"`               
    Name    *string  `yaml:"name,omitempty" json:"name,omitempty"`
    Entity  *string  `yaml:"entity,omitempty" json:"entity,omitempty"`
    Tags    []string `yaml:"tags,omitempty" json:"tags,omitempty"`
}