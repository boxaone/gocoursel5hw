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

func main() {

	fmt.Printf("Ми маємо: %v грн взагалі\n", WholeMoney)
	buy1 := Buy{Apples: 9, Pears: 8}
	fmt.Printf("Скільки грошей треба витратити, щоб купити %v яблук та %v груш?", buy1.Apples, buy1.Pears)
	fmt.Printf("Щоб купити %v яблук та %v груш треба %v грн\n", buy1.Apples, buy1.Pears, buySum(&buy1))
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Printf("Ми можемо купити %v груш\n", hrnToCop(WholeMoney)/hrnToCop(PearPrice))
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Printf("Ми можемо купити %v яблук\n", hrnToCop(WholeMoney)/hrnToCop(ApplePrice))
	buy4 := Buy{Pears: 2, Apples: 2}
	fmt.Printf("Чи ми можемо купити %v груші та %v яблука? %s!\n", buy4.Pears, buy4.Apples, boolToUkrStr(WholeMoney >= buySum(&buy4)))
}
