package create

import (
	"fmt"

	"github.com/rossheat/openai-tune/option"
)

func Create(options option.Create) error {
	fmt.Printf("Create func received options: %v", options)
	return nil
}