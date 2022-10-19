package main

import (
	"fmt"
	"syscall"
)

func sleep_nano(ns int64) (syscall.Timespec, error) {
	dur := syscall.Timespec{
		Sec:  ns / 1000000000,
		Nsec: ns % 1000000000,
	}
	rem := syscall.Timespec{}

	err := syscall.Nanosleep(&dur, &rem)

	return rem, err
}

func sleep_sec(s int64) (syscall.Timespec, error) {
	return sleep_nano(s * 1000000000)
}

func main() {
	fmt.Println(sleep_nano(1 * 1000000000))
	fmt.Println("done")

	fmt.Println(sleep_sec(10))
	fmt.Println("done")
}
