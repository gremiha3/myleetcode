package main

import (
	"fmt"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	dateLayout := "2006-01-02"               //устанавливаем шаблон для парвера даты
	d1, err := time.Parse(dateLayout, date1) //дата1 из "строки" в тип "дата"
	if err != nil {                          //если неверный формат даты - говорим об этом
		fmt.Printf("Неверная дата1 - %v\n", err)
	}
	d2, err := time.Parse(dateLayout, date2) //дата2 из "строки" в тип "дата"
	if err != nil {                          //если неверный формат даты - говорим об этом.
		fmt.Printf("Неверная дата2 - %v\n", err)
	}
	delta := int((d2.Unix() - d1.Unix()) / 3600 / 24) //вычисляем разницу в датах переведя
	//даты в формат Unix (секунды от 1970го года) и, затем, переведя секунды в Дни.
	//Результат приводим к типу int32 из int64
	if delta < 0 { //если разница отрицательная - исправляем это
		delta *= -1
	}
	return delta
}
func main() {
	date1 := "2020-01-15"
	date2 := "2019-12-31"
	fmt.Printf("Разница во времени - %d\n", daysBetweenDates(date1, date2))

}
