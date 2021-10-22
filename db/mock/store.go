// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gaku101/my-portfolio/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(arg0 context.Context, arg1 db.CreateEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), arg0, arg1)
}

// CreateFollow mocks base method.
func (m *MockStore) CreateFollow(arg0 context.Context, arg1 db.CreateFollowParams) (db.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFollow", arg0, arg1)
	ret0, _ := ret[0].(db.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFollow indicates an expected call of CreateFollow.
func (mr *MockStoreMockRecorder) CreateFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFollow", reflect.TypeOf((*MockStore)(nil).CreateFollow), arg0, arg1)
}

// CreateNote mocks base method.
func (m *MockStore) CreateNote(arg0 context.Context, arg1 db.CreateNoteParams) (db.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNote", arg0, arg1)
	ret0, _ := ret[0].(db.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNote indicates an expected call of CreateNote.
func (mr *MockStoreMockRecorder) CreateNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNote", reflect.TypeOf((*MockStore)(nil).CreateNote), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), arg0, arg1)
}

// CreatePostCategory mocks base method.
func (m *MockStore) CreatePostCategory(arg0 context.Context, arg1 db.CreatePostCategoryParams) (db.PostCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.PostCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostCategory indicates an expected call of CreatePostCategory.
func (mr *MockStoreMockRecorder) CreatePostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostCategory", reflect.TypeOf((*MockStore)(nil).CreatePostCategory), arg0, arg1)
}

// CreatePostFavorite mocks base method.
func (m *MockStore) CreatePostFavorite(arg0 context.Context, arg1 db.CreatePostFavoriteParams) (db.PostFavorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.PostFavorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostFavorite indicates an expected call of CreatePostFavorite.
func (mr *MockStoreMockRecorder) CreatePostFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostFavorite", reflect.TypeOf((*MockStore)(nil).CreatePostFavorite), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeleteComments mocks base method.
func (m *MockStore) DeleteComments(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComments indicates an expected call of DeleteComments.
func (mr *MockStoreMockRecorder) DeleteComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComments", reflect.TypeOf((*MockStore)(nil).DeleteComments), arg0, arg1)
}

// DeleteFollow mocks base method.
func (m *MockStore) DeleteFollow(arg0 context.Context, arg1 db.DeleteFollowParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFollow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFollow indicates an expected call of DeleteFollow.
func (mr *MockStoreMockRecorder) DeleteFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFollow", reflect.TypeOf((*MockStore)(nil).DeleteFollow), arg0, arg1)
}

// DeleteFollows mocks base method.
func (m *MockStore) DeleteFollows(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFollows", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFollows indicates an expected call of DeleteFollows.
func (mr *MockStoreMockRecorder) DeleteFollows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFollows", reflect.TypeOf((*MockStore)(nil).DeleteFollows), arg0, arg1)
}

// DeleteMyFavoritePosts mocks base method.
func (m *MockStore) DeleteMyFavoritePosts(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMyFavoritePosts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMyFavoritePosts indicates an expected call of DeleteMyFavoritePosts.
func (mr *MockStoreMockRecorder) DeleteMyFavoritePosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMyFavoritePosts", reflect.TypeOf((*MockStore)(nil).DeleteMyFavoritePosts), arg0, arg1)
}

// DeleteNote mocks base method.
func (m *MockStore) DeleteNote(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNote", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockStoreMockRecorder) DeleteNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockStore)(nil).DeleteNote), arg0, arg1)
}

// DeleteNotes mocks base method.
func (m *MockStore) DeleteNotes(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNotes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNotes indicates an expected call of DeleteNotes.
func (mr *MockStoreMockRecorder) DeleteNotes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNotes", reflect.TypeOf((*MockStore)(nil).DeleteNotes), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockStore) DeletePost(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockStoreMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockStore)(nil).DeletePost), arg0, arg1)
}

// DeletePostCategory mocks base method.
func (m *MockStore) DeletePostCategory(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostCategory indicates an expected call of DeletePostCategory.
func (mr *MockStoreMockRecorder) DeletePostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostCategory", reflect.TypeOf((*MockStore)(nil).DeletePostCategory), arg0, arg1)
}

// DeletePostFavorite mocks base method.
func (m *MockStore) DeletePostFavorite(arg0 context.Context, arg1 db.DeletePostFavoriteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostFavorite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostFavorite indicates an expected call of DeletePostFavorite.
func (mr *MockStoreMockRecorder) DeletePostFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostFavorite", reflect.TypeOf((*MockStore)(nil).DeletePostFavorite), arg0, arg1)
}

// DeletePostFavorites mocks base method.
func (m *MockStore) DeletePostFavorites(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostFavorites", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostFavorites indicates an expected call of DeletePostFavorites.
func (mr *MockStoreMockRecorder) DeletePostFavorites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostFavorites", reflect.TypeOf((*MockStore)(nil).DeletePostFavorites), arg0, arg1)
}

// DeletePostTx mocks base method.
func (m *MockStore) DeletePostTx(arg0 context.Context, arg1 db.DeletePostTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostTx indicates an expected call of DeletePostTx.
func (mr *MockStoreMockRecorder) DeletePostTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostTx", reflect.TypeOf((*MockStore)(nil).DeletePostTx), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserTx mocks base method.
func (m *MockStore) DeleteUserTx(arg0 context.Context, arg1 db.DeleteUserTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserTx indicates an expected call of DeleteUserTx.
func (mr *MockStoreMockRecorder) DeleteUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserTx", reflect.TypeOf((*MockStore)(nil).DeleteUserTx), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetComment mocks base method.
func (m *MockStore) GetComment(arg0 context.Context, arg1 int64) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockStoreMockRecorder) GetComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockStore)(nil).GetComment), arg0, arg1)
}

