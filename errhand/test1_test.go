package tests

import (
	"errors"
	"testing"
)

// // filename ends with .TXT
// func TestHasTXTExtension(t *testing.T) {
//     filename := "example.TXT"
//     err := HasTXTExtension(filename)
//     if err != nil {
//         t.Errorf("Expected nil error, got: %v", err)
//     }
// }

// filename has multiple extensions, but ends with .txt
func TestHasTXTExtension_MultipleExtensionsWithTXT(t *testing.T) {
	filename := "file.tar.gz.txt"
	err := HasTXTExtension(filename)
	if err != nil {
		t.Errorf("Expected nil error, got: %v", err)
	}
}

// filename has spaces before or after, but ends with .txt
func TestHasTXTExtensionWithSpacesBeforeOrAfter(t *testing.T) {
	filename := "  test.txt  "
	err := HasTXTExtension(filename)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// filename has special characters, but ends with .txt
func TestHasTXTExtensionWithSpecialCharacters(t *testing.T) {
	filename := "file@name!.txt"
	err := HasTXTExtension(filename)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// filename is empty
func TestHasTXTExtensionEmptyFilename(t *testing.T) {
	err := HasTXTExtension("")
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else {
		expected := "not .txt"
		if err.Error() != expected {
			t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
		}
	}
}

// filename has no extension
func TestHasTXTExtension_NoExtension(t *testing.T) {
	filename := "file"
	err := HasTXTExtension(filename)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expected := "not .txt"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

// filename has extension, but not .txt
func TestHasExtensionButNotTXT(t *testing.T) {
	filename := "file.doc"
	err := HasTXTExtension(filename)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expected := "not .txt"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

// filename has multiple extensions, but not .txt
func TestHasMultipleExtensionsNotTXT(t *testing.T) {
	filename := "file.docx.pdf"
	err := HasTXTExtension(filename)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expected := "not .txt"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

// filename has spaces before or after, and not .txt
func TestHasSpacesBeforeOrAfterNotTXT(t *testing.T) {
	filename := "  file.txt  "
	err := HasTXTExtension(filename)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else {
		expected := "not .txt"
		if err.Error() != expected {
			t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
		}
	}
}

// filename has special characters, and not .txt
func TestHasTXTExtension_SpecialCharacters(t *testing.T) {
	filename := "file@name"
	err := HasTXTExtension(filename)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := errors.New("not .txt")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error message '%s', got '%s'", expectedErr.Error(), err.Error())
	}
}

// returns error message when filename is empty
func Test_empty_filename(t *testing.T) {
	err := HasTXTExtension("")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	expected := "not .txt"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', but got '%s'", expected, err.Error())
	}
}

// returns error message when filename is too long
func Test_filename_too_long(t *testing.T) {
	longFilename := "this_is_a_very_long_filename_that_exceeds_the_maximum_allowed_length.txt"
	err := HasTXTExtension(longFilename)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	expectedErrorMsg := "not .txt"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMsg, err.Error())
	}
}

// returns error message when filename is a directory
func Test_has_txt_extension_with_directory_filename(t *testing.T) {
	err := HasTXTExtension("/path/to/directory")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	expected := "not .txt"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', but got '%s'", expected, err.Error())
	}
}

// returns error message when filename is a binary file
func Test_has_txt_extension_with_binary_file(t *testing.T) {
	filename := "file.bin"
	err := HasTXTExtension(filename)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else {
		expected := "not .txt"
		if err.Error() != expected {
			t.Errorf("Expected error message '%s', but got '%s'", expected, err.Error())
		}
	}
}
