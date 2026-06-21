package consumer

type FactoryHandler interface {
	CreateStrategy(event *IntegrationEvent) (StrategyHandler, error)
}
