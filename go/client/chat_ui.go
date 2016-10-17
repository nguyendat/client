// Copyright 2016 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/chat1"
)

type ChatUI struct {
	libkb.Contextified
	terminal libkb.TerminalUI
	noOutput bool
}

func (c *ChatUI) ChatAttachmentUploadStart(context.Context, int) error {
	if c.noOutput {
		return nil
	}
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment upload "+ColorString("green", "starting")+"\n")
	return nil
}

func (c *ChatUI) ChatAttachmentUploadProgress(ctx context.Context, arg chat1.ChatAttachmentUploadProgressArg) error {
	if c.noOutput {
		return nil
	}
	percent := (100 * arg.BytesComplete) / arg.BytesTotal
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment upload progress %d%% (%d of %d bytes uploaded)\n", percent, arg.BytesComplete, arg.BytesTotal)
	return nil
}

func (c *ChatUI) ChatAttachmentUploadDone(context.Context, int) error {
	if c.noOutput {
		return nil
	}
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment upload "+ColorString("magenta", "finished")+"\n")
	return nil
}

func (c *ChatUI) ChatAttachmentPreviewUploadStart(context.Context, int) error {
	if c.noOutput {
		return nil
	}
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment preview upload "+ColorString("green", "starting")+"\n")
	return nil
}

func (c *ChatUI) ChatAttachmentPreviewUploadDone(context.Context, int) error {
	if c.noOutput {
		return nil
	}
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment preview upload "+ColorString("magenta", "finished")+"\n")
	return nil
}

func (c *ChatUI) ChatAttachmentDownloadStart(context.Context, int) error {
	if c.noOutput {
		return nil
	}
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment download "+ColorString("green", "starting")+"\n")
	return nil
}

func (c *ChatUI) ChatAttachmentDownloadProgress(context.Context, chat1.ChatAttachmentDownloadProgressArg) error {
	if c.noOutput {
		return nil
	}
	return nil
}

func (c *ChatUI) ChatAttachmentDownloadDone(context.Context, int) error {
	if c.noOutput {
		return nil
	}
	w := c.terminal.ErrorWriter()
	fmt.Fprintf(w, "Attachment download "+ColorString("magenta", "finished")+"\n")
	return nil
}
