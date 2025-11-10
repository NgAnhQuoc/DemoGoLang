package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Food struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

var arrDataFood = []Food{
	{ID: 1, Name: "Apple", Price: 10000, Description: "Apple is a fruit"},
	{ID: 2, Name: "Banana", Price: 20000, Description: "Banana is a fruit"},
	{ID: 3, Name: "Cherry", Price: 30000, Description: ""},
	{ID: 4, Name: "Orange", Price: 40000, Description: "Orange is a fruit"},
	{ID: 5, Name: "Pineapple", Price: 50000, Description: "Pineapple is a fruit"},
	{ID: 6, Name: "Strawberry", Price: 60000, Description: "Strawberry is a fruit"},
	{ID: 7, Name: "Watermelon", Price: 70000, Description: "Watermelon is a fruit"},
	{ID: 8, Name: "Melon", Price: 80000, Description: "Melon is a fruit"},
	{ID: 9, Name: "Grape", Price: 90000, Description: "Grape is a fruit"},
	{ID: 10, Name: "Pear", Price: 100000, Description: "Pear is a fruit"},
}

// Lấy danh sách hoặc tìm kiếm theo name
func getFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result []Food

	searchName := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("name")))
	keySearch := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("search")))

	if keySearch != "" {
		fmt.Println("keySearch", keySearch)
		for _, food := range arrDataFood {
			if strings.Contains(strings.ToLower(food.Name), keySearch) {
				result = append(result, food)
			}
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	if searchName != "" {
		for _, food := range arrDataFood {
			if strings.Contains(strings.ToLower(food.Name), searchName) {
				result = append(result, food)
			}
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	json.NewEncoder(w).Encode(arrDataFood)
	return
}

// Lấy food theo ID
func getFoodById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Cắt phần /food/ ra để lấy id
	idStr := strings.TrimPrefix(r.URL.Path, "/food/")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, food := range arrDataFood {
		if food.ID == id {
			json.NewEncoder(w).Encode(food)
			return
		}
	}

	http.Error(w, "Food not found", http.StatusNotFound)
}


func main() {
	fmt.Println("Server is running on port 8080")

	http.HandleFunc("/food", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getFood(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/food/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getFoodById(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
