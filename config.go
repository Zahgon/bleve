package bleve

import (
	"expvar"
	"io"
	"log"
	"time"

	"github.com/blevesearch/bleve/v2/index/scorch"
	"github.com/blevesearch/bleve/v2/index/upsidedown/store/gtreap"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight/highlighter/html"
	index "github.com/blevesearch/bleve_index_api"
)

var bleveExpVar = expvar.NewMap("bleve")

type configuration struct {
	Cache                  *registry.Cache
	DefaultHighlighter     string
	DefaultKVStore         string
	DefaultMemKVStore      string
	DefaultIndexType       string
	SlowSearchLogThreshold time.Duration
	analysisQueue          *index.AnalysisQueue
}

func (c *configuration) SetAnalysisQueueSize(n int) { _ = "STUB: not implemented"; return }

func (c *configuration) Shutdown() { _ = "STUB: not implemented"; return }

func newConfiguration() *configuration { _ = "STUB: not implemented"; return nil }

var Config *configuration

func init() {
	bootStart := time.Now()

	Config = newConfiguration()

	Config.DefaultHighlighter = html.Name

	Config.DefaultKVStore = ""

	Config.DefaultMemKVStore = gtreap.Name

	Config.DefaultIndexType = scorch.Name

	bootDuration := time.Since(bootStart)
	bleveExpVar.Add("bootDuration", int64(bootDuration))
	indexStats = NewIndexStats()
	bleveExpVar.Set("indexes", indexStats)

	initDisk()
}

var logger = log.New(io.Discard, "bleve", log.LstdFlags)

func SetLog(l *log.Logger) { _ = "STUB: not implemented"; return }
