package queues

import (
	"sync"

	feetypes "github.com/smartcontractkit/chainlink/v2/common/fee/types"
	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/common/types"
)

// priorityHeap is a priority queue of transactions prioritized by creation time. The oldest transaction is at the front of the queue.
// It implements the heap interface in the container/heap package and is safe for concurrent access.
type priorityHeap[
	CHAIN_ID types.ID,
	ADDR, TX_HASH, BLOCK_HASH types.Hashable,
	R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH],
	SEQ types.Sequence,
	FEE feetypes.Fee,
] struct {
	sync.RWMutex
	txs       []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	idToIndex map[int64]int
}

// newPriorityHeap returns a new priorityHeap instance
func NewPriorityHeap[
	CHAIN_ID types.ID,
	ADDR, TX_HASH, BLOCK_HASH types.Hashable,
	R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH],
	SEQ types.Sequence,
	FEE feetypes.Fee,
](capacity int) *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE] {
	if capacity == 0 {
		panic("priority_heap: capacity must be greater than 0")
	}

	pq := priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{
		txs:       make([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], 0, capacity),
		idToIndex: make(map[int64]int),
	}

	return &pq
}

// Close clears the queue
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Close() {
	pq.Lock()
	defer pq.Unlock()

	clear(pq.txs)
	clear(pq.idToIndex)
	pq.txs = nil
	pq.idToIndex = nil
}

// FindIndexByID returns the index of the transaction with the given ID
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindIndexByID(id int64) int {
	pq.RLock()
	defer pq.RUnlock()

	i, ok := pq.idToIndex[id]
	if !ok {
		return -1
	}
	return i
}

// Peek returns the next transaction to be processed
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Peek() *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE] {
	pq.RLock()
	defer pq.RUnlock()

	if len(pq.txs) == 0 {
		return nil
	}
	return pq.txs[0]
}

// Cap returns the capacity of the queue
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Cap() int {
	pq.RLock()
	defer pq.RUnlock()

	if pq.txs == nil {
		return 0
	}
	return cap(pq.txs)
}

// Len, Less, Swap, Push, and Pop methods implement the heap interface
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Len() int {
	pq.RLock()
	defer pq.RUnlock()

	if pq.txs == nil {
		return 0
	}
	return len(pq.txs)
}
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Less(i, j int) bool {
	pq.RLock()
	defer pq.RUnlock()
	// We want Pop to give us the oldest, not newest, transaction based on creation time
	return pq.txs[i].CreatedAt.Before(pq.txs[j].CreatedAt)
}
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Swap(i, j int) {
	pq.Lock()
	defer pq.Unlock()

	pq.txs[i], pq.txs[j] = pq.txs[j], pq.txs[i]
	pq.idToIndex[pq.txs[i].ID] = j
	pq.idToIndex[pq.txs[j].ID] = i
}
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Push(tx any) {
	pq.Lock()
	defer pq.Unlock()

	if len(pq.txs) == cap(pq.txs) {
		// make room if cap is reached
		delete(pq.idToIndex, pq.txs[0].ID)
		copy(pq.txs, pq.txs[1:])
		pq.txs = pq.txs[:len(pq.txs)-1]
	}

	pq.txs = append(pq.txs, tx.(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]))
	pq.idToIndex[tx.(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]).ID] = len(pq.txs)
}
func (pq *priorityHeap[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Pop() any {
	pq.Lock()
	defer pq.Unlock()

	old := pq.txs
	n := len(old)
	tx := old[n-1]
	old[n-1] = nil // avoid memory leak
	pq.txs = old[0 : n-1]
	delete(pq.idToIndex, tx.ID)
	return tx
}