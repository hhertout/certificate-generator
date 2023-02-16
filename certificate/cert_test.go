package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data shloud be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert shloud be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid, expected='GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Error("Error should be return on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklm"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error, the name of the course is too long, error should be returned (course=%s, len=%v)", course, len(course))
	}
}

func TestNameEmpty(t *testing.T) {
	name := ""
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("An error must be returned, name is empty")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "azertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklm"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned, name length is too long (name=%v, len=%v)", name, len(name))
	}
}
