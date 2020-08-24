package client

import (
	"errorDemo/octaviusErrors"
	"errors"
	"log"

	"github.com/sony/sonyflake"
)

func PostMetadata(str string) (uint64, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	newerr := octaviusErrors.New(1, errors.New("demo error"))
	//fmt.Printf("newerr: %v %T\n", newerr, newerr)
	return id, newerr
}
