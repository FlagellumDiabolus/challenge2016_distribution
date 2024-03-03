package main

import (
	"challenge2016_distribution/distribution"
	"challenge2016_distribution/distribution/permissions"
	"challenge2016_distribution/location"
	"challenge2016_distribution/main/validator"
	"challenge2016_distribution/movies"
	"challenge2016_distribution/producer"
	"fmt"
	"strings"
)

func main() {
	var distributors []distribution.Distributor
	var films []movies.Movie
	var producers []producer.Producer
	movie := movies.CreateMovie("Tenet")
	prod := producer.CreateProducer("Nolan", distributors)
	movie.AddProducer(prod)

	groupedData := location.ReadLocations() // Parsing the CSV file containing city data
	for {
		choice := PromptMenu() // Asking the user for their choice
		switch choice {
		case "Create a movie":
			movie, prod := PromptMovie()
			films = append(films, movie)
			producers = append(producers, prod)
		case "Create a new distributor":
			distributorData := PromptDistributorData(false)
			// Getting data for a new distributor and send false for this the distributor Data
			errorRes := validator.ValidateDistributorData(distributorData, groupedData, distributors, false) // Validating distributor data
			if len(errorRes) > 0 {
				fmt.Println(strings.Join(errorRes, "\n"))
				continue
			}
			distributors = append(distributors, distribution.Distributor{Info: distributorData}) // Appending the new distributor to the list
		case "Create a sub distributor":
			subDistributorData := PromptDistributorData(true)                                                  // Getting data for a new sub-distributor and send true for this the sub-distributor Data
			errorRes := validator.ValidateDistributorData(subDistributorData, groupedData, distributors, true) // Validating sub-distributor data
			if len(errorRes) > 0 {
				fmt.Println(strings.Join(errorRes, "\n"))
				continue
			}
			distributors = append(distributors, distribution.Distributor{Info: subDistributorData}) // Appending the new sub-distributor to the list
		case "Check permission for a distributor":
			checkPermissionData := PromptCheckPermissionData()                                                // Getting data to check permission
			errorRes := validator.ValidateCheckPermissionData(checkPermissionData, groupedData, distributors) // Validating permission check data
			if len(errorRes) > 0 {
				fmt.Println(strings.Join(errorRes, "\n"))
				continue
			}
			dist := distribution.Distributor{}
			for _, distributor := range distributors {
				if checkPermissionData.DistributorName == distributor.Info.Name {
					dist = distributor
				}
			}

			for _, region := range checkPermissionData.Regions {
				permissions.CheckPermission(dist, region)
			}
		case "Add distributor to a producer":
			producerName, distributorNames := PromptAddDistributorToProducer()
			fmt.Println("Only adding those distributors which are already present")
			for _, pro := range producers {
				if pro.Name == producerName {
					for _, name := range distributorNames {
						for _, distributor := range distributors {
							if name == distributor.Info.Name {
								pro.Distributors = append(pro.Distributors, distributor)
							}
						}
					}
				}
			}

		case "View Distributors information":
			for _, distributor := range distributors {
				fmt.Printf("Name: %s, Include: %v, Exclude: %v, Parent: %s\n", distributor.Info.Name, distributor.Info.Permissions.Include, distributor.Info.Permissions.Exclude, distributor.Info.Parent)
			}

		case "Exit":
			fmt.Println("Exiting the program")
			return // Exiting the program
		}
	}
}
