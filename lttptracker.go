package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	//"io/ioutil"

	"github.com/fsnotify/fsnotify"
	"github.com/puellanivis/breton/lib/files"
	"github.com/puellanivis/breton/lib/files/json"
	_ "github.com/puellanivis/breton/lib/files/plugins"
	"github.com/puellanivis/breton/lib/glog"
	flag "github.com/puellanivis/breton/lib/gnuflag"
	"github.com/puellanivis/breton/lib/util"

	"./sram"
)

type buffer struct{
	*bytes.Buffer
}

func (b *buffer) Close() error {
	return nil
}

var out files.Writer

func PrintSRAM(ctx context.Context, fname string) {
	f, err := files.Open(ctx, fname)
	if err != nil {
		glog.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Seek(sram.Base, io.SeekStart); err != nil {
		glog.Fatal(err)
	}

	var data sram.ZeldaData

	if err := binary.Read(f, binary.LittleEndian, &data); err != nil {
		glog.Fatal(err)
	}

	b := &buffer{new(bytes.Buffer)}

	json.WriteTo(b, data)
	fmt.Fprintln(out, b)
}

func main() {
	finish, ctx := util.Init("lttp-tracking", 0, 1)
	defer finish()

	ctx, cancel := context.WithCancel(ctx)

	w, err := files.Create(ctx, "-")
	if err != nil {
		glog.Fatal(err)
	}
	defer w.Close()
	out = w

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
	    glog.Fatal(err)
	}
	defer watcher.Close()

	fname := flag.Arg(0)

	PrintSRAM(ctx, fname)

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return

			case event := <-watcher.Events:
				fmt.Fprintf(out, "%#v\n", event)
				if event.Op & fsnotify.Write != 0 {
					PrintSRAM(ctx, fname)
				}

			case err := <-watcher.Errors:
				glog.Error(err)
			}
		}
	}()
	
	if err := watcher.Add(fname); err != nil {
		glog.Fatal(err)
	}

	<-ctx.Done()
}
