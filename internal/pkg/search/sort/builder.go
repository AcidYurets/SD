package sort

type Builder interface {
	AddSort(field string, asc bool)
}
