package searchinput

type CancelSearch struct{}
type UpdateSearch struct {
	FilteredItemsIndices []int
}
type SubmitSearch struct {
	FilteredItemsIndices []int
}
