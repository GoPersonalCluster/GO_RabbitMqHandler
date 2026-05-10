package interfaces

type AbstractFactoryHandler interface {
	GetFactory() (StrategyHandler, error)
}
