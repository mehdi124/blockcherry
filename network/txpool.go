package network

import (
	"sort"
	"sync"

	"github.com/mehdi124/blockcherry/core"
	"github.com/mehdi124/blockcherry/types"
)

type TxMapSorter struct {
	transactions []*core.Transaction
}

func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	txx := make([]*core.Transaction, len(txMap))

	i := 0
	for _, val := range txMap {
		txx[i] = val
		i++
	}

	s := &TxMapSorter{txx}
	sort.Sort(s)

	return s
}

func (s *TxMapSorter) Len() int {
	return len(s.transactions)
}

func (s *TxMapSorter) Swap(i, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}

func (s *TxMapSorter) Less(i, j int) bool {
	return s.transactions[i].FirstSeen() < s.transactions[j].FirstSeen()
}

type TxPool struct {
	all     *TxSortedMap
	pending *TxSortedMap

	maxLength int
}

func NewTxPool() *TxPool {
	return &TxPool{
		all:       NewSortedMap(),
		pending:   NewSortedMap(),
		maxLength: maxLength,
	}
}

func (p *TxPool) Add(tx *core.Transaction) error {

	if p.all.Count() == p.maxLength {
		oldest := p.all.First()
		p.all.Remove(oldes.Hash(core.TxHasher{}))
	}

	if !p.all.Contains(tx.Hash(core.TxHasher{})) {
		p.all.Add(tx)
		p.pending.Add(tx)
	}
}

func (p *TxPool) Contains(hash types.Hash) bool {
	return p.all.Contains(hash)
}

func (p *TxPool) Pending() []*core.Transaction {
	return p.pending.txx.Data
}

func (p *TxPool) ClearPending() {
	p.pending.Clear()
}

func (p *TxPool) PendingCount() int {
	return p.pending.Count()
}

type TxSortedMap struct {
	lock   sync.RWMutex
	lookup map[types.Hash]*core.Transaction
	txx    types.List[*core.Transaction]
}

func NewTxSortedMap() *TxSortedMap {
	return &TxSortedMap{
		lookup: make(map[types.Hash]*core.Transaction),
		txx:    types.NewList[*core.Transaction](),
	}
}

func (tx *TxSortedMap) First() *core.Transaction {
	t.lock.RLock()
	defer t.lock.RUnlock()

	first := t.txx.Get(0)
	return t.lookup[first.Hash(core.TxHasher{})]
}
