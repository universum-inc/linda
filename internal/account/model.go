package account

import "time"

type Account struct {
	DateOfBirth time.Time `db:"date_of_birth"`
	Username    string    `db:"username"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	ID          int64     `db:"id"`
}

type ContactInfo struct {
	ContactType    string `db:"contact_type"`
	ContactDetails string `db:"contact_details"`
	ID             int64  `db:"id"`
	AccountID      int64  `db:"account_id"`
}

type Photo struct {
	CreatedAt time.Time `db:"created_at"`
	Link      string    `db:"link"`
	ID        int64     `db:"id"`
	AccountID int64     `db:"account_id"`
}

type Video struct {
	CreatedAt time.Time `db:"created_at"`
	Link      string    `db:"link"`
	ID        int64     `db:"id"`
	AccountID int64     `db:"account_id"`
}
