package data

type Patient struct {
	id         int
	first_name string
	last_name  string
	state      string // timezone may be more to the point?
	country    string
	created_at string
}
