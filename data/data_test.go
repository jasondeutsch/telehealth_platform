package data

import "testing"

func TestAdminGetPatients(t *testing.T) {
	var admin *Admin
	t.Error(admin.GetAllPatients())
}
