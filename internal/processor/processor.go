package processor

import (
	"fmt"
	"strings"

	"github.com/hanswang/clv/internal/aggregator"
	"github.com/hanswang/clv/internal/parser"
)

type Processor struct {
	Parser     parser.ParserManager
	Aggregator aggregator.AggregatorManager
}

func NewProcessor(parser parser.ParserManager, aggregator aggregator.AggregatorManager) *Processor {
	return &Processor{
		Parser: parser,
		Aggregator: aggregator,
	}
}

func (p *Processor) Run(body []byte) error {
	rows := strings.Split(string(body), "\n")
	if len(rows) < 2 {
		return fmt.Errorf("CSV file has no content excluding header with row count: %d", len(rows))
	}
	_, rows = rows[0], rows[1:]

	entities, err := p.Parser.ParseCSV(rows)
	if err != nil {
		return fmt.Errorf("CSV parse error: %w", err)
	}

	analyses := p.Aggregator.GenerateReport(entities)

	ViewGenerate(analyses)

	return nil
}
