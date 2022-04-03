package grid

type IGridCell interface {
	SetHolder(interface{})
	GetHolder() interface{}
}
