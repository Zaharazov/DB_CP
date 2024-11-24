package routers

import (
	"DB_CP/internal/services"
	"strings"
)

var gym_routes = Routes{

	// Route{
	// 	"CreateGym",
	// 	strings.ToUpper("Post"), // Gyms
	// 	"/v1/gyms/",
	// 	services.CreateGym,
	// },

	// Route{
	// 	"DeleteGymById",
	// 	strings.ToUpper("Delete"),
	// 	"/v1/gyms/{gym_id}/",
	// 	services.DeleteGymById,
	// },

	// Route{
	// 	"GetGymById",
	// 	strings.ToUpper("Get"),
	// 	"/v1/gyms/{gym_id}/",
	// 	services.GetUserById,
	// },

	Route{
		"GetAllGyms",
		strings.ToUpper("Get"),
		"/v1/gyms",
		services.GetAllGyms,
	},

	// Route{
	// 	"EditGymById",
	// 	strings.ToUpper("Put"),
	// 	"/v1/gyms/{gym_id}/",
	// 	services.EditGymById,
	// },
}
