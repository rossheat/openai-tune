package upload

import (
	"fmt"

	"github.com/rossheat/openai-tune/options"
)

func Upload(options options.Upload) {
	fmt.Printf("Upload function received file=%v and purpose=%v\n", options.File, options.Purpose)
}
