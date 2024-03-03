package validator

import (
	"challenge2016_distribution/distribution"
	"challenge2016_distribution/distribution/permissions"
	"challenge2016_distribution/distribution/permissions/models"
	"challenge2016_distribution/location"
	"challenge2016_distribution/movies"
	"challenge2016_distribution/producer"
	"strings"
)

// The function `ValidateDistributorData` validates the data of a sub-distributor, checking for
// errors such as empty fields, duplicate names, invalid regions, and incorrect parent distributor
// name.
func ValidateDistributorData(data distribution.Info, groupedData []location.Country, distributors []distribution.Distributor, subDistributor bool) []string {
	var errorMsg []string

	if strings.TrimSpace(data.Name) == "" {
		errorMsg = append(errorMsg, "Distributor Name must not be empty, please enter a valid distributor name")
	} else if ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.Name)), distributors) {
		errorMsg = append(errorMsg, "Distributor Name already exists")
	}

	if len(data.Permissions.Include) == 0 {
		errorMsg = append(errorMsg, "Include Regions must not be empty, please enter valid regions")
	} else {
		for _, region := range data.Permissions.Include {
			if ValidateRegion(region, groupedData) {
				errorMsg = append(errorMsg, "Include Region '"+region+"' is not present in csv, please enter a valid region")
			}
		}
	}

	if len(data.Permissions.Exclude) > 0 {
		for _, region := range data.Permissions.Exclude {
			if ValidateRegion(region, groupedData) {
				errorMsg = append(errorMsg, "Exclude Region '"+region+"' is not present in csv, please enter a valid region")
			}
		}
	}

	if len(data.Permissions.Exclude) > 0 && len(data.Permissions.Include) > 0 {
		for _, ExcludeRegion := range data.Permissions.Exclude {
			for _, IncludeRegion := range data.Permissions.Include {
				if strings.EqualFold(ExcludeRegion, IncludeRegion) {
					errorMsg = append(errorMsg, "Exclude Region and Include Region should not be Same, please enter a valid region")
				}

			}

		}
	}

	if subDistributor {
		if strings.TrimSpace(data.Parent) == "" {
			errorMsg = append(errorMsg, "Parent distributor Name must not be empty, please enter a valid parent distributor name")
		} else if !ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.Parent)), distributors) {
			errorMsg = append(errorMsg, "Parent distributor Name does not exist, please enter an existing parent distributor name")
		}

		if len(errorMsg) == 0 {
			InputData := append(data.Permissions.Include, data.Permissions.Exclude...)
			checkPermissionWithParent := permissions.CheckPermissionWithAll(strings.TrimSpace(data.Parent), InputData, "subDistributionCreation", distributors)
			if len(checkPermissionWithParent) > 0 {
				errorMsg = append(errorMsg, checkPermissionWithParent...)
			}
		}
	}

	return errorMsg
}

// The function "ValidateDistributorName" checks if a given distributor name exists in a list of
// distributor information.
func ValidateDistributorName(distributorName string, distributorInformation []distribution.Distributor) bool {
	for _, distributor := range distributorInformation {
		if strings.EqualFold(distributor.Info.Name, distributorName) {
			return true
		}
	}
	return false
}

// The function `ValidateRegion` checks if a given region is valid based on a list of grouped data.
func ValidateRegion(reg string, groupedData []location.Country) bool {
	splitTestData := strings.Split(reg, ",")
	for _, region := range splitTestData {
		InputData := strings.Split(strings.ToUpper(region), "-")

		if len(InputData) > 0 && len(InputData) < 4 {
			switch len(InputData) {
			case 1:
				for _, country := range groupedData {
					if strings.EqualFold(country.Name, InputData[0]) {
						return false
					}
				}
			case 2:
				for _, country := range groupedData {
					if strings.EqualFold(country.Name, InputData[1]) {
						for _, state := range country.States {
							if strings.EqualFold(state.Name, InputData[0]) {
								return false
							}
						}
					}
				}
			case 3:
				for _, country := range groupedData {
					if strings.EqualFold(country.Name, InputData[2]) {
						for _, state := range country.States {
							if strings.EqualFold(state.Name, InputData[1]) {
								for _, city := range state.Cities {
									if strings.EqualFold(city.Name, InputData[0]) {
										return false
									}
								}
							}
						}
					}
				}
			default:
				return true
			}
		} else {
			return true
		}
	}
	return true
}

// The function `ValidateCheckPermissionData` validates the `CheckPermissionData` object by checking if
// the distributor name is not empty and exists in the distributor information, and if all the regions
// in the data exist in the grouped data.
func ValidateCheckPermissionData(data models.CheckPermissionData, groupedData []location.Country, distributorInformation []distribution.Distributor) []string {
	var errorMsg []string

	if strings.TrimSpace(data.DistributorName) == "" {
		errorMsg = append(errorMsg, "Distributor Name must not be empty, please enter a valid distributor name")
	} else if !ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.DistributorName)), distributorInformation) {
		errorMsg = append(errorMsg, "Distributor name does not exist")
	}

	for _, region := range data.Regions {
		if ValidateRegion(region, groupedData) {
			errorMsg = append(errorMsg, strings.ToUpper(region)+" does not exist in the csv file, please enter a valid region")
		}
	}

	return errorMsg
}

func ValidateMovie(movie movies.Movie, prod producer.Producer, allMovies []movies.Movie, producers []producer.Producer) []string {
	var errorMsg []string

	if strings.TrimSpace(movie.Name) == "" {
		errorMsg = append(errorMsg, "Producer Name must not be empty, please enter a valid producers name")
	}
	for _, mov := range allMovies {
		if mov.Name == movie.Name {
			errorMsg = append(errorMsg, "Movie already exists "+movie.Name)
		}
	}

	return errorMsg

}

func ValidateProducerName(ProducerName string, producers []producer.Producer) bool {
	for _, producerr := range producers {
		if strings.EqualFold(producerr.Name, ProducerName) {
			return true
		}
	}
	return false
}
