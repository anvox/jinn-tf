package configure

import (
	"log"
	"os"
)

func ListStacks(cfg InitConfig) []string {
	f, err := os.Open(cfg.StackDir)
	if err != nil {
		log.Fatal("[ERROR] Cannot find stacks path")
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		log.Fatal("[ERROR] Failed to read stacks path")
	}

	var stacks []string
	for _, file := range files {
		if file.IsDir() {
			_, reserved := reserved_stack_name()[file.Name()]

			if !reserved {
				stacks = append(stacks, file.Name())
			}
		}
	}

	return stacks
}

func reserved_stack_name() map[string]bool {
	result := make(map[string]bool)

	result["modules"] = true

	return result
}
