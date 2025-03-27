package main

import (
	"fmt"

	"myproject/magazine"
)

func printInfo(sub *magazine.Subscriber) {
	fmt.Println("Имя:", sub.Name)
	fmt.Println("Цена:", sub.Rate)
	fmt.Println("Активность:", sub.Active)
	fmt.Println("Адрес:", sub.Address)
}

func disc(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func defaultSub(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func main() {
	sub1 := defaultSub("Карина")
	sub1.Address.City = "Москва"
	sub1.Address.Street = "Дмитровское шоссе"
	sub1.Address.Index = 127411
	printInfo(sub1)
	sub2 := defaultSub("Буся")
	disc(sub2)
	sub2.City = "Москва"
	sub2.Street = "ЖК Город"
	sub2.Index = 123456
	printInfo(sub2)

	var emp1 magazine.Employee
	emp1.Name = "Артём"
	emp1.Salary = 500000
	emp1.City = "Москва"
	emp1.Street = "ЖК Город"
	emp1.Index = 555555
	fmt.Println("Имя:", emp1.Name)
	fmt.Println("Зарплата:", emp1.Salary)
	fmt.Println("Адрес:", emp1.Address)
}
