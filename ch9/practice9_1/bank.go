package practice9_1

type withdrawRequest struct {
	amount int
	result chan bool
}

var balances = make(chan int)
var deposits = make(chan int)
var withdraws = make(chan withdrawRequest)

func Deposit(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- withdrawRequest{
		amount,
		ch,
	}
	return <-ch
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case request := <-withdraws:
			tmp := balance - request.amount
			res := tmp > 0
			if res {
				balance = tmp
			}
			request.result <- res
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
