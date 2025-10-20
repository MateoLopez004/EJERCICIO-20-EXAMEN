package main

import (
	"encoding/json"
	"log"
	"net/http"

	"Ejercicio_20/internal/config"
	"Ejercicio_20/internal/jwt"
	"Ejercicio_20/internal/metrics"
)

func main() {
	config.LoadKeys()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.EmitToken("usuario123")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte(token))
	})

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		tok := r.Header.Get("Authorization")
		if tok == "" {
			http.Error(w, "Falta token", 400)
			return
		}
		token, err := jwt.ParseToken(tok)
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido: "+err.Error(), 401)
			return
		}
		w.Write([]byte("Token válido ✅"))
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]int64{
			"HS256_tokens": metrics.HS256Count,
			"RS256_tokens": metrics.RS256Count,
		})
	})

	log.Println("Servidor corriendo en :8080")
	http.ListenAndServe(":8080", nil)
}
