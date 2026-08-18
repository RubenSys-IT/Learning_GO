package structure

type PortResult struct {
	Port    int
	IsOpen  bool
	Service string
	Risk    string
}

// Array de tamaño fijo
var cono = [3]int{22, 80, 443}

// 2. FUNCIÓN DE ANÁLISIS
func AnalyzePort(target *PortResult) {
	// Caso 1: Si está cerrado
	if !target.IsOpen {
		target.Risk = "None"
		target.Service = "None"
		return
	}

	// Creación de la vista / sub-slice (primeros 2 elementos)
	vistaPuertos := cono[0:2]

	// Comprobamos si el puerto está en la vista rápida
	esCritico := false
	for _, puerto := range vistaPuertos {
		if puerto == target.Port {
			esCritico = true
			break
		}
	}

	// Evaluamos con un switch
	switch {
	case esCritico:
		target.Service = "Critical System"
		target.Risk = "High"
	default:
		target.Service = "Generic Service"
		target.Risk = "Low"
	}
}
