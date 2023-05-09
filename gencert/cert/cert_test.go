package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid expected ='GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azertyzeofkneizejfieezpfjezipofjzpfjzeipfzifjzinvdsn"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course = %s", course)
	}
}

func TestCourseEmptyName(t *testing.T) {
	_, err := New("Golang", "", "2018-05-31")
	if err == nil {
		t.Errorf("error should be returned on an empty name")
	}

}

func TestNameTooLong(t *testing.T) {
	name := "BOBlenomestrrop long logn logn trop trop"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long name name (name = %s", name)
	}
}

func TestDateParseOk(t *testing.T) {
	date := "05-15-2002"
	c, err := New("Golang", "Bob", date)
	if err == nil {
		t.Errorf("Error should be returned, date in wrong format, expected 'YYYY-MM-DD' got=%v, formated=%v", date, c.Date)
	}

	date = "1996-05-15"
	_, err = New("Golang", "Bob", date)
	if err != nil {
		t.Errorf("Error on date was returned but the format is valid")
	}
}
