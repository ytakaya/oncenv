package oncenv_test

import (
	"testing"

	"github.com/ytakaya/oncenv"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestFromFileSystem(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, oncenv.Analyzer, "a")
}
