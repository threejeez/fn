// Code generated by "stringer -type AggType"; DO NOT EDIT.

package view

import "strconv"

const _AggType_name = "AggTypeNoneAggTypeCountAggTypeSumAggTypeMeanAggTypeDistribution"

var _AggType_index = [...]uint8{0, 11, 23, 33, 44, 63}

func (i AggType) String() string {
	if i < 0 || i >= AggType(len(_AggType_index)-1) {
		return "AggType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AggType_name[_AggType_index[i]:_AggType_index[i+1]]
}
