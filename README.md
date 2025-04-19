# Red Cetario API - Go Clone

Esta es una implementación en Go de la API Red Cetario, usando **Gin** para el enrutamiento y **Gorm** para ORM.

## Requisitos

- Go 1.20+
- MySQL

## Instalación

1. Descarga o clona este repositorio.
2. Ajusta el DSN en `main.go` (usuario, contraseña, host y nombre de base de datos).
3. Ejecuta:
   ```bash
   go mod download
   go run main.go
   ```
4. El servidor correrá en `http://localhost:8080`.

## Endpoints

- `POST /api/clientes` – Registrar nuevo cliente  
- `POST /api/login` – Autenticación  
- `GET /api/recetas` – Listar recetas  
- `GET /api/receta/:id` – Detalles de una receta  
- `POST /api/recetas/:id/comentario` – Agregar comentario  
- `GET /api/notificaciones` – Listar notificaciones  
- `PUT /api/notificaciones/:id/fecha-visto` – Marcar como visto  
- `DELETE /api/notificacion/:id` – Eliminar notificación  
- `DELETE /api/notificaciones/cliente/:id` – Eliminar por cliente  
- `PUT /api/clientes/:id` – Actualizar cliente  

## Estructura del proyecto

- `main.go` – Inicialización y arranque  
- `models/` – Definición de modelos  
- `controllers/` – Lógica de negocio  
- `routes/` – Configuración de rutas  
