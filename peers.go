package groupXcache

import (
	"context"
	pb "groupXcache/groupXcachepb"
)

// ProtoGetter is the interface that must be implemented by a peer.
type ProtoGetter interface {
	Get(ctx context.Context, in *pb.GetRequest, out *pb.GetResponse) error
}

