package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
)

func ParseCommand(s string) (string, []string) {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return "", make([]string, 0)
	}

	parts := strings.Fields(s)

	return parts[0], parts[1:]
}

func HashPassword(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func OpenAssetsFile(fileName string) (*os.File, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(cwd, "assets", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
