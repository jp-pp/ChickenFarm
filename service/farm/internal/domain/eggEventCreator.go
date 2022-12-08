package domain

import (
	"context"
	"fmt"
	"time"
)

type EggEventCreator struct {
	sleepSec 		uint
	eggCount		uint
	eventCreator 	string
}

func NewEggEventCreator(sleepSec uint, eggCount uint, eventCreator string) *EggEventCreator {
	return &EggEventCreator{sleepSec: sleepSec, eggCount: eggCount, eventCreator: eventCreator}
}

type eggEventWriter interface {
	Write(ee *EggEvent)
}

func (c *EggEventCreator) Run(ctx context.Context, eggEventWriter eggEventWriter) {

	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("%v stop...\n", c.eventCreator)
			return
		default:
			for {
				time.Sleep(time.Duration(c.sleepSec) * time.Second)
				eggEventWriter.Write(&EggEvent{
					event:    c.eventCreator,
					eggCount: c.eggCount,
				})
			}
		}
	}()
}
