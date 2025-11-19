package repositories

import (
	"errors"

	"github.com/namduong/project-layout/internal/models"
)

type AdminRepository struct {
	admins []models.Admin
}

func (r *AdminRepository) FindByUsername(username string) (*models.Admin, error) {
	for _, admin := range r.admins {
		if admin.UserName == username {
			return &admin, nil
		}
	}
	return nil, errors.New("admin not found")
}
