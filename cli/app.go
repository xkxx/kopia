// Package cli implements command-line commands for the Kopia.
package cli

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/kopia/kopia/internal/kopialogging"
	"github.com/kopia/kopia/internal/serverapi"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/block"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var log = kopialogging.Logger("kopia/cli")

var (
	app = kingpin.New("kopia", "Kopia - Online Backup").Author("http://kopia.github.io/")

	_ = app.Flag("help-full", "Show help for all commands, including hidden").Action(helpFullAction).Bool()

	repositoryCommands = app.Command("repository", "Commands to manipulate repository.").Alias("repo")
	cacheCommands      = app.Command("cache", "Commands to manipulate local cache").Hidden()
	snapshotCommands   = app.Command("snapshot", "Commands to manipulate snapshots.").Alias("snap")
	policyCommands     = app.Command("policy", "Commands to manipulate snapshotting policies.").Alias("policies")
	serverCommands     = app.Command("server", "Commands to control HTTP API server.")
	manifestCommands   = app.Command("manifest", "Low-level commands to manipulate manifest items.").Hidden()
	blockCommands      = app.Command("block", "Commands to manipulate virtual blocks in repository.").Alias("blk").Hidden()
	blobCommands       = app.Command("blob", "Commands to manipulate BLOBs.").Hidden()
	blockIndexCommands = app.Command("blockindex", "Commands to manipulate block index.").Hidden()
	benchmarkCommands  = app.Command("benchmark", "Commands to test performance of algorithms.").Hidden()
)

func helpFullAction(ctx *kingpin.ParseContext) error {
	_ = app.UsageForContextWithTemplate(ctx, 0, kingpin.DefaultUsageTemplate)
	os.Exit(0)
	return nil
}

func noRepositoryAction(act func(ctx context.Context) error) func(ctx *kingpin.ParseContext) error {
	return func(_ *kingpin.ParseContext) error {
		return act(context.Background())
	}
}

func serverAction(act func(ctx context.Context, cli *serverapi.Client) error) func(ctx *kingpin.ParseContext) error {
	return func(_ *kingpin.ParseContext) error {
		return act(context.Background(), serverapi.NewClient(*serverAddress, http.DefaultClient))
	}
}

func repositoryAction(act func(ctx context.Context, rep *repo.Repository) error) func(ctx *kingpin.ParseContext) error {
	return func(kpc *kingpin.ParseContext) error {
		ctx := context.Background()
		ctx = block.UsingBlockCache(ctx, *enableCaching)
		ctx = block.UsingListCache(ctx, *enableListCaching)
		ctx = blob.WithUploadProgressCallback(ctx, func(desc string, progress, total int64) {
			cliProgress.Report("upload '"+desc+"'", progress, total)
		})

		t0 := time.Now()
		rep := mustOpenRepository(ctx, nil)
		repositoryOpenTime := time.Since(t0)

		storageType := rep.Blobs.ConnectionInfo().Type

		reportStartupTime(storageType, rep.Blocks.Format.Version, repositoryOpenTime)

		t1 := time.Now()
		err := act(ctx, rep)
		commandDuration := time.Since(t1)

		reportSubcommandFinished(kpc.SelectedCommand.FullCommand(), err == nil, storageType, rep.Blocks.Format.Version, commandDuration)
		if cerr := rep.Close(ctx); cerr != nil {
			return errors.Wrap(cerr, "unable to close repository")
		}
		return err
	}
}

// App returns an instance of command-line application object.
func App() *kingpin.Application {
	return app
}
