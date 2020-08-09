package swagger

import "time"

// structs used by swagger to define the JSON body and sucess and fail responses

// ADMIN LOGIN
type adminLoginBody struct {
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"password"`
}

// success
type adminLogin200 struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTY3NjI3MTcsInJvbGUiOiJhZG1pbiIsInN1YiI6MX0.BepVrGACE5xvkA08cMDptPR0sz5fcKkPfc4NO-oxZRE"`
	Message      string `json:"message" example:"Authenticated"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjgyOTUxMTcsInN1YiI6MX0.uS2sJ52TTe40JSvqwGTigbig186sDmXRl9FB8OeheWw"`
	Success      bool   `json:"success" example:"true"`
}

// unauthorized
type adminLogin401 struct {
	Message string `json:"message" example:"Please provide valid credentials"`
	Success bool   `json:"success" example:"false"`
}

// forbidden
type adminLogin403 struct {
	Message string `json:"message" example:"Account has been suspended"`
	Success bool   `json:"success" example:"false"`
}

// unprocessable entity
type adminLogin422 struct {
	Message string `json:"message" example:"Required fields either missing or empty"`
	Success bool   `json:"success" example:"false"`
}

// ADMIN LOGOUT
type adminLogout200 struct {
	Message string `json:"message" example:"Logged out"`
	Success bool   `json:"success" example:"true"`
}

// ADMIN REFRESH
type adminRefreshBody struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjgzNjAyOTIsInN1YiI6MX0.5ddhUgEXlnN3NZH0yrHH1uGspHDFnqevA18BVy5dFDY"`
}

// success
type adminRefresh200 struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTY3NjI3MTcsInJvbGUiOiJhZG1pbiIsInN1YiI6MX0.BepVrGACE5xvkA08cMDptPR0sz5fcKkPfc4NO-oxZRE"`
	Message      string `json:"message" example:"Token refreshed"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjgyOTUxMTcsInN1YiI6MX0.uS2sJ52TTe40JSvqwGTigbig186sDmXRl9FB8OeheWw"`
	Success      bool   `json:"success" example:"true"`
}

// unauthorized
type adminRefresh401 struct {
	Message string `json:"message" example:"Invalid token"`
	Success bool   `json:"success" example:"false"`
}

// forbidden
type adminRefresh403 struct {
	Message string `json:"message" example:"Account has been suspended"`
	Success bool   `json:"success" example:"false"`
}

// unprocessable entity
type adminRefresh422 struct {
	Message string `json:"message" example:"Required fields either missing or empty"`
	Success bool   `json:"success" example:"false"`
}

// GET ALL ADMINS
type adminGetAll1 struct {
	ID        uint64    `json:"id" example:"5"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"admin@test.com"`
	Role      string    `json:"role" example:"admin"`
	Status    string    `json:"status" example:"active"`
	CreatedAt time.Time `json:"created_at" example:"2020-08-07T20:12:54Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-08-07T20:12:54Z"`
}

type adminGetAll200 struct {
	Data         []adminGetAll1
	Message      string `json:"message" example:"All admins"`
	Success      bool   `json:"success" example:"true"`
	TotalResults int64  `json:"total_results" example:"1"`
}

// GET ADMIN BY ID
type adminGetByID200 struct {
	Data    adminGetAll1
	Message string `json:"message" example:"Returned admin"`
	Success bool   `json:"success" example:"true"`
}

// not found
type adminGetByID404 struct {
	Message string `json:"message" example:"Admin not found"`
	Success bool   `json:"success" example:"false"`
}

// not found
type adminGetByID422 struct {
	Message string `json:"message" example:"ID in invalid format"`
	Success bool   `json:"success" example:"false"`
}

// CREATE NEW ADMIN
type adminCreateAdminBody struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"admin@test.com"`
	Password  string `json:"password" example:"password"`
	Role      string `json:"role" example:"admin"`
	Status    string `json:"status" example:"active"`
}

type adminCreateAdmin1 struct {
	ID        uint64 `json:"id" example:"5"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"admin@test.com"`
	Role      string `json:"role" example:"admin"`
}

type adminCreateAdmin201 struct {
	Data    adminCreateAdmin1
	Message string `json:"message" example:"Created new admin"`
	Success bool   `json:"success" example:"true"`
}

type adminCreateAdmin422 struct {
	Message string `json:"message" example:"First name is required"`
	Success bool   `json:"success" example:"false"`
}

type adminCreateAdmin409 struct {
	Message string `json:"message" example:"Admin with email already exists"`
	Success bool   `json:"success" example:"false"`
}

// ADMIN DELETE
type adminDeleteAdmin200 struct {
	Message string `json:"message" example:"Deleted admin with ID of 13"`
	Success bool   `json:"success" example:"true"`
}

type adminDeleteAdmin404 struct {
	Message string `json:"message" example:"Active admin not found"`
	Success bool   `json:"success" example:"false"`
}

type adminDeleteAdmin422 struct {
	Message string `json:"message" example:"ID in invalid format"`
	Success bool   `json:"success" example:"false"`
}

// ADMIN EDIT
type adminEditAdminBody struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"admin@test.com"`
	Password  string `json:"password" example:"password"`
	Role      string `json:"role" example:"admin"`
	Status    string `json:"status" example:"active"`
}

