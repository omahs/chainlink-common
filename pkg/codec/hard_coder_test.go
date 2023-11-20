package codec_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

var onChainType = reflect.TypeOf(hardCodedTestStruct{})

func TestHardCoder(t *testing.T) {
	hardCoder, err := codec.NewHardCoder(map[string]any{"A": "Foo", "C": []int32{2, 3}}, map[string]any{"Z": "Bar", "Q": []struct {
		A int
		B string
	}{{1, "a"}, {2, "b"}}})
	require.NoError(t, err)
	replacingHardCoder, err := codec.NewHardCoder(map[string]any{"A": "two"}, map[string]any{"A": int64(2), "Q": []int32{4, 5}})
	require.NoError(t, err)

	t.Run("NewHardCoder returns error if key and subkey are in a map", func(t *testing.T) {
		_, err := codec.NewHardCoder(map[string]any{"A.Z": "Foo", "A": hardCodedTestStruct{A: "Z"}}, map[string]any{})
		assert.True(t, errors.Is(err, types.ErrInvalidConfig))

		_, err = codec.NewHardCoder(map[string]any{}, map[string]any{"A.Z": "Foo", "A": hardCodedTestStruct{A: "Z"}})
		assert.True(t, errors.Is(err, types.ErrInvalidConfig))
	})

	t.Run("RetypeForOffChain adds fields to struct", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(onChainType)
		require.NoError(t, err)
		assertBasicHardCodedType(t, offChainType)
	})

	t.Run("RetypeForOffChain adds fields to pointers", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.PointerTo(onChainType))
		require.NoError(t, err)
		assert.Equal(t, reflect.Ptr, offChainType.Kind())
		assertBasicHardCodedType(t, offChainType.Elem())
	})

	t.Run("RetypeForOffChain adds fields to pointers of non-structs", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.PointerTo(reflect.SliceOf(onChainType)))
		require.NoError(t, err)
		assert.Equal(t, reflect.Pointer, offChainType.Kind())
		assert.Equal(t, reflect.Slice, offChainType.Elem().Kind())
		assertBasicHardCodedType(t, offChainType.Elem().Elem())
	})

	t.Run("RetypeForOffChain adds fields to slices", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.SliceOf(onChainType))
		require.NoError(t, err)
		assert.Equal(t, reflect.Slice, offChainType.Kind())
		assertBasicHardCodedType(t, offChainType.Elem())
	})

	t.Run("RetypeForOffChain adds fields to arrays", func(t *testing.T) {
		anyArrayLen := 3
		offChainType, err := hardCoder.RetypeForOffChain(reflect.ArrayOf(anyArrayLen, onChainType))
		require.NoError(t, err)
		assert.Equal(t, reflect.Array, offChainType.Kind())
		assert.Equal(t, anyArrayLen, offChainType.Len())
		assertBasicHardCodedType(t, offChainType.Elem())
	})

	t.Run("RetypeForOffChain replaces already existing field", func(t *testing.T) {
		offChainType, err := replacingHardCoder.RetypeForOffChain(onChainType)
		require.NoError(t, err)
		require.Equal(t, onChainType.NumField()+1, offChainType.NumField())
		for i := 0; i < onChainType.NumField(); i++ {
			if onChainType.Field(i).Name == "A" {
				continue
			}
			require.Equal(t, cleanStructField(onChainType.Field(i)), cleanStructField(offChainType.Field(i)))
		}

		a, ok := offChainType.FieldByName("A")
		require.True(t, ok)
		assert.Equal(t, reflect.TypeOf(int64(0)), a.Type)

		extra := offChainType.Field(onChainType.NumField())
		assert.Equal(t, reflect.StructField{Name: "Q", Type: reflect.TypeOf([]int32{})}, cleanStructField(extra))
	})

	t.Run("RetypeForOffChain returns error is existing field type is changed and not hard coded both ways", func(t *testing.T) {
		invalidHardCoder, err := codec.NewHardCoder(map[string]any{}, map[string]any{"A": int64(2), "Q": []int32{4, 5}})
		require.NoError(t, err)
		_, err = invalidHardCoder.RetypeForOffChain(onChainType)
		assert.True(t, errors.Is(err, types.ErrInvalidType))
	})

	t.Run("TransformForOnChain and TransformForOffChain works on structs", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(onChainType)
		require.NoError(t, err)

		iInput := reflect.Indirect(reflect.New(offChainType))
		iInput.FieldByName("B").SetInt(1)

		actual, err := hardCoder.TransformForOnChain(iInput.Interface())
		require.NoError(t, err)

		expected := hardCodedTestStruct{
			A: "Foo",
			B: 1,
			C: []int32{2, 3},
		}
		assert.Equal(t, expected, actual)

		actual, err = hardCoder.TransformForOffChain(expected)
		require.NoError(t, err)
		iInput.FieldByName("A").SetString("Foo")
		iInput.FieldByName("C").Set(reflect.ValueOf([]int32{2, 3}))
		iInput.FieldByName("Z").SetString("Bar")
		q := iInput.FieldByName("Q")
		q.Set(reflect.MakeSlice(q.Type(), 2, 2))
		elm := q.Index(0)
		elm.FieldByName("A").SetInt(1)
		elm.FieldByName("B").SetString("a")
		elm = q.Index(1)
		elm.FieldByName("A").SetInt(2)
		elm.FieldByName("B").SetString("b")
		assert.Equal(t, iInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain returns error if input type was not from TransformForOnChain", func(t *testing.T) {
		_, err := hardCoder.TransformForOnChain(hardCodedTestStruct{})
		assert.True(t, errors.Is(err, types.ErrInvalidType))
	})

	t.Run("TransformForOnChain and TransformForOffChain works on pointers", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.PointerTo(onChainType))
		require.NoError(t, err)

		rInput := reflect.New(offChainType.Elem())
		iInput := reflect.Indirect(rInput)
		iInput.FieldByName("B").SetInt(1)

		actual, err := hardCoder.TransformForOnChain(rInput.Interface())
		require.NoError(t, err)

		expected := &hardCodedTestStruct{
			A: "Foo",
			B: 1,
			C: []int32{2, 3},
		}
		assert.Equal(t, expected, actual)

		actual, err = hardCoder.TransformForOffChain(expected)
		require.NoError(t, err)
		addOffChainAndOnChainHardCodedValues(iInput)
		assert.Equal(t, rInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain works on slices", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.SliceOf(onChainType))
		require.NoError(t, err)

		rInput := reflect.MakeSlice(offChainType, 2, 2)
		iElm := rInput.Index(0)
		iElm.FieldByName("B").SetInt(1)
		iElm = rInput.Index(1)
		iElm.FieldByName("B").SetInt(2)

		actual, err := hardCoder.TransformForOnChain(rInput.Interface())
		require.NoError(t, err)

		expected := []hardCodedTestStruct{
			{
				A: "Foo",
				B: 1,
				C: []int32{2, 3},
			},
			{
				A: "Foo",
				B: 2,
				C: []int32{2, 3},
			},
		}
		assert.Equal(t, expected, actual)

		actual, err = hardCoder.TransformForOffChain(expected)
		require.NoError(t, err)

		addOffChainAndOnChainHardCodedValues(rInput.Index(0))
		addOffChainAndOnChainHardCodedValues(rInput.Index(1))
		assert.Equal(t, rInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain works on slices of slices", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.SliceOf(reflect.SliceOf(onChainType)))
		require.NoError(t, err)

		rInput := reflect.MakeSlice(offChainType, 2, 2)
		iOuter := rInput.Index(0)
		iOuter.Set(reflect.MakeSlice(iOuter.Type(), 2, 2))
		iElm := iOuter.Index(0)
		iElm.FieldByName("B").SetInt(1)
		iElm = iOuter.Index(1)
		iElm.FieldByName("B").SetInt(2)
		iOuter = rInput.Index(1)
		iOuter.Set(reflect.MakeSlice(iOuter.Type(), 2, 2))
		iElm = iOuter.Index(0)
		iElm.FieldByName("B").SetInt(10)
		iElm = iOuter.Index(1)
		iElm.FieldByName("B").SetInt(20)

		actual, err := hardCoder.TransformForOnChain(rInput.Interface())
		require.NoError(t, err)

		expected := [][]hardCodedTestStruct{
			{
				{
					A: "Foo",
					B: 1,
					C: []int32{2, 3},
				},
				{
					A: "Foo",
					B: 2,
					C: []int32{2, 3},
				},
			},
			{
				{
					A: "Foo",
					B: 10,
					C: []int32{2, 3},
				},
				{
					A: "Foo",
					B: 20,
					C: []int32{2, 3},
				},
			},
		}
		assert.Equal(t, expected, actual)

		actual, err = hardCoder.TransformForOffChain(expected)
		require.NoError(t, err)

		addOffChainAndOnChainHardCodedValues(rInput.Index(0).Index(0))
		addOffChainAndOnChainHardCodedValues(rInput.Index(0).Index(1))
		addOffChainAndOnChainHardCodedValues(rInput.Index(1).Index(0))
		addOffChainAndOnChainHardCodedValues(rInput.Index(1).Index(1))
		assert.Equal(t, rInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain works on pointers to non structs", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.PointerTo(reflect.SliceOf(onChainType)))
		require.NoError(t, err)

		rInput := reflect.New(offChainType.Elem())
		sInput := reflect.MakeSlice(offChainType.Elem(), 2, 2)
		iElm := sInput.Index(0)
		iElm.FieldByName("B").SetInt(1)
		iElm = sInput.Index(1)
		iElm.FieldByName("B").SetInt(2)
		reflect.Indirect(rInput).Set(sInput)

		actual, err := hardCoder.TransformForOnChain(rInput.Interface())
		require.NoError(t, err)

		expected := &[]hardCodedTestStruct{
			{
				A: "Foo",
				B: 1,
				C: []int32{2, 3},
			},
			{
				A: "Foo",
				B: 2,
				C: []int32{2, 3},
			},
		}
		assert.Equal(t, expected, actual)

		actual, err = hardCoder.TransformForOffChain(expected)
		require.NoError(t, err)

		addOffChainAndOnChainHardCodedValues(sInput.Index(0))
		addOffChainAndOnChainHardCodedValues(sInput.Index(1))
		assert.Equal(t, rInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain works on arrays", func(t *testing.T) {
		offChainType, err := hardCoder.RetypeForOffChain(reflect.ArrayOf(2, onChainType))
		require.NoError(t, err)

		rInput := reflect.New(offChainType).Elem()
		iElm := rInput.Index(0)
		iElm.FieldByName("B").SetInt(1)
		iElm = rInput.Index(1)
		iElm.FieldByName("B").SetInt(2)

		actual, err := hardCoder.TransformForOnChain(rInput.Interface())
		require.NoError(t, err)

		expected := [2]hardCodedTestStruct{
			{
				A: "Foo",
				B: 1,
				C: []int32{2, 3},
			},
			{
				A: "Foo",
				B: 2,
				C: []int32{2, 3},
			},
		}
		assert.Equal(t, expected, actual)

		actual, err = hardCoder.TransformForOffChain(expected)
		require.NoError(t, err)

		addOffChainAndOnChainHardCodedValues(rInput.Index(0))
		addOffChainAndOnChainHardCodedValues(rInput.Index(1))
		assert.Equal(t, rInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain works on nested fields", func(t *testing.T) {
		nestedHardCoder, err := codec.NewHardCoder(map[string]any{
			"A":   "Top",
			"B.A": "Foo",
			"B.C": []int32{2, 3},
			"C.A": "Foo",
			"C.C": []int32{2, 3},
		}, map[string]any{
			"B.Z": "Bar",
			"B.Q": []struct {
				A int
				B string
			}{{1, "a"}, {2, "b"}},
			"C.Z": "Bar",
			"C.Q": []struct {
				A int
				B string
			}{{1, "a"}, {2, "b"}},
		})
		require.NoError(t, err)

		offChainType, err := nestedHardCoder.RetypeForOffChain(reflect.TypeOf(nestedHardCodedTestStruct{}))
		require.NoError(t, err)

		iInput := reflect.Indirect(reflect.New(offChainType))
		iB := iInput.FieldByName("B")
		iB.FieldByName("B").SetInt(1)
		iC := iInput.FieldByName("C")
		iC.Set(reflect.MakeSlice(iC.Type(), 2, 2))
		iC.Index(0).FieldByName("B").SetInt(2)
		iC.Index(1).FieldByName("B").SetInt(3)
		iInput.FieldByName("D").SetInt(1)

		actual, err := nestedHardCoder.TransformForOnChain(iInput.Interface())
		require.NoError(t, err)

		expected := nestedHardCodedTestStruct{
			A: "Top",
			B: hardCodedTestStruct{
				A: "Foo",
				B: 1,
				C: []int32{2, 3},
			},
			C: []hardCodedTestStruct{
				{
					A: "Foo",
					B: 2,
					C: []int32{2, 3},
				},
				{
					A: "Foo",
					B: 3,
					C: []int32{2, 3},
				},
			},
			D: 1,
		}

		assert.Equal(t, expected, actual)

		actual, err = nestedHardCoder.TransformForOffChain(expected)
		require.NoError(t, err)

		iInput.FieldByName("A").SetString("Top")
		addOffChainAndOnChainHardCodedValues(iB)
		addOffChainAndOnChainHardCodedValues(iC.Index(0))
		addOffChainAndOnChainHardCodedValues(iC.Index(1))
		assert.Equal(t, iInput.Interface(), actual)
	})

	t.Run("TransformForOnChain and TransformForOffChain works for replaced type", func(t *testing.T) {
		offChainType, err := replacingHardCoder.RetypeForOffChain(onChainType)
		require.NoError(t, err)
		iInput := reflect.Indirect(reflect.New(offChainType))
		iInput.FieldByName("B").SetInt(1)

		actual, err := replacingHardCoder.TransformForOnChain(iInput.Interface())
		require.NoError(t, err)

		expected := hardCodedTestStruct{
			A: "two",
			B: 1,
		}

		assert.Equal(t, expected, actual)

		actual, err = replacingHardCoder.TransformForOffChain(expected)
		require.NoError(t, err)
		iInput.FieldByName("A").SetInt(2)
		iInput.FieldByName("Q").Set(reflect.ValueOf([]int32{4, 5}))

		assert.Equal(t, iInput.Interface(), actual)
	})

	t.Run("TransformForOnChain respect hooks", func(t *testing.T) {
		var hook mapstructure.DecodeHookFunc = func(from, to reflect.Kind, val interface{}) (any, error) {
			if to == reflect.Int32 {
				return int32(123), nil
			}
			return val, nil
		}
		hookedHardCoder, err := codec.NewHardCoder(map[string]any{"B": "Z"}, map[string]any{}, hook)
		require.NoError(t, err)

		offChainType, err := hookedHardCoder.RetypeForOffChain(onChainType)
		require.NoError(t, err)

		offChain := reflect.Indirect(reflect.New(offChainType)).Interface()
		onChain, err := hookedHardCoder.TransformForOnChain(offChain)
		require.NoError(t, err)

		assert.Equal(t, hardCodedTestStruct{B: 123}, onChain)
	})
}

// Since we're using the on-chain values that have their hard-coded values set to
// transform back to the off-chain values, the on-chain values will be set in the off-chain
// as well unless it's overwritten.
func addOffChainAndOnChainHardCodedValues(iInput reflect.Value) {
	iInput.FieldByName("A").SetString("Foo")
	iInput.FieldByName("C").Set(reflect.ValueOf([]int32{2, 3}))
	iInput.FieldByName("Z").SetString("Bar")
	q := iInput.FieldByName("Q")
	q.Set(reflect.MakeSlice(q.Type(), 2, 2))
	elm := q.Index(0)
	elm.FieldByName("A").SetInt(1)
	elm.FieldByName("B").SetString("a")
	elm = q.Index(1)
	elm.FieldByName("A").SetInt(2)
	elm.FieldByName("B").SetString("b")
}

func assertBasicHardCodedType(t *testing.T, offChainType reflect.Type) {
	require.Equal(t, onChainType.NumField()+2, offChainType.NumField())
	for i := 0; i < onChainType.NumField(); i++ {
		require.Equal(t, onChainType.Field(i), offChainType.Field(i))
	}

	fn1 := offChainType.Field(onChainType.NumField())
	fn2 := offChainType.Field(onChainType.NumField() + 1)
	var z, q *reflect.StructField
	switch fn1.Name {
	case "Z":
		z = &fn1
	case "Q":
		q = &fn1
	}
	switch fn2.Name {
	case "Z":
		z = &fn2
	case "Q":
		q = &fn2
	}
	require.NotNil(t, z)
	assert.Equal(t, reflect.TypeOf("string"), z.Type)
	require.NotNil(t, q)
	require.Equal(t, reflect.Slice, q.Type.Kind())
	qe := q.Type.Elem()
	require.Equal(t, reflect.Struct, qe.Kind())
	assert.Equal(t, 2, qe.NumField())
	a, ok := qe.FieldByName("A")
	require.True(t, ok)
	assert.Equal(t, reflect.TypeOf(0), a.Type)
	b, ok := qe.FieldByName("B")
	require.True(t, ok)
	assert.Equal(t, reflect.TypeOf("string"), b.Type)
}

type hardCodedTestStruct struct {
	A string
	B int32
	C []int32
}

func cleanStructField(field reflect.StructField) reflect.StructField {
	field.Index = nil
	field.Offset = uintptr(0)
	return field
}

type nestedHardCodedTestStruct struct {
	A string
	B hardCodedTestStruct
	C []hardCodedTestStruct
	D int32
}