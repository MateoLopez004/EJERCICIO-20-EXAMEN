# 🧩 Ejercicio 20 — Migración de algoritmo supervisada (HS256 → RS256)

## 🧠 Descripción

Servicio HTTP que:

- Emite tokens **JWT** con algoritmo **RS256** (`/login`)
- Valida tokens aceptando **HS256** y **RS256** durante la transición (`/verify`)
- Mide adopción por algoritmo (`/metrics`)
- Permite cortar **HS256** al finalizar la migración

---

## ⚙️ Requisitos

- **Go 1.21+** (recomendado: 1.22)
- **Claves RSA** (`private.pem` y `public.pem`)
- **(Windows)** PowerShell o Git Bash con OpenSSL

> 📝 Nota:  
> Este proyecto usa el módulo `Ejercicio_20`.  
> Si cambias el nombre del folder o del módulo, ajusta los **imports** o el `module` en `go.mod`.

---

## 📁 Estructura del proyecto

```bash
Ejercicio_20/
├─ cmd/server/main.go
├─ internal/config/config.go
├─ internal/jwt/jwt.go
├─ internal/metrics/metrics.go
├─ go.mod
├─ private.pem         # (se genera)
└─ public.pem          # (se genera)
```
---
## 🧩 1) Instalar dependencias
Desde la raíz del proyecto:
```bash
go mod tidy
```
---
## 🔐 2) Generar claves RSA (RS256)
Si no existen los archivos private.pem y public.pem, puedes generarlos así:
## 🅰️ Opción A — Git Bash (recomendada)
```bash
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```
## 🅱️ Opción B — PowerShell (con OpenSSL instalado)
```bash
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```
---
🚀 3) Ejecutar el servidor
Desde la raíz del proyecto (donde están los .pem):
```bash
go run ./cmd/server
```
Salida esperada: 
```bash
Servidor corriendo en :8080
```
Deja esta terminal abierta mientras realizas las pruebas.
El servidor seguirá escuchando peticiones.

---
## 🧪 4) Pruebas reales paso a paso
4.1 Obtener un token RS256 (nuevo login)
PowerShell
```bash
$t = Invoke-RestMethod http://localhost:8080/login
$t
```
Git Bash / CMD (curl)
```bash
curl -s http://localhost:8080/login
```

Ejemplo de salida (token JWT):
```bash
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...
```
## 4.2 Verificar el token

El endpoint /verify espera el token en el header Authorization (solo el token, sin “Bearer ”).

PowerShell
```bash
Invoke-RestMethod http://localhost:8080/verify -Headers @{ Authorization = $t }
```
Git Bash / CMD
```bash
curl -H "Authorization: <TU_TOKEN_AQUI>" http://localhost:8080/verify
```
Respuesta esperada:
```bash
Token válido ✅
```
---
 ## 4.3 Ver métricas de adopción
PowerShell
``` bash
Invoke-RestMethod http://localhost:8080/metrics
```
Git Bash / CMD
```bash
curl http://localhost:8080/metrics
```

Ejemplo de salida:
```bash
{"HS256_tokens": 0, "RS256_tokens": 1}
```
Cada vez que verificas un token, se incrementa el contador del algoritmo correspondiente
y se registra un log en consola ([HS256] o [RS256]).
