package agent

import (
	"io"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/lib"
	"github.com/hashicorp/serf/serf"
)

// delegate defines the interface shared by both
// consul.Client and consul.Server.
type commonDelegate interface {
	Encrypted() bool
	GetLANCoordinate() (lib.CoordinateSet, error)
	Leave() error
	LANMembers() []serf.Member
	LANMembersAllSegments() ([]serf.Member, error)
	LANSegmentMembers(segment string) ([]serf.Member, error)
	LocalMember() serf.Member
	JoinLAN(addrs []string) (n int, err error)
	RemoveFailedNode(node string) error
	RPC(method string, args interface{}, reply interface{}) error
	SnapshotRPC(args *structs.SnapshotRequest, in io.Reader, out io.Writer, replyFn structs.SnapshotReplyFn) error
	Shutdown() error
	Stats() map[string]map[string]string
}
