package main

import "fmt"

func usuariosUnicos(login []string) []string {
	// PASO A: Crear el mapa que funcionará como Set
	conjunto := map[string]bool{}

	// PASO B: Llenar el mapa recorriendo 'logs'.
	for _, ip := range login {
		conjunto[ip] = true
	}

	// PASO C: Crear un slice vacío para guardar los logins únicas
	var resultado []string

	// PASO D: Recorrer el mapa con 'range' y añadir (append) cada clave al slice 'resultado'
	for ip := range conjunto {
		resultado = append(resultado, ip)
	}

	return resultado
}

func main() {
	intentosLogin := []string{"admin", "root", "admin", "invitado", "root", "pepe"}

	// 1. Guardamos el resultado de la función en la variable 'unicos'
	unicos := usuariosUnicos(intentosLogin)

	// 2. Usamos fmt.Printf con %v para imprimir el slice resultado
	fmt.Printf("Los usuarios únicos son: %v\n", unicos)

	// 3. (Opcional) Podemos ver la cantidad usando len()
	fmt.Printf("Total de usuarios distintos: %d\n", len(unicos))
}
