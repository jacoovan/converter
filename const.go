package converter

import "errors"

var (
	ErrKnownOriginUnit = errors.New("UNKNOWN(ORIGIN_SYMBOL)")
	ErrKnownDestUnit   = errors.New("UNKNOWN(DEST_SYMBOL)")
)

var (
	unit_storage = []Unit{
		{
			BaseNum: 1,
			Symbol:  `Byte`,
		},
		{
			BaseNum: 1024,
			Symbol:  `KB`,
		},
		{
			BaseNum: 1024 * 1024,
			Symbol:  `MB`,
		},
		{
			BaseNum: 1024 * 1024 * 1024,
			Symbol:  `GB`,
		},
		{
			BaseNum: 1024 * 1024 * 1024 * 1024,
			Symbol:  `TB`,
		},
	}

	unit_hz = []Unit{
		{
			BaseNum: 1,
			Symbol:  `hz`,
		},
		{
			BaseNum: 1000,
			Symbol:  `Khz`,
		},
		{
			BaseNum: 1000 * 1000,
			Symbol:  `Mhz`,
		},
		{
			BaseNum: 1000 * 1000 * 1000,
			Symbol:  `Ghz`,
		},
	}
)
