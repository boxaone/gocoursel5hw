package main

import (
	"fmt"
	"strings"
)

type Cop int     // Тип копійок
type Hrn float64 // Тип гривень
type Buy struct {
	Apples      int
	Pears       int
	Question    string
	Answer      string
	NeedBoolean bool
}

// Грошові константи
const (
	ApplePrice Hrn = 5.99  // Ціна яблук
	PearPrice  Hrn = 7.00  // Ціна груш
	WholeMoney Hrn = 23.00 // Грошей взагалі в наявності
)

// Змінні покупок
var (
	buys = []*Buy{
		{Apples: 9, Pears: 8, Question: "Скільки грошей треба витратити, щоб купити %s?", Answer: "Щоб купити %s треба %s грн"},
		{Pears: int(hrnToCop(WholeMoney) / hrnToCop(PearPrice)), Question: "Скільки груш ми можемо купити?", Answer: "Ми можемо купити %s"},
		{Apples: int(hrnToCop(WholeMoney) / hrnToCop(ApplePrice)), Question: "Скільки яблук ми можемо купити?", Answer: "Ми можемо купити %s"},
		{Pears: 2, Apples: 2, Question: "Чи ми можемо купити %s?", Answer: "%s", NeedBoolean: true}}
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

// Фікс для видалення додаткових параметрів
func extraParam(s string, pc int) string {
	var ep string
	count := strings.Count(s, "%s") + strings.Count(s, "%v")

	for i := count; i < pc; i++ {
		ep = ep + "%.0s"
	}

	return ep
}

func main() {

	fmt.Printf("Ми маємо: %v грн взагалі\n", WholeMoney)
	for _, buy := range buys {
		fmt.Printf(buy.Question+extraParam(buy.Question, 1)+"\n", buyToS(buy))
		var bs = buySum(buy)
		var bss = fmt.Sprintf("%v", bs)
		if buy.NeedBoolean {
			fmt.Printf(buy.Answer+"\n", boolToUkrStr(WholeMoney >= bs))
		} else {
			fmt.Printf(buy.Answer+extraParam(buy.Answer, 2)+"\n", buyToS(buy), bss)
		}
	}
}
