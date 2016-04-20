package fileutils
import (
	"testing"
	"os"
)


func TestCopyDir(t *testing.T){
	src := "/Users/vanja/code/go/src/github.com/tornyak"
	info, err := os.Stat(src)
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}
	CopyDir(src, src, info.Mode().Perm())
}
