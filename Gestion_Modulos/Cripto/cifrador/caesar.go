package cifrador

func OfuscarPayload(payload *string, clave int) {
	textobyte := []byte(*payload)

	for i := 0; i < len(textobyte); i++ {
		textobyte[i] = textobyte[i] + byte(clave)

	}
	*payload = string(textobyte)
}
