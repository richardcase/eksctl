package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/kubicorn/kubicorn/pkg/namer"
)

func ClusterName() string {
	return fmt.Sprintf("%s-%d", namer.RandomName(), time.Now().Unix())
}

// FileExists checks to see if a file exists.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
