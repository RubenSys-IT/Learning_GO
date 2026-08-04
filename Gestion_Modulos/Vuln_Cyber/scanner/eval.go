package scanner

type Servicio struct {
	Nombre  string
	Puerto  int
	Abierto bool
}

func EvaluarServicios(servicios []Servicio, alertas *int) map[string]string {
	estadosServicios := make(map[string]string)

	for _, s := range servicios {
		if (s.Puerto == 21 || s.Puerto == 23 || s.Puerto == 80) && s.Abierto {
			*alertas++
			estadosServicios[s.Nombre] = "RIESGO ALTO"
		} else {
			estadosServicios[s.Nombre] = "OK"
		}
	}

	return estadosServicios
}
