package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteAndReadJson(t *testing.T) {
	filename := "/tmp/person.proto"
	src := Homer
	copy := new(Person)

	var err error
	err = writeProto(filename, src)
	assert.NoError(t, err, "failed writing proto")

	err = readProto(filename, copy)
	assert.NoError(t, err, "failed reading proto")

	assert.Equal(t, src, copy, "orig and copy not equal")
}

func TestWriteJson(t *testing.T) {
	var err error

	err = writeProto("/tmp/person.proto", Homer)
	assert.NoError(t, err, "failed writing proto")
}

func TestWriteAndReadComplexStruct(t *testing.T) {
	filename := "/tmp/family.proto"
	src := Family
	copy := new(People)

	var err error
	err = writeProto(filename, src)
	assert.NoError(t, err, "failed writing proto")

	err = readProto(filename, copy)
	assert.NoError(t, err, "failed reading proto")

	assert.Equal(t, src, copy, "orig and copy not equal")
}
