// Package vault provides vault discovery and path resolution.
package vault

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrNoVault is returned by Resolve when no vault can be located through any tier.
var ErrNoVault = errors.New("no vault found: set MAGPIE_VAULT, add default_vault to config, or run from inside a vault")

// Resolve returns the vault root for the given working directory using a
// three-tier lookup:
//
//  1. The MAGPIE_VAULT environment variable, if set.
//  2. An upward directory walk from cwd to the filesystem root, stopping at
//     the first directory that contains a .magpie subdirectory.
//  3. The defaultVault string, if non-empty.
//
// Returns ErrNoVault if no vault is found through any tier.
func Resolve(cwd string, defaultValue string) (string, error) {
	vaultRoot, exists := os.LookupEnv("MAGPIE_VAULT")
	if exists && vaultRoot != "" {
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
