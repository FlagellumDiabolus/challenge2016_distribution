package location

import (
	"log"
)

// ReadLocations to be called in the driver function
func ReadLocations() []Country {
	regions, err := ParseCSVFile("location/cities.csv")
	if err != nil {
		log.Fatalf("Error parsing locations file: %v", err) // Exiting if there's an error in parsing CSV
	}
	return regions
}
