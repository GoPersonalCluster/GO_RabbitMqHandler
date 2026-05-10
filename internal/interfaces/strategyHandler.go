package interfaces

type StrategyHandler interface {
	Start() (string, error)
}
