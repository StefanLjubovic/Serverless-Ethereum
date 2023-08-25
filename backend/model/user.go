package model

type User struct {
	ID             string        `dynamodbav:"id"`
	Name           string        `dynamodbav:"name"`
	Surname        string        `dynamodbav:"surname"`
	Email          string        `dynamodbav:"email"`
	Image          string        `dynamodbav:"image"`
	UsersCourses   []UsersCourse `dynamodbav:"users_courses"`
	CreatedCourses []int         `dynamodbav:"created_courses"`
	Username       string        `dynamodbav:"username"`
}
