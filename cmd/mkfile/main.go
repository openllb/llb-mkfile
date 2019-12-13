package main

import (
	"os"

	"github.com/moby/buildkit/client/llb"
)

func main() {
	st := llb.Scratch().File(
		llb.Mkfile("/out", 0600, []byte("out")),
	)

	def, err := st.Marshal()
	if err != nil {
		panic(err)
	}

	llb.WriteTo(def, os.Stdout)
}
