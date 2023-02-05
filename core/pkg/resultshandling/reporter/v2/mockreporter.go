package reporter

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/kubescape/kubescape/v2/core/cautils"
)

type ReportMock struct {
	query   string
	message string
}

func NewReportMock(query, message string) *ReportMock {
	return &ReportMock{
		query:   query,
		message: message,
	}
}
func (reportMock *ReportMock) Submit(_ context.Context, opaSessionObj *cautils.OPASessionObj) error {
	return nil
}

func (reportMock *ReportMock) SetCustomerGUID(customerGUID string) {
}

func (reportMock *ReportMock) SetClusterName(clusterName string) {
}

func (reportMock *ReportMock) GetURL() string {
	u, err := url.Parse(reportMock.query)
	if err != nil || u.String() == "" {
		return ""
	}

	q := u.Query()
	q.Add("utm_source", "GitHub")
	q.Add("utm_medium", "CLI")
	q.Add("utm_campaign", "Submit")

	u.RawQuery = q.Encode()

	return u.String()
}

func (reportMock *ReportMock) DisplayReportURL() {

	sep := "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	message := sep + "\n"
	message += "Scan results have not been submitted: " + reportMock.message + "\n"
	if link := reportMock.GetURL(); link != "" {
		message += "For more details: " + link + "\n"
	}
	message += sep + "\n"
	cautils.InfoTextDisplay(os.Stderr, fmt.Sprintf("\n%s\n", message))
}
