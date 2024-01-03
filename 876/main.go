package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) *ListNode {
	current := head
	for current != nil {
		fmt.Printf("# = %v\n", current)
		current = current.Next
	}
	return head
}
func middleNode(head *ListNode) *ListNode {
	current := head      //делаем первый элемент текущим
	middle := head       //делаем первый элемент началом результирующего половинчатого списка
	ind := 1             //индекс, который будем инкрементировать каждую итерацию
	for current != nil { //Выходим из цикла если достигли конца списка
		if ind&1 == 0 { //Если шаг четный, то:
			middle = middle.Next //начало результирующего списка делаем следующий элемент
		} //текущего результирующего списка
		current = current.Next //берем следующий элемент входного списка
		ind++                  //инкрементируем индекс
	}
	return middle //возвращаем результирующий список из функции
}

func main() {
	len := 0                   //счетчик элементов списка
	lst := &ListNode{}         //создаем объект "Односвязный список" (в виде адреса его первого элемента)
	current := lst             //делаем первый элемент - текущим
	cnt := 13                  //количество заполняемых элементов
	for i := 0; i < cnt; i++ { //Заполняем значениями связный список
		current.Val = i + 1 //Присваиваем значение текущему элементу списка
		if i < (cnt - 1) {  //Если номер элемента меньше заданного количества элементов в списке,
			current.Next = &ListNode{} //то создаем следующий элемент списка и записываем его адрес в текущий элемент.
		}
		current = current.Next //Делаем новый созданый элемент - текущим элементом
		len++                  //инкрементируем счетчик элементов
	}

	current = lst                 //Делаем первый элемент списка - текущим
	fmt.Printf("len = %d\n", len) //Выводим количество элементов в списке
	for i := 0; i < (len); i++ {  //Выводим все значения списка:
		fmt.Printf("index = %v, ", i)         //Выводим индекс и
		fmt.Printf("Val = %d\n", current.Val) //Выводим значение текущего элемента
		current = current.Next                //Делаем следующий элемент текущим
	} //Повторяем до конца списка
	fmt.Println("--Исходный односвязный список--")
	printList(lst) //Распечатать весь список
	fmt.Println("-----Решение-----")
	mN := middleNode(lst) // Делаем задачу
	printList(mN)         //Распечатываем решение
}
