package png

import (
	"testing"
)


func bytesEqual(a []byte,b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}


func testingFuncion(t *testing.T) {
}


func TestNewPNG(t *testing.T) {
	testCases := []struct {
		testName string
		bytesContent []byte
		fileName string
		valExpected *PNG
		errExpected string
	}{
		{       
			"Test for nil bytes",
			nil,
			"something",
			nil,
			"bytesContent cannot be empty or nil",

		},
		{
			"Test for empty file name",
			[]byte{1, 2, 3},
			"",
			nil,
			"fileName cannot be empty",
		},
		{
			"Test for success",
			[]byte{1, 2, 3}, 
			"something.name",
			&PNG{bytesContent: []byte{1, 2, 3}, fileName: "something.name"},
			"",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			gotPng, gotErr := NewPNG(testCase.bytesContent, testCase.fileName)
			if gotErr != nil && testCase.errExpected == "" {
				t.Errorf("error not expected: %s", gotErr.Error())	
			} 
			if gotErr == nil && testCase.errExpected != "" {
				t.Errorf("did not raise an error when expected %s", testCase.errExpected)
			}
			if gotErr != nil && testCase.errExpected != "" {
				if gotErr.Error() != testCase.errExpected {
					t.Errorf("errors do not match: \n expected: %s \n got: %s", testCase.errExpected, gotErr.Error())
				}
			}
			if gotPng != nil && testCase.valExpected == nil {
				t.Errorf("png expected to be nil but got not nil")
			}
			if gotPng == nil && testCase.valExpected != nil {
				t.Errorf("png expected not to be nil but got nil")
			}
			if gotPng != nil && testCase.valExpected != nil {
				if gotPng.fileName != testCase.valExpected.fileName {
					t.Errorf("filenames do not match, expected %s got %s", testCase.valExpected.fileName, gotPng.fileName)
				}
				if !bytesEqual(gotPng.bytesContent,testCase.valExpected.bytesContent) {
					t.Errorf("png bytesContent do not match, expected %v got %v", testCase.valExpected.bytesContent, gotPng.bytesContent)
				}
			}
		})
	}
}
