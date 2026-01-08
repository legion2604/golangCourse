package bankImitation

import "sync"

type Account struct {
	ID      int
	Balance int
	mu      sync.Mutex
}
type Bank struct {
	Accounts []*Account
}
