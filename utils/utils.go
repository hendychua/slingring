package utils

import (
  "log"
)

// Check checks if err is not nil and calls log.Fatal(err).
func Check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}