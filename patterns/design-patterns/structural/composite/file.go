package composite

type FileSystemNode interface {
	GetSize() int
}

type File struct {
	Size int
}

func (f *File) GetSize() int {
	return f.Size
}

type Directory struct {
	children []FileSystemNode
}

func (d *Directory) GetSize() int {
	size := 0

	for _, child := range d.children {
		size += child.GetSize()
	}

	return size
}

func (d *Directory) AddChild(child FileSystemNode) {
	d.children = append(d.children, child)
}
