package domain

import (
	"context"
	"fmt"
)

type farm struct {
	farmerEC 		*EggEventCreator
	chickenECList 	[]*EggEventCreator
	eggET
	eggRepository
}

func NewFarm(farmerEC *EggEventCreator, chickenECList []*EggEventCreator, eggET eggET, eggRepository eggRepository) *farm {
	return &farm{farmerEC: farmerEC, chickenECList: chickenECList, eggET: eggET, eggRepository: eggRepository}
}


type eggET interface {
	Write(ee *EggEvent)
	Read() <-chan *EggEvent
}

type eggRepository interface {
	PutAnEgg(egg uint) (err error)
	TakeEggs(eggs uint) (err error)
	CountEggs() (eggs uint)
}


func (f *farm) FarmStart(ctx context.Context) {

	var (
		err 	error
	)

	fctx, cancel := context.WithCancel(ctx)
	defer cancel()

	f.farmerEC.Run(fctx, f.eggET)
	for _, ec := range f.chickenECList {
		ec.Run(fctx, f.eggET)
	}

	for event := range f.eggET.Read() {
		switch event.Event() {
		case ChickenEvent:
			err = f.eggRepository.PutAnEgg(event.EggCount())
			fmt.Printf("%v put: %v; Eggs: %v\n", event.Event(), event.EggCount(), f.CountEggs())
		case FarmerEvent:
			err = f.eggRepository.TakeEggs(event.EggCount())
			fmt.Printf("%v take: %v; Eggs: %v\n", event.Event(), event.EggCount(), f.CountEggs())
		}
		if err != nil {
			cancel()
		}
	}
}




