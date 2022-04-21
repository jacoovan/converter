package converter

type Converter interface {
	Best(originNum float64, originSymbol string) (bestNum float64, bestSymbol string)
	To(originNum float64, originSymbol, destSymbol string) (destNum float64, err error)
}

var (
	UNIT_STORAGE Converter = NewUnitConverter(unit_storage)
	UNIT_HZ      Converter = NewUnitConverter(unit_hz)
)
