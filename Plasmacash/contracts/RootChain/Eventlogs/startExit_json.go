// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eventlog

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*startExitMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s StartExitEvent) MarshalJSON() ([]byte, error) {
	type StartExitEvent struct {
		Exiter       common.Address `json:"exiter"         gencodec:"required"`
		DepositIndex hexutil.Uint64 `json:"depositIndex"   gencodec:"required"`
		Denomination hexutil.Uint64 `json:"denomination"   gencodec:"required"`
		TokenID      hexutil.Uint64 `json:"tokenID"        gencodec:"required"`
		TS           hexutil.Uint64 `json:"timestamp"      gencodec:"required"`
	}
	var enc StartExitEvent
	enc.Exiter = s.Exiter
	enc.DepositIndex = hexutil.Uint64(s.DepositIndex)
	enc.Denomination = hexutil.Uint64(s.Denomination)
	enc.TokenID = hexutil.Uint64(s.TokenID)
	enc.TS = hexutil.Uint64(s.TS)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *StartExitEvent) UnmarshalJSON(input []byte) error {
	type StartExitEvent struct {
		Exiter       *common.Address `json:"exiter"         gencodec:"required"`
		DepositIndex *hexutil.Uint64 `json:"depositIndex"   gencodec:"required"`
		Denomination *hexutil.Uint64 `json:"denomination"   gencodec:"required"`
		TokenID      *hexutil.Uint64 `json:"tokenID"        gencodec:"required"`
		TS           *hexutil.Uint64 `json:"timestamp"      gencodec:"required"`
	}
	var dec StartExitEvent
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Exiter == nil {
		return errors.New("missing required field 'exiter' for StartExitEvent")
	}
	s.Exiter = *dec.Exiter
	if dec.DepositIndex == nil {
		return errors.New("missing required field 'depositIndex' for StartExitEvent")
	}
	s.DepositIndex = uint64(*dec.DepositIndex)
	if dec.Denomination == nil {
		return errors.New("missing required field 'denomination' for StartExitEvent")
	}
	s.Denomination = uint64(*dec.Denomination)
	if dec.TokenID == nil {
		return errors.New("missing required field 'tokenID' for StartExitEvent")
	}
	s.TokenID = uint64(*dec.TokenID)
	if dec.TS == nil {
		return errors.New("missing required field 'timestamp' for StartExitEvent")
	}
	s.TS = uint64(*dec.TS)
	return nil
}
