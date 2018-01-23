package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"hash/fnv"
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

type buffer struct {
	*bytes.Buffer
}

func (b *buffer) Close() error {
	return nil
}

var out files.Writer

var lastSum uint32
var last = make([]byte, 256)

type Diff struct{
	Offset byte
	Old int
	New int
}

func (d *Diff) String() string {
	return fmt.Sprintf("%02X:%d→%d", d.Offset, d.Old, d.New)
}

func PrintSRAM(ctx context.Context, fname string) {
	f, err := files.Open(ctx, fname)
	if err != nil {
		glog.Fatalf("%q: %s", fname, err)
	}
	defer f.Close()

	if _, err := f.Seek(sram.Base, io.SeekStart); err != nil {
		glog.Fatal(err)
	}

	buf := make([]byte, 256)
	if _, err := f.Read(buf); err != nil {
		glog.Fatal(err)
	}

	if buf[0] == 0x60 && buf[1] == 0x60 {
		fmt.Fprintln(out, "state uninitialized")
		return
	}

	h := fnv.New32()
	h.Write(buf)
	sum := h.Sum32()

	if sum == lastSum {
		glog.Infof("no change from hash: %08X", sum)
		return
	}

	r := bytes.NewReader(buf)

	var data sram.ZeldaData

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		glog.Fatal(err)
	}

	b := &buffer{new(bytes.Buffer)}

	json.WriteTo(b, data)
	fmt.Fprintln(out, b)

	var diffs []*Diff
	for i := range buf[:0xD3] {
		if last[i] != buf[i] {
			diffs = append(diffs, &Diff{
				Offset: byte(i),
				Old: int(last[i]),
				New: int(buf[i]),
			})
		}
	}

	fmt.Fprintf(out, "%v\n", diffs)

	copy(last, buf)
	lastSum = sum
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
				//fmt.Fprintf(out, "%#v\n", event)
				if event.Op&fsnotify.Write != 0 {
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
