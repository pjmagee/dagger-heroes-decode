// A dagger module of Heroes Decode by HeroesToolChest that decodes Heroes of the Storm replays.

package main

import (
	"context"
	"dagger/dagger-heroes-decode/internal/dagger"
	"fmt"
)

type DaggerHeroesDecode struct {
}

func (m *DaggerHeroesDecode) Decode(
	ctx context.Context,
	// +optional
	// The replay file to decode
	file *dagger.File,
	// +optional
	// Additional arguments to pass to the decoder
	args []string,

) (*Container, error) {

	repo := dag.Git("https://github.com/HeroesToolChest/HeroesDecode.git")
	dir := repo.Tag("v1.4.0").Tree()

	replayPath := fmt.Sprintf("/app/%s", "replay.StormReplay")

	build := dag.Container().
		From("mcr.microsoft.com/dotnet/sdk:8.0").
		WithWorkdir("/app").
		WithDirectory("/app", dir.Directory("HeroesDecode")).
		WithExec([]string{"dotnet", "publish", "-c", "Release"})

	app := dag.Container().
		From("mcr.microsoft.com/dotnet/runtime:8.0").
		WithWorkdir("/app").
		WithDirectory("/app", build.Directory("/app/bin/Release/net8.0/publish"))

	cmd := []string{"./HeroesDecode"}

	if file != nil {
		replay := []string{"--replay-path", replayPath}
		app.WithFile(replayPath, file)
		cmd = append(cmd, replay...)
	}

	if args != nil {
		cmd = append(cmd, args...)
	}

	return app.
		WithExec(cmd).
		Sync(ctx)
}
