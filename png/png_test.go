package png


import (
	"testing"
	"github.com/swims-hjkl/iden-colour/utils"
)



func TestNewPNG(t *testing.T) {
	var headerBytes = []byte{0X89, 0X50, 0X4E, 0X47, 0X0D, 0X0A, 0X1A, 0x0A}
	testCases := []struct {
		testName string
		bytesContent []byte
		fileName string
		valExpected *PNG
		errExpected string
	}{
		{
			"Test for success",
			headerBytes, 
			"something.name",
			&PNG{bytesContent: headerBytes, fileName: "something.name"},
			"",
		},
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
			"Test for lesser number of bytes than header",
			[]uint8{1, 2, 3}, 
			"something.name",
			nil,
			"png header is not valid",
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
				if !utils.BytesEqual(gotPng.bytesContent,testCase.valExpected.bytesContent) {
					t.Errorf("png bytesContent do not match, expected %v got %v", testCase.valExpected.bytesContent, gotPng.bytesContent)
				}
			}
		})
	}
}
