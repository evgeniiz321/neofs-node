/*
Package state collects functionality for verifying states of network map members.

NetMapCandidateValidator type provides an interface for checking the network
map candidates.
*/
package state

import (
	"errors"

	"github.com/nspcc-dev/neofs-sdk-go/netmap"
)

// NetMapCandidateValidator represents tool which checks state of nodes which
// are going to register in the NeoFS network (enter the network map).
//
// NetMapCandidateValidator can be instantiated using built-in var declaration
// and currently doesn't require any additional initialization.
//
// NetMapCandidateValidator implements
// github.com/nspcc-dev/neofs-node/pkg/innerring/processors/netmap.NodeValidator.
type NetMapCandidateValidator struct {
}

// VerifyAndUpdate checks state of the network map candidate described by
// netmap.NodeInfo parameter. Returns no error if status is correct, otherwise
// returns an error describing a violation of the rules:
//
//	status MUST be ONLINE
//
// VerifyAndUpdate does not mutate the parameter in a binary format.
//
// See also netmap.NodeInfo.IsOnline/SetOnline.
func (x *NetMapCandidateValidator) VerifyAndUpdate(node *netmap.NodeInfo) error {
	if node.IsOnline() {
		return nil
	}

	return errors.New("invalid status: MUST be ONLINE")
}
