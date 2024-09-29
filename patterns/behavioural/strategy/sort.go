package strategy

type Sorter interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (s *BubbleSort) Sort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

type InsertionSort struct{}

func (s *InsertionSort) Sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

type Context struct {
	sorter Sorter
}

func (c *Context) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

func (c *Context) ExecuteSort(array []int) []int {
	return c.sorter.Sort(array)
}
