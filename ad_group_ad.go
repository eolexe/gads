package gads

import (
	"encoding/xml"
	"fmt"
)

type AdGroupAdService struct {
	Auth
}

type AppUrl struct {
	Url    string `xml:"url"`
	OsType string `xml:"osType"` // "OS_TYPE_IOS", "OS_TYPE_ANDROID", "UNKNOWN"
}

type TextAd struct {
	AdGroupId            int64             `xml:"-"`
	Id                   int64             `xml:"id,omitempty"`
	Url                  string            `xml:"url,omitempty"`
	DisplayUrl           string            `xml:"displayUrl,omitempty"`
	FinalUrls            []string          `xml:"finalUrls,omitempty"`
	FinalMobileUrls      []string          `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls         []AppUrl          `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate  string            `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters  *CustomParameters `xml:"urlCustomParameters,omitempty"`
	DevicePreference     int64             `xml:"devicePreference,omitempty"`
	Headline             string            `xml:"headline,omitempty"`
	Description1         string            `xml:"description1,omitempty"`
	Description2         string            `xml:"description2,omitempty"`
	Status               string            `xml:"-"`
	ApprovalStatus       string            `xml:"-"`
	DisapprovalReasons   []string          `xml:"-"`
	TrademarkDisapproved bool              `xml:"-"`
	Labels               []Label           `xml:"-"`
}

type ImageAd struct {
	AdGroupId            int64             `xml:"-"`
	Id                   int64             `xml:"id,omitempty"`
	Url                  string            `xml:"url,omitempty"`
	DisplayUrl           string            `xml:"displayUrl,omitempty"`
	FinalUrls            []string          `xml:"finalUrls,omitempty"`
	FinalMobileUrls      []string          `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls         []AppUrl          `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate  string            `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters  *CustomParameters `xml:"urlCustomParameters,omitempty"`
	DevicePreference     int64             `xml:"devicePreference,omitempty"`
	Image                Media             `xml:"image,omitempty"`
	Name                 string            `xml:"name,omitempty"`
	AdToCopyImageFrom    int64             `xml:"adToCopyImageFrom,omitempty"`
	Status               string            `xml:"-"`
	ApprovalStatus       string            `xml:"-"`
	DisapprovalReasons   []string          `xml:"-"`
	TrademarkDisapproved bool              `xml:"-"`
	Labels               []Label           `xml:"-"`
}

type MobileAd struct {
	AdGroupId            int64             `xml:"-"`
	Id                   int64             `xml:"id,omitempty"`
	Url                  string            `xml:"url"`
	DisplayUrl           string            `xml:"displayUrl"`
	FinalUrls            []string          `xml:"finalUrls,omitempty"`
	FinalMobileUrls      []string          `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls         []AppUrl          `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate  string            `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters  *CustomParameters `xml:"urlCustomParameters,omitempty"`
	DevicePreference     int64             `xml:"devicePreference,omitempty"`
	Headline             string            `xml:"headline"`
	Description          string            `xml:"description"`
	MarkupLanguages      []string          `xml:"markupLanguages"`
	MobileCarriers       []string          `xml:"mobileCarriers"`
	BusinessName         string            `xml:"businessName"`
	CountryCode          string            `xml:"countryCode"`
	PhoneNumber          string            `xml:"phoneNumber"`
	Status               string            `xml:"-"`
	ApprovalStatus       string            `xml:"-"`
	DisapprovalReasons   []string          `xml:"-"`
	TrademarkDisapproved bool              `xml:"-"`
	Labels               []Label           `xml:"-"`
}

type TemplateElementField struct {
	Name       string `xml:"name"`
	Type       string `xml:"type"`
	FieldText  string `xml:"fieldText"`
	FieldMedia string `xml:"fieldMedia"`
}

type TemplateElement struct {
	UniqueName string                 `xml:"uniqueName"`
	Fields     []TemplateElementField `xml:"fields"`
}

type TemplateAd struct {
	AdGroupId            int64             `xml:"-"`
	Id                   int64             `xml:"id,omitempty"`
	Url                  string            `xml:"url"`
	DisplayUrl           string            `xml:"displayUrl"`
	FinalUrls            []string          `xml:"finalUrls,omitempty"`
	FinalMobileUrls      []string          `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls         []AppUrl          `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate  string            `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters  *CustomParameters `xml:"urlCustomParameters,omitempty"`
	DevicePreference     int64             `xml:"devicePreference,omitempty"`
	TemplateId           int64             `xml:"templateId"`
	AdUnionId            int64             `xml:"adUnionId"`
	TemplateElements     []TemplateElement `xml:"templateElements"`
	Dimensions           []Dimensions      `xml:"dimensions"`
	Name                 string            `xml:"name"`
	Duration             int64             `xml:"duration"`
	originAdId           *int64            `xml:"originAdId"`
	Status               string            `xml:"-"`
	ApprovalStatus       string            `xml:"-"`
	DisapprovalReasons   []string          `xml:"-"`
	TrademarkDisapproved bool              `xml:"-"`
	Labels               []Label           `xml:"-"`
}

