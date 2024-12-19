package main

import (
	"fmt"
	"strings"
	"reflect"
)

var cache = NewCache()

func findByID(records []SquirrelData, id string) (SquirrelData, error) {
	for _, data := range records {
		if data.UniqueSquirrelID == id {
			return data, nil
		}
	}
	return SquirrelData{}, fmt.Errorf("Unable to find data by ID")
}

func findByAge(records []SquirrelData, age string) ([]SquirrelData, error) {
	// Check cache first
	if cachedResult, exists := cache.Get(age); exists {
		return cachedResult, nil
	}
	
	var result []SquirrelData
	for _, data := range records {
		if data.Age == age {
			result = append(result, data)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("Unable to find data by Age")
	}
	
	cache.Set(age, result)
	
	return result, nil
}

func findByPrimaryFurColor(records []SquirrelData, color string) ([]SquirrelData, error) {
	
	key := fmt.Sprintf("primary:%s", color)
	
	// Check cache first
	if cachedResult, exists := cache.Get(key); exists {
		return cachedResult, nil
	}
	
	var result []SquirrelData
	for _, data := range records {
		if data.PrimaryFurColor == color {
			result = append(result, data)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Unable find data by Primary Fur Color")
	}

	cache.Set(key, result)
	return result, nil
}

func findByHighlightFurColor(records []SquirrelData, color string) ([]SquirrelData, error) {
	
	key := fmt.Sprintf("highlight:%s", color)
	
	// Check cache first
	if cachedResult, exists := cache.Get(key); exists {
		return cachedResult, nil
	}
	
	var result []SquirrelData
	for _, data := range records {
		if data.HighlightFurColor == color {
			result = append(result, data)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Unable find data by Highlight Fur Color")
	}

	cache.Set(key, result)
	return result, nil
}


func findByLocation(records []SquirrelData, loc string) ([]SquirrelData, error) {
	
	// Check cache first
	if cachedResult, exists := cache.Get(loc); exists {
		return cachedResult, nil
	}
	
	var result []SquirrelData
	for _, data := range records {
		if data.Location == loc {
			result = append(result, data)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Unable find squirrel at %s", loc)
	}

	cache.Set(loc, result)
	return result, nil
}

func findSquirrelByCondition(records []SquirrelData, cond string) ([]SquirrelData, error) {
	cond = strings.ToLower(cond)
	
	// Check cache first
	if cachedResult, exists := cache.Get(cond); exists {
		return cachedResult, nil
	}
	
	var result []SquirrelData
	for _, data := range records {
		v := reflect.ValueOf(data)
		field := v.FieldByNameFunc(func(name string) bool {
			return strings.ToLower(name) == cond
		})

		if field.IsValid() && field.Bool() {
			result = append(result, data)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Unable find squirrel with condition: %s", cond)
	}

	cache.Set(cond, result)
	return result, nil
}