package arraysslicesrevisited

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

func Reduce[T, U any](collection []T, accumulator func(U, T) U, initialValue U) U {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
    return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
    Name string
    Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
    return Reduce(
        transactions,
        applyTransaction,
        account,
    )
}

func applyTransaction(a Account, transaction Transaction) Account {
    if transaction.From == a.Name {
        a.Balance -= transaction.Sum
    }
    if transaction.To == a.Name {
        a.Balance += transaction.Sum
    }
        return a
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
    for _, v := range items {
        if predicate(v) {
            return v, true
        }
    }
    return
}
