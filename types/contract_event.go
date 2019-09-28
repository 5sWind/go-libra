package types

import (
	"errors"
	"fmt"

	"github.com/the729/go-libra/crypto/sha3libra"
	"github.com/the729/go-libra/generated/pbtypes"
	"github.com/the729/go-libra/types/proof"
	"github.com/the729/lcs"
)

type EventKey []byte

type EventHandle struct {
	Count uint64
	Key   EventKey
}

// ContractEvent is a output event of transaction
type ContractEvent struct {
	Key            EventKey
	SequenceNumber uint64
	Data           []byte
}

// EventList is a list of events
type EventList []*ContractEvent

// EventProof is a chain of proof that a event is included in the ledger
type EventProof struct {
	// LedgerInfoToTransactionInfoProof is a Merkle Tree accumulator to prove that TransactionInfo
	// is included in the ledger.
	LedgerInfoToTransactionInfoProof *proof.Accumulator

	// TransactionInfo is the info of the transaction that leads to this version of the ledger.
	*TransactionInfo

	// TransactionInfoToEventProof is an accumulator proof from event root hash in TransactionInfo
	// to actual event.
	TransactionInfoToEventProof *proof.Accumulator
}

// EventWithProof is an event with proof
type EventWithProof struct {
	TransactionVersion uint64
	EventIndex         uint64
	Event              *ContractEvent
	Proof              *EventProof
}

// ProvenEvent is an event proven to be included in the ledger.
type ProvenEvent struct {
	proven     bool
	txnVersion uint64
	eventIndex uint64
	event      *ContractEvent
	ledgerInfo *ProvenLedgerInfo
}

// Clone deep clones this struct.
func (eh *EventHandle) Clone() *EventHandle {
	out := &EventHandle{}
	out.Key = cloneBytes(eh.Key)
	out.Count = eh.Count
	return out
}

// FromProto parses a protobuf struct into this struct.
func (e *ContractEvent) FromProto(pb *pbtypes.Event) error {
	if pb == nil {
		return ErrNilInput
	}
	e.Key = pb.Key
	e.SequenceNumber = pb.SequenceNumber
	e.Data = pb.EventData

	return nil
}

// Hash ouptuts the hash of this struct, using the appropriate hash function.
func (e *ContractEvent) Hash() sha3libra.HashValue {
	hasher := sha3libra.NewContractEvent()
	if err := lcs.NewEncoder(hasher).Encode(e); err != nil {
		panic(err)
	}
	return hasher.Sum([]byte{})
}

// Clone deep clones this struct.
func (e *ContractEvent) Clone() *ContractEvent {
	out := &ContractEvent{}
	out.Key = cloneBytes(e.Key)
	out.SequenceNumber = e.SequenceNumber
	out.Data = cloneBytes(e.Data)
	return out
}

// Hash ouptuts the hash of this struct, using the appropriate hash function.
func (el EventList) Hash() sha3libra.HashValue {
	nodeHasher := sha3libra.NewEventAccumulator()
	hasher := sha3libra.NewAccumulator(nodeHasher)
	for _, e := range el {
		hasher.Write(e.Hash())
	}
	return hasher.Sum([]byte{})
}

// Clone deep clones this struct.
func (el EventList) Clone() EventList {
	if el == nil {
		return nil
	}
	out := make([]*ContractEvent, 0, len(el))
	for _, e := range el {
		out = append(out, e.Clone())
	}
	return out
}

// FromProto parses a protobuf struct into this struct.
func (ep *EventProof) FromProto(pb *pbtypes.EventProof) error {
	var err error
	if pb == nil {
		return ErrNilInput
	}

	ep.LedgerInfoToTransactionInfoProof = &proof.Accumulator{Hasher: sha3libra.NewTransactionAccumulator()}
	err = ep.LedgerInfoToTransactionInfoProof.FromProto(pb.LedgerInfoToTransactionInfoProof)
	if err != nil {
		return err
	}
	ep.TransactionInfo = &TransactionInfo{}
	err = ep.TransactionInfo.FromProto(pb.TransactionInfo)
	if err != nil {
		return err
	}
	ep.TransactionInfoToEventProof = &proof.Accumulator{Hasher: sha3libra.NewEventAccumulator()}
	err = ep.TransactionInfoToEventProof.FromProto(pb.TransactionInfoToEventProof)
	if err != nil {
		return err
	}
	return nil
}

// FromProto parses a protobuf struct into this struct.
func (ep *EventWithProof) FromProto(pb *pbtypes.EventWithProof) error {
	if pb == nil {
		return ErrNilInput
	}
	ep.TransactionVersion = pb.TransactionVersion
	ep.EventIndex = pb.EventIndex
	ep.Event = &ContractEvent{}
	if err := ep.Event.FromProto(pb.Event); err != nil {
		return err
	}
	ep.Proof = &EventProof{}
	if err := ep.Proof.FromProto(pb.Proof); err != nil {
		return err
	}
	return nil
}

// Verify the proof of the event, and output a ProvenEvent if successful.
func (ep *EventWithProof) Verify(provenLedgerInfo *ProvenLedgerInfo) (*ProvenEvent, error) {
	var err error

	eventHash := ep.Event.Hash()
	err = ep.Proof.TransactionInfoToEventProof.Verify(ep.EventIndex, eventHash, ep.Proof.EventRootHash)
	if err != nil {
		return nil, fmt.Errorf("cannot verify event from transaction info: %v", err)
	}

	if ep.TransactionVersion > provenLedgerInfo.GetVersion() {
		return nil, errors.New("event txn version > ledger version")
	}

	err = ep.Proof.LedgerInfoToTransactionInfoProof.Verify(
		ep.TransactionVersion, ep.Proof.TransactionInfo.Hash(),
		provenLedgerInfo.GetTransactionAccumulatorHash(),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot verify transaction info from ledger info: %v", err)
	}

	return &ProvenEvent{
		proven:     true,
		txnVersion: ep.TransactionVersion,
		eventIndex: ep.EventIndex,
		event:      ep.Event.Clone(),
	}, nil
}

func (pe *ProvenEvent) GetTransactionVersion() uint64 {
	if !pe.proven {
		panic("not valid proven event")
	}
	return pe.txnVersion
}

func (pe *ProvenEvent) GetEventIndex() uint64 {
	if !pe.proven {
		panic("not valid proven event")
	}
	return pe.eventIndex
}

func (pe *ProvenEvent) GetEvent() *ContractEvent {
	if !pe.proven {
		panic("not valid proven event")
	}
	return pe.event.Clone()
}
