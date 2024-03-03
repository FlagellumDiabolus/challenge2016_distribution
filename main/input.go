package main

import (
	"challenge2016_distribution/distribution"
	"challenge2016_distribution/distribution/permissions/models"
	"challenge2016_distribution/movies"
	"challenge2016_distribution/producer"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"strings"
)

// The `PromptMenu` function in Go displays a menu for selecting different choices and returns the
// selected option.
func PromptMenu() string {
	fmt.Println("Please specify the regions you wish to include or exclude for this distributor (use hyphens for specifying location hierarchy, e.g., Chennai-Tamil Nadu-India, Karnataka-India)")
	prompt := promptui.Select{
		Label: "Select one of the below choices",
		Items: []string{
			"Add a new movie",
			"Create a new distributor",
			"Create a sub distributor",
			"Check permission for a distributor",
			"Add distributors to a producer",
			"View Distributors information",
			"Exit",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

func PromptMovie() (movies.Movie, producer.Producer) {
	var movie movies.Movie
	var crew []movies.Crew

	promptName := promptui.Prompt{
		Label:       "Enter movie name:",
		HideEntered: true,
	}
	name, _ := promptName.Run()
	movie.Name = name
	fmt.Println(promptName.Label, name)

	promptCrew := promptui.Prompt{
		Label:       "Enter the Crews roles like technician, makeup, etc",
		HideEntered: true,
	}
	crewInput, _ := promptCrew.Run()
	crew = make([]movies.Crew, 0)

	for _, role := range strings.Split(crewInput, ",") {
		person := movies.Person{Role: role}
		crew = append(crew, person)
	}
	movie.Crew = crew
	fmt.Println(promptCrew.Label, crewInput)

	promptProducer := promptui.Prompt{
		Label:       "Enter the Producer Name like Nolan or Spielberg or Guru Dutt",
		HideEntered: true,
	}
	prodInput, _ := promptProducer.Run()
	prod := producer.CreateProducer(prodInput, nil)
	movie.AddProducer(prod)
	fmt.Println(promptProducer.Label, prodInput)
	return movie, prod
}

// The function `PromptDistributorData` in Go prompts the user to enter distributor data including
// name, regions to include and exclude, and optionally the parent distributor name.
func PromptDistributorData(subDistributor bool) distribution.Info {
	var distributor distribution.Distributor

	promptName := promptui.Prompt{
		Label:       "Enter distributor name:",
		HideEntered: true,
	}
	name, _ := promptName.Run()
	distributor.Info.Name = name
	fmt.Println(promptName.Label, name)

	promptInclude := promptui.Prompt{
		Label:       "Enter the regions you want to include for this distributor (comma separated): ",
		HideEntered: true,
	}
	includeInput, _ := promptInclude.Run()
	distributor.Info.Permissions.Include = strings.Split(includeInput, ",")
	fmt.Println(promptInclude.Label, includeInput)

	promptExclude := promptui.Prompt{
		Label:       "Enter the regions you want to exclude for this distributor (comma separated): ",
		HideEntered: true,
	}
	excludeInput, _ := promptExclude.Run()
	distributor.Info.Permissions.Exclude = strings.Split(excludeInput, ",")
	fmt.Println(promptExclude.Label, excludeInput)

	if subDistributor {
		promptParent := promptui.Prompt{
			Label:       "Enter the name of the parent distributor: ",
			HideEntered: true,
		}
		parent, _ := promptParent.Run()
		distributor.Info.Parent = parent
		fmt.Println(promptParent.Label, parent)
	}

	return distributor.Info
}

func PromptAddDistributorToProducer() (string, []string) {
	promptProducerName := promptui.Prompt{
		Label:       "Enter Producer name to whom distributors are to be added:",
		HideEntered: true,
	}
	prodName, _ := promptProducerName.Run()
	fmt.Println(promptProducerName.Label, prodName)

	promptDistributors := promptui.Prompt{
		Label:       "Enter distributor names that are to be added (comma separated):",
		HideEntered: true,
	}
	distInput, _ := promptDistributors.Run()
	distributors := strings.Split(distInput, ",")
	fmt.Println(promptDistributors.Label, distributors)
	return prodName, distributors
}

// The `PromptCheckPermissionData` function in Go prompts the user to enter a distributor name and
// regions for permission checking.
func PromptCheckPermissionData() models.CheckPermissionData {
	var data models.CheckPermissionData

	promptName := promptui.Prompt{
		Label:       "Enter distributor name that needs to be checked:",
		HideEntered: true,
	}
	data.DistributorName, _ = promptName.Run()
	fmt.Println(promptName.Label, data.DistributorName)

	promptRegions := promptui.Prompt{
		Label:       "Enter distributor name that needs to be checked (comma separated):",
		HideEntered: true,
	}
	regionsInput, _ := promptRegions.Run()
	data.Regions = strings.Split(regionsInput, ",")
	fmt.Println(promptRegions.Label, data.Regions)
	return data
}
