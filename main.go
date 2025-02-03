package main

import (
 "bufio"
 "fmt"
 "os"
 "regexp"
 "strconv"
)

func main() {
 scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Введите выражение: ")
 var input string
 scanner.Scan()
 input = scanner.Text()
 re := regexp.MustCompile(`^"([^"]+)"\s*([+\-*\/])\s*("([^"]+)"|(\d+))$`)
 matches := re.FindStringSubmatch(input)
 if matches == nil {
  panic("Ошибка: Неверный формат ввода.")
 }

 a := matches[1]
 operator := matches[2]
 b := matches[4]
 numstr := matches[5]

 var result string

 switch operator {
 case "+":
  result = a + b
 case "-":
  result = computation.Strminus(a, b)
 case "*":
  num, err := strconv.Atoi(numstr)
  if err != nil  num < 1  num > 10 {
   panic("Число должно быть от 1 до 10 включительно.")
  }
  result = computation.Umnostr(a, num)
 case "/":
  num, err := strconv.Atoi(numstr)
  if err != nil  num < 1  num > 10 {
   panic("Число должно быть от 1 до 10 включительно.")
  }
  result = computation.Delstr(a, num)
 default:
  panic("Недопустимый оператор. Используйте +, -, *, /.")
 }

 if len(result) > 40 {
  result = result[:40] + "..."
 }

 fmt.Printf("Результат: %q\n", result)
}
