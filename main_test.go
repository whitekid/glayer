package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	require.NoError(t, playURL(ctx, "https://www.sample-videos.com/audio/mp3/crowd-cheering.mp3"))
}
