Ejercicio 20 — Migración de algoritmo supervisada (HS256 → RS256)

Servicio HTTP que:

Emite tokens JWT con RS256 (/login)

Valida tokens aceptando HS256 y RS256 durante la transición (/verify)

Mide adopción por algoritmo (/metrics)

Permite cortar HS256 al finalizar la migración

Requisitos

Go 1.21+ (recomendado 1.22)

Claves RSA (private.pem y public.pem)

(Windows) PowerShell o Git Bash

Nota: Este proyecto usa el módulo Ejercicio_20. Si cambias el nombre del folder o del módulo, ajusta los imports o el module en go.mod.

Estructura
Ejercicio_20/
├─ cmd/server/main.go
├─ internal/config/config.go
├─ internal/jwt/jwt.go
├─ internal/metrics/metrics.go
├─ go.mod
├─ private.pem         (se genera)
└─ public.pem          (se genera)
´´´´
1) Instalar dependencias

En la carpeta del proyecto:

go mod tidy

2) Generar claves RSA (RS256) (SI NO ESTAN GENERADAS)


    Opción A (fácil en Windows): Git Bash
    openssl genrsa -out private.pem 2048
    openssl rsa -in private.pem -pubout -out public.pem
 

    Opción B: PowerShell con OpenSSL instalado
    openssl genrsa -out private.pem 2048
    openssl rsa -in private.pem -pubout -out public.pem


Si no tienes OpenSSL, puedes generar claves con un script Go (pídelo si lo necesitas).

3) Ejecutar el servidor


Desde la raíz del proyecto (donde están los .pem):

go run .\cmd\server


Salida esperada:

Servidor corriendo en :8080


Deja esa terminal abierta (el servidor sigue escuchando).


4) Probar en una situación real (paso a paso)

   4.1 Obtener un token RS256 (nuevo login)

PowerShell

$t = Invoke-RestMethod http://localhost:8080/login

$t


curl (Git Bash/CMD)

curl -s http://localhost:8080/login


Verás un JWT (string con 2 puntos), p.ej.:

eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...

4.2 Verificar el token

El endpoint espera el token en el header Authorization (solo el token, sin “Bearer ”).

PowerShell

Invoke-RestMethod http://localhost:8080/verify -Headers @{ Authorization = $t }


curl

curl -H "Authorization: <PEGA_AQUI_TU_TOKEN>" http://localhost:8080/verify


Respuesta esperada:

Token válido ✅

4.3 Ver métricas de adopción

PowerShell

Invoke-RestMethod http://localhost:8080/metrics


curl

curl http://localhost:8080/metrics


Ejemplo de salida:

{"HS256_tokens":0,"RS256_tokens":1}


Cada vez que verificas un token, se incrementa el contador del algoritmo correspondiente y se escribe un log en consola [HS256] o [RS256].