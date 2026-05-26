package main

import "fmt"

func main() {
	protocolo := "HTTP"
	switch protocolo {
	case "FTP", "SCP", "SFTP":
		fmt.Println("Tráfico de transferencia de archivos")
	case "HTTP", "HTTPS":
		fmt.Println("Conexión WEB")
	case "RDP":
		fmt.Println("RDP")
	default:
		fmt.Println("No existe")
	}
}
