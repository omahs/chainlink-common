package encodings_test

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/smartcontractkit/libocr/bigbigendian"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"
	"github.com/smartcontractkit/chainlink-common/pkg/codec/encodings"
	"github.com/smartcontractkit/chainlink-common/pkg/codec/encodings/bigendian"
	encodingtestutils "github.com/smartcontractkit/chainlink-common/pkg/codec/encodings/testutils"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/testutils"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	. "github.com/smartcontractkit/chainlink-common/pkg/types/interfacetests"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

// TestCodecFromTypeCodecs tests all the codecs in packages under pkg/codec/encodings
// Implementations outside of this repository should run the interface tests themselves.
// The ones in this package are tested here for code coverage to count.
func TestCodecFromTypeCodecs(t *testing.T) {
	t.Parallel()
	biit := &bigEndianInterfaceTester{}
	RunCodecWithStrictArgsInterfaceTest(t, biit)
	RunCodecWithStrictArgsInterfaceTest(t, testutils.WrapCodecTesterForLoop(biit))
	t.Run("GetMaxEncodingSize delegates to Size", func(t *testing.T) {
		testCodec := &encodingtestutils.TestTypeCodec{
			Value: []int{55, 11},
			Bytes: []byte{1, 2},
		}

		c := encodings.CodecFromTypeCodec{"test": testCodec}

		actual, err := c.GetMaxEncodingSize(tests.Context(t), 50, "test")
		require.NoError(t, err)

		expected, err := testCodec.Size(50)
		require.NoError(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("GetMaxEncodingSize delegates to SizeTopLevel if present", func(t *testing.T) {
		testCodec := &encodingtestutils.TestTypeCodec{
			Value: []int{55, 11},
			Bytes: []byte{1, 2},
		}

		structCodec, err := encodings.NewStructCodec([]encodings.NamedTypeCodec{
			{Name: "Test", Codec: testCodec},
			{Name: "Test2", Codec: testCodec},
		})

		c := encodings.CodecFromTypeCodec{"test": structCodec}

		actual, err := c.GetMaxEncodingSize(tests.Context(t), 50, "test")
		require.NoError(t, err)

		singleItemSize, err := testCodec.Size(50)
		require.NoError(t, err)

		assert.Equal(t, singleItemSize*2, actual)
	})

	t.Run("GetMaxDecodingSize delegates to Size", func(t *testing.T) {
		testCodec := &encodingtestutils.TestTypeCodec{
			Value: []int{55, 11},
			Bytes: []byte{1, 2},
		}

		c := encodings.CodecFromTypeCodec{"test": testCodec}

		actual, err := c.GetMaxDecodingSize(tests.Context(t), 50, "test")
		require.NoError(t, err)

		expected, err := testCodec.Size(50)
		require.NoError(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("GetMaxDecodingSize delegates to SizeTopLevel if present", func(t *testing.T) {
		testCodec := &encodingtestutils.TestTypeCodec{
			Value: []int{55, 11},
			Bytes: []byte{1, 2},
		}

		structCodec, err := encodings.NewStructCodec([]encodings.NamedTypeCodec{
			{Name: "Test", Codec: testCodec},
			{Name: "Test2", Codec: testCodec},
		})

		c := encodings.CodecFromTypeCodec{"test": structCodec}

		actual, err := c.GetMaxDecodingSize(tests.Context(t), 50, "test")
		require.NoError(t, err)

		singleItemSize, err := testCodec.Size(50)
		require.NoError(t, err)

		assert.Equal(t, singleItemSize*2, actual)
	})
}

type interfaceTesterBase struct{}

func (*interfaceTesterBase) Setup(_ *testing.T) {}

func (*interfaceTesterBase) GetAccountBytes(i int) []byte {
	ib := byte(i)
	return []byte{ib, ib + 1, ib + 2, ib + 3, ib + 4, ib + 5, ib + 6, ib + 7}
}

type bigEndianInterfaceTester struct {
	interfaceTesterBase
}

func (b *bigEndianInterfaceTester) Name() string {
	return "big endian"
}

func (b *bigEndianInterfaceTester) EncodeFields(t *testing.T, request *EncodeRequest) []byte {
	switch request.TestOn {
	case TestItemType, TestItemArray1Type:
		return b.encode(t, nil, &request.TestStructs[0], request)
	case TestItemSliceType:
		bytes := []byte{byte(len(request.TestStructs))}
		for _, ts := range request.TestStructs {
			bytes = b.encode(t, bytes, &ts, request)
		}
		return bytes
	case TestItemArray2Type:
		return b.encode(t, b.encode(t, nil, &request.TestStructs[0], request), &request.TestStructs[1], request)
	}

	require.Fail(t, "unknown test type")
	return nil
}

func (b *bigEndianInterfaceTester) encode(t *testing.T, bytes []byte, ts *TestStruct, request *EncodeRequest) []byte {
	bytes = binary.BigEndian.AppendUint32(bytes, uint32(*ts.Field))
	bytes = append(bytes, byte(len(ts.DifferentField)))
	bytes = append(bytes, []byte(ts.DifferentField)...)
	bytes = append(bytes, byte(ts.OracleID))
	for _, oid := range ts.OracleIDs {
		bytes = append(bytes, byte(oid))
	}
	bytes = append(bytes, byte(len(ts.Account)))
	bytes = append(bytes, ts.Account...)
	bytes = append(bytes, byte(len(ts.Accounts)))
	for _, account := range ts.Accounts {
		bytes = append(bytes, byte(len(account)))
		bytes = append(bytes, account...)
	}
	bs, err := bigbigendian.SerializeSigned(4, ts.BigField)
	require.NoError(t, err)
	bytes = append(bytes, bs...)
	bytes = append(bytes, ts.NestedStruct.FixedBytes[:]...)
	bytes = binary.BigEndian.AppendUint64(bytes, uint64(ts.NestedStruct.Inner.I))
	bytes = append(bytes, byte(len(ts.NestedStruct.Inner.S)))
	bytes = append(bytes, []byte(ts.NestedStruct.Inner.S)...)
	if request.ExtraField {
		bytes = append(bytes, 5)
	}

	if request.MissingField {
		bytes = bytes[:len(bytes)-1]
	}

	return bytes
}

func newTestStructCodec(t *testing.T) encodings.TypeCodec {
	sCodec, err := bigendian.NewString(127)
	innerTestStruct, err := encodings.NewStructCodec([]encodings.NamedTypeCodec{
		{Name: "I", Codec: bigendian.Uint64{}},
		{Name: "S", Codec: sCodec},
	})
	require.NoError(t, err)
	arr2, err := encodings.NewArray(2, bigendian.Uint8{})
	require.NoError(t, err)

	midCodec, err := encodings.NewStructCodec([]encodings.NamedTypeCodec{
		{Name: "FixedBytes", Codec: arr2},
		{Name: "Inner", Codec: innerTestStruct},
	})
	require.NoError(t, err)

	oIds, err := encodings.NewArray(32, bigendian.OracleID{})
	require.NoError(t, err)

	size, err := bigendian.GetIntTypeCodecByByteSize(1)
	require.NoError(t, err)
	acc, err := encodings.NewSlice(bigendian.Uint8{}, size)
	require.NoError(t, err)

	accs, err := encodings.NewSlice(acc, size)
	require.NoError(t, err)

	bi, err := bigendian.NewBigInt(4, true)
	require.NoError(t, err)

	ts, err := encodings.NewStructCodec([]encodings.NamedTypeCodec{
		{Name: "Field", Codec: &encodings.NotNilPointer{Elm: bigendian.Uint32{}}},
		{Name: "DifferentField", Codec: sCodec},
		{Name: "OracleID", Codec: bigendian.OracleID{}},
		{Name: "OracleIDs", Codec: oIds},
		{Name: "Account", Codec: acc},
		{Name: "Accounts", Codec: accs},
		{Name: "BigField", Codec: bi},
		{Name: "NestedStruct", Codec: midCodec},
	})
	require.NoError(t, err)
	return ts
}

func (b *bigEndianInterfaceTester) GetCodec(t *testing.T) types.Codec {
	testStruct := newTestStructCodec(t)
	size, err := bigendian.GetIntTypeCodecByByteSize(1)
	require.NoError(t, err)
	slice, err := encodings.NewSlice(testStruct, size)
	require.NoError(t, err)
	arr1, err := encodings.NewArray(1, testStruct)
	require.NoError(t, err)
	arr2, err := encodings.NewArray(2, testStruct)
	require.NoError(t, err)

	ts := CreateTestStruct(0, b)

	c := &encodings.CodecFromTypeCodec{
		TestItemType:            testStruct,
		TestItemSliceType:       slice,
		TestItemArray1Type:      arr1,
		TestItemArray2Type:      arr2,
		TestItemWithConfigExtra: testStruct,
		NilType:                 encodings.Empty{},
	}
	require.NoError(t, err)

	mod, err := codec.NewHardCoder(map[string]any{
		"BigField": ts.BigField.String(),
		"Account":  ts.Account,
	}, map[string]any{"ExtraField": AnyExtraValue}, codec.BigIntHook)
	require.NoError(t, err)

	byTypeMod, err := codec.NewByItemTypeModifier(map[string]codec.Modifier{
		TestItemType:            codec.MultiModifier{},
		TestItemSliceType:       codec.MultiModifier{},
		TestItemArray1Type:      codec.MultiModifier{},
		TestItemArray2Type:      codec.MultiModifier{},
		TestItemWithConfigExtra: mod,
		NilType:                 codec.MultiModifier{},
	})
	require.NoError(t, err)

	modCodec, err := codec.NewModifierCodec(c, byTypeMod, codec.BigIntHook)
	require.NoError(t, err)

	_, err = mod.RetypeToOffChain(reflect.PointerTo(testStruct.GetType()), TestItemWithConfigExtra)
	require.NoError(t, err)

	return modCodec
}

func (b *bigEndianInterfaceTester) IncludeArrayEncodingSizeEnforcement() bool {
	return true
}
