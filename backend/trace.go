package backend

import (
	"encoding/json"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/shopspring/decimal"
)

type CallFrame struct {
	Action              CallAction   `json:"action"`
	BlockHash           *common.Hash `json:"blockHash,omitempty"`
	BlockNumber         uint64       `json:"blockNumber"`
	Error               string       `json:"error,omitempty"`
	Result              *CallResult  `json:"result,omitempty"`
	Subtraces           int          `json:"subtraces"`
	TraceAddress        []int        `json:"traceAddress"`
	TransactionHash     *common.Hash `json:"transactionHash"`
	TransactionPosition uint64       `json:"transactionPosition"`
	Type                string       `json:"type"`
}

type CallAction struct {
	Author         *common.Address `json:"author,omitempty"`
	RewardType     string          `json:"rewardType,omitempty"`
	SelfDestructed *common.Address `json:"address,omitempty"`
	Balance        *big.Int        `json:"balance,omitempty"`
	CallType       string          `json:"callType,omitempty"`
	CreationMethod string          `json:"creationMethod,omitempty"`
	From           *common.Address `json:"from,omitempty"`
	Gas            *uint64         `json:"gas,omitempty"`
	Init           *[]byte         `json:"init,omitempty"`
	Input          *[]byte         `json:"input,omitempty"`
	RefundAddress  *common.Address `json:"refundAddress,omitempty"`
	To             *common.Address `json:"to,omitempty"`
	Value          *big.Int        `json:"value,omitempty"`
}

// MarshalJSON marshals as JSON.
func (f CallAction) MarshalJSON() ([]byte, error) {
	type flatCallAction struct {
		Author         *common.Address `json:"author,omitempty"`
		RewardType     string          `json:"rewardType,omitempty"`
		SelfDestructed *common.Address `json:"address,omitempty"`
		Balance        *hexutil.Big    `json:"balance,omitempty"`
		CallType       string          `json:"callType,omitempty"`
		CreationMethod string          `json:"creationMethod,omitempty"`
		From           *common.Address `json:"from,omitempty"`
		Gas            *hexutil.Uint64 `json:"gas,omitempty"`
		Init           *hexutil.Bytes  `json:"init,omitempty"`
		Input          *hexutil.Bytes  `json:"input,omitempty"`
		RefundAddress  *common.Address `json:"refundAddress,omitempty"`
		To             *common.Address `json:"to,omitempty"`
		Value          *hexutil.Big    `json:"value,omitempty"`
	}
	var enc flatCallAction
	enc.Author = f.Author
	enc.RewardType = f.RewardType
	enc.SelfDestructed = f.SelfDestructed
	enc.Balance = (*hexutil.Big)(f.Balance)
	enc.CallType = f.CallType
	enc.CreationMethod = f.CreationMethod
	enc.From = f.From
	enc.Gas = (*hexutil.Uint64)(f.Gas)
	enc.Init = (*hexutil.Bytes)(f.Init)
	enc.Input = (*hexutil.Bytes)(f.Input)
	enc.RefundAddress = f.RefundAddress
	enc.To = f.To
	enc.Value = (*hexutil.Big)(f.Value)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *CallAction) UnmarshalJSON(input []byte) error {
	type flatCallAction struct {
		Author         *common.Address `json:"author,omitempty"`
		RewardType     *string         `json:"rewardType,omitempty"`
		SelfDestructed *common.Address `json:"address,omitempty"`
		Balance        *hexutil.Big    `json:"balance,omitempty"`
		CallType       *string         `json:"callType,omitempty"`
		CreationMethod *string         `json:"creationMethod,omitempty"`
		From           *common.Address `json:"from,omitempty"`
		Gas            *hexutil.Uint64 `json:"gas,omitempty"`
		Init           *hexutil.Bytes  `json:"init,omitempty"`
		Input          *hexutil.Bytes  `json:"input,omitempty"`
		RefundAddress  *common.Address `json:"refundAddress,omitempty"`
		To             *common.Address `json:"to,omitempty"`
		Value          *hexutil.Big    `json:"value,omitempty"`
	}
	var dec flatCallAction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Author != nil {
		f.Author = dec.Author
	}
	if dec.RewardType != nil {
		f.RewardType = *dec.RewardType
	}
	if dec.SelfDestructed != nil {
		f.SelfDestructed = dec.SelfDestructed
	}
	if dec.Balance != nil {
		f.Balance = (*big.Int)(dec.Balance)
	}
	if dec.CallType != nil {
		f.CallType = *dec.CallType
	}
	if dec.CreationMethod != nil {
		f.CreationMethod = *dec.CreationMethod
	}
	if dec.From != nil {
		f.From = dec.From
	}
	if dec.Gas != nil {
		f.Gas = (*uint64)(dec.Gas)
	}
	if dec.Init != nil {
		f.Init = (*[]byte)(dec.Init)
	}
	if dec.Input != nil {
		f.Input = (*[]byte)(dec.Input)
	}
	if dec.RefundAddress != nil {
		f.RefundAddress = dec.RefundAddress
	}
	if dec.To != nil {
		f.To = dec.To
	}
	if dec.Value != nil {
		f.Value = (*big.Int)(dec.Value)
	}
	return nil
}

