package main

import "fmt"

func aplicaRebaja(stock []int) {
	for i, v := range stock {
		stock[i] = v / 2
	}
}

func main() {
	stock := []int{50, 30, 20, 100, 80}
	copiaStock := make([]int, len(stock))
	copy(copiaStock, stock)

	// Cambiamos [3:4] por [3:] o [3:5]
	oferta := stock[3:]

	aplicaRebaja(oferta)

	fmt.Println("Original:    ", stock)       // [50 30 20 50 40]
	fmt.Println("Vista oferta:", oferta)      // [50 40]
	fmt.Println("Copia intacta:", copiaStock) // [50 30 20 100 80]
}
