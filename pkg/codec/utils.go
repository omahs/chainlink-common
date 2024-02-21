package codec

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

// FitsInNBitsSigned returns if the [*math/big.Int] can fit in n bits as a signed integer.
// Namely, if it's in the range [-2^(n-1), 2^(n-1) - 1]
func FitsInNBitsSigned(n int, bi *big.Int) bool {
	if bi.Sign() < 0 {
		bi = new(big.Int).Neg(bi)
		bi.Sub(bi, big.NewInt(1))
	}
	return bi.BitLen() <= n-1
}

// BigIntHook is a mapstructure hook that converts number types to *[math/big.Int] and vice versa.
// Float values are cast to an int64 before being converted to a *[math/big.Int].
func BigIntHook(_, to reflect.Type, data any) (any, error) {
	if to == reflect.TypeOf(&big.Int{}) {
		bigInt := big.NewInt(0)

		switch v := data.(type) {
		case float64:
			bigInt.SetInt64(int64(v))
		case float32:
			bigInt.SetInt64(int64(v))
		case int:
			bigInt.SetInt64(int64(v))
		case int8:
			bigInt.SetInt64(int64(v))
		case int16:
			bigInt.SetInt64(int64(v))
		case int32:
			bigInt.SetInt64(int64(v))
		case int64:
			bigInt.SetInt64(v)
		case uint:
			bigInt.SetUint64(uint64(v))
		case uint8:
			bigInt.SetUint64(uint64(v))
		case uint16:
			bigInt.SetUint64(uint64(v))
		case uint32:
			bigInt.SetUint64(uint64(v))
		case uint64:
			bigInt.SetUint64(v)
		case string:
			_, ok := bigInt.SetString(v, 10)
			if !ok {
				return nil, fmt.Errorf("%w: cannot decode %s as big int", types.ErrInvalidType, v)
			}
		default:
			return data, nil
		}

		return bigInt, nil
	}

	if bi, ok := data.(*big.Int); ok {
		switch to {
		case reflect.TypeOf(0):
			if FitsInNBitsSigned(strconv.IntSize, bi) {
				return int(bi.Int64()), nil
			}
			return nil, fmt.Errorf("%w: can not %s fit into int", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(int8(0)):
			if FitsInNBitsSigned(8, bi) {
				return int8(bi.Int64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into int8", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(int16(0)):
			if FitsInNBitsSigned(16, bi) {
				return int16(bi.Int64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into int16", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(int32(0)):
			if FitsInNBitsSigned(32, bi) {
				return int32(bi.Int64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into int32 ", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(int64(0)):
			if FitsInNBitsSigned(64, bi) {
				return bi.Int64(), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into int64 ", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(uint(0)):
			if bi.Sign() >= 0 && bi.BitLen() <= strconv.IntSize {
				return uint(bi.Uint64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into uint", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(uint8(0)):
			if bi.Sign() >= 0 && bi.BitLen() <= 8 {
				return uint8(bi.Uint64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into uint8", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(uint16(0)):
			if bi.Sign() >= 0 && bi.BitLen() <= 16 {
				return uint16(bi.Uint64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into uint16", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(uint32(0)):
			if bi.Sign() >= 0 && bi.BitLen() <= 32 {
				return uint32(bi.Uint64()), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into uint32", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(uint64(0)):
			if bi.Sign() >= 0 && bi.BitLen() <= 64 {
				return bi.Uint64(), nil
			}
			return nil, fmt.Errorf("%w: cannot fit %s into uint64", types.ErrInvalidType, bi.String())
		case reflect.TypeOf(""):
			return bi.String(), nil
		default:
			return data, nil
		}
	}

	return data, nil
}

// SliceToArrayVerifySizeHook is a mapstructure hook that verifies if a slice is being assigned to an array,
// it will have the same number of elements. The default in mapstructure is to allow the slice to be smaller
// and will zero out the remaining elements in that case.
func SliceToArrayVerifySizeHook(from reflect.Type, to reflect.Type, data any) (any, error) {
	// By default, if the array is bigger it'll still work. (ie []int{1, 2, 3} => [4]int{} works with 0 at the end
	// [2]int{} would not work. This seems to lenient, but can be discussed.
	if from.Kind() == reflect.Slice && to.Kind() == reflect.Array {
		slice := reflect.ValueOf(data)

		// The use case here is that values may be added later (eg hard-coded mod)
		// Additionally, if you want to zero out in the plugin may not know the size
		// This allows the array to be zeroed out when the slice is empty to account for these use cases.
		if slice.Len() == 0 {
			return reflect.MakeSlice(slice.Type(), to.Len(), to.Len()).Interface(), nil
		}

		if slice.Len() != to.Len() {
			return nil, fmt.Errorf("%w: expected size %v got %v", types.ErrSliceWrongLen, to.Len(), slice.Len())
		}
	}

	return data, nil
}

var i64Type = reflect.TypeOf(int64(0))
var timeType = reflect.TypeOf(time.Time{})
var timePtrType = reflect.PointerTo(timeType)
var biType = reflect.TypeOf(&big.Int{})

// epochToTimeHook is a mapstructure hook that converts a unix epoch to a [time.Time] and vice versa.
// To do this, [time.Unix] and [time.Time.Unix] are used.
func epochToTimeHook(from reflect.Type, to reflect.Type, data any) (any, error) {
	if to == timeType {
		return convertToTime(from, data), nil
	} else if to == timePtrType {
		if t, ok := convertToTime(from, data).(time.Time); ok {
			return &t, nil
		}
	} else if tData, ok := data.(time.Time); ok {
		return convertToEpoch(to, tData), nil
	} else if tData, ok := data.(*time.Time); ok {
		return convertToEpoch(to, *tData), nil
	}

	return data, nil
}

func convertToTime(from reflect.Type, data any) any {
	if from.ConvertibleTo(i64Type) {
		return time.Unix(reflect.ValueOf(data).Convert(i64Type).Int(), 0).UTC()
	} else if from.ConvertibleTo(biType) {
		return time.Unix(reflect.ValueOf(data).Convert(biType).Interface().(*big.Int).Int64(), 0).UTC()
	} else if from.Kind() == reflect.Pointer {
		rData := reflect.ValueOf(data)
		if !rData.IsNil() {
			return convertToTime(from.Elem(), rData.Elem().Interface())
		}
	}

	return data
}

func convertToEpoch(to reflect.Type, data time.Time) any {
	if to.ConvertibleTo(i64Type) {
		return reflect.ValueOf(data.Unix()).Convert(to).Interface()
	} else if to.ConvertibleTo(biType) {
		return reflect.ValueOf(big.NewInt(data.Unix())).Convert(to).Interface()
	}

	return data
}

// getMapsFromPath takes a valueMap that represents a struct in a map.
// The key is the field name and the value is either the raw value of the field or a map[string]any representing a nested struct.
// The path is the fields to navigate to.  If any field in the path is a slice or array, multiple maps will be returned.
// It is expected that the final field represents a struct, or a slice/array/pointer to one.
// Example:
// valueMap = map[string]any{"A": map[string]any{"B" : []Foo{{C : 10, D : 100}, {C : 20, D : 200}}}}
// path = []string{"A", "B"}
// returns []map[string]any{{"C" : 10, "D" : 100}, {"C" : 20, "D" : 200}}, nil
func getMapsFromPath(valueMap map[string]any, path []string) ([]map[string]any, error) {
	extractMaps := []map[string]any{valueMap}
	for n, p := range path {
		tmp := make([]map[string]any, 0, len(extractMaps))
		for _, extractMap := range extractMaps {
			item, ok := extractMap[p]
			if !ok {
				return nil, fmt.Errorf("%w: cannot find %s", types.ErrInvalidType, strings.Join(path[:n+1], "."))
			}

			iItem := reflect.ValueOf(item)
			switch iItem.Kind() {
			case reflect.Array, reflect.Slice:
				length := iItem.Len()
				maps := make([]map[string]any, length)
				for i := 0; i < length; i++ {
					if err := mapstructure.Decode(iItem.Index(i).Interface(), &maps[i]); err != nil {
						return nil, fmt.Errorf("%w: %w", types.ErrInvalidType, err)
					}
				}
				extractMap[p] = maps
				tmp = append(tmp, maps...)
			default:
				var m map[string]any
				if err := mapstructure.Decode(item, &m); err != nil {
					return nil, fmt.Errorf("%w: %w", types.ErrInvalidType, err)
				}
				extractMap[p] = m
				tmp = append(tmp, m)
			}
		}
		extractMaps = tmp
	}
	return extractMaps, nil
}

func addr(value reflect.Value) reflect.Value {
	if value.CanAddr() {
		return value.Addr()
	}
	tmp := reflect.New(value.Type())
	tmp.Elem().Set(value)
	return tmp
}
