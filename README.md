# Coverity API Client for Go

## About

Coverity is a brand of software development products from Synopsys, consisting primarily of static code analysis tools and dynamic code analysis services. The tools enable engineers and security teams to find defects and security vulnerabilities in custom source code written in C, C++, Java, C#, JavaScript and more.

## Installation

    go get github.com/denglitong/gocoverity

## Usage

### configuration service

```go

import "github.com/denglitong/gocoverity"

const (
    ConfigurationServiceURL = "http://my.cov_connect.domain:8080/ws/v9/configurationservice"
    DefectServiceURL        = "http://my.cov_connect.domain:8080/ws/v9/defectservice"
    CovUsername             = "denglitong"
    CovPassword             = "123456"
    WSUId                   = "XWSSGID-1349973313023-787497544" // explore request xml from soapui you can found wsu:Id="XWSSGID-1349973313023-787497544"
    SoapEnvMustUnderstand   = "1" // explore request xml from soapui you can found soapenv:mustUnderstand="1"
)

configClient := gocoverity.NewClient(ConfigurationServiceURL, CovUsername, CovPassword, 
    WSUId, SoapEnvMustUnderstand, configuration.Namespace)
configurationService := configuration.NewConfigurationService(configClient)

pageSize, sortAscending, startIndex := 10, true, 0
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

```

### defect service

```go

defectClient := gocoverity.NewClient(DefectServiceURL, CovUsername, CovPassword, 
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

```

## Testing

    go test

## Contribute

All Contributions are welcome. The todo list is on the bottom of this README. Feel free to send a pull request.

## TODO

* Add more testcase