package main

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Palindrome(s []string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	charCount := make(map[rune]int)
	for _, c := range s1 {
		charCount[c]++
	}
	for _, c := range s2 {
		charCount[c]--
		if charCount[c] < 0 {
			return false
		}
	}
	return true
}

func RemoveDigits(s string) string {
	var result []rune
	for _, c := range s {
		if c < '0' || c > '9' {
			result = append(result, c)
		}
	}
	return string(result)
}

func ReplaceDigits(s string, r string) string {
	var result []rune
	replacement := []rune(r)
	if len(replacement) != 1 {
		panic("replacement should be a single character")
	}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = append(result, replacement[0])
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

type Student interface {
	Name() string
}

type Course interface {
	Name() string
	EnrollStudent(s Student) error
}

type DataSource interface {
	ReadStudent(studentID int) (Student, error)
	ReadCourse(courseID int) (Course, error)
}

func EnrollStudentToCourse(dataSource DataSource, sID, cID int) error {
	student, err := dataSource.ReadStudent(sID)
	if err != nil {
		return err
	}
	course, err := dataSource.ReadCourse(cID)
	if err != nil {
		return err
	}
	return course.EnrollStudent(student)
}