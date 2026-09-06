package metrics

import (
	"io"

	"github.com/blevesearch/go-metrics"
)

var timerPercentiles = []float64{0.5, 0.75, 0.95, 0.99, 0.999}

func TimerMap(timer metrics.Timer) map[string]interface{} { _ = "STUB: not implemented"; return nil }

func isNanOrInf(v float64) bool { _ = "STUB: not implemented"; return false }

func WriteTimerJSON(w io.Writer, timer metrics.Timer) { _ = "STUB: not implemented"; return }

func WriteTimerCSVHeader(w io.Writer, prefix string) { _ = "STUB: not implemented"; return }

func WriteTimerCSV(w io.Writer, timer metrics.Timer) { _ = "STUB: not implemented"; return }
