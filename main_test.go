package main

import (
	"errors"
	"fmt"
	"testing"
)


func TestPromise_Resolve(t *testing.T) {
	p := NewPromise()

	var fulfilledValue interface{}
	var finallyCalled bool

	onFulfilled := func(value interface{}) interface{} {
		fulfilledValue = value
		return value
	}

	onFinally := func() interface{} {
		finallyCalled = true
		return nil
	}

	p.Then(onFulfilled, nil).Finally(onFinally)
	p.Resolve("Hello, World!")

	if fulfilledValue != "Hello, World!" {
		t.Errorf("Unexpected fulfilled value: got %v, want %v", fulfilledValue, "Hello, World!")
	}

	if !finallyCalled {
		t.Error("Expected finally callback to be called, but it wasn't")
	}
}

func TestPromise_Reject(t *testing.T) {
	p := NewPromise()

	var rejectedError error
	var finallyCalled bool

	onRejected := func(err error) interface{} {
		rejectedError = err
		return nil
	}

	onFinally := func() interface{} {
		finallyCalled = true
		return nil
	}

	p.Catch(onRejected).Finally(onFinally)
	p.Reject(errors.New("Something went wrong"))

	if rejectedError == nil {
		t.Error("Expected rejected error, got nil")
	}

	if !finallyCalled {
		t.Error("Expected finally callback to be called, but it wasn't")
	}
}

func TestPromise_Then(t *testing.T) {
	p := NewPromise()

	var onFulfilledCalled bool

	onFulfilled := func(value interface{}) interface{} {
		onFulfilledCalled = true
		return value
	}

	p.Then(onFulfilled, nil)
	p.Resolve("Hello, World!")

	if !onFulfilledCalled {
		t.Error("Expected onFulfilled callback to be called, but it wasn't")
	}
}

func TestPromise_Catch(t *testing.T) {
	p := NewPromise()

	var onRejectedCalled bool

	onRejected := func(err error) interface{} {
		onRejectedCalled = true
		return nil
	}

	p.Catch(onRejected)
	p.Reject(errors.New("Something went wrong"))

	if !onRejectedCalled {
		t.Error("Expected onRejected callback to be called, but it wasn't")
	}
}

func TestPromise_Finally(t *testing.T) {
	p := NewPromise()

	var onFinallyCalled bool

	onFinally := func() interface{} {
		onFinallyCalled = true
		return nil
	}

	p.Finally(onFinally)
	p.Resolve("Hello, World!")

	if !onFinallyCalled {
		t.Error("Expected onFinally callback to be called, but it wasn't")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running tests...")
	m.Run()
	fmt.Println("Tests finished.")
}
