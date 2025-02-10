package requests

type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=2"`
    Email    string `json:"email" validate:"required,email"`
    Phone    string `json:"phone" validate:"required,e164"`  // E.164 format
    Password string `json:"password" validate:"required,min=6"`
} 