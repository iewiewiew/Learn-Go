package main

import (
	"errors"
	"fmt"
)

/**
 * @author       weimenghua
 * @time         2024-07-22 17:22
 * @description
 */

// BankAccount 定义 BankAccount 结构体
type BankAccount struct {
	AccountNumer string
	Balance      float64
}

// Deposit 存款方法，使用指针接收器
func (ba *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("存款金额必须大于零")
		return
	}
	ba.Balance += amount
	fmt.Printf("成功存入 %.2f 元，当前余额: %.2f 元\n", amount, ba.Balance)
}

// Withdrow 取款方法，使用指针接收器
func (ba *BankAccount) Withdrow(amount float64) error {
	if amount <= 0 {
		fmt.Println("取款金额必须大于零")
		return nil
	}

	if amount > ba.Balance {
		return errors.New("余额不足")
	}

	ba.Balance -= amount
	fmt.Printf("成功取出 %.2f 元，当前余额：%.2f 元\n", amount, ba.Balance)

	return nil
}

// CheckBalance 查询余额方法，使用值接收器
func (ba BankAccount) CheckBalance() float64 {
	return ba.Balance
}

func test() {
	// 声明 BankAccount 变量，自动初始化其字段为零值
	var BA BankAccount

	// 初始化字段
	BA.AccountNumer = "1234567890"
	BA.Balance = 1000.00

	// 存款
	BA.Deposit(500.00)

	// 取款
	err := BA.Withdrow(200.00)
	if err != nil {
		fmt.Println("取款失败:", err)
	}

	// 查询余额
	balance := BA.CheckBalance()
	fmt.Printf("账户余额: %.2f 元\n", balance)

	// 尝试取款超额
	err = BA.Withdrow(2000.00)
	if err != nil {
		fmt.Println("取款失败:", err)
	}
}

/*
*
Withdraw 方法：使用指针接收器 ba *BankAccount，以便在取款时修改账户的实际余额，并进行错误处理。
CheckBalance 方法：使用值接收器 ba BankAccount，因为该方法只是读取余额，不需要修改结构体的值。
*/
func main() {
	// 创建一个 BankAccount 实例
	account := BankAccount{AccountNumer: "1234567890", Balance: 1000.00}

	// 存款
	account.Deposit(500.00)

	// 取款
	err := account.Withdrow(200.00)
	if err != nil {
		fmt.Println("取款失败:", err)
	}

	// 查询余额
	balance := account.CheckBalance()
	fmt.Printf("账户余额: %.2f 元\n", balance)

	// 尝试取款超额
	err = account.Withdrow(2000.00)
	if err != nil {
		fmt.Println("取款失败:", err)
	}

	fmt.Println()
	test()
}
