package response

type ResponseError struct {
	//Code    int      `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

type SimpleResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SimpleResponseList struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   *int64      `json:"total"`
}

type Paging struct {
	Page  int    `form:"page" json:"page" binding:"required"`
	Limit int    `form:"limit" json:"limit" binding:"required"`
	Sort  string `form:"sort" json:"sort" binding:"required"`
}

type GetByIDsRequest struct {
	IDs []int64 `json:"ids"`
}
