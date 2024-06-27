package chainreader_test

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/relayer/pluginprovider/chainreader"
	chainreadertest "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/relayer/pluginprovider/chainreader/test"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	. "github.com/smartcontractkit/chainlink-common/pkg/types/interfacetests" //nolint
)

func TestVersionedBytesFunctions(t *testing.T) {
	const unsupportedVer = 25913
	t.Run("chainreader.EncodeVersionedBytes unsupported type", func(t *testing.T) {
		invalidData := make(chan int)

		_, err := chainreader.EncodeVersionedBytes(invalidData, chainreader.JSONEncodingVersion2)

		assert.True(t, errors.Is(err, types.ErrInvalidType))
	})

	t.Run("chainreader.EncodeVersionedBytes unsupported encoding version", func(t *testing.T) {
		expected := fmt.Errorf("%w: unsupported encoding version %d for data map[key:value]", types.ErrInvalidEncoding, unsupportedVer)
		data := map[string]interface{}{
			"key": "value",
		}

		_, err := chainreader.EncodeVersionedBytes(data, unsupportedVer)
		if err == nil || err.Error() != expected.Error() {
			t.Errorf("expected error: %s, but got: %v", expected, err)
		}
	})

	t.Run("chainreader.DecodeVersionedBytes", func(t *testing.T) {
		var decodedData map[string]interface{}
		expected := fmt.Errorf("unsupported encoding version %d for versionedData [97 98 99 100 102]", unsupportedVer)
		versionedBytes := &pb.VersionedBytes{
			Version: unsupportedVer, // Unsupported version
			Data:    []byte("abcdf"),
		}

		err := chainreader.DecodeVersionedBytes(&decodedData, versionedBytes)
		if err == nil || err.Error() != expected.Error() {
			t.Errorf("expected error: %s, but got: %v", expected, err)
		}
	})
}

func TestChainReaderInterfaceTests(t *testing.T) {
	t.Parallel()

	chainreadertest.TestAllEncodings(t, func(version chainreader.EncodingVersion) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			fake := &fakeChainReader{}
			RunChainReaderInterfaceTests(
				t,
				chainreadertest.WrapChainReaderTesterForLoop(
					&fakeChainReaderInterfaceTester{impl: fake},
					chainreadertest.WithChainReaderLoopEncoding(version),
				),
			)
		}
	})
}

func TestBind(t *testing.T) {
	t.Parallel()

	chainreadertest.TestAllEncodings(t, func(version chainreader.EncodingVersion) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			es := &errChainReader{}
			errTester := chainreadertest.WrapChainReaderTesterForLoop(
				&fakeChainReaderInterfaceTester{impl: es},
				chainreadertest.WithChainReaderLoopEncoding(version),
			)

			errTester.Setup(t)
			chainReader := errTester.GetChainReader(t)

			for _, errorType := range errorTypes {
				es.err = errorType
				t.Run("Bind unwraps errors from server "+errorType.Error(), func(t *testing.T) {
					ctx := tests.Context(t)
					err := chainReader.Bind(ctx, []types.BoundContract{{Name: "Name", Address: "address"}})
					assert.True(t, errors.Is(err, errorType))
				})
			}
		}
	})
}

