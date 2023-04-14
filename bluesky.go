package main

import (
	"context"
	"fmt"
	"time"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/gen2brain/beeep"
)

func BskyFeedPost(xrpcc *xrpc.Client, text string) (string, error) {
	err := beeep.Notify("まぜそば大陸", fmt.Sprintf("BskyFeedPost: %s", text), "")
	if err != nil {
		return "", err
	}

	// return "<MOCK URI>", nil // TEMP

	resp, err := comatproto.RepoCreateRecord(context.TODO(), xrpcc, &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       xrpcc.Auth.Did,
		Record: &lexutil.LexiconTypeDecoder{&appbsky.FeedPost{
			Text:      text,
			CreatedAt: time.Now().Format("2006-01-02T15:04:05.000Z"),
		}},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create post: %w", err)
	}
	return resp.Uri, nil
}
