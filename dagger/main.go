// A generated module for DaggerHeroesDecode functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-heroes-decode/internal/dagger"
	"fmt"
)

type DaggerHeroesDecode struct {
	StormReplay *dagger.File
	Args        []string
}

func (m *DaggerHeroesDecode) WithStormReplay(ctx context.Context, path *dagger.File) (*DaggerHeroesDecode, error) {
	m.StormReplay = path
	return m, nil
}

func (m *DaggerHeroesDecode) WithArgs(ctx context.Context, args []string) (*DaggerHeroesDecode, error) {
	m.Args = args
	return m, nil
}

func (m *DaggerHeroesDecode) Decode(ctx context.Context) (*Container, error) {

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

	cmd := []string{"./HeroesDecode", "--replay-path", replayPath}
	cmd = append(cmd, m.Args...)

	return app.
		WithFile(replayPath, m.StormReplay).
		WithExec(cmd).
		Sync(ctx)
}
