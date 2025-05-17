package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TheTitanrain/w32"
	"github.com/micmonay/keybd_event"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	workLoop()
}

func workLoop() {
	for {
		sessionDuration := time.Duration(rand.Intn(120)+60) * time.Second
		sessionEnd := time.Now().Add(sessionDuration)

		fmt.Println("🟢 Начало", sessionDuration)

		for time.Now().Before(sessionEnd) {
			doRandomActivity()
			time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
		}

		breakTime := time.Duration(rand.Intn(60)+30) * time.Second
		fmt.Println("🟡 Перерыв на", breakTime)
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
	}
}

func moveMouseRandom() {
	x := rand.Intn(800) + 100
	y := rand.Intn(600) + 100
	w32.SetCursorPos(x, y)
	fmt.Println("м:", x, y)
}

func typeRandomKey() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println("keyboard error:", err)
		return
	}

	keys := []int{
		keybd_event.VK_A,
		keybd_event.VK_S,
		keybd_event.VK_D,
		keybd_event.VK_F,
		keybd_event.VK_SPACE,
		keybd_event.VK_TAB,
		keybd_event.VK_ENTER,
	}
	k := keys[rand.Intn(len(keys))]
	kb.SetKeys(k)

	if err := kb.Launching(); err != nil {
		fmt.Println("key press error:", err)
	}
	fmt.Println("⌨", k)
}
