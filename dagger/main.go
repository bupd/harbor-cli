package main

import (
	"context"
	"dagger/harbor-cli/internal/dagger"
	"fmt"
	"log"
	"strings"
)

const (
	GOLANGCILINT_VERSION = "v1.61.0"
	GO_VERSION           = "1.22.5"
	SYFT_VERSION         = "v1.9.0"
	GORELEASER_VERSION   = "v2.3.2"
)

func New(
	// Local or remote directory with source code, defaults to "./"
	// +optional
	// +defaultPath="./"
	source *dagger.Directory,
) *HarborCli {
	return &HarborCli{Source: source}
}

type HarborCli struct {
	// Local or remote directory with source code, defaults to "./"
	Source *dagger.Directory
}

// Create build of Harbor CLI for local testing and development
func (m *HarborCli) BuildDev(
	ctx context.Context,
	platform string,
) *dagger.File {
	fmt.Println("🛠️  Building Harbor-Cli with Dagger...")
	// Define the path for the binary output
	os, arch, err := parsePlatform(platform)
	if err != nil {
		log.Fatalf("Error parsing platform: %v", err)
	}
	builder := dag.Container().
		From("golang:"+GO_VERSION+"-alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-"+GO_VERSION)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-"+GO_VERSION)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithMountedDirectory("/src", m.Source). // Ensure the source directory with go.mod is mounted
		WithWorkdir("/src").
		WithEnvVariable("GOOS", os).
		WithEnvVariable("GOARCH", arch).
		WithExec([]string{"go", "build", "-o", "bin/harbor-cli", "cmd/harbor/main.go"})
	return builder.File("bin/harbor-cli")
}

// Return list of containers for list of oses and arches
//
// FIXME: there is a bug where you cannot return a list of containers right now
// this function works as expected because it is only called by other functions but
// calling it via the CLI results in an error. That is why this into a private function for
// now so that no one calls this https://github.com/dagger/dagger/issues/8202#issuecomment-2317291483
func (m *HarborCli) build(
	ctx context.Context,
) []*dagger.Container {
	platformVariants := make([]*dagger.Container, 0, 6)
	fmt.Println("🛠️  Building with Dagger...")
	oses := []string{"linux", "darwin", "windows"}
	arches := []string{"amd64", "arm64"}
	for _, goos := range oses {
		for _, goarch := range arches {
			bin_path := fmt.Sprintf("build/%s/%s/", goos, goarch)
			builder := dag.Container().
				From("golang:"+GO_VERSION+"-alpine").
				WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-"+GO_VERSION)).
				WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
				WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-"+GO_VERSION)).
				WithEnvVariable("GOCACHE", "/go/build-cache").
				WithMountedDirectory("/src", m.Source).
				WithWorkdir("/src").
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch).
				WithExec([]string{"go", "build", "-o", bin_path + "harbor", "/src/cmd/harbor/main.go"}).
				WithWorkdir(bin_path).
				WithExec([]string{"ls"}).
				WithEntrypoint([]string{"./harbor"})

			outputDir := builder.Directory(".")

			// wrap the output directory in a new empty container marked
			// with the platform
			platform := fmt.Sprintf("%s/%s", goos, goarch)
			cont := dag.
				Container(dagger.ContainerOpts{Platform: dagger.Platform(platform)}).
				WithRootfs(outputDir).WithEntrypoint([]string{"./harbor"})

			platformVariants = append(platformVariants, cont)
		}
	}
	return platformVariants
}

// Run linter golangci-lint and write the linting results to a file golangci-lint-report.txt
func (m *HarborCli) LintReport(ctx context.Context) *dagger.File {
	report := "golangci-lint-report.sarif"
	return m.lint(ctx).WithExec([]string{
		"golangci-lint", "run",
		"--out-format", "sarif:" + report,
		"--issues-exit-code", "0",
	}).File(report)
}

// Run linter golangci-lint
func (m *HarborCli) Lint(ctx context.Context) (string, error) {
	return m.lint(ctx).WithExec([]string{"golangci-lint", "run"}).Stderr(ctx)
}

func (m *HarborCli) lint(ctx context.Context) *dagger.Container {
	fmt.Println("👀 Running linter and printing results to file golangci-lint.txt.")
	linter := dag.Container().
		From("golangci/golangci-lint:"+GOLANGCILINT_VERSION+"-alpine").
		WithMountedCache("/lint-cache", dag.CacheVolume("/lint-cache")).
		WithEnvVariable("GOLANGCI_LINT_CACHE", "/lint-cache").
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src")
	return linter
}

