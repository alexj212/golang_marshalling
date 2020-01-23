package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteAndReadGob(t *testing.T) {
	filename := "/tmp/person.gob"
	src := Homer
	copy := new(Person)

	var err error
	err = writeGob(filename, src)
	assert.NoError(t, err, "failed writing gob")

	err = readGob(filename, copy)
	assert.NoError(t, err, "failed reading gob")

	assert.Equal(t, src, copy, "orig and copy not equal")
}

func TestWriteGob(t *testing.T) {

	var err error

	err = writeGob("/tmp/person.gob", Homer)
	assert.NoError(t, err, "failed writing gob")
}

func TestWriteAndReadComplexStruct(t *testing.T) {
	filename := "/tmp/family.gob"
	src := Family
	copy := new(People)

	var err error
	err = writeGob(filename, src)
	assert.NoError(t, err, "failed writing gob")

	err = readGob(filename, copy)
	assert.NoError(t, err, "failed reading gob")

	assert.Equal(t, src, copy, "orig and copy not equal")
}
