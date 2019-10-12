package gomock

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Seed random. Setting seed to 0 will use time.Now().UnixNano()
func Seed(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(seed)
	}
}

func Bool() bool {
	return rand.Intn(1) == 1
}

// Int64 will generate a random int64 value
func Int64() int64 {
	return rand.Int63n(math.MaxInt64) + math.MinInt64
}

// Generate random integer between min and max
func randIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

func randFloat32Range(min, max float32) float32 {
	if min == max {
		return min
	}
	return rand.Float32()*(max-min) + min
}

/*func Number(min int, max int) int {
	return randIntRange(min, max)
}*/

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return randIntRange(min, max)
}

//生成占位的数字
func NumberS(min int, max int) string {
	num := randIntRange(min, max)
	str := strconv.Itoa(num)
	length := len(strconv.Itoa(max))
	if len(str) < length {
		str = strings.Repeat("0", length-len(str)) + str
	}
	return str
}
