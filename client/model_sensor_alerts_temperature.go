/*
 * SensorPush Public API
 *
 * This is a swagger definition for the SensorPush public REST API. Download the definition file [here](https://api.sensorpush.com/api/v1/support/swagger/swagger-v1.json).
 *
 * API version: v1.0.20200621
 * Contact: support@sensorpush.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SensorAlertsTemperature alert settings for temperature
type SensorAlertsTemperature struct {
	// Is enabled?
	Enabled bool `json:"enabled,omitempty"`
	// Max threshold for alert
	Max float32 `json:"max,omitempty"`
	// Min threshold for alert
	Min float32 `json:"min,omitempty"`
}
