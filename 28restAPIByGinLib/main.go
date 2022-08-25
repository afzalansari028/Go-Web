package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json"fullname"'`
	Website  string `json"website"`
}

//fake DB
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

func main() {
	router := gin.Default()

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Sonu Raj", Website: "sonu.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 399, Author: &Author{FullName: "Vikram Ingawale", Website: "vikram.dev"}})

	router.GET("/test", Test)
	router.GET("/courses", GetAllCourses)
	router.GET("/course/:id", GetOneCourse)
	router.POST("/course", CreateOneCourse)
	router.PUT("/course/:id", UpdateOneCourse)
	router.DELETE("/course/:id", DeleteOneCourse)

	router.Run() //port:8080

}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to gin"})
}

func GetAllCourses(c *gin.Context) {
	fmt.Println("Get all courses")
	c.JSON(http.StatusOK, courses)
	// return
}

func GetOneCourse(c *gin.Context) {
	fmt.Println("Get One course")

	for _, course := range courses {
		if course.CourseId == c.Param("id") {
			c.JSON(200, course) //statusOk
			return
		}
	}
	c.String(http.StatusNotFound, "Course not available") //404
}

func CreateOneCourse(c *gin.Context) {
	fmt.Println("Create one course")
	var course Course
	// jsonDataBytes, _ := ioutil.ReadAll(c.Request.Body)
	// json.Unmarshal(jsonDataBytes, &course)
	json.NewDecoder(c.Request.Body).Decode(&course)

	// validate for empty object{}
	if course.IsEmpty() {
		c.String(200, "Please send content") //204
		// c.JSON(404,"Please send content")
		// log.Print("user sending empty object...")
		return
	}

	//validate for empty courseName
	if course.CourseName == "" {
		c.String(http.StatusConflict, "Please send with course name") //204
		// log.Print("testing...........")
		return
	}

	// validate duplicate data client sending
	for _, courseT := range courses {
		if courseT.CourseName == course.CourseName {
			c.String(http.StatusConflict, "This course is already exist") //409
			return
		}
	}
	//  append course into courses
	courses = append(courses, course)
	c.JSON(200, course)
}

func UpdateOneCourse(c *gin.Context) {
	fmt.Println("Update one course")
	//loop , match id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == c.Param("id") {
			courses = append(courses[:index], courses[index+1:]...)
			var NewCourse Course
			json.NewDecoder(c.Request.Body).Decode(&NewCourse)
			courses = append(courses, NewCourse)
			c.JSON(200, NewCourse)
		}
	}
}

func DeleteOneCourse(c *gin.Context) {
	fmt.Println("Delete one course")
	//iterate, match id, delete
	for index, course := range courses {
		if course.CourseId == c.Param("id") {
			courses = append(courses[:index], courses[index+1:]...)
			c.JSON(200, course)
			break
		}
	}
	c.String(404, "Please send the valid id")
}
