package upload

import "fmt"

func Upload(file *string, purpose *string) {
	fmt.Printf("Upload function received file=%v and purpose=%v\n", *file, *purpose)
}
