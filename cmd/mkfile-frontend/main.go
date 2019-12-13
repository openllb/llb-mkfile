package main

import (
	"log"

	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
	mkfile "github.com/openllb/llb-mkfile"
)

func main() {
	if err := grpcclient.RunFromEnvironment(appcontext.Context(), mkfile.Run); err != nil {
		log.Printf("fatal error: %+v", err)
		panic(err)
	}
}
