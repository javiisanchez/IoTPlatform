package domain

import "time"

// Entidades de negocio principales de la app

type Device struct {
	ID        string
	Name      string
	Location  string
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
