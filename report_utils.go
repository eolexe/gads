package gads

import "encoding/xml"

const (
	DATE_RANGE_TODAY               = "TODAY"
	DATE_RANGE_YESTERDAY           = "YESTERDAY"
	DATE_RANGE_LAST_7_DAYS         = "LAST_7_DAYS"
	DATE_RANGE_LAST_WEEK           = "LAST_WEEK"
	DATE_RANGE_LAST_BUSINESS_WEEK  = "LAST_BUSINESS_WEEK"
	DATE_RANGE_THIS_MONTH          = "THIS_MONTH"
	DATE_RANGE_LAST_MONTH          = "LAST_MONTH"
	DATE_RANGE_ALL_TIME            = "ALL_TIME"
	DATE_RANGE_CUSTOM_DATE         = "CUSTOM_DATE"
	DATE_RANGE_LAST_14_DAYS        = "LAST_14_DAYS"
	DATE_RANGE_LAST_30_DAYS        = "LAST_30_DAYS"
	DATE_RANGE_THIS_WEEK_SUN_TODAY = "THIS_WEEK_SUN_TODAY"
	DATE_RANGE_THIS_WEEK_MON_TODAY = "THIS_WEEK_MON_TODAY"
	DATE_RANGE_LAST_WEEK_SUN_SAT   = "LAST_WEEK_SUN_SAT"
)

const (
	SEGMENT_DATE_DAY     = "Date"
	SEGMENT_DATE_WEEK    = "Week"
	SEGMENT_DATE_MONTH   = "Month"
	SEGMENT_DATE_QUARTER = "Quarter"
	SEGMENT_DATE_YEAR    = "Year"
)

const (
	DOWNLOAD_FORMAT_XML = "XML"
)

type CampaignReport struct {
	XMLName xml.Name       `xml:"report"`
	Row     []*CampaignRow `xml:"table>row"`
}

type CampaignRow struct {
	XMLName     xml.Name `xml:"row"`
	AvgCPC      int64    `xml:"avgCPC,attr"`
	AvgCPM      int64    `xml:"avgCPM,attr"`
	CampaignID  int64    `xml:"campaignID,attr"`
	Clicks      int64    `xml:"clicks,attr"`
	Cost        int64    `xml:"cost,attr"`
	Impressions int64    `xml:"impressions,attr"`
}

type BudgetReport struct {
	XMLName xml.Name     `xml:"report"`
	Row     []*BudgetRow `xml:"table>row"`
}

type BudgetRow struct {
	XMLName     xml.Name `xml:"row"`
	AvgCPC      int64    `xml:"avgCPC,attr"`
	AvgCPM      int64    `xml:"avgCPM,attr"`
	CampaignID  int64    `xml:"campaignID,attr"`
	Clicks      int64    `xml:"clicks,attr"`
	Cost        int64    `xml:"cost,attr"`
	Impressions int64    `xml:"impressions,attr"`
	Conversions int64    `xml:"conversions,attr"`
}

type ReportDefinition struct {
	XMLName                xml.Name `xml:"reportDefinition"`
	Id                     string   `xml:"id,omitempty"`
	Selector               Selector `xml:"selector"`
	ReportName             string   `xml:"reportName"`
	ReportType             string   `xml:"reportType"`
	HasAttachment          string   `xml:"hasAttachment,omitempty"`
	DateRangeType          string   `xml:"dateRangeType"`
	CreationTime           string   `xml:"creationTime,omitempty"`
	DownloadFormat         string   `xml:"downloadFormat"`
	IncludeZeroImpressions bool     `xml:"includeZeroImpressions"`
}

//Magic that sets downloadFormat automaticaly
func (c *ReportDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	c.DownloadFormat = DOWNLOAD_FORMAT_XML

	type Alias ReportDefinition

	start.Name = xml.Name{
		"", "reportDefinition",
	}

	e.EncodeElement((*Alias)(c), start)
	return nil
}

type ReportUtils struct {
	Auth
}

func NewReportUtils(auth *Auth) *ReportUtils {
	return &ReportUtils{Auth: *auth}
}

func (s *ReportUtils) DownloadCampaignPerformaceReport(reportDefinition *ReportDefinition) (report CampaignReport, err error) {
	reportDefinition.ReportType = "CAMPAIGN_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)
	if err != nil {
		return report, err
	}

	report = CampaignReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadBudgetPerformanceReport(reportDefinition *ReportDefinition) (report BudgetReport, err error) {
	reportDefinition.ReportType = "BUDGET_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	//	respBody = `<?xml version='1.0' encoding='UTF-8' standalone='yes'?>
	//	<report>
	//		<report-name name="Report #553f5265b3d84"/>
	//		<date-range date="All Time"/>
	//		<table>
	//			<columns>
	//				<column display="Campaign ID" name="campaignID"/>
	//				<column display="Avg. CPC" name="avgCPC"/>
	//				<column display="Avg. CPM" name="avgCPM"/>
	//				<column display="Cost" name="cost"/>
	//				<column display="Clicks" name="clicks"/>
	//				<column display="Impressions" name="impressions"/>
	//			</columns>
	//			<row avgCPC="1" campaignID="246257700" clicks="4" cost="1" impressions="7"/>
	//			<row avgCPC="2" campaignID="246257700" clicks="4" cost="1" impressions="7"/>
	//			<row avgCPC="3" campaignID="246257700" clicks="4" cost="1" impressions="7"/>
	//			<row avgCPC="4" campaignID="246257700" clicks="4" cost="1" impressions="7"/>
	//		</table>
	//	</report>`

	if err != nil {
		return report, err
	}

	report = BudgetReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}
