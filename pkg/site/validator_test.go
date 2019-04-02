package site

import (
	"testing"
)

const passMark = "\u2713"
const failMark = "\u2717"

func assertResponseEqual(t *testing.T, expected bool, actual bool) {
	t.Helper() // comment this line to see tests fail due to 'if expected != actual'
	if expected != actual {
		t.Errorf("%t != %t %s", expected, actual, failMark)
	} else {
		t.Logf("%t == %t %s", expected, actual, passMark)
	}
}

func TestIsValidRequestURL(t *testing.T) {

	for _, testCase := range []struct {
		Name     string
		URL      string
		Expected bool
	}{
		{"Valid HTTP URL", "http://www.google.com", true},
		{"Valid HTTPS URL", "https://www.google.com", true},
		{"Invalid HTTPS URL", "https", false},
		{"Invalid URL - Without protocol", "www.example.txt", false},
		{"Invalid URL - empty string", "", false},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := IsValidRequestURL(testCase.URL)

			assertResponseEqual(t, testCase.Expected, actual)
		})
	}

}
