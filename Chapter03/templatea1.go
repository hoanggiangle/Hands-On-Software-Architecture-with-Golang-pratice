package main

import "fmt"

type masterAlgorithm struct {
	template  TemplateA
	templateB TemplateB
}

//TemplateA interface to resolve problem A
type TemplateA interface {
	StepA1()
	StepA
	
	x=?=2 10$-6
	2 4 7 3    k = 4
	0 1 1 0 
	1 2 3 1
	
	5 12 3 5 3 12
//TemplateB abc
type TemplateB interface {
	StepB1()
	StepB2()
}

//VariantA action
type VariantA struct{}

//StepA1 abc
func (v *VariantA) StepA1() {
	fmt.Println("StepA1")
}

//StepA2 abc
func (v *VariantA) StepA2() {
	fmt.Println("StepA2")
}

//VariantB abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
type VariantB struct{}

//StepB1 abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
func (n *VariantB) StepB1() {
	fmt.Println("StepB1")
}

//StepB2 a
func (n *VariantB) StepB2() {
	fmt.Println("StepB2")
}
func (t *masterAlgorithm) RunTemplateA() {
	t.template.StepA1()
	t.template.StepA2()
}

func (t *masterAlgorithm) RunTemplateB() {
	t.templateB.StepB2()
	t.templateB.StepB1()
}
func main() {
	var myZoo = masterAlgorithm{template: new(VariantA), templateB: new(VariantB)}
	myZoo.RunTemplateA()

	myZoo.RunTemplateB()
}