// GetCommentsId mocks base method.
func (m *MockStore) GetCommentsId(arg0 context.Context, arg1 int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsId", arg0, arg1)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsId indicates an expected call of GetCommentsId.
func (mr *MockStoreMockRecorder) GetCommentsId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsId", reflect.TypeOf((*MockStore)(nil).GetCommentsId), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), arg0, arg1)
}

// GetFollow mocks base method.
func (m *MockStore) GetFollow(arg0 context.Context, arg1 db.GetFollowParams) (db.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollow", arg0, arg1)
	ret0, _ := ret[0].(db.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollow indicates an expected call of GetFollow.
func (mr *MockStoreMockRecorder) GetFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollow", reflect.TypeOf((*MockStore)(nil).GetFollow), arg0, arg1)
}

// GetMyFavoritePost mocks base method.
func (m *MockStore) GetMyFavoritePost(arg0 context.Context, arg1 db.GetMyFavoritePostParams) (db.PostFavorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyFavoritePost", arg0, arg1)
	ret0, _ := ret[0].(db.PostFavorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyFavoritePost indicates an expected call of GetMyFavoritePost.
func (mr *MockStoreMockRecorder) GetMyFavoritePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyFavoritePost", reflect.TypeOf((*MockStore)(nil).GetMyFavoritePost), arg0, arg1)
}

// GetNote mocks base method.
func (m *MockStore) GetNote(arg0 context.Context, arg1 int64) (db.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", arg0, arg1)
	ret0, _ := ret[0].(db.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockStoreMockRecorder) GetNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockStore)(nil).GetNote), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockStore) GetPost(arg0 context.Context, arg1 int64) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockStoreMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockStore)(nil).GetPost), arg0, arg1)
}

// GetPostCategory mocks base method.
func (m *MockStore) GetPostCategory(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostCategory indicates an expected call of GetPostCategory.
func (mr *MockStoreMockRecorder) GetPostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostCategory", reflect.TypeOf((*MockStore)(nil).GetPostCategory), arg0, arg1)
}

// GetPostFavorite mocks base method.
func (m *MockStore) GetPostFavorite(arg0 context.Context, arg1 int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostFavorite", arg0, arg1)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostFavorite indicates an expected call of GetPostFavorite.
func (mr *MockStoreMockRecorder) GetPostFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostFavorite", reflect.TypeOf((*MockStore)(nil).GetPostFavorite), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetUserImage mocks base method.
func (m *MockStore) GetUserImage(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserImage", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserImage indicates an expected call of GetUserImage.
func (mr *MockStoreMockRecorder) GetUserImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserImage", reflect.TypeOf((*MockStore)(nil).GetUserImage), arg0, arg1)
}

// ListCategories mocks base method.
func (m *MockStore) ListCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockStoreMockRecorder) ListCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockStore)(nil).ListCategories), arg0)
}

