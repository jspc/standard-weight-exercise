package main

import (
	"fmt"
	"strconv"

	"github.com/dixonwille/wmenu/v5"
)

// There are two age ranges:
//
//  1. LTE35: less than or equal to 35 years old
//  2. GT35: greater than 35
type AgeGroup int

const (
	LTE35 AgeGroup = iota // 0
	GT35                  // 1
)

type Gender int

const (
	Male   Gender = iota // 0
	Female               // 1
)

type Input struct {
	AgeGroup AgeGroup
	Gender   Gender
	Height   float64
	Weight   float64
}

func GetInput() (i Input, err error) {
	err = i.getAge()
	if err != nil {
		return
	}

	err = i.getGender()
	if err != nil {
		return
	}

	err = i.getHeight()
	if err != nil {
		return
	}

	err = i.getWeight()

	return
}

func (i *Input) getAge() error {
	menu := wmenu.NewMenu("Which of these best describes your age?")

	menu.Option("35 or younger", nil, true, func(_ wmenu.Opt) error { i.AgeGroup = LTE35; return nil })
	menu.Option("36 or older", nil, false, func(_ wmenu.Opt) error { i.AgeGroup = GT35; return nil })

	return menu.Run()
}

func (i *Input) getGender() error {
	menu := wmenu.NewMenu("Which of these best describes you?")

	menu.Option("Male", nil, true, func(_ wmenu.Opt) error { i.Gender = Male; return nil })
	menu.Option("Female", nil, false, func(_ wmenu.Opt) error { i.Gender = Female; return nil })

	return menu.Run()
}

func (i *Input) getHeight() (err error) {
	i.Height, err = floatQuestion("What is your height in cm?")

	return
}

func (i *Input) getWeight() (err error) {
	i.Weight, err = floatQuestion("What is your weight in KG?")

	return
}

func floatQuestion(q string) (float64, error) {
	var resp string

	fmt.Println(q)
	fmt.Scanln(&resp)

	return strconv.ParseFloat(resp, 64)
}
