package service

type OuterService interface {
	Hello() string
}

type OuterHandlerImpl struct{}

func NewOuterHandler() OuterService {
	return &OuterHandlerImpl{}
}

func (oh *OuterHandlerImpl) Hello() string {
	return "Hello, this is outer service."
}
