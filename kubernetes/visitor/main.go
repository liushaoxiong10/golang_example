package main

import "fmt"

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func() error

type VisitorList []Visitor

func (l VisitorList) Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(func() error {
			fmt.Println("in visitor list before fn")
			fn()
			fmt.Println("in visitor list after fn")
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

type Visitor1 struct{}

func (v Visitor1) Visit(fn VisitorFunc) error {
	fmt.Println("in visitor1 before fn")
	fn()
	fmt.Println("in visitor1 after fn")
	return nil
}

type Visitor2 struct {
	visitor Visitor
}

func (v Visitor2) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("in visitor2 before fn")
		fn()
		fmt.Println("in visitor2 after fn")
		return nil
	})
	return nil
}

type Visitor3 struct {
	visitor Visitor
}

func (v Visitor3) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("in visitor3 before fn")
		fn()
		fmt.Println("in visitor3 after fn")
		return nil
	})
	return nil
}

func main() {
	var visitor Visitor
	var visitors VisitorList

	visitor = Visitor1{}
	visitors = append(visitors, visitor)
	visitor = Visitor2{VisitorList(visitors)}
	visitor = Visitor3{visitor}
	visitor.Visit(func() error {
		fmt.Println("in visitor func")
		return nil
	})

}
