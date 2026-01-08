package bankImitation

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

var (
	bank            *Bank
	wg              sync.WaitGroup
	mu              sync.Mutex
	initialAmount   int
	finalAmount     int
	completedTasks  int
	unfinishedTasks int
	minBalance      int
	maxBalance      int
)

func fillOutBank() {
	bank = &Bank{
		Accounts: make([]*Account, 0, 100),
	}
	for i := 0; i < 100; i++ {
		bank.Accounts = append(bank.Accounts, &Account{ID: i + 1, Balance: 1000})
		initialAmount += 1000
	}
}

func (b *Bank) Transfer(fromID, toID int, amount int) bool {
	if fromID == toID {
		return false
	}
	first, second := fromID, toID
	if fromID > toID {
		first, second = toID, fromID
	}
	b.Accounts[first-1].mu.Lock()
	b.Accounts[second-1].mu.Lock()
	defer b.Accounts[first-1].mu.Unlock()
	defer b.Accounts[second-1].mu.Unlock()

	if b.Accounts[fromID-1].Balance >= amount {
		b.Accounts[fromID-1].Balance -= amount
		b.Accounts[toID-1].Balance += amount
		return true
	}
	return false
}

func Start() {
	fillOutBank()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10000; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				res := bank.Transfer(bank.Accounts[rand.IntN(100)].ID, bank.Accounts[rand.IntN(100)].ID, rand.IntN(100))
				mu.Lock()
				if res {
					completedTasks++
				} else {
					unfinishedTasks++
				}
				mu.Unlock()

			}()
		}
	}
	wg.Wait()
	minBalance = math.MaxInt
	for i := 0; i < len(bank.Accounts); i++ {
		maxBalance = int(math.Max(float64(maxBalance), float64(bank.Accounts[i].Balance)))
		minBalance = int(math.Min(float64(minBalance), float64(bank.Accounts[i].Balance)))
		finalAmount += bank.Accounts[i].Balance
	}

	fmt.Println("Начальная сумма денег: ", initialAmount)
	fmt.Println("Финальная сумма денег: ", finalAmount)
	fmt.Println("Количество заверщённых задач: ", completedTasks)
	fmt.Println("Количество незаверщённых задач: ", unfinishedTasks)
	fmt.Println("Минимальный баланс: ", minBalance)
	fmt.Println("Максимальный баланс: ", maxBalance)
}
