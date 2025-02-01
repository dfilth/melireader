# MeliFileReader

# Go MELI ITEMS FILE READER API

## Descripción
Este es un servicio API RESTFUL simple para la lectura de archivos de ítems MELI, escrito en Go utilizando el framework web Gin y una base de datos MongoDB, implementando la arquitectura hexagonal.

## Contacto
**Desarrollador:** Ing. Daniel Torres  
**GitHub:** [https://github.com/dfilth/melireader)

## Versiones
- **Versión API:** 1.0

## Uso
La API está disponible en el siguiente endpoint base:

localhost:8080/v1/

## ENVS
Añadir archivo .env en la raíz del proyecto con lo siguiente

```bash
APP_NAME="file-reader"
APP_ENV="development"

HTTP_URL="0.0.0.0"
HTTP_PORT="8080"
HTTP_ALLOWED_ORIGINS="http://localhost:3000,http://localhost:5173"

DB_CONNECTION="mongodb://root:secret@mongodb:27017/meli?authSource=admin"
DB_HOST="mongodb"
DB_PORT="27017"
DB_NAME="meli"
DB_USER="root"
DB_PASSWORD=secret

DOCUMENT_DB_CONNECTION="mongodb"
DOCUMENT_DB_NAME="meli"
```

## Esquemas soportados
- HTTP
- HTTPS

## Ejecución con Docker
Para ejecutar el proyecto mediante Docker, sigue estos pasos:

Ejecuta el siguiente comando para levantar los contenedores de MongoDB y la aplicación:
```bash
docker compose up -d
```
Comprueba que los contenedores estén activos:
```bash
docker ps
```
Luego, utiliza el siguiente endpoint para subir un archivo:
```bash
curl --request POST \
  --url http://localhost:8080/v1/items/file-upload \
  --header 'Content-Type: multipart/form-data' \
  --header 'User-Agent: insomnia/8.6.1' \
  --form file=@/home/deicide/Downloads/test.csv
```

Si deseas consultar los registros, utiliza el siguiente endpoint pasando la página y el tamaño por página como parámetros:
```bash
curl --request GET \
  --url 'http://localhost:8080/v1/items?page=1&pageSize=10'
```

Opcionalmente, puedes ver en la consola el estado del servicio:
```bash
docker logs -f melireader-app-1
```

Tener en cuenta que podemos conectarnos al mongo a traves del gestor de BD
```bash
con las credenciales: usr: root, pwd: secret
url: mongodb://localhost:27017/meli?authSource=admin
```

hola