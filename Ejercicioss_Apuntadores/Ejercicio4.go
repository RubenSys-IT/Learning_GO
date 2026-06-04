package main

import "fmt"

func main() {
	firewallActivo := true
	ptrFw := &firewallActivo

	//Valor modificado
	*ptrFw = false

	if firewallActivo == true {
		fmt.Println(" Firewall activado")
	} else {
		fmt.Println("Firewall desactivado")
	}
}
