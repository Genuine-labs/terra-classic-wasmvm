package api

import "testing"

func TestAdd(t *testing.T) {
	res := Add(5, 7)
	if res != 12 {
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestGreet(t *testing.T) {
	res := string(Greet([]byte("Fred")))
	if res != "Hello, Fred" {
		t.Fatalf("Bad greet: %s", res)
	}

	res = string(Greet(nil))
	if res != "Hello, <nil>" {
		t.Fatalf("Bad greet: %s", res)
	}
}

func TestDivide(t *testing.T) {
	res, err := Divide(15, 3)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
	e := getError()
	if e != nil {
		t.Errorf("getError() should return nil, got %s", e.Error())
	}
	if res != 5 {
		t.Fatalf("Unexpected result: %d", res)
	}

	res, err = Divide(6, 0)
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
	errMsg := err.Error()
	if errMsg != "Cannot divide by zero" {
		t.Fatalf("Unexpected error msg: %s", errMsg)
	}
	if res != 0 {
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestRandomMessage(t *testing.T) {
	// this should pass
	res, err := RandomMessage(123)
	if err != nil {
		t.Fatalf("Expected no err, got %s", err)
	}
	if e2 := getError(); e2 != nil {
		t.Fatalf("Expected no getError, got %s", err)
	}
	if res != "You are a winner!" {
		t.Fatalf("Unexpected result: %s", res)
	}

	// this should error (normal)
	res, err = RandomMessage(-20)
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
	if errMsg := err.Error(); errMsg != "Too low" {
		t.Fatalf("Unexpected error msg: %s", errMsg)
	}
	if res != "" {
		t.Fatalf("Unexpected result: %s", res)
	}

	// this should panic
	res, err = RandomMessage(0)
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
	if errMsg := err.Error(); errMsg != "Caught panic" {
		t.Fatalf("Unexpected error msg: %s", errMsg)
	}
	if res != "" {
		t.Fatalf("Unexpected result: %s", res)
	}
	// make sure error cleared after read
	if e2 := getError(); e2 != nil {
		t.Fatalf("Expected no getError, got %s", err)
	}

	// this should pass (again)
	res, err = RandomMessage(789)
	if err != nil {
		t.Fatalf("Expected no err, got %s", err)
	}
	if res != "You are a winner!" {
		t.Fatalf("Unexpected result: %s", res)
	}
}