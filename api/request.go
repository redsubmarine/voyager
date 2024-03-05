package api

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type updateUserRequest struct {
	ID       int64  `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type deleteUserRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
}
