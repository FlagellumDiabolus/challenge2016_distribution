package permissions

import (
	"challenge2016_distribution/distribution"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

func CheckPermission(distributor distribution.Distributor, region string) bool {
	for _, exclude := range distributor.Info.Permissions.Exclude {
		if strings.HasPrefix(region, exclude) {
			return false
		}
	}

	for _, include := range distributor.Info.Permissions.Include {
		if strings.HasPrefix(region, include) {
			return true
		}
	}
	fmt.Printf("Permission not there for" + region)
	return false
}

// The CheckPermission function checks if a distributor has access to certain test data based on their
// inclusion and exclusion lists.
func CheckPermissionWithAll(distributorName string, InputData []string, origin string, distributorInformation []distribution.Distributor) []string {
	var validationResult []string
	var errorMsg []string

	// Get distributor data by name
	var distributorData distribution.Distributor
	for _, distributor := range distributorInformation {
		if strings.EqualFold(distributor.Info.Name, distributorName) {
			distributorData = distributor
			break
		}
	}

	for _, data := range InputData {
		switch len(strings.Split(data, "-")) {
		case 1:
			if slices.Contains(distributorData.Info.Permissions.Include, data) {
				validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
			} else {
				validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
				errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
			}
		case 2:
			newTestData := strings.Split(data, "-")
			if slices.Contains(distributorData.Info.Permissions.Include, newTestData[1]) {
				if slices.Contains(distributorData.Info.Permissions.Exclude, data) {
					validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
					errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
				} else {
					validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
				}
			} else if slices.Contains(distributorData.Info.Permissions.Include, data) {
				validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
			} else {
				validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
				errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
			}
		case 3:
			newTestData := strings.Split(data, "-")
			if slices.Contains(distributorData.Info.Permissions.Include, newTestData[2]) {
				if slices.Contains(distributorData.Info.Permissions.Include, newTestData[1]+"-"+newTestData[2]) {
					if slices.Contains(distributorData.Info.Permissions.Include, data) {
						validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
					} else {
						if slices.Contains(distributorData.Info.Permissions.Exclude, data) {
							validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
							errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
						} else {
							validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
						}
					}
				} else {
					if slices.Contains(distributorData.Info.Permissions.Exclude, newTestData[1]+"-"+newTestData[2]) {
						validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
						errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
					} else {
						if slices.Contains(distributorData.Info.Permissions.Exclude, data) {
							validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
							errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
						} else {
							validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
						}
					}
				}
			} else {
				if slices.Contains(distributorData.Info.Permissions.Include, newTestData[1]+"-"+newTestData[2]) {
					if slices.Contains(distributorData.Info.Permissions.Exclude, data) {
						validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
						errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
					} else {
						validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
					}
				} else if slices.Contains(distributorData.Info.Permissions.Include, data) {
					validationResult = append(validationResult, distributorData.Info.Name+" has access to "+data)
				} else {
					validationResult = append(validationResult, distributorData.Info.Name+" does not have access to "+data)
					errorMsg = append(errorMsg, distributorData.Info.Name+" does not have access to "+data)
				}
			}
		}
	}

	if origin == "subDistributionCreation" {
		return errorMsg
	}
	return validationResult
}
