package brokennode

import "math/bits"

const (
	Broken  uint64 = 1
	Working uint64 = 0
)

func FindBrokenNodes(brokenNodes int, reports []bool) string {
	var config, max uint64
	result := make([]byte, len(reports))

	config = 1<<uint8(brokenNodes) - 1
	max = 1 << uint8(len(reports))

	return getResult(config, max, reports, result)
}

func getResult(config uint64, max uint64, reports []bool, result []byte) string {
	var fForExit bool
	for {
		if config, fForExit = checkThisConfigByEnum(config, max, reports); fForExit {
			return string(result)
		}
		merge(result, config)

		if config = nextPerm(config); config >= max {
			return string(result)
		}
	}
}

func checkThisConfigByEnum(config uint64, max uint64, reports []bool) (configOut uint64, fForExit bool) {
	for !checkConf(config, reports) {
		if config = nextPerm(config); config >= max {
			return config, true
		}
	}
	return config, false
}

func checkConf(conf uint64, reports []bool) bool {
	c := duplicateFirst(conf, len(reports))

	for i := range reports {
		var want uint64

		if c&1 == Broken {
			c >>= 1
			continue
		}
		if reports[i] {
			want = Working
		} else {
			want = Broken
		}
		c >>= 1
		if c&1 != want {
			return false
		}
	}
	return true
}

func merge(result []byte, conf uint64) {
	for i := range result {
		if result[i] == 0 {
			result[i] = returnChar(conf)
		}
		if result[i] != returnChar(conf) {
			result[i] = '?'
		}
		conf >>= 1
	}
}

func returnChar(conf uint64) byte {
	if conf&1 == Broken {
		return 'B'
	}
	return 'W'
}

func nextPerm(v uint64) uint64 {
	t := v | (v - 1)
	return (t + 1) | ((^t&-^t - 1) >> (uint8(bits.TrailingZeros64(v)) + 1))
}

func duplicateFirst(conf uint64, len int) uint64 {
	return conf | ((conf & 1) << uint8(len))
}
