package main

import "fmt"

type Cop int32   // Кількість цілих копійок
type Hrn float64 // Кількість цілих копійок

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
	fmt.Printf("Щоб купити 9 яблук та 8 груш треба %v грн\n", copToHrn(9*hrnToCop(ApplePrice)+8*hrnToCop(PearPrice)))
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Printf("Ми можемо купити %v груш\n", hrnToCop(WholeMoney)/hrnToCop(PearPrice))
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Printf("Ми можемо купити %v яблук\n", hrnToCop(WholeMoney)/hrnToCop(ApplePrice))
	fmt.Printf("Чи ми можемо купити 2 груші та 2 яблука? %s!\n", boolToUkrStr(hrnToCop(WholeMoney) > 2*hrnToCop(PearPrice)+2*hrnToCop(ApplePrice)))
}
