4.3 Ver métricas de adopción
PowerShell
Invoke-RestMethod http://localhost:8080/metrics

Git Bash / CMD
curl http://localhost:8080/metrics


Ejemplo de salida:

{"HS256_tokens": 0, "RS256_tokens": 1}


Cada vez que verificas un token, se incrementa el contador del algoritmo correspondiente
y se registra un log en consola ([HS256] o [RS256]).
