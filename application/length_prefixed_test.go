package application

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestLengthPrefixedCodecRoundTrip(t *testing.T) {
	var wire bytes.Buffer
	codec := LengthPrefixedCodec{MaxFrameSize: 32}
	if err := codec.WriteFrame(&wire, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	if got := wire.Bytes(); !bytes.Equal(got, []byte{0, 5, 'h', 'e', 'l', 'l', 'o'}) {
		t.Fatalf("wire=%v", got)
	}
	payload, err := codec.ReadFrame(&wire)
	if err != nil {
		t.Fatal(err)
	}
	if string(payload) != "hello" {
		t.Fatalf("payload=%q", payload)
	}
}

func TestLengthPrefixedCodecRejectsOversizedFrame(t *testing.T) {
	codec := LengthPrefixedCodec{MaxFrameSize: 4}
	if err := codec.WriteFrame(&bytes.Buffer{}, []byte("hello")); !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("err=%v, want %v", err, ErrFrameTooLarge)
	}
	_, err := codec.ReadFrame(bytes.NewReader([]byte{0, 5, 'h', 'e', 'l', 'l', 'o'}))
	if !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("err=%v, want %v", err, ErrFrameTooLarge)
	}
}

func TestLengthPrefixedCodecDefaultLimitMatchesWireFormat(t *testing.T) {
	codec := LengthPrefixedCodec{}
	if err := codec.WriteFrame(&bytes.Buffer{}, make([]byte, DefaultMaxFrameSize)); err != nil {
		t.Fatalf("maximum frame: %v", err)
	}
	if err := codec.WriteFrame(&bytes.Buffer{}, make([]byte, DefaultMaxFrameSize+1)); !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("err=%v, want %v", err, ErrFrameTooLarge)
	}
}

func TestLengthPrefixedCodecCompletesShortWrites(t *testing.T) {
	writer := &shortWriter{limit: 1}
	codec := LengthPrefixedCodec{MaxFrameSize: 8}
	if err := codec.WriteFrame(writer, []byte("ping")); err != nil {
		t.Fatal(err)
	}
	want := []byte{0, 4, 'p', 'i', 'n', 'g'}
	if !bytes.Equal(writer.Bytes(), want) {
		t.Fatalf("wire=%v, want=%v", writer.Bytes(), want)
	}
}

func TestLengthPrefixedCodecWritesFrameIntoReusableBuffer(t *testing.T) {
	codec := LengthPrefixedCodec{MaxFrameSize: 8}
	dst := make([]byte, 6)
	payload := dst[2:]
	copy(payload, "ping")
	writer := &countingWriter{limit: len(dst)}
	if err := codec.WriteFrameInto(writer, payload, dst); err != nil {
		t.Fatal(err)
	}
	want := []byte{0, 4, 'p', 'i', 'n', 'g'}
	if !bytes.Equal(writer.Bytes(), want) {
		t.Fatalf("wire=%v, want=%v", writer.Bytes(), want)
	}
	if writer.writes != 1 {
		t.Fatalf("writes=%d, want 1", writer.writes)
	}

	writer = &countingWriter{limit: 2}
	if err := codec.WriteFrameInto(writer, payload, dst); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(writer.Bytes(), want) || writer.writes != 3 {
		t.Fatalf("short writes=%d wire=%v, want 3 writes and %v", writer.writes, writer.Bytes(), want)
	}
}

func TestLengthPrefixedCodecWriteFrameIntoRejectsShortBuffer(t *testing.T) {
	codec := LengthPrefixedCodec{MaxFrameSize: 8}
	err := codec.WriteFrameInto(&bytes.Buffer{}, []byte("ping"), make([]byte, 5))
	if !errors.Is(err, io.ErrShortBuffer) {
		t.Fatalf("err=%v, want %v", err, io.ErrShortBuffer)
	}
}

func TestLengthPrefixedCodecReadsIntoReusableBuffer(t *testing.T) {
	codec := LengthPrefixedCodec{MaxFrameSize: 8}
	dst := make([]byte, 8)
	payload, err := codec.ReadFrameInto(bytes.NewReader([]byte{0, 4, 'p', 'o', 'n', 'g'}), dst)
	if err != nil {
		t.Fatal(err)
	}
	if string(payload) != "pong" || &payload[0] != &dst[0] {
		t.Fatalf("payload=%q does not reuse destination", payload)
	}

	_, err = codec.ReadFrameInto(bytes.NewReader([]byte{0, 5, 'h', 'e', 'l', 'l', 'o'}), dst[:4])
	if !errors.Is(err, io.ErrShortBuffer) {
		t.Fatalf("err=%v, want %v", err, io.ErrShortBuffer)
	}
}

type shortWriter struct {
	bytes.Buffer
	limit int
}

func (w *shortWriter) Write(payload []byte) (int, error) {
	if len(payload) > w.limit {
		payload = payload[:w.limit]
	}
	return w.Buffer.Write(payload)
}

type countingWriter struct {
	bytes.Buffer
	limit  int
	writes int
}

func (w *countingWriter) Write(payload []byte) (int, error) {
	w.writes++
	if len(payload) > w.limit {
		payload = payload[:w.limit]
	}
	return w.Buffer.Write(payload)
}
