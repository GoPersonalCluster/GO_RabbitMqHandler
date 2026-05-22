package interfaces

type Publisher interface {
	Publish(message []byte) error
}
