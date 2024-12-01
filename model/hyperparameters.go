package model

type Hyperparameters struct {
	BatchSize              interface{} `yaml:"batch_size,omitempty" json:"batch_size,omitempty"`
	LearningRateMultiplier interface{} `yaml:"learning_rate_multiplier,omitempty" json:"learning_rate_multiplier,omitempty"`
	NEpochs                interface{} `yaml:"n_epochs,omitempty" json:"n_epochs,omitempty"`
}
