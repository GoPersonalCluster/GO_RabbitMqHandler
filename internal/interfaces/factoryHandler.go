package interfaces

type FactoryHandler interface {
	GetStrategy() (StrategyHandler, error)
}
