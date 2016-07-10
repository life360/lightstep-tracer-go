package lightstep

import (
	"testing"

	"github.com/life360/basictracer-go"
)

func TestMaxBufferSize(t *testing.T) {
	recorder := NewRecorder(Options{
		AccessToken: "0987654321",
	}).(*Recorder)

	checkCapSize := func(spanLen, spanCap int) {
		recorder.lock.Lock()
		defer recorder.lock.Unlock()

		if recorder.buffer.cap() != spanCap {
			t.Error("Unexpected buffer size")
		}
		if recorder.buffer.len() != spanLen {
			t.Error("Unexpected buffer size")
		}
	}

	checkCapSize(0, defaultMaxSpans)

	spans := make([]basictracer.RawSpan, defaultMaxSpans)
	for _, span := range spans {
		recorder.RecordSpan(span)
	}

	checkCapSize(defaultMaxSpans, defaultMaxSpans)

	spans = append(spans, make([]basictracer.RawSpan, defaultMaxSpans)...)
	for _, span := range spans {
		recorder.RecordSpan(span)
	}

	checkCapSize(defaultMaxSpans, defaultMaxSpans)

	maxBuffer := 10
	recorder = NewRecorder(Options{
		AccessToken:      "0987654321",
		MaxBufferedSpans: maxBuffer,
	}).(*Recorder)

	checkCapSize(0, maxBuffer)

	spans = append(spans, make([]basictracer.RawSpan, 100*defaultMaxSpans)...)
	for _, span := range spans {
		recorder.RecordSpan(span)
	}

	checkCapSize(maxBuffer, maxBuffer)
}
