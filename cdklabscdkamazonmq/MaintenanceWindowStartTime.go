package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Start time of the weekly, 2-hours time window to apply pending updates or patches to the broker.
// Experimental.
type MaintenanceWindowStartTime struct {
	// The day of the week.
	// Experimental.
	DayOfWeek DayOfWeek `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The time, in 24-hour format.
	// Experimental.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// The time zone.
	// Experimental.
	TimeZone awscdk.TimeZone `field:"required" json:"timeZone" yaml:"timeZone"`
}

