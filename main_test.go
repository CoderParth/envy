package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var expected = map[string]bool{
	"DB_URL=":                true,
	"AWS_ACCESS_KEY_ID=":     true,
	"AWS_ACCESS_KEY_SECRET=": true,
	"SENDGRID_API_KEY=":      true,
	"FROM_EMAIL=":            true,
	"API_KEY=":               true,
	"API_SECRET=":            true,
	"JWT_SECRET=":            true,
	"REDIS_URL=":             true,
	"COOKIE_SECRET=":         true,
}

func TestMainProgram(t *testing.T) {
	// Build the program
	cmd := exec.Command("go", "build")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build program: %v", err)
	}

	// Run the program
	cmd = exec.Command("./envy")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run program: %v", err)
	}

	// Check if the .env.example file was created
	_, err = os.Stat(".env.example")
	if os.IsNotExist(err) {
		t.Fatalf(".env.example file was not created")
	}

	// Read the .env.example file
	data, err := os.ReadFile(".env.example")
	if err != nil {
		t.Fatalf("Failed to read .env.example: %v", err)
	}

	strData := string(data)
	fmt.Println(strData)

	// Check if the file contains the expected content
	for k := range expected {
		if !strings.Contains(strData, k) {
			t.Fatalf(".env.example does not contain the expected content")
		}
	}
}
