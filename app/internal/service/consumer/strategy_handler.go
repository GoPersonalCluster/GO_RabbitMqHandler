package consumer

type StrategyHandler interface {
	Start() ([]byte, error)
}
