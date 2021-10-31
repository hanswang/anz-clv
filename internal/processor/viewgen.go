package processor

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hanswang/clv/internal/types"
)

var Out io.Writer = os.Stdout

func dfsRenderReportBreakdown(report *types.Report) bool {
	breached := false
	if (report.Allocation < report.Usage) {
		breached = true
		fmt.Fprintf(
			Out,
			"  Limit breach at %v (limit = %d, direct utilisation = %d, combined utilisation = %d).\n",
			report.Name, report.Allocation, report.DirectUsage, report.Usage,
		)
	}
	if (report.Allocation < report.SubTotalLimit) {
		fmt.Fprintf(
			Out,
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

func ViewGenerate(reports []*types.Report) {
	for i, report := range reports {
		fmt.Fprintf(Out, "Entities: %v:\n", strings.Join(report.Entries, "/"))
		breached := dfsRenderReportBreakdown(report)
		if (!breached) {
			fmt.Fprintf(Out, "  No limit breaches\n")
		}
		if i != len(reports) -1 {
			fmt.Fprintln(Out)
		}
	}
}