package main

import (
	"fmt"
	"github.com/denglitong/gocoverity/client"
	"github.com/denglitong/gocoverity/configuration"
	"github.com/denglitong/gocoverity/defect"
)

const (
	ConfigurationServiceURL = "http://my.cov_connect.domain:8080/ws/v9/configurationservice" // configurationservice wsdl url
	DefectServiceURL        = "http://my.cov_connect.domain:8080/ws/v9/defectservice"        // defectservice wsdl url
	CovUsername             = "denglitong"                                                   // coverity server account username
	CovPassword             = "123456"                                                       // coverity server account password
	WSUId                   = "XWSSGID-1349973313023-787497544"                              // explore request xml from soapui you can found wsu:Id="XWSSGID-1349973313023-787497544"
	SoapEnvMustUnderstand   = "1"                                                            // explore request xml from soapui you can found soapenv:mustUnderstand="1"
)

func main() {
	configClient := client.NewClient(ConfigurationServiceURL, CovUsername, CovPassword, WSUId, SoapEnvMustUnderstand, configuration.Namespace)
	configurationService := configuration.NewConfigurationService(configClient)

	pageSize, sortAscending, startIndex := 100, true, 0
	response, err := configurationService.GetUsers(&configuration.GetUsers{
		PageSpec: &configuration.PageSpecDataObj{
			PageSize:      &pageSize,
			SortAscending: &sortAscending,
			StartIndex:    &startIndex,
		},
	})

	if err != nil {
		panic(fmt.Sprintf("GetProjects error: %v", err))
	}

	fmt.Println("users len", len(response.Return.Users))
	for _, user := range response.Return.Users {
		fmt.Printf("%v", *user.Username)
		if user.Email != nil {
			fmt.Printf(" %v", *user.Email)
		}
		fmt.Println()
	}

	fmt.Println()

	defectClient := client.NewClient(DefectServiceURL, CovUsername, CovPassword,
		WSUId, SoapEnvMustUnderstand, configuration.Namespace)
	defectService := defect.NewDefectService(defectClient)

	projectName := "tunas-java-demo-dlt_test-v9.git.n.xiaomi.com"

	metricsProjectResp, err := defectService.GetComponentMetricsForProject(&defect.GetComponentMetricsForProject{
		ProjectId: &defect.ProjectIdDataObj{
			Name: &projectName,
		},
	})

	if err != nil {
		panic(fmt.Sprintf("GetProjects error: %v", err))
	}

	for _, componentMetric := range metricsProjectResp.Return {
		fmt.Println("codeLineCount:", *componentMetric.CodeLineCount)
		fmt.Println("blankLineCount:", *componentMetric.BlankLineCount)
		fmt.Println("commentLineCount:", *componentMetric.CommentLineCount)
		fmt.Println("dismissedCount:", *componentMetric.DismissedCount)
		fmt.Println("fixedCount:", *componentMetric.FixedCount)
		fmt.Println("metricsDate:", *componentMetric.MetricsDate)
		fmt.Println("newCount:", *componentMetric.NewCount)
		fmt.Println("outstandingCount:", *componentMetric.OutstandingCount)
		fmt.Println("totalCount:", *componentMetric.TotalCount)
		fmt.Println("triagedCount:", *componentMetric.TriagedCount)
	}
}
