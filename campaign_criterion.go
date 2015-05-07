package gads

import (
	//  "strings"
	//  "strconv"
	"encoding/xml"
	"fmt"
)

type CampaignCriterionService struct {
	Auth
}

func NewCampaignCriterionService(auth *Auth) *CampaignCriterionService {
	return &CampaignCriterionService{Auth: *auth}
}

type CampaignCriterion struct {
	CampaignId  int64     `xml:"campaignId"`
	Criterion   Criterion `xml:"criterion"`
	BidModifier float64   `xml:"bidModifier,omitempty"`
	Errors      []error   `xml:"-"`
}

func (cc CampaignCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	isNegative := false
	//fmt.Printf("processing -> %#v\n",ncc)
	start.Attr = append(
		start.Attr,
		xml.Attr{
			xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"},
			"CampaignCriterion",
		},
	)
	e.EncodeToken(start)
	e.EncodeElement(&cc.CampaignId, xml.StartElement{Name: xml.Name{"", "campaignId"}})
	e.EncodeElement(&isNegative, xml.StartElement{Name: xml.Name{"", "isNegative"}})
	if err := criterionMarshalXML(cc.Criterion, e); err != nil {
		return err
	}
	e.EncodeToken(start.End())
	return nil
}

type NegativeCampaignCriterion struct {
	CampaignId  int64     `xml:"campaignId"`
	Criterion   Criterion `xml:"criterion"`
	BidModifier *float64   `xml:"bidModifier,omitempty"`
	Errors      []error   `xml:"-"`

}

type CampaignCriterions []interface{}
type CampaignCriterionOperations map[string]CampaignCriterions

func (ncc NegativeCampaignCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	isNegative := true
	//fmt.Printf("processing -> %#v\n",ncc)
	start.Attr = append(
		start.Attr,
		xml.Attr{
			xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"},
			"NegativeCampaignCriterion",
		},
	)
	e.EncodeToken(start)
	e.EncodeElement(&ncc.CampaignId, xml.StartElement{Name: xml.Name{"", "campaignId"}})
	e.EncodeElement(&isNegative, xml.StartElement{Name: xml.Name{"", "isNegative"}})
	criterionMarshalXML(ncc.Criterion, e)
	e.EncodeToken(start.End())
	return nil
}

func (ccs *CampaignCriterions) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	cc := NegativeCampaignCriterion{}
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			switch start.Name.Local {
			case "campaignId":
				if err := dec.DecodeElement(&cc.CampaignId, &start); err != nil {
					return err
				}
			case "criterion":
				criterion, err := criterionUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				cc.Criterion = criterion
			case "bidModifier":
				if err := dec.DecodeElement(&cc.BidModifier, &start); err != nil {
					return err
				}
			}
		}
	}
	*ccs = append(*ccs, cc)
	return nil
}

/*
func NewNegativeCampaignCriterion(campaignId int64, bidModifier float64, criterion interface{}) CampaignCriterion {
  return CampaignCriterion{
    CampaignId: campaignId,
    Criterion: criterion,
    BidModifier: bidModifier
  }
  switch c := criterion.(type) {
  case AdScheduleCriterion:
  case AgeRangeCriterion:
  case ContentLabelCriterion:
  case GenderCriterion:
  case KeywordCriterion:
  case LanguageCriterion:
  case LocationCriterion:
  case MobileAppCategoryCriterion:
  case MobileApplicationCriterion:
  case MobileDeviceCriterion:
  case OperatingSystemVersionCriterion:
  case PlacementCriterion:
  case PlatformCriterion:
  case ProductCriterion:
  case ProximityCriterion:
  case UserInterestCriterion:
    cc.Criterion = criterion
  case UserListCriterion:
    cc.Criterion = criterion
  case VerticalCriterion:
  }
}
*/