func TestGetLatestValue(t *testing.T) {
	t.Parallel()

	chainreadertest.TestAllEncodings(t, func(version chainreader.EncodingVersion) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			es := &errChainReader{}
			errTester := chainreadertest.WrapChainReaderTesterForLoop(
				&fakeChainReaderInterfaceTester{impl: es},
				chainreadertest.WithChainReaderLoopEncoding(version),
			)

			errTester.Setup(t)
			chainReader := errTester.GetChainReader(t)

			t.Run("nil reader should return unimplemented", func(t *testing.T) {
				t.Parallel()

				ctx := tests.Context(t)

				nilTester := chainreadertest.WrapChainReaderTesterForLoop(&fakeChainReaderInterfaceTester{impl: nil})
				nilTester.Setup(t)
				nilCr := nilTester.GetChainReader(t)

				err := nilCr.GetLatestValue(ctx, types.NewBoundContract("", "", "method"), "anything", "anything")
				assert.Equal(t, codes.Unimplemented, status.Convert(err).Code())
			})

			for _, errorType := range errorTypes {
				es.err = errorType
				t.Run("GetLatestValue unwraps errors from server "+errorType.Error(), func(t *testing.T) {
					ctx := tests.Context(t)
					err := chainReader.GetLatestValue(ctx, types.NewBoundContract("", "", "method"), nil, "anything")
					assert.True(t, errors.Is(err, errorType))
				})
			}

			// make sure that errors come from client directly
			es.err = nil
			t.Run("GetLatestValue returns error if type cannot be encoded in the wire format", func(t *testing.T) {
				ctx := tests.Context(t)
				err := chainReader.GetLatestValue(ctx, types.NewBoundContract("", "", "method"), &cannotEncode{}, &TestStruct{})
				assert.True(t, errors.Is(err, types.ErrInvalidType))
			})
		}
	})
}

func TestQueryKey(t *testing.T) {
	t.Parallel()

	chainreadertest.TestAllEncodings(t, func(version chainreader.EncodingVersion) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			impl := &protoConversionTestChainReader{}
			crTester := chainreadertest.WrapChainReaderTesterForLoop(&fakeChainReaderInterfaceTester{impl: impl}, chainreadertest.WithChainReaderLoopEncoding(version))
			crTester.Setup(t)
			cr := crTester.GetChainReader(t)

			es := &errChainReader{}
			errTester := chainreadertest.WrapChainReaderTesterForLoop(&fakeChainReaderInterfaceTester{impl: es})
			errTester.Setup(t)
			chainReader := errTester.GetChainReader(t)

			t.Run("nil reader should return unimplemented", func(t *testing.T) {
				ctx := tests.Context(t)

				nilTester := chainreadertest.WrapChainReaderTesterForLoop(&fakeChainReaderInterfaceTester{impl: nil})
				nilTester.Setup(t)
				nilCr := nilTester.GetChainReader(t)

				_, err := nilCr.QueryKey(ctx, types.NewBoundContract("", "", ""), query.KeyFilter{}, query.LimitAndSort{}, &[]interface{}{nil})
				assert.Equal(t, codes.Unimplemented, status.Convert(err).Code())
			})

			for _, errorType := range errorTypes {
				es.err = errorType
				t.Run("QueryKey unwraps errors from server "+errorType.Error(), func(t *testing.T) {
					ctx := tests.Context(t)
					_, err := chainReader.QueryKey(ctx, types.NewBoundContract("", "", ""), query.KeyFilter{}, query.LimitAndSort{}, &[]interface{}{nil})
					assert.True(t, errors.Is(err, errorType))
				})
			}

			t.Run("test QueryKey proto conversion", func(t *testing.T) {
				for _, tc := range generateQueryFilterTestCases(t) {
					impl.expectedQueryFilter = tc
					filter, err := query.Where(tc.Key, tc.Expressions...)
					require.NoError(t, err)
					_, err = cr.QueryKey(tests.Context(t), types.NewBoundContract("", "", ""), filter, query.LimitAndSort{}, &[]interface{}{nil})
					require.NoError(t, err)
				}
			})
		}
	})
}

var encoder = makeEncoder()

func makeEncoder() cbor.EncMode {
	opts := cbor.CoreDetEncOptions()
	opts.Sort = cbor.SortCanonical
	e, _ := opts.EncMode()
	return e
}

type fakeChainReaderInterfaceTester struct {
	interfaceTesterBase
	impl types.ContractReader
}

func (it *fakeChainReaderInterfaceTester) Setup(_ *testing.T) {
	fake, ok := it.impl.(*fakeChainReader)
	if ok {
		fake.stored = []TestStruct{}
		fake.triggers = []TestStruct{}
	}
}

func (it *fakeChainReaderInterfaceTester) GetChainReader(_ *testing.T) types.ContractReader {
	return it.impl
}

