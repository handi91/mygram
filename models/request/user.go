package request

type RegisterUser struct {
	Username string `json:"username" valid:"required~Username is required"`
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string `json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `json:"age" valid:"required~Age is required,range(8|10000)~Minimum age is 8"`
}

type LoginUser struct {
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string `json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type UpdateUser struct {
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email format"`
	Username string `json:"username" valid:"required~Username is required"`
}
