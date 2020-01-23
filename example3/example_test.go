package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteAndReadJson(t *testing.T) {
	filename := "/tmp/person.json"
	src := Alex
	copy := new(Person)

	var err error
	err = writeJson(filename, src)
	assert.NoError(t, err, "failed writing json")

	err = readJson(filename, copy)
	assert.NoError(t, err, "failed reading json")

	assert.Equal(t, src, copy, "orig and copy not equal")
}

func TestWriteJson(t *testing.T) {
	var err error

	err = writeJson("/tmp/person.json", Alex)
	assert.NoError(t, err, "failed writing json")
}

func TestWriteAndReadComplexStruct(t *testing.T) {
	filename := "/tmp/family.json"
	src := Family
	copy := new(People)

	var err error
	err = writeJson(filename, src)
	assert.NoError(t, err, "failed writing json")

	err = readJson(filename, copy)
	assert.NoError(t, err, "failed reading json")

	assert.Equal(t, src, copy, "orig and copy not equal")
}
