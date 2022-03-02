package entity

type Users struct {
	ID          int
	ScreenName  string
	DisplayName string
	Password    string
	Email       *string
	CreatedAt   int64
	UpdatedAt   int64
}
