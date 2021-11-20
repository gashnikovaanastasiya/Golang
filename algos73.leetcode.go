package FirstProject
import (
	"fmt"
)
type Set struct {
	Elements map[int]struct{}
}
func NewSet() *Set {
	var set Set
	set.Elements=make(map[int]struct{})
	return &set
}
func (s *Set) Add(elem int) {
	s.Elements[elem]=struct{}{}
}
func (s *Set) Contains(elem int) bool {
	_, exists := s.Elements[elem]
	return exists
}
 func FindZero(matrix [][]int) {
	 rows := NewSet()
	 columns := NewSet()
	 n := len(matrix)
	 m := len(matrix[0])
	 for i := 0; i < n; i++ {
		 for j := 0; j < m; j++ {
			 if matrix[i][j] == 0 {
				 rows.Add(i)
				 columns.Add(j)
			 }
		 }
	 }
	 for i := 0; i < n; i++ {
		 for j := 0; j < m; j++ {
			 if rows.Contains(i) || columns.Contains(j) {
				 matrix[i][j] = 0
			 }
		 }
	 }
	 for i := 0; i < n; i++ {
		 for j := 0; j < m; j++ {
			 fmt.Println(matrix[i][j])
		 }
	 }
 }
