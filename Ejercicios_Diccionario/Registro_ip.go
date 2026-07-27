package main

import "fmt"

func main() {
	// Declaración e inicialización limpia del map
	servidores := map[string]string{
		"firewall": "192.168.1.1",
		"proxmox":  "192.168.1.10",
		"ansible":  "192.168.1.29",
	}

	// 1. Imprimir la IP de un servidor específico
	fmt.Printf("La IP del firewall es: %s\n", servidores["firewall"])

	// 2. Comprobar si existe la clave "vpn"
	busqueda := "vpn"
	ip, ok := servidores[busqueda] // Usamos 'ip' para no sobrescribir el map

	if ok {
		fmt.Printf("Existe %s con IP %s\n", busqueda, ip)
	} else {
		fmt.Printf("No se encontró el servidor %s\n", busqueda)
	}
}
