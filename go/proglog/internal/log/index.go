package log

import (
	"github.com/tysonmote/gommap"
	"io"
	"os"
)

var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	entWidth        = offWidth + posWidth
)

type index struct {
	file *os.File
	size uint64
	mmap gommap.MMap
}

func newIndex(f *os.File, c Config) (*index, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	if err := os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}
	mmap, err := gommap.Map(f.Fd(), gommap.PROT_READ|gommap.PROT_WRITE, gommap.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	return &index{
		file: f,
		size: uint64(fi.Size()),
		mmap: mmap,
	}, nil
}

func (i *index) Close() error {
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	if err := i.file.Sync(); err != nil {
		return err
	}
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

func (i *index) Read(in int64) (uint32, uint64, error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}

	var out uint32
	if in == -1 {
		out = uint32((i.size / entWidth) - 1)
	} else {
		out = uint32(in)
	}

	pos := uint64(out) * entWidth
	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}

	out = enc.Uint32(i.mmap[pos : pos+offWidth])
	pos = enc.Uint64(i.mmap[pos+offWidth : pos+entWidth])

	return out, pos, nil
}

func (i *index) Write(off uint32, pos uint64) error {
	if uint64(len(i.mmap)) < i.size+entWidth {
		return io.EOF
	}
	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += entWidth
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
