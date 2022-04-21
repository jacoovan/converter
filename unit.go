package converter

import (
	"math"
	"strings"
)

type Unit struct {
	BaseNum float64
	Symbol  string
}

type UnitConvert struct {
	sortedUnits []Unit
}

// assume param `sortedUnits` is sorted
// and order is low to high
func NewUnitConverter(sortedUnits []Unit) Converter {
	c := &UnitConvert{
		sortedUnits: sortedUnits,
	}
	return c
}

func (c *UnitConvert) Best(originNum float64, originSymbol string) (bestNum float64, bestSymbol string) {
	bestNum, bestSymbol = originNum, originSymbol
	var pos int = -1
	for i, v := range c.sortedUnits {
		if strings.EqualFold(v.Symbol, originSymbol) {
			pos = i
			bestSymbol = v.Symbol
			break
		}
	}
	if pos == -1 {
		return bestNum, bestSymbol
	}

	for i := pos; i < len(c.sortedUnits)-1; i++ {
		nextMulti := c.sortedUnits[i+1].BaseNum / c.sortedUnits[i].BaseNum
		if bestNum < nextMulti {
			break
		}
		bestNum = bestNum / nextMulti
		bestSymbol = c.sortedUnits[i+1].Symbol
	}
	return math.Floor(bestNum*1000) / 1000, bestSymbol
}

func (c *UnitConvert) To(originNum float64, originSymbol, destSymbol string) (destNum float64, err error) {
	if strings.EqualFold(originSymbol, destSymbol) {
		return originNum, nil
	}

	var destPos int = -1
	for i, v := range c.sortedUnits {
		if strings.EqualFold(v.Symbol, destSymbol) {
			destPos = i
			break
		}
	}
	if destPos == -1 {
		return 0, ErrKnownDestUnit
	}

	var originPos int = -1
	for i, v := range c.sortedUnits {
		if strings.EqualFold(v.Symbol, originSymbol) {
			originPos = i
			break
		}
	}
	if originPos == -1 {
		return 0, ErrKnownOriginUnit
	}

	destNum = originNum
	if originPos > destPos {
		for i := originPos; i > destPos; i-- {
			nextMulti := c.sortedUnits[i].BaseNum / c.sortedUnits[i-1].BaseNum
			destNum = destNum * nextMulti
		}
	} else {
		for i := originPos; i < destPos; i++ {
			nextMulti := c.sortedUnits[i+1].BaseNum / c.sortedUnits[i].BaseNum
			destNum = destNum / nextMulti
		}
	}
	return math.Floor(destNum*1000) / 1000, nil
}
