package types

import (
	"fmt"

	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	sdkerrors "cosmossdk.io/errors"
)

const (
	// ackErrorString defines a string constant included in error acknowledgements
	// NOTE: Changing this const is state machine breaking as acknowledgements are written into state
	ackErrorString = "error handling packet on destination chain: see events for details"
)

// NewErrorAcknowledgement returns a deterministic error string which may be used in
// the packet acknowledgement.
func NewErrorAcknowledgement(err error) channeltypes.Acknowledgement {
	// the ABCI code is included in the abcitypes.ResponseDeliverTx hash
	// constructed in Tendermint and is therefore deterministic
	_, code, _ := sdkerrors.ABCIInfo(err, false) // discard non-deterministic codespace and log values

	errorString := fmt.Errorf("ABCI code: %d: %s", code, ackErrorString)

	return channeltypes.NewErrorAcknowledgement(errorString)
}
