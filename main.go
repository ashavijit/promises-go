package main

import (
        "errors"
        "fmt"
)

type Promise struct {
        onFulfilled func(interface{}) interface{}
        onRejected  func(error) interface{}
        onFinally   func() interface{}
}

func NewPromise() *Promise {
        return &Promise{}
}

func (p *Promise) Then(onFulfilled func(interface{}) interface{}, onRejected func(error) interface{}) *Promise {
        p.onFulfilled = onFulfilled
        p.onRejected = onRejected
        return p
}

func (p *Promise) Catch(onRejected func(error) interface{}) *Promise {
        p.onRejected = onRejected
        return p
}

func (p *Promise) Finally(onFinally func() interface{}) *Promise {
        p.onFinally = onFinally
        return p
}

func (p *Promise) Resolve(value interface{}) {
        if p.onFulfilled != nil {
                value = p.onFulfilled(value)
        }

        if p.onFinally != nil {
                p.onFinally()
        }
}

func (p *Promise) Reject(err error) {
        if p.onRejected != nil {
                value := p.onRejected(err)
                if value != nil {
                        fmt.Println("Unhandled promise rejection:", value)
                }
        }

        if p.onFinally != nil {
                p.onFinally()
        }
}

func main() {
        p := NewPromise()

        p.Then(func(value interface{}) interface{} {
                fmt.Println("Fulfilled:", value)
                return value
        }, nil).
                Catch(func(err error) interface{} {
                        fmt.Println("Rejected:", err)
                        return nil
                }).
                Finally(func() interface{} {
                        fmt.Println("Finally")
                        return nil
                })

        p.Resolve("Hello, World!")
        p.Reject(errors.New("Something went wrong"))
}
