package api

import (
	"fmt"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gaku101/my-portfolio/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// New Server creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/users/:username", server.getUser)
	authRoutes.PUT("/users", server.updateUser)

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts/by/:owner", server.getAccountByOwner)

	authRoutes.POST("/transfers", server.createTransfer)

	authRoutes.POST("/posts", server.createPost)
	authRoutes.GET("/posts/:id", server.getPost)
	authRoutes.GET("/posts", server.listMyPosts)
	authRoutes.GET("/posts/list", server.listPosts)
	authRoutes.PUT("/posts", server.updatePost)
	authRoutes.DELETE("/posts/:id", server.deletePost)

	authRoutes.POST("/category", server.createCategory)
	authRoutes.GET("/category", server.listCategories)

	authRoutes.GET("/post-favorite/list/:userId", server.listFavoritePosts)
	authRoutes.GET("/post-favorite/:postId", server.getPostFavorite)
	authRoutes.POST("/post-favorite", server.createPostFavorite)

	authRoutes.POST("/follow", server.createFollow)
	authRoutes.GET("/follow/:followingId", server.getFollow)
	authRoutes.GET("/follow", server.listFollow)
	authRoutes.DELETE("/follow/:followingId", server.deleteFollow)

	authRoutes.POST("/comments", server.createComment)
	authRoutes.GET("/comments/:postId", server.listComments)
	authRoutes.DELETE("/comments/:id", server.deleteComment)

	authRoutes.POST("/images/:username", server.uploadImage)

	authRoutes.POST("/notes", server.createNote)


	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
