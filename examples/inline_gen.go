// Code generated by github.com/alextanhongpin/getter, DO NOT EDIT.

package examples

func (i InlineParent) Name() string {
	return i.InlineChildren.Name
}

func (i InlineParent) Age() int64 {
	return i.InlineChildren.age
}
