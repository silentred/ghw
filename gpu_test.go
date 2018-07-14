//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"os"
	"testing"
)

func TestGPU(t *testing.T) {
	if _, ok := os.LookupEnv("GHW_TESTING_SKIP_GPU"); ok {
		t.Skip("Skipping GPU tests.")
	}
	info, err := GPU()
	if err != nil {
		t.Fatalf("Expected no error creating GPUInfo, but got %v", err)
	}

	if len(info.GraphicsCards) == 0 {
		t.Fatalf("Expected >0 GPU cards, but found 0.")
	}
}
