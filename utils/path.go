package utils

import (
	"os"
	"path/filepath"
)

// get golang project absolute path
// just return: $GOPATH/src/domain/username/projectName(delimiter is depend on OS)
func GetProjAbsPath(domain, username, projectName string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", domain, username, projectName)
}
