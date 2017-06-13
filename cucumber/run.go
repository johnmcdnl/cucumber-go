package cucumber

import (
	"path/filepath"
	"os"
	"strings"
	"fmt"
)

const (
	featureExtension = ".feature"
)

func Run(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, featureExtension) {
			fmt.Println(path)
		}
		return err
	})
	return err

}
