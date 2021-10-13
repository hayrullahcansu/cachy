package utility

import (
	"io/ioutil"
	"os"

	"github.com/hayrullahcansu/cachy/cross"
	"github.com/hayrullahcansu/cachy/framework/logging"
)

func ReadFileAsString(filePath string) string {
	b, err := ReadFile(filePath)
	if err != nil {
		logging.Infof("Cannot read file: %s%s", filePath, cross.NewLine)
		return ""
	}
	return string(b)
}

func ReadFile(filePath string) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logging.Errorf("open file error: %v", err)
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
