package cmd

import (
	"backend-test/helpers"
	"backend-test/internal/api"
	"backend-test/internal/interfaces"
	"backend-test/internal/repository"
	"backend-test/internal/services"
)

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthcheckAPI interfaces.IHealthcheckHandler
	UserAPI        interfaces.IUserHandler
	InventoryAPI   interfaces.IInventoryHandler
	IngredientAPI  interfaces.IIngredientHandler
	RecipeAPI      interfaces.IRecipeHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	// user
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	userSvc := &services.RegisterService{
		UserRepo: userRepo,
	}

	userAPI := &api.IUserHandler{
		UserService: userSvc,
	}

	// inventory
	invRepo := &repository.InventoryRepository{
		DB: helpers.DB,
	}

	invSvc := &services.InventoryService{
		InventoryRepo: invRepo,
	}

	invAPI := &api.InventoryAPI{
		InventoryService: invSvc,
	}

	// recipe
	repoRecipe := &repository.RecipeRepository{
		DB: helpers.DB,
	}

	svcRecipe := &services.ReceipeService{
		RepoReceipe: repoRecipe,
	}

	recipeAPI := &api.ReceipeAPi{
		SvcRecipe: svcRecipe,
	}

	// ingredient
	ingredientRepo := &repository.IngredientRepository{
		DB: helpers.DB,
	}

	ingredientSvc := &services.IngredientService{
		RepoIngredient: ingredientRepo,
	}

	ingredientAPI := &api.IngredientAPI{
		SvcIngredient: ingredientSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		HealthcheckAPI: healthcheckAPI,
		UserAPI:        userAPI,
		InventoryAPI:   invAPI,
		IngredientAPI:  ingredientAPI,
		RecipeAPI:      recipeAPI,
	}
}
