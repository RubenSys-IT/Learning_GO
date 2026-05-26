//Este es un selector de canales, a lo TV

package main

import "fmt"

func main() {
	canal := 1
	switch canal {
	case 1:
		fmt.Println("Viendo noticias")
	case 2:
		fmt.Println("Viendo Deportes")
	case 3:
		fmt.Println("Viendo Películas")
	default:
		fmt.Println("No estas viendo nada.")
	}
}
