package gorusgender

import "testing"

func TestGetGender(t *testing.T) {
	tests := []struct {
		name       string
		firstName  string
		middleName string
		lastName   string
		want       Gender
	}{
		{
			name:       "Male by first name",
			firstName:  "Александр",
			middleName: "",
			lastName:   "",
			want:       GenderMale,
		},
		{
			name:       "Female by first name",
			firstName:  "Екатерина",
			middleName: "",
			lastName:   "",
			want:       GenderFemale,
		},
		{
			name:       "Male by middle name",
			firstName:  "",
			middleName: "Иванович",
			lastName:   "",
			want:       GenderMale,
		},
		{
			name:       "Female by middle name",
			firstName:  "",
			middleName: "Сергеевна",
			lastName:   "",
			want:       GenderFemale,
		},
		{
			name:       "Male by last name ending",
			firstName:  "",
			middleName: "",
			lastName:   "Иванов",
			want:       GenderMale,
		},
		{
			name:       "Female by last name ending",
			firstName:  "",
			middleName: "",
			lastName:   "Иванова",
			want:       GenderFemale,
		},
		{
			name:       "Male with mixed male indicators",
			firstName:  "Александр",
			middleName: "Иванович",
			lastName:   "Иванов",
			want:       GenderMale,
		},
		{
			name:       "Female with mixed female indicators",
			firstName:  "Екатерина",
			middleName: "Сергеевна",
			lastName:   "Иванова",
			want:       GenderFemale,
		},
		{
			name:       "Unknown when conflicting",
			firstName:  "Александр",
			middleName: "Сергеевна",
			lastName:   "",
			want:       GenderUnknown,
		},
		{
			name:       "Unknown when no matches",
			firstName:  "Qwerty",
			middleName: "Asdfg",
			lastName:   "Zxcvb",
			want:       GenderUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetGender(tt.firstName, tt.middleName, tt.lastName)
			if got != tt.want {
				t.Errorf("GetGender(%q, %q, %q) = %v, want %v",
					tt.firstName, tt.middleName, tt.lastName, got, tt.want)
			}
		})
	}
}
