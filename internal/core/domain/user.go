package domain


type User struct {
	ID          string
	Version     int

	FirstName   string
	LastName    string
	PhoneNumber *string

	CreatedAt   string
	UpdatedAt   string
}
