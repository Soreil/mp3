package mp3

import (
	"os"
	"testing"
)

const dataDir = "inputData/"

type testCase struct {
	isMP3         bool
	inputFilename string
}

var cases = []testCase{
	{false, "exampleFalse.mp3"},
	{true, "exampleTrue.mp3"},
}

func TestChecker(t *testing.T) {
	if err := os.Chdir(dataDir); err != nil {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.Chdir(dataDir); err != nil {
			t.Fatal(err)
		}
	}
	for _, test := range cases {
		if IsMP3(test.inputFilename) != test.isMP3 {
			t.Fatal(test.inputFilename+":", "Wrong assumption about whether or not this file is an MP3 file")
		}
	}
}
