package login

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	IsStatus int    `json:"istatus"`
	// Nik           int         `json:"nik"`
	// Email         string      `json:"email"`
	// IsActive      int         `json:"isactive"`
	// UserTypeID    int         `json:"usertypeid"`
	// UserTypeCode  string      `json:"usertypecode"`
	// UserTypeName  string      `json:"usertypename"`
	// PlantSlocItem interface{} `json:"plantsloc"`
}
