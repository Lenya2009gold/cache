package main

import (
	"awesomeProject5/internal/cache"
	"fmt"
	"time"
)

func main() {
	lruCache := cache.New(3)

	lruCache.Set("key1", "value1")
	lruCache.Set("key2", "value2")
	lruCache.Set("key3", "value3")

	time.Sleep(2 * time.Second)
	lruCache.Remove("key1")
	if value, found := lruCache.Get("key1"); found {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 не найден")
	}

	if value, found := lruCache.Get("key2"); found {
		fmt.Println("key2:", value)
	} else {
		fmt.Println("key2 не найден")
	}

	// Добавляем новый элемент, что вызовет удаление наименее используемого элемента
	lruCache.Set("key4", "value4")

	// Попробуем получить старые элементы
	if value, found := lruCache.Get("key3"); found {
		fmt.Println("key3:", value)
	} else {
		fmt.Println("key3 не найден")
	}

}
