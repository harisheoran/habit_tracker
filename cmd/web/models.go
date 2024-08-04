package main

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	Title    string
	Action   string
	Duration string
}
