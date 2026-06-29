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

func (iE *IntegrationEvent) GetNextQueue() (string, error) {
	return "", nil
}

func (iE *IntegrationEvent) AddMetaHeader(
	source string,
	eventName string,

) {
	iE.MetaHeader = append(iE.MetaHeader, MetaHeader{
		Source:    source,
		EventName: eventName,
	})
}

func (mH *MetaHeader) AddArgs(
	key string,
	value string,
) {
	mH.Args = append(mH.Args, Args{
		Key:   key,
		Value: value,
	})
}

func (eP *IntegrationEvent) ExchangePayload(
	payload []byte,
) {
	eP.Payload = payload
}
