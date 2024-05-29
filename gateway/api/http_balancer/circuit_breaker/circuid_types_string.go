// Code generated by "stringer -type=CircuitStatus -output=./circuid_types_string.go"; DO NOT EDIT.

package circuit_breaker

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Closed-0]
	_ = x[Opened-1]
	_ = x[HalfOpened-2]
}

const _CircuitStatus_name = "ClosedOpenedHalfOpened"

var _CircuitStatus_index = [...]uint8{0, 6, 12, 22}

func (i CircuitStatus) String() string {
	if i < 0 || i >= CircuitStatus(len(_CircuitStatus_index)-1) {
		return "CircuitStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CircuitStatus_name[_CircuitStatus_index[i]:_CircuitStatus_index[i+1]]
}
