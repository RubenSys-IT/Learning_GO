package checker

import "os"

type AuditTarget struct {
	Path         string
	ExpectedPerm string
	IsCritical   bool
	Status       string // "OK", "MISSING", "RISK_EXPOSED"
	RiskLevel    string // "HIGH", "MEDIUM", "LOW"
}

// Array estático global
var RutasCriticas = [3]string{"/etc/shadow", "/etc/passwd", "/etc/sudoers"}

func AuditSystem(target *AuditTarget) {
	// 1. Creación de la vista / sub-slice (primeros 2 elementos)
	vistaCriticas := RutasCriticas[0:2]

	// Comprobamos si la ruta está en la vista rápida
	esMuyCritico := false
	for _, ruta := range vistaCriticas {
		if target.Path == ruta {
			esMuyCritico = true
			break
		}
	}

	// 2. Comprobación física con os.Stat
	_, err := os.Stat(target.Path)
	if err != nil {
		target.Status = "MISSING"
	} else {
		target.Status = "OK"
	}

	// 3. Evaluamos el nivel de riesgo
	switch {
	case target.Status == "MISSING" && target.IsCritical:
		target.RiskLevel = "HIGH"
	case esMuyCritico:
		target.RiskLevel = "MEDIUM"
	default:
		target.RiskLevel = "LOW"
	}
}
