package dto

type CreateUserDTO struct {
	Name  *string
	Email *string
	Age   *int64
}

type UpdateUserDTO struct {
	Name  *string
	Email *string
	Age   *int64
}
