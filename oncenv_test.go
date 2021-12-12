package oncenv_test

import (
	"testing"

	"github.com/ytakaya/oncenv"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestOncenv(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, oncenv.Analyzer, "a")
}
