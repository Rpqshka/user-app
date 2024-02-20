package user_app

type User struct {
	Id            string `gorm:"primaryKey"`
	FirstName     string `gorm:"column:first_name"`
	LastName      string `gorm:"column:last_name"`
	Age           int    `gorm:"column:age"`
	RecordingDate int64  `gorm:"column:recording_date"`
}

type SearchInput struct {
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	MaxAge    int   `json:"max_age"`
	MinAge    int   `json:"min_age"`
}