type AdGroupAdOperations map[string]AdGroupAds

func NewAdGroupAdService(auth *Auth) *AdGroupAdService {
	return &AdGroupAdService{Auth: *auth}
}

func NewTextAd(adGroupId int64, url, displayUrl, headline, description1, description2, status string) TextAd {
	return TextAd{
		AdGroupId:    adGroupId,
		Url:          url,
		DisplayUrl:   displayUrl,
		Headline:     headline,
		Description1: description1,
		Description2: description2,
		Status:       status,
	}
}

type AdGroupAdLabel struct {
	AdGroupAdId int64 `xml:"adGroupAdId"`
	LabelId     int64 `xml:"labelId"`
}

type AdGroupAdLabelOperations map[string][]AdGroupAdLabel

type AdUrlUpgrade struct {
	AdId                int64  `xml:"adId"`
	FinalUrl            string `xml:"finalUrl"`
	FinalMobileUrl      string `xml:"finalMobileUrl"`
	TrackingUrlTemplate string `xml:"trackingUrlTemplate"`
}

// Get returns an array of ad's and the total number of ad's matching
// the selector.
//
// Example
//
//   ads, totalCount, err := adGroupAdService.Get(
//     gads.Selector{
//       Fields: []string{
//         "AdGroupId",
//         "Status",
//         "AdGroupCreativeApprovalStatus",
//         "AdGroupAdDisapprovalReasons",
//         "AdGroupAdTrademarkDisapproved",
//       },
//       Predicates: []gads.Predicate{
//         {"AdGroupId", "EQUALS", []string{adGroupId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "AdGroupId", "Id", "Url", "DisplayUrl", "CreativeFinalUrls", "CreativeFinalMobileUrls",
//   "CreativeFinalAppUrls", "CreativeTrackingUrlTemplate", "CreativeUrlCustomParameters",
//   "DevicePreference", "Status", "AdGroupCreativeApprovalStatus", "AdGroupAdDisapprovalReasons"
//   "AdGroupAdTrademarkDisapproved", "Labels"
//
//   TextAd
//     "Headline", "Description1", "Description2"
//
//   ImageAd
//     "ImageCreativeName"
//
// filterable fields are
//   "AdGroupId", "Id", "Url", "DisplayUrl", "CreativeFinalUrls", "CreativeFinalMobileUrls",
//   "CreativeFinalAppUrls", "CreativeTrackingUrlTemplate", "CreativeUrlCustomParameters",
//   "DevicePreference", "Status", "AdGroupCreativeApprovalStatus", "AdGroupAdDisapprovalReasons"
//   "Labels"
//
//   TextAd specific fields
//     "Headline", "Description1", "Description2"
//
//   ImageAd specific fields
//     "ImageCreativeName"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#get
//
func (s AdGroupAdService) Get(selector Selector) (adGroupAds AdGroupAds, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupAdServiceUrl,
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
		return adGroupAds, totalCount, err
	}

	//	respBody := `<getResponse xmlns="https://adwords.google.com/api/adwords/cm/v201409"><rval><totalNumEntries>1</totalNumEntries><Page.Type>AdGroupAdPage</Page.Type><entries><adGroupId>16946665260</adGroupId><ad xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="ImageAd"><id>61965159780</id><url>http://example.com</url><displayUrl>example.com</displayUrl><Ad.Type>ImageAd</Ad.Type><image><type>IMAGE</type><referenceId>-2469457768847769600</referenceId><dimensions><key>SHRUNKEN</key><value><width>300</width><height>239</height></value></dimensions><dimensions><key>FULL</key><value><width>300</width><height>250</height></value></dimensions><dimensions><key>PREVIEW</key><value><width>120</width><height>100</height></value></dimensions><urls><key>SHRUNKEN</key><value>https://tpc.googlesyndication.com/pageadimg/imgad?id=CICAgKDjp67d3QEQrAIY7wEoATIIr-a3ag4_bPY</value></urls><urls><key>FULL</key><value>https://tpc.googlesyndication.com/pageadimg/imgad?id=CICAgKDjp67d3QEQrAIY-gEoATIImuEGs1sFFh8</value></urls><urls><key>PREVIEW</key><value>https://tpc.googlesyndication.com/pageadimg/imgad?id=CICAgKDjp67d3QEQeBhkKAEyCMbjS9Yi6TAi</value></urls><mimeType>IMAGE_JPEG</mimeType><fileSize>0</fileSize><Media.Type>Image</Media.Type></image><name>gfdhfghfgh</name></ad><status>ENABLED</status><approvalStatus>UNCHECKED</approvalStatus><trademarkDisapproved>false</trademarkDisapproved></entries></rval></getResponse>`
	getResp := struct {
		Size       int64      `xml:"rval>totalNumEntries"`
		AdGroupAds AdGroupAds `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	return getResp.AdGroupAds, getResp.Size, err
}

// Mutate allows you to add, modify and remove ads, returning the
// modified ads.
//
// Example
//
//  ads, err := adGroupAdService.Mutate(
//    gads.AdGroupAdOperations{
//      "ADD": {
//        gads.NewTextAd(
//          adGroup.Id,
//          "https://classdo.com/en",
//          "classdo.com",
//          "test headline",
//          "test line one",
//          "test line two",
//          "PAUSED",
//        ),
//      },
//      "SET": {
//        modifiedAd,
//      },
//      "REMOVE": {
//        adNeedingRemoval,
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#mutate
//
func (s *AdGroupAdService) Mutate(adGroupAdOperations AdGroupAdOperations) (adGroupAds AdGroupAds, err error) {
	type adGroupAdOperation struct {
		Action    string     `xml:"operator"`
		AdGroupAd AdGroupAds `xml:"operand"`
	}
	operations := []adGroupAdOperation{}
	for action, adGroupAds := range adGroupAdOperations {
		for _, adGroupAd := range adGroupAds {
			operations = append(operations,
				adGroupAdOperation{
					Action:    action,
					AdGroupAd: []interface{}{adGroupAd},
				},
			)
		}
	}

	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupAdOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}

	respBody, err := s.Auth.request(adGroupAdServiceUrl, "mutate", mutation)
	if err != nil {
		return adGroupAds, err
	}
	mutateResp := struct {
		AdGroupAds AdGroupAds `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupAds, err
	}
	return mutateResp.AdGroupAds, err
}

