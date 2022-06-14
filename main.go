package main

import "fmt"

type Cop int32 // Кількість цілих копійок

const (
	ApplePrice = 5.99  // Ціна яблук
	PearPrice  = 7.00  // Ціна груш
	WholeMoney = 23.00 // Грошей взагалі в наявності
)

// Переводимо гривні(раціональні float64) в копійки(цілі)
func fToCop(grn float64) Cop {
	return (Cop)(grn * 100)
}

// Переводимо копійки(цілі) в гривні(раціональні float64)
func copToF(cop Cop) float64 {
	return (float64)(cop) / 100
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
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Printf("Щоб купити 9 яблук та 8 груш треба %v грн\n", copToF(fToCop(PearPrice)*9+fToCop(ApplePrice)*8))
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Printf("Ми можемо купити %v груш\n", fToCop(WholeMoney)/fToCop(PearPrice))
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Printf("Ми можемо купити %v яблук\n", fToCop(WholeMoney)/fToCop(ApplePrice))
	fmt.Printf("Чи ми можемо купити 2 груші та 2 яблука? %s!\n", boolToUkrStr(fToCop(WholeMoney) > 2*fToCop(ApplePrice)+2*fToCop(PearPrice)))
}
