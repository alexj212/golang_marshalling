package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteAndReadXml(t *testing.T) {
	var err error

	err = writeXml("/tmp/person.xml", Homer)
	assert.NoError(t, err, "failed writing xml")

	copy := &Person{}
	err = readXml("/tmp/person.xml", copy)
	assert.NoError(t, err, "failed reading xml")

	assert.Equal(t, Homer, copy, "orig and copy not equal")
}

func TestWriteXml(t *testing.T) {
	var err error

	err = writeXml("/tmp/person.xml", Homer)
	assert.NoError(t, err, "failed writing xml")
}

func TestWriteAndReadComplexStruct(t *testing.T) {
	filename := "/tmp/family.xml"
	src := Family
	copy := new(People)

	var err error

	err = writeXml(filename, src)
	assert.NoError(t, err, "failed writing xml")

	err = readXml(filename, copy)
	assert.NoError(t, err, "failed reading xml")

	assert.Equal(t, src, copy, "orig and copy not equal")
}
