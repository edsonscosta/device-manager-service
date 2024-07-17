package repository

import (
	"database/sql"
	"device-manager-service/internal/device/model"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type DeviceRepository struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) DeviceRepository {
	return DeviceRepository{
		db: db,
	}
}

func (r DeviceRepository) Get(deviceID uuid.UUID) (*model.Device, error) {
	var device model.Device
	err := r.db.QueryRow("SELECT * FROM devices WHERE device_id = $1", deviceID).Scan(
		&device.DeviceID,
		&device.Name,
		&device.Brand,
		&device.IsActive,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &device, nil
}

func (r DeviceRepository) GetAll(limit int, offset int, brand string) ([]*model.Device, error) {
	var devices []*model.Device
	statement := "SELECT * FROM devices WHERE 1=1 "
	if brand != "" {
		statement = statement + fmt.Sprintf("AND brand = '%s' ", brand)
	}
	statement = statement + "LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(statement, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		device := model.Device{}
		err := rows.Scan(
			&device.DeviceID,
			&device.Name,
			&device.Brand,
			&device.IsActive,
			&device.CreatedAt,
			&device.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		devices = append(devices, &device)
	}

	return devices, nil
}

func (r DeviceRepository) Create(device *model.Device) error {
	_, err := r.db.Exec("INSERT INTO devices "+
		"(device_id, name, brand, is_active, created_at, updated_at) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING *",
		device.DeviceID,
		device.Name,
		device.Brand,
		device.IsActive,
		device.CreatedAt,
		device.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (r DeviceRepository) Update(device *model.Device) error {
	_, err := r.db.Exec("UPDATE devices "+
		"SET name = $1, "+
		"brand = $2, "+
		"is_active = $3, "+
		"updated_at = $4 "+
		"WHERE device_id = $5",
		device.Name, device.Brand, device.IsActive, time.Now(), device.DeviceID)
	return err
}

func (r DeviceRepository) UpdateStatus(deviceID uuid.UUID, status bool) error {
	_, err := r.db.Exec("UPDATE devices "+
		"SET status = $1, "+
		"updated_at = $2 "+
		"WHERE device_id = $3",
		status, time.Now(), deviceID)
	return err
}

func (r DeviceRepository) Delete(deviceID uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM devices "+
		"WHERE device_id = $1", deviceID)
	return err
}