type adminEditAdmin200 struct {
	Data    adminCreateAdmin1
	Message string `json:"message" example:"Updated admin"`
	Success bool   `json:"success" example:"true"`
}

type adminEditAdmin404 struct {
	Message string `json:"message" example:"Active admin not found"`
	Success bool   `json:"success" example:"false"`
}

type adminEditAdmin409 struct {
	Message string `json:"message" example:"Email address already in use"`
	Success bool   `json:"success" example:"false"`
}

type adminEditAdmin422 struct {
	Message string `json:"message" example:"Email is in an invalid format"`
	Success bool   `json:"success" example:"false"`
}

// ADMIN GET SELF
type adminGetSelf200 struct {
	Data    []adminGetAll1
	Message string `json:"message" example:"Returned admin"`
	Success bool   `json:"success" example:"true"`
}

// ADMIN EDIT SELF
type adminEditSelfBody struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"admin@test.com"`
	Password  string `json:"password" example:"password"`
}

// CREATE NEW CLIENT
type adminCreateClientBody struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"client@test.com"`
	Phone     string `json:"phone" example:"(702) 867-5309"`
	Notes     string `json:"notes" example:"Random text about client"`
	Password  string `json:"password" example:"password"`
}

type adminCreateClient1 struct {
	ID        uint64 `json:"id" example:"5"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"client@test.com"`
	Phone     string `json:"phone" example:"(702) 867-5309"`
	Notes     string `json:"notes" example:"Random text about client"`
	Status    string `json:"status" example:"active"`
}

type adminCreateClient201 struct {
	Data    adminCreateClient1
	Message string `json:"message" example:"Created new client"`
	Success bool   `json:"success" example:"true"`
}

type adminCreateClient422 struct {
	Message string `json:"message" example:"First name is required"`
	Success bool   `json:"success" example:"false"`
}

type adminCreateClient409 struct {
	Message string `json:"message" example:"Client with email already exists"`
	Success bool   `json:"success" example:"false"`
}

// Client DELETE
type adminDeleteClient200 struct {
	Message string `json:"message" example:"Deleted client with ID of 13"`
	Success bool   `json:"success" example:"true"`
}

type adminDeleteClient404 struct {
	Message string `json:"message" example:"Active client not found"`
	Success bool   `json:"success" example:"false"`
}

type adminDeleteClient422 struct {
	Message string `json:"message" example:"ID in invalid format"`
	Success bool   `json:"success" example:"false"`
}

// GET ALL CLIENTS
type adminGetAllClients1 struct {
	ID        uint64    `json:"id" example:"5"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"admin@test.com"`
	Phone     string    `json:"phone" example:"(702) 867-5309"`
	Notes     string    `json:"notes" example:"This is a test note full of random test that is related to this client."`
	Status    string    `json:"status" example:"active"`
	CreatedAt time.Time `json:"created_at" example:"2020-08-07T20:12:54Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-08-07T20:12:54Z"`
}

type adminGetAllClients200 struct {
	Data         []adminGetAll1
	Message      string `json:"message" example:"All clients"`
	Success      bool   `json:"success" example:"true"`
	TotalResults int64  `json:"total_results" example:"1"`
}

// GET CLIENT BY ID
type adminGetClientByID200 struct {
	Data    adminGetAllClients1
	Message string `json:"message" example:"Returned client"`
	Success bool   `json:"success" example:"true"`
}

// not found
type adminGetClientByID404 struct {
	Message string `json:"message" example:"Client not found"`
	Success bool   `json:"success" example:"false"`
}

// invalid format
type adminGetClientByID422 struct {
	Message string `json:"message" example:"ID in invalid format"`
	Success bool   `json:"success" example:"false"`
}

// CLIENT EDIT
type adminEditClientBody struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"admin@test.com"`
	Password  string `json:"password" example:"password"`
	Notes     string `json:"notes" example:"This is a test note full of random test that is related to this client."`
	Phone     string `json:"phone" example:"(702) 867-5309"`
	Status    string `json:"status" example:"active"`
}

type adminEditClient200 struct {
	Data    adminCreateClient1
	Message string `json:"message" example:"Updated client"`
	Success bool   `json:"success" example:"true"`
}

type adminEditClient404 struct {
	Message string `json:"message" example:"Active client not found"`
	Success bool   `json:"success" example:"false"`
}

type adminEditClient409 struct {
	Message string `json:"message" example:"Email address already in use"`
	Success bool   `json:"success" example:"false"`
}

type adminEditClient422 struct {
	Message string `json:"message" example:"Email is in an invalid format"`
	Success bool   `json:"success" example:"false"`
}

// CLIENT GET SELF
type clientGetSelf200 struct {
	Data    clientEditSelfBody
	Message string `json:"message" example:"Returned own details"`
	Success bool   `json:"success" example:"true"`
}

// CLIENT EDIT SELF
type clientEditSelfBody struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Email     string `json:"email" example:"client@test.com"`
	Password  string `json:"password" example:"password"`
	Phone     string `json:"phone" example:"(702) 867-5309"`
}

type clientEditSelf200 struct {
	Data    clientEditSelfBody
	Message string `json:"message" example:"Updated own details"`
	Success bool   `json:"success" example:"true"`
}
