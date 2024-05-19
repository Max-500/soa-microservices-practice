# Requerimientos

1. Tener instalado correctamente NodeJs, TypeScript y Docker.
2. Tener instalado correctamente Golang usando gomodules y no GOPATH.
3. Tener instalado MySQL y MongoDB.
4. Tener una base de datos en MongoDB o MySQL llamada `mydatabase` para el servicio de productos.
5. Tener una base de datos en MySQL llamada `orders` para el servicio de pedidos.
6. Asegurar que MySQL y MongoDB estén corriendo en los puertos predefinidos.
7. Tener libre los puertos 3000, 3001 y 8080.
8. Ejecutar el siguiente comando `docker pull rabbitmq:3-management`
9. Ejecutar el siguiente comando `docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management`
10. Entrar a la interfaz de usuario de RabbitMq en esta ruta `http://localhost:15672/` con las credencias `user: guest` y `password: guest`
11. Crear 3 colas, todas que sean con el atributo o característica llamada `durable` y con los nombres de `receive_get_products_queue`, `send_get_products_queue`, y`update_stock_queue`.


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