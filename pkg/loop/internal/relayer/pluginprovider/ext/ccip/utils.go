package ccip

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	ccippb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/ccip"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

type GRPCResourceCloser interface {
	Close(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

// shutdownGRPCServer is a helper function to release server resources
// created by a grpc client.
func shutdownGRPCServer(ctx context.Context, rc GRPCResourceCloser) error {
	_, err := rc.Close(ctx, &emptypb.Empty{})
	// due to the handler in the server, it may shutdown before it sends a response to client
	// in that case, we expect the client to receive an Unavailable or Internal error
	if status.Code(err) == codes.Unavailable || status.Code(err) == codes.Internal {
		return nil
	}
	return err
}

func sequenceNumbersToPB(seqNums []cciptypes.SequenceNumberRange) ([]*ccippb.SequenceNumberRange, error) {
	res := make([]*ccippb.SequenceNumberRange, len(seqNums))
	for i, seqNum := range seqNums {
		res[i] = &ccippb.SequenceNumberRange{
			Min: seqNum.Min,
			Max: seqNum.Max,
		}
	}
	return res, nil
}

func sequenceNumbersPBToSlice(seqNums []*ccippb.SequenceNumberRange) ([]cciptypes.SequenceNumberRange, error) {
	res := make([]cciptypes.SequenceNumberRange, len(seqNums))
	for i, seqNum := range seqNums {
		res[i] = cciptypes.SequenceNumberRange{
			Min: seqNum.Min,
			Max: seqNum.Max,
		}
	}
	return res, nil
}
