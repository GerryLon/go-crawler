package utils

import (
	"fmt"
	"testing"
)

// go test -v xxx will get more output
func TestGetProjAbsPath(t *testing.T) {
	projPath := GetProjAbsPath("github.com", "GerryLon", "go-crawler")
	t.Log(projPath)
	fmt.Println("projPath is:", projPath)
}
