package domain

const (
	FarmerEvent = "f"
	ChickenEvent = "c"
)


type eggEventTransmitter struct {
	eggEventChanel  chan *EggEvent
}

func NewEggEventTransmitter() *eggEventTransmitter {
	return &eggEventTransmitter{eggEventChanel: make(chan *EggEvent)}
}

func (e *eggEventTransmitter) Write(ee *EggEvent) {
	e.eggEventChanel<-ee
}

func (e *eggEventTransmitter) Read() <-chan *EggEvent {
	return e.eggEventChanel
}

func (e *eggEventTransmitter) Close() {
	close(e.eggEventChanel)
}

type EggEvent struct {
	event 		string
	eggCount	uint
}

func (ee *EggEvent) Event() string {
	return ee.event
}

func (ee *EggEvent) EggCount() uint {
	return ee.eggCount
}


