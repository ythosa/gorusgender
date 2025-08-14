package gorusgender

import "strings"

func GetGender(firstName string, middleName string, lastName string) Gender {
	partsGenders := []Gender{
		getGenderByFirstName(firstName),
		getGenderByMiddleName(middleName),
		getGenderByLastName(lastName),
	}

	var (
		isMale   bool
		isFemale bool
	)

	for _, gender := range partsGenders {
		switch gender {
		case GenderMale:
			isMale = true
		case GenderFemale:
			isFemale = true
		}
	}

	switch {
	case isMale && !isFemale:
		return GenderMale
	case !isMale && isFemale:
		return GenderFemale
	default:
		return GenderUnknown
	}
}

func getGenderByFirstName(firstName string) Gender {
	firstName = normalizeNamePart(firstName)

	if _, ok := namesByGender[GenderMale][firstName]; ok {
		return GenderMale
	}

	if _, ok := namesByGender[GenderFemale][firstName]; ok {
		return GenderFemale
	}

	return GenderUnknown
}

func getGenderByLastName(lastName string) Gender {
	lastName = normalizeNamePart(lastName)

	switch {
	case isCorrectCompletion(lastName, lastNameCompletionsByGender[GenderMale]):
		return GenderMale
	case isCorrectCompletion(lastName, lastNameCompletionsByGender[GenderFemale]):
		return GenderFemale
	default:
		return GenderUnknown
	}
}

func getGenderByMiddleName(middleName string) Gender {
	middleName = normalizeNamePart(middleName)

	switch {
	case isCorrectCompletion(middleName, middleNameCompletionsByGender[GenderMale]):
		return GenderMale
	case isCorrectCompletion(middleName, middleNameCompletionsByGender[GenderFemale]):
		return GenderFemale
	default:
		return GenderUnknown
	}
}

func isCorrectCompletion(namePart string, completions []string) bool {
	for _, completion := range completions {
		if strings.HasSuffix(namePart, completion) {
			return true
		}
	}

	return false
}

func normalizeNamePart(namePart string) string {
	return strings.ToLower(strings.ReplaceAll(namePart, " ", ""))
}
