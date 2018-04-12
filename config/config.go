package config

import "github.com/tchalo/iselia"

func NewIselia() (*iselia.IseliaService, error) {
	return &iselia.IseliaService{}, nil
}
