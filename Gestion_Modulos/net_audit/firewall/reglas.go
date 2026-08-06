package firewall

type Regla struct {
	Puerto int
	Activa bool
}

func DesactivarPuerto(reglas *[]Regla, puertoObjetivo int, cambios *int) {
	// Recorremos el slice usando la longitud del slice desreferenciado (*reglas)
	for i := 0; i < len(*reglas); i++ {

		// Comprobamos si el puerto coincide Y la regla está activa
		if (*reglas)[i].Puerto == puertoObjetivo && (*reglas)[i].Activa {
			// Modificamos el valor directamente en memoria
			(*reglas)[i].Activa = false

			// Incrementamos el contador externo a través de su puntero
			*cambios++
		}
	}
}
