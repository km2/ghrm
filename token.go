package ghrm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func DefaultTokenPath() string {
	return filepath.Join(xdg.DataHome, "ghrm", "token.json")
}

func ReadToken(tokenPath string) (*Token, error) {
	bytes, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var token Token
	if err := json.Unmarshal(bytes, &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &token, nil
}
