package money

import "sync"

var moneySyncPool = sync.Pool{
	New: func() interface{} {
		return new(Money)
	},
}

func acquireMoney() *Money {
	return moneySyncPool.Get().(*Money)
}

func releaseMoney(m *Money) {
	m.amountCents = 0
	m.currency = USD
	moneySyncPool.Put(m)
}
