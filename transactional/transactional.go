package main

import "fmt"
import "errors"

func helloWolrd() (err error) {
    rollback := func(rollback func()) {
        if err != nil {
            rollback()
        }
    }

    err = initA()
    if err != nil {
        return err
    }
    defer rollback(rollbackInitA)

    err = initB()
    if err != nil {
        return err
    }
    defer rollback(rollbackInitB)

    err = initC()
    if err != nil {
        return err
    }
    defer rollback(rollbackInitC)

    return nil
}

func main(){
    helloWolrd()
}

func initA() (err error) {
    fmt.Println("init A")
    return nil
}

func initB() (err error) {
    fmt.Println("init B")
    return nil
}

func initC() (err error) {
    fmt.Println("init C")
    return errors.New("roll back test")
}

func rollbackInitA() {
    fmt.Println("rollback Init A")
}

func rollbackInitB() {
    fmt.Println("rollback Init B")
}

func rollbackInitC() {
    fmt.Println("rollback Init C")
}