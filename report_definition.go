package gads

import (
	"encoding/xml"
)

type ReportDefinitionField struct {
	FieldName           string `xml:"fieldName"`
	DisplayFieldName    string `xml:"displayFieldName"`
	XmlAttributeName    string `xml:"xmlAttributeName"`
	FieldType           string `xml:"fieldType"`
	FieldBehavior       string `xml:"fieldBehavior"`
	EnumValues          string `xml:"enumValues"`
	CanSelect           string `xml:"canSelect"`
	CanFilter           string `xml:"canFilter"`
	IsEnumType          string `xml:"isEnumType"`
	IsBeta              string `xml:"isBeta"`
	IsZeroRowCompatible string `xml:"isZeroRowCompatible"`
	EnumValuePairs      string `xml:"enumValuePairs"`
}

type ReportDefinitionService struct {
	Auth
}

func NewReportDefinitionService(auth *Auth) *ReportDefinitionService {
	return &ReportDefinitionService{Auth: *auth}
}

func (s *ReportDefinitionService) Get(selector Selector) (reports []ReportDefinitionField, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		reportDefinitionServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return reports, totalCount, err
	}
	getResp := struct {
		Size    int64                   `xml:"rval>totalNumEntries"`
		Reports []ReportDefinitionField `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return reports, totalCount, err
	}
	return getResp.Reports, getResp.Size, err
}
