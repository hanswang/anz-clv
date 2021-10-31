package processor

import (
	"fmt"
	"strings"

	"github.com/hanswang/clv/internal/types"
)

func dfsRenderReportBreakdown(report *types.Report) bool {
	breached := false
	if (report.Allocation < report.Usage) {
		breached = true
		fmt.Printf(
			"  Limit breach at %v (limit = %d, direct utilisation = %d, combined utilisation = %d).\n",
			report.Name, report.Allocation, report.DirectUsage, report.Usage,
		)
	}
	if (report.Allocation < report.SubTotalLimit) {
		fmt.Printf(
			"  Warning for limit at %v (limit = %d, combined sub-entity limit = %d).\n",
			report.Name, report.Allocation, report.SubTotalLimit,
		)
	}

	for _, subReport := range report.SubReports {
		subBreach := dfsRenderReportBreakdown(subReport)
		if subBreach {
			breached = true
		}
	}

	return breached
}

func viewGenerate(reports []*types.Report) {
	for i, report := range reports {
		fmt.Printf("Entities: %v:\n", strings.Join(report.Entries, ","))
		breached := dfsRenderReportBreakdown(report)
		if (!breached) {
			fmt.Println("  No limit breaches")
		}
		if i != len(reports) -1 {
			fmt.Println()
		}
	}
}