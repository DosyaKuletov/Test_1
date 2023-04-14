package Test_1

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	testCases := []struct {
		testCaseName   string
		input          int
		expectedResult bool
		expectedMsg    string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, tc := range testCases {
		actualResult, msg := isPrime(tc.input)
		if tc.expectedResult && !actualResult {
			t.Errorf("%s: expected true but got false", tc.testCaseName)
		}

		if !tc.expectedResult && actualResult {
			t.Errorf("%s: expected false but got true", tc.testCaseName)
		}

		if tc.expectedMsg != msg {
			t.Errorf("%s: expected %s but got %s", tc.testCaseName, tc.expectedMsg, msg)
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	testTable := []struct {
		name         string
		testValue    string
		expectedDone bool
		expectedMsg  string
	}{
		{"sending integer", "7", false, "7 is a prime number!"},
		{"sending non-integer", "8.5", false, "Please enter a whole number!"},
		{"sending quit command", "q", true, ""},
	}

	for _, e := range testTable {
		buf := bytes.NewBufferString(e.testValue)
		msg, resultDone := checkNumbers(bufio.NewScanner(buf))
		if e.expectedDone && !resultDone {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expectedDone && resultDone {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.expectedMsg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.expectedMsg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	err := w.Close()
	if err != nil {
		t.Fatalf("error was returned during the test: %s", err)
	}

	out, _ := io.ReadAll(r)

	expected := "-> "
	actual := string(out)
	if actual != expected {
		t.Errorf("prompt: expected %s but got %s", expected, actual)
	}
}

func Test_intro(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	err := w.Close()
	if err != nil {
		t.Fatalf("error was returned during the test: %s", err)
	}

	out, _ := io.ReadAll(r)
	expected :=
		"Is it Prime?\n" +
			"------------\n" +
			"Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n" +
			"-> "
	actual := string(out)
	if actual != expected {
		t.Errorf("prompt: expected %s but got %s", expected, actual)
	}
}

func Test_readUserInput(t *testing.T) {
	input := "7\nq"
	expected := "7 is a prime number!\n" +
		"-> "

	r, w, _ := os.Pipe()
	os.Stdout = w
	in := strings.NewReader(input)
	doneChan := make(chan bool)

	go readUserInput(in, doneChan)

	<-doneChan

	close(doneChan)

	err := w.Close()
	if err != nil {
		t.Fatalf("error was returned during the test: %s", err)
	}

	out, _ := io.ReadAll(r)
	actual := string(out)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
