package vault

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrNoVault = errors.New("no vault found: set MAGPIE_VAULT, add default_vault to config, or run from inside a vault")

func Resolve(cwd string, defaultValue string) (string, error) {
	vaultRoot, exists := os.LookupEnv("MAGPIE_VAULT")
	if exists {
		return vaultRoot, nil
	}

	for dir := cwd; ; dir = filepath.Dir(dir) {
		if isVaultRoot(dir) {
			return dir, nil
		}
		if dir == filepath.Dir(dir) {
			break
		}
	}

	if defaultValue != "" {
		return defaultValue, nil
	}

	return "", ErrNoVault
}

// isVaultRoot reports whether dir is a vault root (i.e. contains a .magpie directory)
func isVaultRoot(dir string) bool {
	info, err := os.Stat(filepath.Join(dir, ".magpie"))
	if err != nil {
		return false
	}
	return info.IsDir()
}
