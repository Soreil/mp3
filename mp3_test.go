package mp3

import (
	"bytes"
	"os"
	"testing"

	"github.com/Soreil/imager"
)

const dataDir = "inputData/"

type testCase struct {
	isMP3         bool
	inputFilename string
}

var cases = []testCase{
	{false, "exampleFalse.mp3"},
	{true, "exampleTrue.mp3"},
	{true, "exampleImage.mp3"},
	{true, "examplePNG.mp3"},
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
		if IsMP3(test.inputFilename) {
			if b, err := ExtractImage(test.inputFilename); err == nil {
				t.Log(test.inputFilename, "has an image inside of it")
				_, inFmt, outFmt, err := imager.Thumbnail(bytes.NewReader(b), imager.Sharp)
				if err != nil {
					t.Fatal(err)
				}
				t.Log("Input format:", inFmt)
				t.Log("Output format:", outFmt)
			} else {
				t.Log(err, test)
			}
		}
	}
}
