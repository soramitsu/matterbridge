package helper

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testLineLength = 64

var lineSplittingTestCases = map[string]struct {
	input          string
	splitOutput    []string
	nonSplitOutput []string
}{
	"Short single-line message": {
		input:          "short",
		splitOutput:    []string{"short"},
		nonSplitOutput: []string{"short"},
	},
	"Long single-line message": {
		input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		splitOutput: []string{
			"Lorem ipsum dolor sit amet, consectetur adipis <clipped message>",
			"cing elit, sed do eiusmod tempor incididunt ut <clipped message>",
			" labore et dolore magna aliqua.",
		},
		nonSplitOutput: []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
	},
	"Short multi-line message": {
		input: "I\ncan't\nget\nno\nsatisfaction!",
		splitOutput: []string{
			"I",
			"can't",
			"get",
			"no",
			"satisfaction!",
		},
		nonSplitOutput: []string{
			"I",
			"can't",
			"get",
			"no",
			"satisfaction!",
		},
	},
	"Long multi-line message": {
		input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\n" +
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.\n" +
			"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.\n" +
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		splitOutput: []string{
			"Lorem ipsum dolor sit amet, consectetur adipis <clipped message>",
			"cing elit, sed do eiusmod tempor incididunt ut <clipped message>",
			" labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercita <clipped message>",
			"tion ullamco laboris nisi ut aliquip ex ea com <clipped message>",
			"modo consequat.",
			"Duis aute irure dolor in reprehenderit in volu <clipped message>",
			"ptate velit esse cillum dolore eu fugiat nulla <clipped message>",
			" pariatur.",
			"Excepteur sint occaecat cupidatat non proident <clipped message>",
			", sunt in culpa qui officia deserunt mollit an <clipped message>",
			"im id est laborum.",
		},
		nonSplitOutput: []string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
			"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		},
	},
	"Message ending with new-line.": {
		input:          "Newline ending\n",
		splitOutput:    []string{"Newline ending"},
		nonSplitOutput: []string{"Newline ending"},
	},
	"Long message containing UTF-8 multi-byte runes": {
		input: "???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????",
		splitOutput: []string{
			"????????????????????????????????????????????? <clipped message>",
			"????????????????????????????????????????????? <clipped message>",
			"????????????????????????????????????????????? <clipped message>",
			"????????????????????????????????????????????? <clipped message>",
			"???????????????????????????",
		},
		nonSplitOutput: []string{"???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????"},
	},
}

func TestGetSubLines(t *testing.T) {
	for testname, testcase := range lineSplittingTestCases {
		splitLines := GetSubLines(testcase.input, testLineLength, "")
		assert.Equalf(t, testcase.splitOutput, splitLines, "'%s' testcase should give expected lines with splitting.", testname)
		for _, splitLine := range splitLines {
			byteLength := len([]byte(splitLine))
			assert.True(t, byteLength <= testLineLength, "Splitted line '%s' of testcase '%s' should not exceed the maximum byte-length (%d vs. %d).", splitLine, testcase, byteLength, testLineLength)
		}

		nonSplitLines := GetSubLines(testcase.input, 0, "")
		assert.Equalf(t, testcase.nonSplitOutput, nonSplitLines, "'%s' testcase should give expected lines without splitting.", testname)
	}
}

func TestConvertWebPToPNG(t *testing.T) {
	if os.Getenv("LOCAL_TEST") == "" {
		t.Skip()
	}

	input, err := ioutil.ReadFile("test.webp")
	if err != nil {
		t.Fail()
	}

	d := &input
	err = ConvertWebPToPNG(d)
	if err != nil {
		t.Fail()
	}

	err = ioutil.WriteFile("test.png", *d, 0o644) // nolint:gosec
	if err != nil {
		t.Fail()
	}
}
