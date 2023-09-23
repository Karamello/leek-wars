package main

import (
	"fmt"
	"sync"

	"github.com/Karamello/leek-wars/leekwars"
)

func main() {

	lw := leekwars.New("https://leekwars.com/api/")

	leeks := lw.GetLeekOpponents(99838)
	var wg sync.WaitGroup

	for _, leek := range leeks {
		wg.Add(1)
		leek := leek
		go func() {
			defer wg.Done()
			detailed := lw.GetLeekInfo(leek.ID)
			fmt.Println(fmt.Sprintf("Name: %s, Level: %d, Talent: %d", leek.Name, leek.Level, leek.Talent))
			info := fmt.Sprintf("HP: %d, Strength: %d, Resistence: %d, Magic %d", detailed.Life, detailed.Strength, detailed.Resistance, detailed.Magic)
			fmt.Println(info)
		}()

	}
	wg.Wait()

}
