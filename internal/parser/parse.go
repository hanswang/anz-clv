package parser

//go:generate mockgen -destination=../mocks/mock_parser.go -package=mocks github.com/hanswang/clv/internal/parser ParserManager

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hanswang/clv/internal/types"
	log "github.com/sirupsen/logrus"
)

type Parser struct {
}

type ParserManager interface {
	ParseCSV(rows []string) (*map[string]types.Entity, error)
}

const colume_count = 4

func (p *Parser) ParseCSV(rows []string) (*map[string]types.Entity, error) {
	log.Debug("Start input parsing")
	start := time.Now()
	entities := make(map[string]types.Entity, len(rows))

	// assume valid csv passed in - header removed
	for _, r := range rows {
		cols := strings.Split(r, ",")
		if len(cols) != colume_count {
			return nil, fmt.Errorf("found mismatch pattern for csv row: %v", r)
		}

		limit, e := strconv.Atoi(cols[2])
		if e != nil {
			return nil, fmt.Errorf("found non-numeric value in limit col for row %v: %w", r, e)
		}
		utilised, e := strconv.Atoi(cols[3])
		if e != nil {
			return nil, fmt.Errorf("found non-numeric value in utilised col for row %v: %w", r, e)
		}

		entity := types.Entity{
			Name:     cols[0],
			Limit:    limit,
			Utilised: utilised,
		}

		if len(cols[1]) != 0 {
			entity.Parent = &cols[1]
		}
		entities[cols[0]] = entity
	}

	for _, ent := range entities {
		if ent.Parent == nil {
			continue
		}
		p := *ent.Parent
		if _, ok := entities[p]; !ok {
			return nil, fmt.Errorf("found entity pointing to non-existent parent: %v -> %v", ent.Name, p)
		}
	}

	log.Debugf("Done CSV parsing - took %s", time.Since(start))
	return &entities, nil
}
