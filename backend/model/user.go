package model

type User struct {
	ID             string        `dynamodbav:"id"`
	Name           string        `dynamodbav:"name"`
	Surname        string        `dynamodbav:"surname"`
	Email          string        `dynamodbav:"email"`
	Image          string        `dynamodbav:"image"`
	UsersCourses   []UsersCourse `dynamodbav:"users_courses"`
	CreatedCourses []Course      `dynamodbav:"created_courses"`
	Username       string        `dynamodbav:"username"`
	Password       string        `dynamodbav:"password"`
}
