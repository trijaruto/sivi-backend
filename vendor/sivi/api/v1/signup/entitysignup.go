package signup

type SignUpRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID         int    `json:"id"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	IsStatus   int    `json:"istatus"`
	CreateTime string `json:"createtime"`
	CreatedBy  string `json:"createdby"`
	UpdateTime string `json:"updatetime"`
	UpdatedBy  string `json:"updatedby"`
	UsertypeID int    `json:"usertype_id"`
}