type CallResult struct {
	Address *common.Address `json:"address,omitempty"`
	Code    *[]byte         `json:"code,omitempty"`
	GasUsed *uint64         `json:"gasUsed,omitempty"`
	Output  *[]byte         `json:"output,omitempty"`
}

// MarshalJSON marshals as JSON.
func (f CallResult) MarshalJSON() ([]byte, error) {
	type flatCallResult struct {
		Address *common.Address `json:"address,omitempty"`
		Code    *hexutil.Bytes  `json:"code,omitempty"`
		GasUsed *hexutil.Uint64 `json:"gasUsed,omitempty"`
		Output  *hexutil.Bytes  `json:"output,omitempty"`
	}
	var enc flatCallResult
	enc.Address = f.Address
	enc.Code = (*hexutil.Bytes)(f.Code)
	enc.GasUsed = (*hexutil.Uint64)(f.GasUsed)
	enc.Output = (*hexutil.Bytes)(f.Output)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *CallResult) UnmarshalJSON(input []byte) error {
	type flatCallResult struct {
		Address *common.Address `json:"address,omitempty"`
		Code    *hexutil.Bytes  `json:"code,omitempty"`
		GasUsed *hexutil.Uint64 `json:"gasUsed,omitempty"`
		Output  *hexutil.Bytes  `json:"output,omitempty"`
	}
	var dec flatCallResult
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address != nil {
		f.Address = dec.Address
	}
	if dec.Code != nil {
		f.Code = (*[]byte)(dec.Code)
	}
	if dec.GasUsed != nil {
		f.GasUsed = (*uint64)(dec.GasUsed)
	}
	if dec.Output != nil {
		f.Output = (*[]byte)(dec.Output)
	}
	return nil
}

type Trace struct {
	Timestamp       time.Time        `json:"block_timestamp" gorm:"column:block_timestamp" example:"2023-01-02 12:00:23"`
	BlockNum        uint64           `json:"blknum" gorm:"column:blknum" example:"14218502"`
	TransactionHash *string          `json:"txhash" gorm:"column:txhash" example:"0xaae3c030ee04b1ef071e00198818a113a3ac20db252fbfba4f78572aa59f5226"`
	TransactionPos  uint64           `json:"txpos" gorm:"column:txpos" example:"0"`
	FromAddress     *string          `json:"from_address" gorm:"column:from_address" example:"0x1dc907d55f1be2bc4370feb0f01fb89324b8941c"`
	ToAddress       *string          `json:"to_address" gorm:"column:to_address" example:"0x7a250d5630b4cf539739df2c5dacb4c659f2488d"`
	Value           *decimal.Decimal `json:"value" gorm:"column:value" example:"250000000000000000"`
	Input           string           `json:"input" gorm:"column:input" example:"0x"`
	Output          string           `json:"output" gorm:"column:output" example:"0x"`
	TraceType       string           `json:"trace_type" gorm:"column:trace_type" example:"call"`
	CallType        string           `json:"call_type" gorm:"column:call_type" example:"call"`
	RewardType      string           `json:"reward_type" gorm:"column:reward_type" example:"block"`
	Gas             *decimal.Decimal `json:"gas" gorm:"column:gas" example:"477212"`
	GasUsed         uint64           `json:"gas_used" gorm:"column:gas_used" example:"91903"`
	SubTraces       int              `json:"sub_traces" gorm:"column:sub_traces" example:"0"`
	TraceAddress    string           `json:"trace_address" gorm:"column:trace_address" example:"[]"`
	Error           string           `json:"error" gorm:"column:error" example:"Reverted"`
	// Status          int              `json:"status" gorm:"column:status" example:"0"`
}

