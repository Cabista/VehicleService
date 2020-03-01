/*
 * VehicleService
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: johnfg2610\"gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Vehicle - Represents a vehicle
type Vehicle struct {

	Id string `json:"id,omitempty"`

	VehicleInformation VehicleType `json:"vehicleInformation,omitempty"`

	PlateNumber float32 `json:"plateNumber,omitempty"`

	Millage int32 `json:"millage,omitempty"`

	VehicleExpiry VehicleExpiry `json:"vehicleExpiry,omitempty"`

	BusinessID string `json:"businessID,omitempty"`

	CreationDate string `json:"creationDate,omitempty"`

	DriverID string `json:"driverID,omitempty"`

	IsActive bool `json:"isActive,omitempty"`
}
