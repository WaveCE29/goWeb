package main

import (
	"encoding/json"
	"fmt"
)

// Course represents a course with its details.
type Course struct {
	CourseID   int     `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      float64 `json:"price"`
	ImageURL   string  `json:"imageurl"`
}

func main() {
	// Example usage
	course := Course{
		CourseID:   101,
		CourseName: "Introduction to Go",
		Price:      29.99,
		ImageURL:   "https://example.com/course-image.jpg",
	}

	// Convert Course struct to JSON
	courseJSON, err := json.Marshal(course)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("Course JSON:", string(courseJSON))

	// Convert JSON to Course struct
	var newCourse Course
	err = json.Unmarshal(courseJSON, &newCourse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("Course struct: %+v\n", newCourse)
}
