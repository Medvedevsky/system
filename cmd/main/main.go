package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	workLoop()
}

func workLoop() {
	for {
		// Эмулируем активность от 1 до 3 минут
		sessionDuration := time.Duration(rand.Intn(120)+60) * time.Second
		sessionEnd := time.Now().Add(sessionDuration)

		println("🟢 Начало", sessionDuration)

		for time.Now().Before(sessionEnd) {
			doRandomActivity()
			time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
		}

		// Перерыв от 30 до 90 секунд
		breakTime := time.Duration(rand.Intn(60)+30) * time.Second
		println("🟡 Перерыв на", breakTime)
		time.Sleep(breakTime)
	}
}

func doRandomActivity() {
	action := rand.Intn(3)

	switch action {
	case 0:
		moveMouseRandom()
	case 1:
		typeRandomKey()
	case 2:
		scrollMouse()
	}
}

func moveMouseRandom() {
	x := rand.Intn(800) + 100
	y := rand.Intn(600) + 100
	robotgo.MoveSmooth(x, y, 0.9, 0.9)
	println("мышь:", x, y)
}

func typeRandomKey() {
	keys := []string{"a", "s", "d", "f", "space", "tab", "enter"}
	k := keys[rand.Intn(len(keys))]
	robotgo.KeyTap(k)
	println("⌨", k)
}

func scrollMouse() {
	scroll := rand.Intn(10) + 1
	direction := 1
	if rand.Intn(2) == 0 {
		direction = -1
	}
	robotgo.Scroll(0, scroll*direction)
	println("sc", scroll*direction)
}
