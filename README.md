# Requerimientos

1. Tener instalado correctamente NodeJs y TypeScript.
2. Tener instalado correctamente Golang usando gomodules y no GOPATH.
3. Tener instalado MySQL y MongoDB.
4. Tener una base de datos en MongoDB o MySQL llamada `mydatabase` para el servicio de productos.
5. Tener una base de datos en MySQL llamada `orders` para el servicio de pedidos.
6. Asegurar que MySQL y MongoDB estén corriendo en los puertos predefinidos.
7. Tener libre los puertos 3000, 3001 y 8080.

# Instalación

Sigue estos pasos para instalar y ejecutar el proyecto:

1. Clonar el proyecto: `git clone https://github.com/Max-500/soa-microservices-practice.git`
2. Ejecutar `npm install`.
3. Ejecutar `npm run dev`.
4. Ejecutar `npm run dev-order`.
5. Navegar al directorio del gestor de productos: `cd src/ProductManagment`.
6. Ejecutar `go mod tidy`.
7. Navegar al directorio src: `cd src`.
8. Ejecutar `go run main.go`.