package session

func PurgarExpiradas(sesiones map[string]bool, eliminadas *int) {
	for token, activa := range sesiones {
		if !activa { // Equivalente a: if activa == false
			delete(sesiones, token) // Elimina la clave del mapa directamente
			*eliminadas++           // Incrementa el contador pasado por puntero
		}
	}
}
