// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategory(ctx context.Context, name string) (Category, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePostCategory(ctx context.Context, arg CreatePostCategoryParams) (PostCategory, error)
	CreatePostFavorite(ctx context.Context, arg CreatePostFavoriteParams) (PostFavorite, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountByOwner(ctx context.Context, owner string) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetFollow(ctx context.Context, arg GetFollowParams) (Follow, error)
	GetMyFavoritePost(ctx context.Context, arg GetMyFavoritePostParams) (PostFavorite, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetPostCategory(ctx context.Context, postID int64) (Category, error)
	GetPostFavorite(ctx context.Context, postID int64) ([]int64, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	GetUserImage(ctx context.Context, username string) (string, error)
	ListCategories(ctx context.Context) ([]Category, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListFavoritePosts(ctx context.Context, arg ListFavoritePostsParams) ([]Post, error)
	ListFollow(ctx context.Context, arg ListFollowParams) ([]ListFollowRow, error)
	ListMyPosts(ctx context.Context, arg ListMyPostsParams) ([]Post, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdatePostCategory(ctx context.Context, arg UpdatePostCategoryParams) (PostCategory, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
