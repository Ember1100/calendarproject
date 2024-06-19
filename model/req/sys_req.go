package req

type PageQuery struct {
	CreateID string `json:"create_id"` //创建人id
	PageSize string `json:"pageSize"`
	PageNum  string `json:"pageNum"`
}

type PostRequest struct {
	CreateID uint `json:"create_id" binding:"required"`

	WarnDate string `json:"warn_date" binding:"required"`

	WarnContext string `json:"warn_content" binding:"required"`
	SendType    string `json:"send_type" binding:"required"`
}
