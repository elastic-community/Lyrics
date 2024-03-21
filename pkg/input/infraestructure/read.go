package infraestructure

import (
	"os"
)

func Read(filePath string) (*os.File, error) {
	return os.Open(filePath)
}
