// Unordered set of 2D points
package percolation

type PointSet struct {
	// The PointSet needs to be able to represent a Point as an int.
	convertFrom1D func(int) (int, int)
	convertTo1D   func(int, int) int
	// data's keys are 1D coordinates covering the grid.  When the value
	// associated with a key is true, that key is part of the point set.
	data map[int]bool
}

// Create a new point set with grid dimensions (Lx, Ly).
func NewPointSet(convertFrom1D func(int) (int, int), convertTo1D func(int, int) int) *PointSet {
	ps := new(PointSet)
	ps.convertFrom1D = convertFrom1D
	ps.convertTo1D = convertTo1D
	ps.data = make(map[int]bool)
	return ps
}

// Return an arbitary point in the set.
// This may be inefficient! Depends on whether range ps.data is lazy or if it
// builds a list of all possible (k, v).
func (ps *PointSet) Point() (int, int) {
	for k, v := range ps.data {
		if v {
			return ps.convertFrom1D(k)
		}
	}
	// there are no points in the set
	return -1, -1
}

// Return the number of points in the set.
func (ps *PointSet) Size() int {
	count := 0
	for _, v := range ps.data {
		if v {
			count++
		}
	}
	return count
}

// Return a slice of all (x, y) points in the set.
func (ps *PointSet) Elements() [][]int {
	elements := [][]int{}
	for k, v := range ps.data {
		if v {
			x, y := ps.convertFrom1D(k)
			elements = append(elements, []int{x, y})
		}
	}
	return elements
}

// Return true if and only if ps and comp are equal point sets.
func (ps *PointSet) Equals(comp *PointSet) bool {
	// different size sets can't be equal
	if ps.Size() != comp.Size() {
		return false
	}
	// check each element
	for _, point := range comp.Elements() {
		x, y := point[0], point[1]
		if !ps.Contains(x, y) {
			return false
		}
	}
	return true
}

// Return true if and only if (x, y) is in the point set.
func (ps *PointSet) Contains(x, y int) bool {
	value, ok := ps.data[ps.convertTo1D(x, y)]
	if ok {
		return value
	}
	return false
}

// Add a point to the set.
func (ps *PointSet) Add(x, y int) {
	ps.data[ps.convertTo1D(x, y)] = true
}

// Remove a point from the set.
func (ps *PointSet) Remove(x, y int) {
	// key is deleted from data
	ps.data[ps.convertTo1D(x, y)] = false, false
}
