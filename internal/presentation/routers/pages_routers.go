package routers

import (
	"DB_CP/internal/services"
	"strings"
)

var pages_routes = Routes{
	Route{
		"GetGymsPages",
		strings.ToUpper("Get"),
		"/v1/gyms_pages",
		services.GetGymsPages,
	},

	Route{
		"GetGymsPageById",
		strings.ToUpper("Get"),
		"/v1/gyms_pages/{gym_id}/",
		services.GetGymsPageById,
	},
}
