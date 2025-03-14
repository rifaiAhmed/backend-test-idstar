package cmd

import (
	"backend-test/helpers"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {

	r := gin.Default()
	r.Use(MiddlewareCORS())

	dependency := dependencyInject()
	r.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	authV1 := r.Group("/auth")
	authV1.POST("/submit-email", dependency.UserAPI.SendMail)
	authV1.GET("/magic-link", dependency.UserAPI.CekSessionByUUID)

	invV1 := r.Group("/inventory")
	invV1.POST("/", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.InsertInv)
	invV1.PUT("/:id", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.UpdateInv)
	invV1.DELETE("/:id", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.DeleteInv)
	invV1.GET("/", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.GetAllInv)

	recipeV1 := r.Group("/recipe")
	recipeV1.POST("/", dependency.MiddlewareValidateAuth, dependency.RicipeAPI.InsertRecipe)
	recipeV1.DELETE("/:id", dependency.MiddlewareValidateAuth, dependency.RicipeAPI.DeleteRecipe)
	recipeV1.GET("/", dependency.MiddlewareValidateAuth, dependency.RicipeAPI.GetAllRecipe)

	ingredientV1 := r.Group("/ingredient")
	ingredientV1.POST("/", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.InsertIngredient)
	ingredientV1.PUT("/:id", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.UpdateIngredient)
	ingredientV1.DELETE("/:id", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.DeleteIngredient)
	ingredientV1.GET("/", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.GetAllIngredient)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

func MiddlewareCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3039"}, // Ganti dengan domain frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
