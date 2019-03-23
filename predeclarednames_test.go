package predeclarednames_test

import (
	"testing"

	"github.com/YoshikiShibata/predeclarednames"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, predeclarednames.Analyzer, "a")
}