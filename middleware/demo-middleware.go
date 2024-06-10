package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	ID         int     `json: "id"`
	Name       string  `json: "name"`
	Price      float64 `json: "price"`
	Instructor string  `json: "instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{
			"id": 1,	
			"name": "Java",
			"price": 1000,
			"instructor": "John"
		},
		{

			"id": 2,	
			"name": "Python",
			"price": 1200,
			"instructor": "Doe"

		}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}

}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Contend-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(bodyBytes, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newCourse.ID = getNextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		log.Println(CourseList)

		return

	}

}

func getNextID() int {
	highestID := -1
	for _, course := range CourseList {
		if course.ID > highestID {
			highestID = course.ID
		}
	}
	return highestID + 1
}

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.ID == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("Course with ID %v not found", ID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		courseJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Contend-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPut:
		var updatedCourse Course
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedCourse.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updatedCourse
		CourseList[listItemIndex] = *course
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("berfore handler")

		handler.ServeHTTP(w, r)
		fmt.Println("finish handler")
	})
}

func main() {
	courseItemHandler := http.HandlerFunc(courseHandler)
	courseListHandler := http.HandlerFunc(coursesHandler)
	http.Handle("/course/", middlewareHandler(courseItemHandler))
	http.Handle("/course", middlewareHandler(courseListHandler))
	http.ListenAndServe(":5000", nil)

}
