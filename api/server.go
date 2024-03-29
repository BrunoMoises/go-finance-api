package api

import (
	db "github.com/BrunoMoises/go-finance-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func CORSConfig() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(CORSConfig())

	//User
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)

	//Category
	router.POST("/category", server.createCategory)
	router.GET("/category/id/:id", server.getCategory)
	router.GET("/category", server.getCategories)
	router.DELETE("/category/:id", server.deleteCategory)
	router.PUT("/category/:id", server.updateCategory)

	//Account
	router.POST("/account", server.createAccount)
	router.GET("/account/id/:id", server.getAccount)
	router.GET("/account", server.getAccounts)
	router.GET("/account/graph/:user_id/:type", server.getAccountGraph)
	router.GET("/account/reports/:user_id/:type", server.getAccountReports)
	router.DELETE("/account/:id", server.deleteAccount)
	router.PUT("/account/:id", server.updateAccount)

	router.POST("/login", server.login)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error: ": err.Error()}
}
