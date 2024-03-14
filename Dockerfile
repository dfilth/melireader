# Establecer la imagen base de Go. Especifica la versión de Go que necesitas.
# Por ejemplo, "1.18" puede ser reemplazado por la versión específica que utilizas.
FROM golang:1.22.1 as builder


# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum para descargar las dependencias.
# Esto aprovecha la caché de capas de Docker para descargar dependencias solo si cambian.
COPY go.mod go.sum ./

# Descargar las dependencias del proyecto.
RUN go mod download

# Copiar el resto del código fuente del proyecto al contenedor.
COPY . .

# Compilar la aplicación. Reemplaza "main.go" con el camino y nombre de tu archivo principal si es necesario.
# También, ajusta el nombre del binario de salida según prefieras.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/http/main.go

# Utilizar una imagen Docker scratch como imagen base para la imagen final. Esto hace la imagen lo más pequeña posible.
# Si necesitas una shell o herramientas adicionales en tu contenedor, considera usar "alpine" en lugar de "scratch".
FROM alpine:latest

# Copiar el binario compilado desde la imagen de construcción a la imagen final.
COPY --from=builder /app/myapp .

# Puerto que tu aplicación usará. Asegúrate de exponer el mismo puerto que tu aplicación escucha.
EXPOSE 8080

# Comando para ejecutar la aplicación.
CMD ["./myapp"]
