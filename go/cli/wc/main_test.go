package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	res := count(b, false, false)

	assert.Equal(t, 4, res)
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	res := count(b, true, false)

	assert.Equal(t, 3, res)
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("0123456789")

	res := count(b, false, true)

	assert.Equal(t, 10, res)
}
