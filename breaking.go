package main

import "fmt"

func Breaking(dataContractFileName string, stableContractUrl string, pathToType []string, pathToSpecification []string) error {
	all, err := GetDifferences(dataContractFileName, stableContractUrl, pathToType, pathToSpecification)
	if err != nil {
		return err
	}

	breaking := breakingDifferences(all)

	PrintDifferences(breaking)

	if len(breaking) != 0 {
		return fmt.Errorf("found breaking differences between the data contracts")
	}

	return nil
}

func breakingDifferences(allDifferences []DatasetDifference) []DatasetDifference {
	var breakingDifferences []DatasetDifference
	for _, difference := range allDifferences {
		if difference.Severity == DatasetDifferenceSeverityBreaking {
			breakingDifferences = append(breakingDifferences, difference)
		}
	}
	return breakingDifferences
}