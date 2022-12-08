package farm

import (
	"context"
	"farm/service/farm/internal/data/repository"
	"farm/service/farm/internal/domain"
	"farm/service/farm/internal/tools/rand"
	"net/http"
)

type app struct {

}

func NewApp() *app {
	return &app{}
}

func (a *app) Start()  {

	var (
		err error
	)

	http.HandleFunc()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	chickenCount := 3
	ctx := context.Background()

	fridge := repository.NewFridge()

	farmer := domain.NewEggEventCreator(rand.RandomUint(10, 15), rand.RandomUint(10, 20), domain.FarmerEvent)
	chickens := make([]*domain.EggEventCreator, 0, chickenCount)

	for i := 0; i < chickenCount; i++ {
		chicken := domain.NewEggEventCreator(rand.RandomUint(2, 10), rand.RandomUint(2, 5), domain.ChickenEvent)
		chickens = append(chickens, chicken)
	}

	eggET := domain.NewEggEventTransmitter()
	defer eggET.Close()

	farm := domain.NewFarm(farmer, chickens, eggET, fridge)

	farm.FarmStart(ctx)
}




