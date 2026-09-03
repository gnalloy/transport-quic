package application

import (
	"encoding/binary"
	"fmt"
	"io"
)

const DefaultMaxFrameSize = 1<<16 - 1

// LengthPrefixedCodec 使用两字节网络序长度前缀封装应用消息。
//
// DNS-over-QUIC、DNS-over-TCP 这类短消息 request-response 协议可直接复用。
type LengthPrefixedCodec struct {
	// MaxFrameSize 限制 payload 最大字节数，0 表示 65536 字节。
	MaxFrameSize int
}

// WriteFrame 写出一帧长度前缀消息。
func (c LengthPrefixedCodec) WriteFrame(w io.Writer, payload []byte) error {
	if w == nil {
		return ErrInvalidConfig
	}
	maxFrameSize := c.maxFrameSize()
	if len(payload) > maxFrameSize || len(payload) > 0xffff {
		return fmt.Errorf("%w: %d > %d", ErrFrameTooLarge, len(payload), maxFrameSize)
	}
	var header [2]byte
	binary.BigEndian.PutUint16(header[:], uint16(len(payload)))
	if err := writeFull(w, header[:]); err != nil {
		return err
	}
	return writeFull(w, payload)
}

// ReadFrame 读取一帧长度前缀消息。
func (c LengthPrefixedCodec) ReadFrame(r io.Reader) ([]byte, error) {
	size, err := c.readFrameSize(r)
	if err != nil {
		return nil, err
	}
	payload := make([]byte, size)
	if _, err = io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// ReadFrameInto 将一帧读取到调用方缓冲区，避免稳定负载下重复分配。
func (c LengthPrefixedCodec) ReadFrameInto(r io.Reader, dst []byte) ([]byte, error) {
	size, err := c.readFrameSize(r)
	if err != nil {
		return nil, err
	}
	if size > len(dst) {
		return nil, fmt.Errorf("%w: need %d bytes, have %d", io.ErrShortBuffer, size, len(dst))
	}
	payload := dst[:size]
	if _, err = io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// WriteRequest 实现 StreamCodec。
func (c LengthPrefixedCodec) WriteRequest(stream Stream, payload []byte) error {
	return c.WriteFrame(stream, payload)
}

// ReadResponse 实现 StreamCodec。
func (c LengthPrefixedCodec) ReadResponse(stream Stream) ([]byte, error) {
	return c.ReadFrame(stream)
}

func (c LengthPrefixedCodec) maxFrameSize() int {
	if c.MaxFrameSize <= 0 {
		return DefaultMaxFrameSize
	}
	return c.MaxFrameSize
}

func (c LengthPrefixedCodec) readFrameSize(r io.Reader) (int, error) {
	if r == nil {
		return 0, ErrInvalidConfig
	}
	var header [2]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return 0, err
	}
	size := int(binary.BigEndian.Uint16(header[:]))
	maxFrameSize := c.maxFrameSize()
	if size > maxFrameSize {
		return 0, fmt.Errorf("%w: %d > %d", ErrFrameTooLarge, size, maxFrameSize)
	}
	return size, nil
}

func writeFull(w io.Writer, payload []byte) error {
	for len(payload) > 0 {
		written, err := w.Write(payload)
		if written > 0 {
			payload = payload[written:]
		}
		if err != nil {
			return err
		}
		if written == 0 {
			return io.ErrShortWrite
		}
	}
	return nil
}
