package aggregator

//go:generate mockgen -destination=../mocks/mock_aggregator.go -package=mocks github.com/hanswang/clv/internal/aggregator AggregatorManager

import (
	"time"

	"github.com/hanswang/clv/internal/types"
	log "github.com/sirupsen/logrus"
)

type Aggregator struct {
}

type AggregatorManager interface {
	GenerateReport(entities *map[string]types.Entity) []*types.Report
}

func dfsAggregateChildren(root *types.Report, entities *map[string][]types.Entity) {
	children, ok := (*entities)[root.Name]
	if !ok {
		return
	}
	for _, entity := range children {
		report := &types.Report{
			Name: entity.Name,
			Entries: []string{entity.Name},
			Allocation: entity.Limit,
			DirectUsage: entity.Utilised,
			Usage: entity.Utilised,
		}
		dfsAggregateChildren(report, entities)

		root.Entries = append(root.Entries, report.Entries...)
		root.Usage += report.Usage
		root.SubTotalLimit += report.SubTotalLimit + report.Allocation
		root.SubReports = append(root.SubReports, report)
	}

}

func (a *Aggregator) GenerateReport(entities *map[string]types.Entity) []*types.Report {
	log.Debug("Start aggregating")
	start := time.Now()
	reps := []*types.Report{}
	chdEntities := map[string][]types.Entity{}
	
	// find lead node & reverse child -> parent relation entity to parent -> children relation entity
	for name, entity := range *entities {
		if entity.Parent == nil {
			r := &types.Report{
				Name: entity.Name,
				Entries: []string{entity.Name},
				Allocation: entity.Limit,
				DirectUsage: entity.Utilised,
				Usage: entity.Utilised,
			}
			reps = append(reps, r)
			delete(*entities, name)
		} else {
			chdEntities[*entity.Parent] = append(chdEntities[*entity.Parent], entity)
			delete(*entities, name)
		}
	}

	for _, lead := range reps {
		dfsAggregateChildren(lead, &chdEntities)
	}

	log.Debugf("Done aggregating - took %s", time.Since(start))
	return reps
}