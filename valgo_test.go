package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
  var name = "Elon Musk"
  assert.False(t, Is(name).Empty().Valid(), "Expected name was not empty")

  name = ""
  assert.True(t, Is(name).Empty().Valid(), "Expected name was empty")
}

func TestBlank(t *testing.T) {

}