func (it *fakeChainReaderInterfaceTester) GetBindings(_ *testing.T) []types.BoundContract {
	return []types.BoundContract{
		{Name: AnyContractName, Address: AnyContractName},
		{Name: AnySecondContractName, Address: AnySecondContractName},
	}
}

func (it *fakeChainReaderInterfaceTester) SetLatestValue(t *testing.T, testStruct *TestStruct) {
	fake, ok := it.impl.(*fakeChainReader)
	assert.True(t, ok)
	fake.SetLatestValue(testStruct)
}

func (it *fakeChainReaderInterfaceTester) TriggerEvent(t *testing.T, testStruct *TestStruct) {
	fake, ok := it.impl.(*fakeChainReader)
	assert.True(t, ok)
	fake.SetTrigger(testStruct)
}

func (it *fakeChainReaderInterfaceTester) MaxWaitTimeForEvents() time.Duration {
	return time.Millisecond * 100
}

type fakeChainReader struct {
	fakeTypeProvider
	stored   []TestStruct
	lock     sync.Mutex
	triggers []TestStruct
}

func (f *fakeChainReader) Start(_ context.Context) error { return nil }

func (f *fakeChainReader) Close() error { return nil }

func (f *fakeChainReader) Ready() error { panic("unimplemented") }

func (f *fakeChainReader) Name() string { panic("unimplemented") }

func (f *fakeChainReader) HealthReport() map[string]error { panic("unimplemented") }

func (f *fakeChainReader) Bind(_ context.Context, _ []types.BoundContract) error {
	return nil
}

func (f *fakeChainReader) Unbind(_ context.Context, _ []types.BoundContract) error {
	return nil
}

func (f *fakeChainReader) SetLatestValue(ts *TestStruct) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.stored = append(f.stored, *ts)
}

func (f *fakeChainReader) GetLatestValue(_ context.Context, contract types.BoundContract, params, returnVal any) error {
	if contract.Method() == MethodReturningUint64 {
		r := returnVal.(*uint64)

		if contract.Contract() == AnyContractName {
			*r = AnyValueToReadWithoutAnArgument
		} else {
			*r = AnyDifferentValueToReadWithoutAnArgument
		}

		return nil
	} else if contract.Method() == MethodReturningUint64Slice {
		r := returnVal.(*[]uint64)
		*r = AnySliceToReadWithoutAnArgument

		return nil
	} else if contract.Method() == MethodReturningSeenStruct {
		pv := params.(*TestStruct)
		rv := returnVal.(*TestStructWithExtraField)

		rv.TestStruct = *pv
		rv.ExtraField = AnyExtraValue
		rv.Account = anyAccountBytes
		rv.BigField = big.NewInt(2)

		return nil
	} else if contract.Method() == EventName {
		f.lock.Lock()
		defer f.lock.Unlock()

		if len(f.triggers) == 0 {
			return types.ErrNotFound
		}

		*returnVal.(*TestStruct) = f.triggers[len(f.triggers)-1]

		return nil
	} else if contract.Method() == EventWithFilterName {
		f.lock.Lock()
		defer f.lock.Unlock()

		param := params.(*FilterEventParams)

		for i := len(f.triggers) - 1; i >= 0; i-- {
			if *f.triggers[i].Field == param.Field {
				*returnVal.(*TestStruct) = f.triggers[i]

				return nil
			}
		}

		return types.ErrNotFound
	} else if contract.Method() == DifferentMethodReturningUint64 {
		r := returnVal.(*uint64)
		*r = AnyDifferentValueToReadWithoutAnArgument

		return nil
	} else if contract.Method() != MethodTakingLatestParamsReturningTestStruct {
		return errors.New("unknown method " + contract.Method())
	}

	f.lock.Lock()
	defer f.lock.Unlock()

	lp := params.(*LatestParams)
	rv := returnVal.(*TestStruct)
	*rv = f.stored[lp.I-1]

	return nil
}

