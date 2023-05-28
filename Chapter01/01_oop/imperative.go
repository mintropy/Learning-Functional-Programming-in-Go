package main

import "fmt"

func imperative() {
	// 슬라이스 자료형이 주어졌을 때 탐색하는 명령형 코드
	var found bool
	carToLookFor := "Blazer"
	cars := []string{"Accord", "IS250", "Blazer"}
	for _, car := range cars {
		if car == carToLookFor {
			found = true
		}
	}
	fmt.Printf("Found? %v", found)
}

func functional() {
	// 슬라이스 자료형에서 탐색을 하는 함수형 프로그래밍
	// contains라는 메서드를 사용한다면 쉽게 탐색할 수 있음
	// cars := []string{"Accord", "IS250", "Blazer"}
	// fmt.Printf("Found? %v", cars.contains("Blazer"))
}

func oop() {
    // 주어진 자료 중 탐색하는 객체지향 프로그래밍
	type Car struct {
		Model string
	}
	accord := &Car{"Accord"}
	is250 := &Car{"is250"}
	blazer := &Car{"blazer"}
	cars := []*Car{is250, accord, blazer}

	var found bool
	carToLookFor := blazer
	for _, car := range cars {
		if car == carToLookFor {
			found = true
		}
	}
	fmt.Printf("Found? %v", found)
}

func main() {
	imperative()
	fmt.Println()
	oop()
}
