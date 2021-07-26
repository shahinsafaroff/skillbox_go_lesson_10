package main

import (
	"fmt"
	"math"
)

func main() {
	var deposit uint64
	var months, years uint16
	const monthsOfYear = 12
	var incomeRate, currentProfit, totalDeposit, creditProfit, debetProfit float64
	fmt.Println("Калкуляция дивидендов от оставленного вклада")
	fmt.Printf("Введите обшую сумму вклада на валюте: ")
	fmt.Scan(&deposit)
	fmt.Printf("Укажите ежемесячный процент по дивидендам: ")
	fmt.Scan(&incomeRate)
	fmt.Printf("Уточните срок вклада на сколко лет оставляете свой бабло у нас:) ")
	fmt.Scan(&years)
	totalDeposit = float64(deposit)
	for months = 1; months <= years*monthsOfYear; months++ {
		// Калкулируем обший процент на остаток по счету
		currentProfit = totalDeposit * incomeRate / 100
		// Получаем заисляемую на депозит сумму
		creditProfit = math.Trunc(currentProfit*100) / 100
		// Остаток отправляем в банк
		debetProfit += currentProfit - creditProfit
		// Итого по счету на текущий интервал
		totalDeposit += creditProfit
	}
	fmt.Printf("Итого за %v лет при сумме вклада %v совокупный доход от вклада составит: %.2f\n", years, deposit, totalDeposit)
	fmt.Printf("Сумма, зачисленная в ползу банка: %.10f\n", debetProfit)
}