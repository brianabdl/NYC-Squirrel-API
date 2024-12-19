package main

import (
	"fmt"
	"strings"
	"reflect"
)

var cache = NewCache()

func generateCacheKey(prefix, value string) string {
	return fmt.Sprintf("%s:%s", prefix, value)
}

func findByID(records []SquirrelData, id string) (SquirrelData, error) {
	for _, data := range records {
		if data.UniqueSquirrelID == id {
			return data, nil
		}
	}
	return SquirrelData{}, fmt.Errorf("Unable to find data by ID")
}

func findByField(records []SquirrelData, prefix string, cond string) ([]SquirrelData, error) {
	key := generateCacheKey(prefix, cond)
	
	if cachedResult, exists := cache.Get(key); exists {
		return cachedResult, nil
	}

	var result []SquirrelData
	for _, data := range records {
		v := reflect.ValueOf(data)
		field := v.FieldByName(prefix)
		if field.IsValid() && field.String() == cond {
			result = append(result, data)
		}
	}
	
	if len(result) == 0 {
		return nil, fmt.Errorf("Unable to find data by %s: %s", prefix, cond)
	}
	
	cache.Set(key, result)
	return result, nil
}

func findByAge(records []SquirrelData, age string) ([]SquirrelData, error) {
	return findByField(records, "Age", age)
}

func findByPrimaryFurColor(records []SquirrelData, color string) ([]SquirrelData, error) {
	return findByField(records, "PrimaryFurColor", color)
}

func findByHighlightFurColor(records []SquirrelData, color string) ([]SquirrelData, error) {
	return findByField(records, "HighlightFurColor", color)
}

func findByLocation(records []SquirrelData, loc string) ([]SquirrelData, error) {
	return findByField(records, "Location", loc)
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