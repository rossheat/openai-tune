package model

type FineTuneConfig struct {
	Model           string           `yaml:"model" json:"model"`
	TrainingFile    string           `yaml:"training_file" json:"training_file"`
	ValidationFile  *string          `yaml:"validation_file,omitempty" json:"validation_file,omitempty"`
	Hyperparameters *Hyperparameters `yaml:"hyperparameters,omitempty" json:"hyperparameters,omitempty"`
	Suffix          *string          `yaml:"suffix,omitempty" json:"suffix,omitempty"`
	Integrations    []Integration    `yaml:"integrations,omitempty" json:"integrations,omitempty"`
	Seed            *int             `yaml:"seed,omitempty" json:"seed,omitempty"`
}
