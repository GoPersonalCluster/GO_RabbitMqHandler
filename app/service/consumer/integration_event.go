package consumer

type RabbitMQHandlerConfig struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Queue           string
	AbstractFactory FactoryHandler
}

type IntegrationEvent struct {
	EventName  string
	Payload    []byte
	OccuredAt  int64
	MetaHeader []MetaHeader
}
type MetaHeader struct {
	Source    string
	EventName string
	Args      []Args
	OccuredAt int64
}

type Args struct {
	Key   string
	Value string
}

func (gNQ *IntegrationEvent) GetNextQueue() (string, error) {
	return "", nil
}

func (eP *IntegrationEvent) AddMetaHeader(
	source string,
	eventName string,
	args []struct {
		Key   string
		Value string
	},

) {

}
func (eP *IntegrationEvent) ExchangePayload(
	payload []byte,
) {
	eP.Payload = payload
}