func (s *CampaignCriterionService) Get(selector Selector) (campaignCriterions CampaignCriterions, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		campaignCriterionServiceUrl,
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

//	respBody := `<getResponse xmlns="https://adwords.google.com/api/adwords/cm/v201409"><rval><totalNumEntries>8</totalNumEntries><Page.Type>CampaignCriterionPage</Page.Type><entries><campaignId>210576900</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Platform"><id>30000</id><type>PLATFORM</type><Criterion.Type>Platform</Criterion.Type></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>210576900</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Platform"><id>30001</id><type>PLATFORM</type><Criterion.Type>Platform</Criterion.Type></criterion><bidModifier>2.319999933242798</bidModifier><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>210576900</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Platform"><id>30002</id><type>PLATFORM</type><Criterion.Type>Platform</Criterion.Type></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>210576900</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Location"><id>1023191</id><type>LOCATION</type><Criterion.Type>Location</Criterion.Type><locationName>New York</locationName><displayType>City</displayType><targetingStatus>ACTIVE</targetingStatus><parentLocations><id>2840</id><Criterion.Type>Location</Criterion.Type></parentLocations><parentLocations><id>21167</id><Criterion.Type>Location</Criterion.Type></parentLocations><parentLocations><id>200501</id><Criterion.Type>Location</Criterion.Type></parentLocations></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>231500820</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Platform"><id>30000</id><type>PLATFORM</type><Criterion.Type>Platform</Criterion.Type></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>231500820</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Platform"><id>30001</id><type>PLATFORM</type><Criterion.Type>Platform</Criterion.Type></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>231500820</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Platform"><id>30002</id><type>PLATFORM</type><Criterion.Type>Platform</Criterion.Type></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries><entries><campaignId>231500820</campaignId><isNegative>false</isNegative><criterion xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="Location"><id>9007534</id><type>LOCATION</type><Criterion.Type>Location</Criterion.Type><locationName>20010</locationName><displayType>Postal Code</displayType><targetingStatus>ACTIVE</targetingStatus><parentLocations><id>2840</id><Criterion.Type>Location</Criterion.Type></parentLocations><parentLocations><id>21140</id><Criterion.Type>Location</Criterion.Type></parentLocations><parentLocations><id>200511</id><Criterion.Type>Location</Criterion.Type></parentLocations><parentLocations><id>1014895</id><Criterion.Type>Location</Criterion.Type></parentLocations><parentLocations><id>9040468</id><Criterion.Type>Location</Criterion.Type></parentLocations></criterion><CampaignCriterion.Type>CampaignCriterion</CampaignCriterion.Type></entries></rval></getResponse>`

	if err != nil {
		return campaignCriterions, totalCount, err
	}
	getResp := struct {
		Size               int64              `xml:"rval>totalNumEntries"`
		CampaignCriterions CampaignCriterions `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return campaignCriterions, totalCount, err
	}
	return getResp.CampaignCriterions, getResp.Size, err
}

func (s *CampaignCriterionService) Mutate(campaignCriterionOperations CampaignCriterionOperations) (campaignCriterions CampaignCriterions, err error) {
	type campaignCriterionOperation struct {
		Action            string      `xml:"operator"`
		CampaignCriterion interface{} `xml:"operand"`
	}
	operations := []campaignCriterionOperation{}
	for action, campaignCriterions := range campaignCriterionOperations {
		for _, campaignCriterion := range campaignCriterions {
			operations = append(operations,
				campaignCriterionOperation{
					Action:            action,
					CampaignCriterion: campaignCriterion,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []campaignCriterionOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}
	respBody, err := s.Auth.request(campaignCriterionServiceUrl, "mutate", mutation)
	if err != nil {
		/*
			    switch t := err.(type) {
			    case *ErrorsType:
				    for action, campaignCriterions := range campaignCriterionOperations {
					    for _, campaignCriterion := range campaignCriterions {
			          campaignCriterions = append(campaignCriterions,campaignCriterion)
			        }
			      }
			      for _, aef := range t.ApiExceptionFaults {
			        for _,e := range aef.Errors {
			          switch et := e.(type) {
			          case CriterionError:
			            offset, err := strconv.ParseInt(strings.Trim(et.FieldPath,"abcdefghijklmnop.]["),10,64)
			            if err != nil {
			              return CampaignCriterions{}, err
			            }
			            cc := campaignCriterions[offset]
			            switch c := cc.(type) {
			            case CampaignCriterion:
			              CampaignCriterion(campaignCriterions[offset]).Errors = append(campaignCriterions[offset].(CampaignCriterion).Errors,fmt.Errorf(et.Reason))
			            case NegativeCampaignCriterion:
			              NegativeCampaignCriterion(campaignCriterions[offset]).Errors = append(NegativeCampaignCriterion(campaignCriterions[offset].Errors),fmt.Errorf(et.Reason))
			            }
			          }
			        }
			      }
			    default:
		*/
		return campaignCriterions, err
		//}
	}
	mutateResp := struct {
		CampaignCriterions CampaignCriterions `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return campaignCriterions, err
	}
	return mutateResp.CampaignCriterions, err
}

func (s *CampaignCriterionService) Query(query string) (campaignCriterions CampaignCriterions, err error) {
	return campaignCriterions, err
}
