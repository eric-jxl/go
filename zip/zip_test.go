package zip

import "testing"

func TestCreateZip(t *testing.T) {
	CreateZip("zip.go","eric.zip")
	t.Log("Success")
}
