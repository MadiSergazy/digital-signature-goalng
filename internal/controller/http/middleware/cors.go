package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	// cfg := config.Get()

	corsProvided := cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins: cfg.CORS.AllowOrigins, // Provide your list of allowed origins here
		// AllowMethods: []string{
		// 	http.MethodGet,
		// 	http.MethodPost,
		// 	http.MethodPut,
		// 	http.MethodDelete,
		// 	http.MethodOptions,
		// },
		// AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// AllowCredentials: true,
	})

	return corsProvided
}

// func corsMiddleware() gin.HandlerFunc {
// 	cfg := config.Get()

// 	corsProvided := cors.New(cors.Config{
// 		AllowOrigins: cfg.CORS.AllowOrigins,
// 		AllowMethods: []string{
// 			http.MethodGet,
// 			http.MethodPost,
// 			http.MethodPut,
// 			http.MethodDelete,
// 			http.MethodOptions,
// 		},
// 		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		AllowCredentials: true,
// 	})

// 	return corsProvided
// }
