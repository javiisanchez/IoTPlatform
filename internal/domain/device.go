package domain

import "time"

// Entidades de negocio principales de la app

type Device struct {
	ID        string     `json:"ID"`
	Name      string     `json:"Name"`
	Location  string     `json:"Location"`
	Status    string     `json:"Status"`
	CreatedAt *time.Time `json:"CreatedAt"`
	UpdatedAt *time.Time `json:"UpdatedAt"`
}

func NewDevice(ID string, Name string, Location string, Status string, CreatedAt *time.Time, UpdatedAt *time.Time) *Device {
	return &Device{
		ID:        ID,
		Name:      Name,
		Location:  Location,
		Status:    Status,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}

func (d *Device) GetID() string {
	return d.ID
}

func (d *Device) GetName() string {
	return d.Name
}

func (d *Device) GetLocation() string {
	return d.Location
}

func (d *Device) GetStatus() string {
	return d.Status
}

func (d *Device) GetCreatedAt() *time.Time {
	return d.CreatedAt
}

func (d *Device) GetUpdatedAt() *time.Time {
	return d.UpdatedAt
}

func (d *Device) SetID(ID string) {
	d.ID = ID
}

func (d *Device) SetName(Name string) {
	d.Name = Name
}

func (d *Device) SetLocation(Location string) {
	d.Location = Location
}

func (d *Device) SetStatus(Status string) {
	d.Status = Status
}

func (d *Device) SetCreatedAt(CreatedAt *time.Time) {
	d.CreatedAt = CreatedAt
}

func (d *Device) SetUpdatedAt(UpdatedAt *time.Time) {
	d.UpdatedAt = UpdatedAt
}
