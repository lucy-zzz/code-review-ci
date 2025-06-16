package repository

import (
	"app/internal"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) GetAverageSpeedByBrand(b string) (speed float64, err error) {
	var sum float64
	var sumCount int64

	for _, v := range r.db {
		if v.Brand == b {
			sum += v.MaxSpeed
			sumCount++
		}
	}

	if sumCount == 0 {
		return speed, fmt.Errorf("")
	}

	speed = sum / float64(sumCount)

	return speed, err
}
