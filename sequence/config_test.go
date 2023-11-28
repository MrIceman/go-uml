package sequence

import (
	"os"
	"testing"
)

func TestDiagramFromYaml(t *testing.T) {
	d, err := DiagramFromYaml("test.yaml")	
	if err != nil {
		t.Fatalf("err %v", err)
	}

	defer os.Remove("test.png")

	d.Render()
}
