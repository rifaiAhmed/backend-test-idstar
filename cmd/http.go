package cmd

import (
	"backend-test/helpers"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	r.Use(MiddlewareCORS())

	r.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	r.POST("/auth/submit-email", dependency.UserAPI.SendMail)
	r.GET("/auth/magic-link", dependency.UserAPI.CekSessionByUUID)

	// inventory
	r.POST("/inventory", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.InsertInv)
	r.PUT("/inventory/:id", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.UpdateInv)
	r.DELETE("/inventory/:id", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.DeleteInv)
	r.GET("/inventory", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.GetAllInv)
	r.GET("/inventory/template", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.GetTemplate)
	r.POST("/inventory/upload", dependency.MiddlewareValidateAuth, dependency.InventoryAPI.UploadExcel)

	// recipe
	r.POST("/recipe", dependency.MiddlewareValidateAuth, dependency.RecipeAPI.InsertRecipe)
	r.PUT("/recipe/:id", dependency.MiddlewareValidateAuth, dependency.RecipeAPI.UpdateRecipe)
	r.DELETE("/recipe/:id", dependency.MiddlewareValidateAuth, dependency.RecipeAPI.DeleteRecipe)
	r.GET("/recipe", dependency.MiddlewareValidateAuth, dependency.RecipeAPI.GetAllRecipe)

	// ingredient
	r.POST("/ingredient", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.InsertIngredient)
	r.PUT("/ingredient/:id", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.UpdateIngredient)
	r.DELETE("/ingredient/:id", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.DeleteIngredient)
	r.GET("/ingredient", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.GetAllIngredient)
	r.GET("/ingredient/recipe", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.GetRecipeIncludeIngredients)
	r.POST("/ingredient/multiple", dependency.MiddlewareValidateAuth, dependency.IngredientAPI.MultipleCreateUpdate)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}
