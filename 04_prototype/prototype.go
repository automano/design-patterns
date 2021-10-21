package prototype

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{f.name + "_clone"}
}

func (f *file) getName() string {
	return f.name
}

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + "\t")
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, child := range f.children {
		copy := child.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
