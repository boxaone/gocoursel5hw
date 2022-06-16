package main

import "fmt"

type Cop int     // Тип копійок
type Hrn float64 // Тип гривень
type Buy struct {
	Apples int
	Pears  int
}

const (
	ApplePrice Hrn = 5.99  // Ціна яблук
	PearPrice  Hrn = 7.00  // Ціна груш
	WholeMoney Hrn = 23.00 // Грошей взагалі в наявності
)

var (
	buy1 = &Buy{Apples: 9, Pears: 8}
	buy2 = &Buy{Pears: int(hrnToCop(WholeMoney) / hrnToCop(PearPrice))}
	buy3 = &Buy{Apples: int(hrnToCop(WholeMoney) / hrnToCop(ApplePrice))}
	buy4 = &Buy{Pears: 2, Apples: 2}
)

// Переводимо гривні(раціональні) в копійки(цілі)
func hrnToCop(grn Hrn) Cop {
	return (Cop)(grn * 100)
}

// Переводимо копійки(цілі) в гривні(раціональні)
func copToHrn(cop Cop) Hrn {
	return (Hrn)(cop) / 100
}

// Функція знаходження суми купівлі
func buySum(toBuy *Buy) Hrn {

	return copToHrn(Cop(toBuy.Apples)*hrnToCop(ApplePrice) + Cop(toBuy.Pears)*hrnToCop(PearPrice))
}

// Виводимо boolean українською
func boolToUkrStr(value bool) string {

	if value {
		return "Так"
	}

	return "Ні"
}

// Виводить список того що купляється
func buyToS(buy *Buy) string {
	var s string
	if buy.Apples > 0 {
		s = s + fmt.Sprintf("%v яблук", buy.Apples)
	}

	if buy.Pears > 0 {
		if buy.Apples > 0 {
			s = s + " та "
		}
		s = s + fmt.Sprintf("%v груш", buy.Pears)
	}
	return s
}

func main() {

	fmt.Printf("Ми маємо: %v грн взагалі\n", WholeMoney)
	fmt.Printf("Скільки грошей треба витратити, щоб купити %s?\n", buyToS(buy1))
	fmt.Printf("Щоб купити %s треба %v грн\n", buyToS(buy1), buySum(buy1))
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Printf("Ми можемо купити %s\n", buyToS(buy2))
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Printf("Ми можемо купити %s\n", buyToS(buy3))
	fmt.Printf("Чи ми можемо купити %s? %s!\n", buyToS(buy4), boolToUkrStr(WholeMoney >= buySum(buy4)))

}