// Create snapshot release with goreleaser
func (m *HarborCli) SnapshotRelease(
	ctx context.Context,
	githubToken *dagger.Secret,
) {
	_, err := m.
		goreleaserContainer(githubToken).
		WithExec([]string{"goreleaser", "release", "--snapshot", "--clean"}).
		Stderr(ctx)
	if err != nil {
		log.Printf("❌ Error occured during snapshot release for the recently merged pull-request: %s", err)
		return
	}
	log.Println("Pull-Request tasks completed successfully 🎉")
}

// Create release with goreleaser
func (m *HarborCli) Release(
	ctx context.Context,
	// Github API token
	githubToken *dagger.Secret,
) {
	goreleaser := m.goreleaserContainer(githubToken).
		WithExec([]string{"ls", "-la"}).
		WithExec([]string{"goreleaser", "release", "--clean"})

	_, err := goreleaser.Stderr(ctx)
	if err != nil {
		log.Printf("Error occured during release: %s", err)
		return
	}
	log.Println("Release tasks completed successfully 🎉")
}

// PublishImage publishes a container image to a registry with a specific tag and signs it using Cosign.
func (m *HarborCli) PublishImage(
	ctx context.Context,
	cosignKey *dagger.Secret,
	cosignPassword *dagger.Secret,
	regUsername string,
	regPassword *dagger.Secret,
	regAddress string,
	publishAddress string,
	tag string,
) string {
	builds := m.build(ctx)

	publisher := dag.Container().WithRegistryAuth(regAddress, regUsername, regPassword)
	// Push the versioned tag
	versionedAddress := fmt.Sprintf("%s:%s", publishAddress, tag)
	addr, err := publisher.Publish(ctx, versionedAddress, dagger.ContainerPublishOpts{PlatformVariants: builds})
	if err != nil {
		panic(err)
	}
	_, err = dag.Cosign().Sign(ctx, cosignKey, cosignPassword, []string{addr}, dagger.CosignSignOpts{RegistryUsername: regUsername, RegistryPassword: regPassword})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully published image to %s 🎉\n", addr)

	return addr
}

// Return the platform of the container
func buildPlatform(ctx context.Context, container *dagger.Container) string {
	platform, err := container.Platform(ctx)
	if err != nil {
		log.Fatalf("error getting platform: %v", err)
	}
	return string(platform)
}

// Return a container with the goreleaser binary mounted and the source directory mounted.
func (m *HarborCli) goreleaserContainer(
	// Github API token
	githubToken *dagger.Secret,
) *dagger.Container {
	// Export the syft binary from the syft container as a file to generate SBOM
	syft := dag.Container().
		From(fmt.Sprintf("anchore/syft:%s", SYFT_VERSION)).
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("syft-gomod")).
		File("/syft")

	return dag.Container().
		From(fmt.Sprintf("goreleaser/goreleaser:%s", GORELEASER_VERSION)).
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-"+GO_VERSION)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-"+GO_VERSION)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithFile("/bin/syft", syft).
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src").
		WithEnvVariable("TINI_SUBREAPER", "true").
		WithSecretVariable("GITHUB_TOKEN", githubToken)
}

// Generate CLI Documentation with doc.go and return the directory containing the generated files
func (m *HarborCli) RunDoc(ctx context.Context) *dagger.Directory {
	return dag.Container().
		From("golang:"+GO_VERSION+"-alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-"+GO_VERSION)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-"+GO_VERSION)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src/doc").
		WithExec([]string{"go", "run", "doc.go"}).
		WithWorkdir("/src").Directory("/src/doc")
}

// Executes Go tests and returns the directory containing the test results
func (m *HarborCli) Test(ctx context.Context) *dagger.Directory {
	return dag.Container().
		From("golang:"+GO_VERSION+"-alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-"+GO_VERSION)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-"+GO_VERSION)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src").
		WithExec([]string{"go", "test", "-v", "./..."}).
		Directory("/src")
}

// Parse the platform string into os and arch
func parsePlatform(platform string) (string, string, error) {
	parts := strings.Split(platform, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid platform format: %s. Should be os/arch. E.g. darwin/amd64", platform)
	}
	return parts[0], parts[1], nil
}
