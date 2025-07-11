package userloginservice


type IGetUserByEmailResponse struct{
	Email string `json:"email"`
    Password string `json:"password"`
	
}