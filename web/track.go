package main

import (
	"fmt"
	"sort"
)

type Track struct {
	Id int
	Year int
}

type CustomSort struct {
	t []*Track
	less func(x, y *Track) bool
}

func (self CustomSort) Len() int { return len(self.t) }
func (self CustomSort) Less(i, j int) bool { return self.less(self.t[i], self.t[j]) }
func (self CustomSort) Swap(i, j int) { self.t[i], self.t[j] = self.t[j], self.t[i] }

func main() {
	tracks := []*Track{
		{1, 2010},
		{2, 2011},
		{3, 2012},
		{4, 2010},
		{5, 2008},
	}
	// sortのSortは引数を破壊的に並べ替えるので注意。並べ替えた結果を返却するわけではない。
	sort.Sort(CustomSort{tracks, func(x, y *Track) bool { 
		if x.Year != y.Year{
			return x.Year < y.Year
		}
		return false
	}})
	
	for _, x := range tracks{
		fmt.Println(x)
		// &{5 2008}
		// &{1 2010}
		// &{4 2010}
		// &{2 2011}
		// &{3 2012}
	}
}