func (t *Trace) AsCallFrame() *CallFrame {
	// common fields
	var txHash common.Hash
	if t.TransactionHash != nil {
		txHash = common.HexToHash(*t.TransactionHash)
	}
	frame := &CallFrame{
		BlockNumber:         t.BlockNum,
		Error:               t.Error,
		Subtraces:           t.SubTraces,
		TransactionHash:     &txHash,
		TransactionPosition: t.TransactionPos,
		Type:                t.TraceType,
	}

	traceAddress := t.TraceAddress
	traceAddress = strings.ReplaceAll(traceAddress, "[", "")
	traceAddress = strings.ReplaceAll(traceAddress, "]", "")
	traceAddress = strings.ReplaceAll(traceAddress, " ", "")
	if traceAddress == "" {
		frame.TraceAddress = []int{}
	} else {
		traceAddresses := strings.Split(traceAddress, ",")
		traceIntAddresses := make([]int, len(traceAddresses))
		for i, s := range traceAddresses {
			pos, e := strconv.Atoi(s)
			if e != nil {
				log.Error("failed to parse traceAddress", "s", s, "e", e)
			}
			traceIntAddresses[i] = pos
		}
		frame.TraceAddress = traceIntAddresses
	}

	var (
		from      common.Address
		to        common.Address
		input, _  = hexutil.Decode(t.Input)
		output, _ = hexutil.Decode(t.Output)
	)
	if t.FromAddress != nil {
		from = common.HexToAddress(*t.FromAddress)
	}
	if t.ToAddress != nil {
		to = common.HexToAddress(*t.ToAddress)
	}

	switch strings.ToUpper(t.TraceType) {
	case vm.CREATE.String(), vm.CREATE2.String():
		gas := uint64(t.Gas.BigInt().Int64())
		gasUsed := t.GasUsed
		frame.Action = CallAction{
			From:  &from,
			Gas:   &gas,
			Value: t.Value.BigInt(),
			Init:  &input,
		}
		frame.Result = &CallResult{
			GasUsed: &gasUsed,
			Address: &to,
			Code:    &output,
		}
	case vm.SELFDESTRUCT.String(), "SUICIDE":
		frame.Action = CallAction{
			SelfDestructed: &from,
			Balance:        t.Value.BigInt(),
			RefundAddress:  &to,
		}
	case vm.CALL.String(), vm.STATICCALL.String(), vm.CALLCODE.String(), vm.DELEGATECALL.String():
		gas := uint64(t.Gas.BigInt().Int64())
		gasUsed := t.GasUsed
		frame.Action = CallAction{
			From:     &from,
			To:       &to,
			Gas:      &gas,
			Value:    t.Value.BigInt(),
			Input:    &input,
			CallType: t.CallType,
		}
		frame.Result = &CallResult{
			GasUsed: &gasUsed,
			Output:  &output,
		}
	default:
		log.Error("unrecognized call frame", "traceType", t.TraceType)
	}

	// Revert output contains useful information (revert reason).
	// Otherwise discard result.
	if t.Error != "" && t.Error != vm.ErrExecutionReverted.Error() {
		frame.Result = nil
	}

	return frame
}