func (f *fakeChainReader) QueryKey(_ context.Context, _ types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, _ any) ([]types.Sequence, error) {
	if filter.Key == EventName {
		f.lock.Lock()
		defer f.lock.Unlock()
		if len(f.triggers) == 0 {
			return []types.Sequence{}, nil
		}

		var sequences []types.Sequence
		for _, trigger := range f.triggers {
			sequences = append(sequences, types.Sequence{Data: trigger})
		}

		if !limitAndSort.HasSequenceSort() {
			sort.Slice(sequences, func(i, j int) bool {
				if sequences[i].Data.(TestStruct).Field == nil || sequences[j].Data.(TestStruct).Field == nil {
					return false
				}
				return *sequences[i].Data.(TestStruct).Field > *sequences[j].Data.(TestStruct).Field
			})
		}

		return sequences, nil
	}
	return nil, nil
}

func (f *fakeChainReader) SetTrigger(testStruct *TestStruct) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.triggers = append(f.triggers, *testStruct)
}

type errChainReader struct {
	err error
}

func (e *errChainReader) Start(_ context.Context) error { return nil }

func (e *errChainReader) Close() error { return nil }

func (e *errChainReader) Ready() error { panic("unimplemented") }

func (e *errChainReader) Name() string { panic("unimplemented") }

func (e *errChainReader) HealthReport() map[string]error { panic("unimplemented") }

func (e *errChainReader) GetLatestValue(_ context.Context, _ types.BoundContract, _, _ any) error {
	return e.err
}

func (e *errChainReader) Bind(_ context.Context, _ []types.BoundContract) error {
	return e.err
}

func (e *errChainReader) Unbind(_ context.Context, _ []types.BoundContract) error {
	return e.err
}

func (e *errChainReader) QueryKey(_ context.Context, _ types.BoundContract, _ query.KeyFilter, _ query.LimitAndSort, _ any) ([]types.Sequence, error) {
	return nil, e.err
}

type protoConversionTestChainReader struct {
	expectedBindings     types.BoundContract
	expectedQueryFilter  query.KeyFilter
	expectedLimitAndSort query.LimitAndSort
}

func (pc *protoConversionTestChainReader) Start(_ context.Context) error { return nil }

func (pc *protoConversionTestChainReader) Close() error { return nil }

func (pc *protoConversionTestChainReader) Ready() error { panic("unimplemented") }

func (pc *protoConversionTestChainReader) Name() string { panic("unimplemented") }

func (pc *protoConversionTestChainReader) HealthReport() map[string]error { panic("unimplemented") }

func (pc *protoConversionTestChainReader) GetLatestValue(_ context.Context, _ types.BoundContract, _, _ any) error {
	return nil
}

func (pc *protoConversionTestChainReader) Bind(_ context.Context, bc []types.BoundContract) error {
	if !reflect.DeepEqual(pc.expectedBindings, bc) {
		return fmt.Errorf("bound contract wasn't parsed properly")
	}
	return nil
}

func (pc *protoConversionTestChainReader) Unbind(_ context.Context, bc []types.BoundContract) error {
	if !reflect.DeepEqual(pc.expectedBindings, bc) {
		return fmt.Errorf("bound contract wasn't parsed properly")
	}

	return nil
}

func (pc *protoConversionTestChainReader) QueryKey(_ context.Context, _ types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, _ any) ([]types.Sequence, error) {
	if !reflect.DeepEqual(pc.expectedQueryFilter, filter) {
		return nil, fmt.Errorf("filter wasn't parsed properly")
	}

	// using deep equal on a slice returns false when one slice is nil and another is empty
	// normalize to nil slices if empty or nil for comparison
	var (
		aSlice []query.SortBy
		bSlice []query.SortBy
	)

	if pc.expectedLimitAndSort.SortBy != nil && len(pc.expectedLimitAndSort.SortBy) > 0 {
		aSlice = pc.expectedLimitAndSort.SortBy
	}

	if limitAndSort.SortBy != nil && len(limitAndSort.SortBy) > 0 {
		bSlice = limitAndSort.SortBy
	}

	if !reflect.DeepEqual(pc.expectedLimitAndSort.Limit, limitAndSort.Limit) || !reflect.DeepEqual(aSlice, bSlice) {
		return nil, fmt.Errorf("limitAndSort wasn't parsed properly")
	}

	return nil, nil
}
