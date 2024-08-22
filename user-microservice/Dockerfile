# Etapa 1: Construcción
FROM gradle:8.2-jdk17 AS builder

# Información del mantenedor
LABEL maintainer="github.com/dieg0code"

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos necesarios para la construcción
COPY gradlew .
COPY gradle gradle
COPY build.gradle .
COPY settings.gradle .
COPY src src

# Construir la aplicación
RUN ./gradlew build --no-daemon

# Etapa 2: Imagen ligera para la ejecución
FROM openjdk:17-jdk-slim

# Información del mantenedor
LABEL maintainer="github.com/dieg0code"

# Exponer el puerto de la aplicación
EXPOSE 8080

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar el archivo JAR generado desde la etapa de construcción
COPY --from=builder /app/build/libs/*.jar app.jar

# Ejecutar la aplicación
ENTRYPOINT ["java", "-Djava.security.egd=file:/dev/./urandom", "-jar", "/app/app.jar"]
