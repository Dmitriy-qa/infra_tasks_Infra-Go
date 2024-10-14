package main

import (
	"sync"
	"testing"
)

func TestRandPhone(t *testing.T) {
	// Проверка диапазона сгенерированных номеров телефонов
	for i := 0; i < 1000; i++ {
		phone := randPhone()
		if phone < 89000000000 || phone > 89999999999 {
			t.Errorf("randPhone() = %d; want a phone number in range [89000000000, 89999999999]", phone)
		}
	}
}

func TestGenerate(t *testing.T) {
	td := testData{}
	var wg sync.WaitGroup
	n := 100

	generate(n, &td, &wg)
	wg.Wait()

	// Проверка, что количество сгенерированных телефонов соответствует ожидаемому
	if len(td.phones) != n {
		t.Errorf("generate() = %d; want %d phones", len(td.phones), n)
	}

	// Дополнительно можно проверить уникальность сгенерированных телефонов
	phoneSet := make(map[int]struct{})
	for _, phone := range td.phones {
		phoneSet[phone] = struct{}{}
	}

	if len(phoneSet) != len(td.phones) {
		t.Error("generate() produced duplicate phone numbers")
	}
}
