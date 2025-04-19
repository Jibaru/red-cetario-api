package routes

import (
	"redcetarioapi/controllers"

	"github.com/gin-gonic/gin"
)

// Setup routes
func Setup(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/clientes", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/recetas", controllers.GetRecetas)
		api.GET("/receta/:id", controllers.GetReceta)
		api.POST("/recetas/:id/comentario", controllers.PostComentario)

		api.GET("/notificaciones", controllers.GetNotificaciones)
		api.PUT("/notificaciones/:id/fecha-visto", controllers.UpdateNotificacion)
		api.DELETE("/notificacion/:id", controllers.DeleteNotificacion)
		api.DELETE("/notificaciones/cliente/:id", controllers.DeleteNotificacionesByCliente)

		api.PUT("/clientes/:id", controllers.UpdateCliente)
	}
}
