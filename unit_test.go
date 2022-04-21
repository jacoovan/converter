package converter_test

import (
	"testing"

	"github.com/jacoovan/converter"
	"github.com/stretchr/testify/assert"
)

func Test_Unit_Converter_Best(t *testing.T) {
	var originNum float64
	var originSymbol string

	var bestNum float64
	var bestSymbol string

	originNum = 1
	originSymbol = `byte`
	bestNum, bestSymbol = converter.UNIT_STORAGE.Best(originNum, originSymbol)
	assert.Equal(t, 1.0, bestNum)
	assert.Equal(t, `Byte`, bestSymbol)

	originNum = 1024 - 1
	originSymbol = `byte`
	bestNum, bestSymbol = converter.UNIT_STORAGE.Best(originNum, originSymbol)
	assert.Equal(t, 1023.0, bestNum)
	assert.Equal(t, `Byte`, bestSymbol)

	originNum = 1024
	originSymbol = `byte`
	bestNum, bestSymbol = converter.UNIT_STORAGE.Best(originNum, originSymbol)
	assert.Equal(t, 1.0, bestNum)
	assert.Equal(t, `KB`, bestSymbol)
}

func Test_Unit_Converter_To(t *testing.T) {
	var originNum float64
	var originSymbol string

	var destNum float64
	var destSymbol string
	var err error

	originNum = 1
	originSymbol = `byte`
	destSymbol = `byte`
	destNum, err = converter.UNIT_STORAGE.To(originNum, originSymbol, destSymbol)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1.0, destNum)

	originNum = 1
	originSymbol = `byte`
	destSymbol = `KB`
	destNum, err = converter.UNIT_STORAGE.To(originNum, originSymbol, destSymbol)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0.0, destNum)

	originNum = 2
	originSymbol = `byte`
	destSymbol = `KB`
	destNum, err = converter.UNIT_STORAGE.To(originNum, originSymbol, destSymbol)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0.001, destNum)

	originNum = 1023
	originSymbol = `byte`
	destSymbol = `KB`
	destNum, err = converter.UNIT_STORAGE.To(originNum, originSymbol, destSymbol)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0.999, destNum)

	originNum = 1
	originSymbol = `KB`
	destSymbol = `Byte`
	destNum, err = converter.UNIT_STORAGE.To(originNum, originSymbol, destSymbol)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1024.0, destNum)
}
