package location

import (
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

func ParseCSVFile(csvFilePath string) ([]Country, error) {
	locationsFile, err := os.OpenFile(csvFilePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer locationsFile.Close()

	locations := []Location{}

	if err := gocsv.UnmarshalFile(locationsFile, &locations); err != nil { // Load clients from file
		return nil, err
	}

	groupedData := make([]Country, 0)

	for _, location := range locations {
		countryName := strings.ToUpper(location.CountryName)
		stateName := strings.ToUpper(location.ProvinceName)
		cityName := strings.ToUpper(location.CityName)

		var countryIndex int
		countryExists := false
		for i, country := range groupedData {
			if strings.EqualFold(country.Name, countryName) {
				countryIndex = i
				countryExists = true
				break
			}
		}

		if !countryExists {
			newCountry := Country{
				Name: countryName,
				States: []State{
					{
						Name:   stateName,
						Cities: []City{{Name: cityName}},
					},
				},
			}
			groupedData = append(groupedData, newCountry)
		} else {
			var stateIndex int
			stateExists := false
			for i, state := range groupedData[countryIndex].States {
				if strings.EqualFold(state.Name, stateName) {
					stateIndex = i
					stateExists = true
					break
				}
			}

			if !stateExists {
				newState := State{
					Name:   stateName,
					Cities: []City{{Name: cityName}},
				}
				groupedData[countryIndex].States = append(groupedData[countryIndex].States, newState)
			} else {
				groupedData[countryIndex].States[stateIndex].Cities = append(groupedData[countryIndex].States[stateIndex].Cities, City{Name: cityName})
			}
		}
	}
	return groupedData, nil
}
