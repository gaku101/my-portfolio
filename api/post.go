package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type postResponse struct {
	Id           int64       `json:"id"`
	Author       string      `json:"author"`
	Title        string      `json:"title"`
	BookAuthor   string      `json:"book_author"`
	BookImage    string      `json:"book_image"`
	BookPage     int16       `json:"book_page"`
	BookPageRead int16       `json:"book_page_read"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
type postRelatedResponse struct {
	Id           int64       `json:"id"`
	Author       string      `json:"author"`
	Title        string      `json:"title"`
	BookAuthor   string      `json:"book_author"`
	BookImage    string      `json:"book_image"`
	BookPage     int16       `json:"book_page"`
	BookPageRead int16       `json:"book_page_read"`
	Category     db.Category `json:"category"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	AuthorImage  string      `json:"authorImage"`
	Favorites    int         `json:"favorites"`
	CommentsNum  int         `json:"commentsNum"`
}

func newPostRelatedResponse(post db.Post, category db.Category, authorImage string, favorites int, commentsNum int) postRelatedResponse {
	return postRelatedResponse{
		Id:           post.ID,
		Author:       post.Author,
		Title:        post.Title,
		Category:     category,
		BookAuthor:   post.BookAuthor,
		BookImage:    post.BookImage,
		BookPage:     post.BookPage,
		BookPageRead: post.BookPageRead,
		CreatedAt:    post.CreatedAt,
		AuthorImage:  authorImage,
		Favorites:    favorites,
		CommentsNum:  commentsNum,
	}
}
func newPostResponse(post db.Post) postResponse {
	return postResponse{
		Id:           post.ID,
		Author:       post.Author,
		Title:        post.Title,
		BookAuthor:   post.BookAuthor,
		BookImage:    post.BookImage,
		BookPage:     post.BookPage,
		BookPageRead: post.BookPageRead,
		CreatedAt:    post.CreatedAt,
	}
}

type createPostRequest struct {
	Author     string `json:"author" binding:"required,alphanum"`
	Title      string `json:"title" binding:"required"`
	BookAuthor string `json:"bookAuthor"`
	BookImage  string `json:"bookImage" `
	BookPage   int16  `json:"bookPage" `
}

type createPostResponse struct {
	Post  postResponse `json:"post"`
	Entry db.Entry     `json:"entry"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, valid := server.validUser(ctx, req.Author)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to create other user's post")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreatePostTxParams{
		UserID:     user.ID,
		Author:     req.Author,
		Title:      req.Title,
		BookAuthor: req.BookAuthor,
		BookImage:  req.BookImage,
		BookPage:   req.BookPage,
		Amount:     3,
	}

	result, err := server.store.CreatePostTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := createPostResponse{
		Post:  newPostResponse(result.Post),
		Entry: result.Entry,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
type getPostResponse struct {
	Post         postResponse `json:"post"`
	PostFavorite int          `json:"postFavorite"`
	Category     db.Category  `json:"category"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.GetPost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	category, err := server.store.GetPostCategory(ctx, post.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("post_id = %v's category not set", post.ID)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	postFavorite := len(server.getFavoriteCount(ctx, post.ID))
	rsp := getPostResponse{
		Post:         newPostResponse(post),
		PostFavorite: postFavorite,
		Category:     category,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type listPostRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listMyPosts(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListMyPostsParams{
		Author: authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	posts, err := server.store.ListMyPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []postRelatedResponse
	for i := range posts {
		post := posts[i]
		category, err := server.store.GetPostCategory(ctx, post.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v's category not set", post.ID)
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

		}
		postFavorite := len(server.getFavoriteCount(ctx, post.ID))
		commentsNum := server.getCommentsCount(ctx, post.ID)

		rsp := newPostRelatedResponse(post, category, "", postFavorite, commentsNum)
		response = append(response, rsp)
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) listPosts(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListPostsParams{
		Author: authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	posts, err := server.store.ListPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []postRelatedResponse
	for i := range posts {
		post := posts[i]
		category, err := server.store.GetPostCategory(ctx, post.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("post_id = %v's category not set", post.ID)
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

		}
		authorImage, err := server.store.GetUserImage(ctx, post.Author)
		preUrl := getPresignedUrl(authorImage)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		favorites := len(server.getFavoriteCount(ctx, post.ID))
		commentsNum := server.getCommentsCount(ctx, post.ID)

		rsp := newPostRelatedResponse(post, category, preUrl, favorites, commentsNum)
		response = append(response, rsp)
	}

	ctx.JSON(http.StatusOK, response)
}

type updatePostRequest struct {
	ID           int64  `json:"id" binding:"required,min=1"`
	Author       string `json:"author" binding:"required,alphanum"`
	BookPageRead int16  `json:"bookPageRead"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, valid := server.validUser(ctx, req.Author)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Username != authPayload.Username {
		err := errors.New("Not allowed to update other user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdatePostParams{
		ID:           req.ID,
		BookPageRead: req.BookPageRead,
	}

	post, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newPostResponse(post)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getFavoriteCount(ctx *gin.Context, postId int64) []int64 {
	postFavorite, err := server.store.GetPostFavorite(ctx, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return postFavorite
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return postFavorite
	}
	return postFavorite
}
func (server *Server) getCommentsCount(ctx *gin.Context, postId int64) int {
	comments, err := server.store.GetCommentsId(ctx, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("post_id = %v's comments not set", postId)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return len(comments)
		}
	}
	return len(comments)
}

type deletePostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	post, err := server.store.GetPost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if post.Author != authPayload.Username {
		err := errors.New("Not allowed to delete other user's post")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.DeletePostTxParams{
		ID: req.ID,
	}

	err = server.store.DeletePostTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
