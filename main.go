package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/spf13/cobra"
)

func playURL(ctx context.Context, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("load failed with status %d", resp.StatusCode)
	}

	streamer, format, err := mp3.Decode(resp.Body)
	if err != nil {
		return err
	}

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Millisecond*100)); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		cancel()
	})))

	<-ctx.Done()

	return nil
}

func main() {
	root := &cobra.Command{
		Use:          "audio_url",
		SilenceUsage: true,
		Short:        "console audio player",
		Args:         cobra.MinimumNArgs(1),
		RunE:         func(cmd *cobra.Command, args []string) error { return playURL(cmd.Context(), args[0]) },
	}

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
