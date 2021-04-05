package main

import (
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	in := strings.NewReader(`folder,favorite,type,name,notes,fields,login_uri,login_username,login_password,login_totp
social,,login,Example,sample note,,https://example.com,examplename,examplepassword,
`)
	expected := `title,website,username,password,notes
Example,example.com,examplename,examplepassword,sample note
`
	out := strings.Builder{}
	convert(in, &out)
	got := out.String()
	if got != expected {
		t.Errorf(`
Expected:
----
%s

Got:
----
%s`, expected, got)
	}
}

func TestAppendCsvToOutFileName(t *testing.T) {
	outFileName := "out"
	expected := "out.csv"
	appendCsvToOutFileName(&outFileName)
	if outFileName != expected {
		t.Errorf("Expected: %s, got: %s", outFileName, expected)
	}
}