// ListComments mocks base method.
func (m *MockStore) ListComments(arg0 context.Context, arg1 db.ListCommentsParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComments indicates an expected call of ListComments.
func (mr *MockStoreMockRecorder) ListComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockStore)(nil).ListComments), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), arg0, arg1)
}

// ListFavoritePosts mocks base method.
func (m *MockStore) ListFavoritePosts(arg0 context.Context, arg1 db.ListFavoritePostsParams) ([]db.ListFavoritePostsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavoritePosts", arg0, arg1)
	ret0, _ := ret[0].([]db.ListFavoritePostsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFavoritePosts indicates an expected call of ListFavoritePosts.
func (mr *MockStoreMockRecorder) ListFavoritePosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavoritePosts", reflect.TypeOf((*MockStore)(nil).ListFavoritePosts), arg0, arg1)
}

// ListFollow mocks base method.
func (m *MockStore) ListFollow(arg0 context.Context, arg1 db.ListFollowParams) ([]db.ListFollowRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollow", arg0, arg1)
	ret0, _ := ret[0].([]db.ListFollowRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollow indicates an expected call of ListFollow.
func (mr *MockStoreMockRecorder) ListFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollow", reflect.TypeOf((*MockStore)(nil).ListFollow), arg0, arg1)
}

// ListMyAllPosts mocks base method.
func (m *MockStore) ListMyAllPosts(arg0 context.Context, arg1 string) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMyAllPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMyAllPosts indicates an expected call of ListMyAllPosts.
func (mr *MockStoreMockRecorder) ListMyAllPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyAllPosts", reflect.TypeOf((*MockStore)(nil).ListMyAllPosts), arg0, arg1)
}

// ListMyPosts mocks base method.
func (m *MockStore) ListMyPosts(arg0 context.Context, arg1 db.ListMyPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMyPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMyPosts indicates an expected call of ListMyPosts.
func (mr *MockStoreMockRecorder) ListMyPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyPosts", reflect.TypeOf((*MockStore)(nil).ListMyPosts), arg0, arg1)
}

// ListNotes mocks base method.
func (m *MockStore) ListNotes(arg0 context.Context, arg1 db.ListNotesParams) ([]db.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotes", arg0, arg1)
	ret0, _ := ret[0].([]db.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNotes indicates an expected call of ListNotes.
func (mr *MockStoreMockRecorder) ListNotes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotes", reflect.TypeOf((*MockStore)(nil).ListNotes), arg0, arg1)
}

// ListPosts mocks base method.
func (m *MockStore) ListPosts(arg0 context.Context, arg1 db.ListPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts.
func (mr *MockStoreMockRecorder) ListPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockStore)(nil).ListPosts), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateNote mocks base method.
func (m *MockStore) UpdateNote(arg0 context.Context, arg1 db.UpdateNoteParams) (db.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNote", arg0, arg1)
	ret0, _ := ret[0].(db.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNote indicates an expected call of UpdateNote.
func (mr *MockStoreMockRecorder) UpdateNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNote", reflect.TypeOf((*MockStore)(nil).UpdateNote), arg0, arg1)
}

// UpdatePoints mocks base method.
func (m *MockStore) UpdatePoints(arg0 context.Context, arg1 db.UpdatePointsParams) (db.UpdatePointsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePoints", arg0, arg1)
	ret0, _ := ret[0].(db.UpdatePointsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePoints indicates an expected call of UpdatePoints.
func (mr *MockStoreMockRecorder) UpdatePoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePoints", reflect.TypeOf((*MockStore)(nil).UpdatePoints), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockStore) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockStoreMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockStore)(nil).UpdatePost), arg0, arg1)
}

// UpdatePostCategory mocks base method.
func (m *MockStore) UpdatePostCategory(arg0 context.Context, arg1 db.UpdatePostCategoryParams) (db.PostCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.PostCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePostCategory indicates an expected call of UpdatePostCategory.
func (mr *MockStoreMockRecorder) UpdatePostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePostCategory", reflect.TypeOf((*MockStore)(nil).UpdatePostCategory), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserImage mocks base method.
func (m *MockStore) UpdateUserImage(arg0 context.Context, arg1 db.UpdateUserImageParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserImage", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserImage indicates an expected call of UpdateUserImage.
func (mr *MockStoreMockRecorder) UpdateUserImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserImage", reflect.TypeOf((*MockStore)(nil).UpdateUserImage), arg0, arg1)
}
