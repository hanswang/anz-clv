package processor

import (
	"fmt"
	"strings"

	"github.com/hanswang/clv/internal/aggregator"
	"github.com/hanswang/clv/internal/parser"
)

func Run(body []byte) error {
	var p parser.ParserManager = parser.New()

	rows := strings.Split(string(body), "\n")
	if len(rows) < 2 {
		return fmt.Errorf("CSV file has no content excluding header with row count: %d", len(rows))
	}
	_, rows = rows[0], rows[1:]

	entities, err := p.ParseCSV(rows)
	if err != nil {
		return fmt.Errorf("CSV parse error: %w", err)
	}

	var a aggregator.AggregatorManager = aggregator.New()

	analyses := a.GenerateReport(entities)

	viewGenerate(analyses)

	return nil
}