// MutateLabel allows you to add and removes labels from ads.
//
// Example
//
//  ads, err := adGroupAdService.MutateLabel(
//    gads.AdGroupAdLabelOperations{
//      "ADD": {
//        gads.AdGroupAdLabel{AdGroupAdId: 3200, LabelId: 5353},
//        gads.AdGroupAdLabel{AdGroupAdId: 4320, LabelId: 5643},
//      },
//      "REMOVE": {
//        gads.AdGroupAdLabel{AdGroupAdId: 3653, LabelId: 5653},
//      },
//    }
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#mutateLabel
//
func (s *AdGroupAdService) MutateLabel(adGroupAdLabelOperations AdGroupAdLabelOperations) (adGroupAdLabels []AdGroupAdLabel, err error) {
	type adGroupAdLabelOperation struct {
		Action         string         `xml:"operator"`
		AdGroupAdLabel AdGroupAdLabel `xml:"operand"`
	}
	operations := []adGroupAdLabelOperation{}
	for action, adGroupAdLabels := range adGroupAdLabelOperations {
		for _, adGroupAdLabel := range adGroupAdLabels {
			operations = append(operations,
				adGroupAdLabelOperation{
					Action:         action,
					AdGroupAdLabel: adGroupAdLabel,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupAdLabelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutateLabel",
		},
		Ops: operations}
	respBody, err := s.Auth.request(adGroupAdServiceUrl, "mutateLabel", mutation)
	if err != nil {
		return adGroupAdLabels, err
	}
	mutateResp := struct {
		AdGroupAdLabels []AdGroupAdLabel `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupAdLabels, err
	}

	return mutateResp.AdGroupAdLabels, err
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#query
//
func (s *AdGroupAdService) Query(query string) (adGroupAds AdGroupAds, totalCount int64, err error) {
	return adGroupAds, totalCount, ERROR_NOT_YET_IMPLEMENTED
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#upgradeUrl
//
func (s *AdGroupAdService) UpgradeUrl(adUrlUpgrades []AdUrlUpgrade) (adGroupAds AdGroupAds, err error) {
	return adGroupAds, ERROR_NOT_YET_IMPLEMENTED
}
