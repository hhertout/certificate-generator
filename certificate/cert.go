package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 30

type Certificate struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Certificate) error
}

func New(course, name, date string) (*Certificate, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Certificate{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is Presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

// Validate the title of the course
func validateCourse(course string) (string, error) {
	course, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(course, " course") {
		course = course + " course"
	}

	return strings.ToTitle(course), nil
}

// Validate the name of person
func validateName(name string) (string, error) {
	name, err := validateStr(name, MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(name), nil
}

// Parse the date to the correct format
func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}

// Validate the format string for the certificate
func validateStr(str string, maxLength int) (string, error) {
	str = strings.TrimSpace(str)
	if len(str) <= 0 {
		return str, fmt.Errorf("invalid string length, got %s, len=%v", str, len(str))
	} else if len(str) > maxLength {
		return str, fmt.Errorf("invalid string lenght, max length is %v, got %v", len(str), maxLength)
	}

	return str, nil
}
