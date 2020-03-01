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

type VehicleType struct {

	Id string `json:"id,omitempty"`

	Make string `json:"make,omitempty"`

	Model string `json:"model,omitempty"`

	Color string `json:"color,omitempty"`

	YearOfManufacture int32 `json:"yearOfManufacture,omitempty"`

	PassengerCount string `json:"passengerCount,omitempty"`

	// CO2Emissions per gram per mile
	CO2Emissions string `json:"CO2Emissions,omitempty"`
}
