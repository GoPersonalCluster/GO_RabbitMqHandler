package consumer

type FactoryHandler interface {
	CreateStrategy(body *[]byte) (StrategyHandler, error)
}
