package error_system

type Error interface {
	Id() string
	Error() string //golang error interface
	Child() Error
	HierarchyId() string
}

type ErrorImpl struct {
	id          string
	description string
	child       Error
	hierarchyId string
}

func (err ErrorImpl) Error() string {
	return err.description
}

func (err ErrorImpl) Id() string {
	return err.id
}

func (err ErrorImpl) Child() Error {
	return err.child
}

func (err ErrorImpl) HierarchyId() string {
	return err.hierarchyId
}

func NewError(id string, description string, child Error) ErrorImpl {
	hierarchyId := id
	desc := description
	if child != nil {
		hierarchyId += child.HierarchyId()
		desc += ": " + child.Error()
	}
	return ErrorImpl{id: id, description: description, child: child, hierarchyId: hierarchyId}
}
