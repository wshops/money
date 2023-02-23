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
	m.amountCents = 0
	m.currency = USD
	moneySyncPool.Put(m)
}
