package database

type Id int64

type NewUser struct {
	Group     Id     `json:"group" db:"group"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	BirthYear uint16 `json:"birth_year" db:"birth_year"`
}

type User struct {
	Id        Id     `json:"id" db:"id"`
	Group     Id     `json:"group" db:"group"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	BirthYear uint16 `json:"birth_year" db:"birth_year"`
}

type NewGroup struct {
	ParentGroup *Id    `json:"parent_group,omitempty" db:"parent_group"`
	Name        string `json:"name" db:"name"`
}

type Group struct {
	Id          Id     `json:"id" db:"id"`
	ParentGroup *Id    `json:"parent_group,omitempty" db:"parent_group"`
	Name        string `json:"name" db:"name"`
}

type GroupWithUserCount struct {
	Id          Id     `json:"id" db:"id"`
	ParentGroup *Id    `json:"parent_group,omitempty" db:"parent_group"`
	Name        string `json:"name" db:"name"`
	UserCount   int    `json:"user_count" db:"user_count"`
}
