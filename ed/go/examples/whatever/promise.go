package main

import (
  "errors"
  "fmt"
  "time"
)

func main() {
  po := new(PurchaseOrder)
  po.Value = 42.27

  SavePO(po, false).Then(func(obj interface{}) error {
    po := obj.(*PurchaseOrder)
    fmt.Printf("Purchase Order saved with id: '%d'\n", po.Number)

    return nil
  }, func(err error) {
    fmt.Printf("Failed to save Purchase Order: " + err.Error() + "\n")
  }).Then(func(obj interface{}) error {
    fmt.Println("Second promise success")
    return nil
  }, func(err error) {
    fmt.Println("Second promise failure: " + err.Error())
  })

  fmt.Scanln()
}

type PurchaseOrder struct {
  Number int
  Value  float64
}

func SavePO(po *PurchaseOrder, shouldFail bool) *Promise {
  result := new(Promise)

  result.successChannel = make(chan interface{}, 1)
  result.failureChannel = make(chan error, 1)

  go func() {
    time.Sleep(3 * time.Second)
    if shouldFail {
      result.failureChannel <- errors.New("Failed to save purchase order")
    } else {
      po.Number = 1234
      result.successChannel <- po
    }
  }()

  return result
}

type Promise struct {
  successChannel chan interface{}
  failureChannel chan error
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
  result := new(Promise)

  result.successChannel = make(chan interface{}, 1)
  result.failureChannel = make(chan error, 1)

  timeout := time.After(1 * time.Second)

  go func() {
    select {
    case obj := <-this.successChannel:
      newErr := success(obj)
      if newErr == nil {
        result.successChannel <- obj
      } else {
        result.failureChannel <- newErr
      }
    case err := <-this.failureChannel:
      failure(err)
      result.failureChannel <- err
    case <-timeout:
      failure(errors.New("Promise timed out"))
    }
  }()

  return result
}
