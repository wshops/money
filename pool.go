package money

import "sync"

var moneySyncPool = sync.Pool{
	New: func() interface{} {
		return new(money)
	},
}

func acquireMoney() *money {
	return moneySyncPool.Get().(*money)
}

func releaseMoney(m *money) {
	m.AmountCents = 0
	m.Currency = USD
	moneySyncPool.Put(m)
}
