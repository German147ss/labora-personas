# Etapa de construcción: Usa la imagen base de Go 
FROM golang:1.22 AS builder

# Configure the Go proxy (if needed)
ENV GOPROXY=direct

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al contenedor
COPY . .

# Compila la aplicación
RUN go build -o main .

# Etapa de ejecución: Usar una imagen alpine para la etapa de ejecución por su tamaño reducido
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el ejecutable de la etapa de construcción al contenedor final
COPY --from=builder /app/main .

# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
