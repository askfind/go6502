package main

import (
	"fmt"
)

// Таб.1 Алфавит троичной симметричной системы счисления
// +--------------------------+-------+-------+-------+
// | Символы 	              |   -   |	  0   |	  +   |
// +--------------------------+-------+-------+-------+
// | Числа                    |  -1   |   0   |   1   |
// +--------------------------+-------+-------+-------+
// | Логика                   | false |  nil  | true  |
// +--------------------------+-------+-------+-------+

// Объявление троичных типов
type trit interface{}
type tryte [6]interface{}

// Преобразование трита в целое число
func trit2int(t trit) int8 {
	if t == true {
		return 1
	}
	if t == false {
		return -1
	}
	return 0
}

// Преобразование целое число в трит
func int2trit(i int8) trit {
	if i > 0 {
		return true
	}
	if i < 0 {
		return false
	}
	return nil
}

// Таб.2 Троичное сложение
// .------------------------.
// |     | -1  |  0  |  1   |
// |------------------------|
// | -1  | -+  |  -  |  0   |
// |------------------------|
// |  0  | -1  |  0  |  1   |
// |------------------------|
// |  1  |  0  |  1  |  +-  |
// .------------------------.

// Троичное сложение двух тритов с переносом
func add_t(a trit, b trit) (c trit, carry trit) {

	if a == false && b == false {
		return true, false
	} else if a == false && b == nil {
		return false, nil
	} else if a == false && b == true {
		return nil, nil
	} else if a == nil && b == false {
		return false, nil
	} else if a == nil && b == nil {
		return nil, nil
	} else if a == nil && b == true {
		return true, nil
	} else if a == true && b == false {
		return nil, nil
	} else if a == true && b == nil {
		return true, nil
	} else if a == true && b == true {
		return false, true
	}
	return nil, nil
}

// Таб.3 Троичное умножение
// .-----------------------.
// |     | -1  |  0  |  1  |
// |-----------------------|
// | -1  |  +  |  0  |  -  |
// |-----------------------|
// |  0  |  0  |  0  |  0  |
// |-----------------------|
// |  1  |  -  |  0  |  +  |
// .-----------------------.
// Троичное сложение двух тритов с переносом
func mul_t(a trit, b trit) trit {

	if a == false && b == false {
		return true
	} else if a == false && b == true {
		return false
	} else if a == true && b == false {
		return false
	} else if a == true && b == true {
		return true
	}
	return nil
}

// ----------
// Main
// ----------
func main() {

	fmt.Printf("Run main() ---------------------\n")

	// Троичные переменные
	//var t trit
	//var tr tryte
	//var s, carry trit


	fmt.Printf("--------------------------------\n")
}
