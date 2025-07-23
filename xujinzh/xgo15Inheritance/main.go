package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (student *Student) ShowInfo() {
	fmt.Printf("name=%q, age=%v, score=%v\n",
		student.Name, student.Age, student.Score)
}

func (student *Student) SetScore(score int) {
	student.Score = score
}

type Pupil struct {
	Student // 匿名结构体
}

func (pupil *Pupil) Testing() {
	fmt.Println("Pupil testing")
}

type Graduate struct {
	Student // 匿名结构体
}

func (graduate *Graduate) Testing() {
	fmt.Println("Graduate testing")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "xiaoming"
	pupil.Student.Age = 8
	pupil.Student.SetScore(99)
	pupil.Testing()
	pupil.Student.ShowInfo()

	pupil1 := &Pupil{}
	pupil1.Name = "xiaohua"
	pupil1.Age = 9
	pupil1.SetScore(100)
	pupil.Testing()
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "yuz"
	graduate.Student.Age = 28
	graduate.Student.SetScore(100)
	graduate.Testing()
	graduate.Student.ShowInfo()
}
