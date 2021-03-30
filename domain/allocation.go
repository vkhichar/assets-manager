package domain

type Allocation struct {
	User_id          int     `db:"user_id"`
	Asset_id         int     `db:"asset_id"`
	Date_Allocated   *string `db:"date_alloc"`
	Date_Deallocated *string `db:"date_dealloc"`
}
