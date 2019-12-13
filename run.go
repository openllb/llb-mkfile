package mkfile

import (
	"context"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/pkg/errors"
)

var (
	CopyOptions = &llb.CopyInfo{
		FollowSymlinks:      true,
		CopyDirContentsOnly: true,
		AttemptUnpack:       false,
		CreateDestPath:      true,
		AllowWildcard:       true,
		AllowEmptyWildcard:  true,
	}
)

func Run(ctx context.Context, c client.Client) (*client.Result, error) {
	var dt []byte

	content, ok := c.BuildOpts().Opts["content"]
	if ok {
		dt = []byte(content)
	}

	st := llb.Scratch().File(
		llb.Mkfile("/out", 0600, dt),
	)

	def, err := st.Marshal()
	if err != nil {
		return nil, err
	}

	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to solve")
	}

	return res, nil
}
