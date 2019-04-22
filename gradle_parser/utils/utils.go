package utils

import (
    "flag"

)

// utils

type isErr error

func isErr(e error) {
    if e != nil {
        panic(e)
    }
}

type isFlagPassed string

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}
