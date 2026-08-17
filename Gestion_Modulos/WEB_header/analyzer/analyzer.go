package analyzer

import "fmt"

// HeaderTarget representa el objetivo a auditar.
type HeaderTarget struct {
	Domain    string
	Headers   map[string]string
	RiskScore int
	Flags     []string
}

// Array de tamaño fijo con firmas de prueba (Vector)
var firmasConocidas = [4]string{"nginx/1.18.0", "Apache/2.4.1", "IIS/10.0", "Unknown"}

// EvaluateTarget recibe un PUNTERO para modificar el struct original.
func EvaluateTarget(target *HeaderTarget) {
	if target == nil || len(target.Headers) == 0 {
		fmt.Println("[-] Target inválido o sin cabeceras.")
		return
	}

	// 1. CREACIÓN DE VISTA / SUB-SLICE
	// Tomamos una vista/sub-slice de los primeros 2 elementos del array fijo
	firmasTop := firmasConocidas[0:2] // Tipo: []string de longitud 2, capacidad 4

	// 2. REVISIÓN DE CABECERA 'Server'
	valServer, existeServer := target.Headers["Server"]
	if existeServer {
		// Recorremos la vista (sub-slice)
		for _, firma := range firmasTop {
			if valServer == firma {
				target.RiskScore += 2
				// append crea/actualiza el slice interno de Flags
				target.Flags = append(target.Flags, "Server disclosure: "+valServer)
				break
			}
		}
	}

	// 3. SWITCH SIN CONDICIÓN PARA COMPROBAR CABECERAS FALTANTES
	switch {
	case target.Headers["X-Frame-Options"] == "":
		target.RiskScore += 3
		target.Flags = append(target.Flags, "Missing X-Frame-Options")
		fallthrough // Pasa a evaluar la siguiente condición obligatoriamente

	case target.Headers["Strict-Transport-Security"] == "":
		target.RiskScore += 5
		target.Flags = append(target.Flags, "Missing HSTS")
	}
}
