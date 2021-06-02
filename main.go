package main

import (
	"fmt"
)

func calculateProfit(gia_thuc_hien float64, gia_thanh_toan float64, gia_CW float64, rate float64, luong_CW int) float64 {
	var profit float64
	if gia_thanh_toan > gia_thuc_hien {
		profit = (float64(float64(luong_CW)/rate) * (gia_thanh_toan - gia_thuc_hien)) - float64(luong_CW)*gia_CW
	} else {
		profit = 0 - float64(luong_CW)*gia_CW
	}
	return profit
}

func float64ToMoneyFormat(floatMoney float64) string {
	stringMoney := fmt.Sprintf("%.2f", floatMoney)
	for i := len(stringMoney) - 3; i >= 0; i -= 3 {
		if (len(stringMoney)-3-i) != 0 && i != 0 && string(stringMoney[i-1]) != "-" {
			stringMoney = stringMoney[:i] + "," + stringMoney[i:]
		}
	}
	return stringMoney
}

func main() {
	var gia_thanh_toan, gia_thuc_hien float64
	var luong_CW int
	var rate float64
	var gia_CW float64 = 1000
	fmt.Print("Giá thanh toán: ")
	fmt.Scanln(&gia_thanh_toan)
	fmt.Print("Giá thực hiện: ")
	fmt.Scanln(&gia_thuc_hien)
	fmt.Print("Giá một CW: ")
	fmt.Scanln(&gia_CW)
	fmt.Print("Lượng CW: ")
	fmt.Scanln(&luong_CW)
	fmt.Print("Rate CW/CK: ")
	fmt.Scanln(&rate)
	fmt.Printf("Profit: %s VND\n", float64ToMoneyFormat(calculateProfit(gia_thuc_hien, gia_thanh_toan, gia_CW, rate, luong_CW)))
}
