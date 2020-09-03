package countname_test

import (
	"testing"

	"github.com/harukitosa/merucari/countname"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, countname.Analyzer, "a")
}

