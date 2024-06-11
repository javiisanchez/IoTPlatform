package repositories

import (
	"IoTPlatform/internal/domain"
	"context"
	"time"

	"IoTPlatform/internal/adapters/storage/postgres"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

/**
 * DeviceRepository implements port.DeviceRepository interface
 * and provides an access to the postgres database
 */
type DeviceRepository struct {
	db *postgres.DB
}

// NewCategoryRepository creates a new category repository instance
func NewDeviceRepository(db *postgres.DB) *DeviceRepository {
	return &DeviceRepository{
		db,
	}
}

// DeviceRepository creates a new Device record in the database
func (cr *DeviceRepository) CreateDevice(ctx context.Context, device *domain.Device) (*domain.Device, error) {
	query := cr.db.QueryBuilder.Insert("device").
		Columns("name").
		Values(device.Name).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&device.ID,
		&device.Name,
		&device.Location,
		&device.Status,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if errCode := cr.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return device, nil
}

// GetCategoryByID retrieves a category record from the database by id
func (cr *DeviceRepository) GetDeviceByID(ctx context.Context, ID string) (*domain.Device, error) {
	var device domain.Device

	query := cr.db.QueryBuilder.Select("*").
		From("device").
		Where(sq.Eq{"id": ID}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&device.ID,
		&device.Name,
		&device.Location,
		&device.Status,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &device, nil
}

// ListCategories retrieves a list of categories from the database
func (cr *DeviceRepository) ListDevices(ctx context.Context, skip, limit uint64) ([]domain.Device, error) {
	var device domain.Device
	var devices []domain.Device

	query := cr.db.QueryBuilder.Select("*").
		From("device").
		OrderBy("id").
		Limit(limit).
		Offset((skip - 1) * limit)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := cr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&device.ID,
			&device.Name,
			&device.Location,
			&device.Status,
			&device.CreatedAt,
			&device.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

// UpdateCategory updates a category record in the database
func (cr *DeviceRepository) UpdateDevice(ctx context.Context, device *domain.Device) (*domain.Device, error) {
	query := cr.db.QueryBuilder.Update("device").
		Set("name", device.Name).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": device.ID}).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&device.ID,
		&device.Name,
		&device.Location,
		&device.Status,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if errCode := cr.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return device, nil
}

// DeleteCategory deletes a category record from the database by id
func (cr *DeviceRepository) DeleteDevice(ctx context.Context, ID string) error {
	query := cr.db.QueryBuilder.Delete("device").
		Where(sq.Eq{"id": ID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = cr.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DeviceRepository) Create(device domain.Device) error {
	//Here your code for login in mongo database
	return nil
}

func (r *DeviceRepository) Delete(ID string) error {
	//Here your code for save in mongo database
	return nil
}
