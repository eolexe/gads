package gads

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type OperationError struct {
	Code      int64  `xml:"OperationError>Code"`
	Details   string `xml:"OperationError>Details"`
	ErrorCode string `xml:"OperationError>ErrorCode"`
	Message   string `xml:"OperationError>Message"`
}

type BudgetError struct {
	Path    string `xml:"fieldPath"`
	String  string `xml:"errorString"`
	Trigger string `xml:"trigger"`
	Reason  string `xml:"reason"`
}

type CriterionError struct {
	FieldPath   string `xml:"fieldPath"`
	Trigger     string `xml:"trigger"`
	ErrorString string `xml:"errorString"`
	Reason      string `xml:"reason"`
}

type TargetError struct {
	FieldPath   string `xml:"fieldPath"`
	Trigger     string `xml:"trigger"`
	ErrorString string `xml:"errorString"`
	Reason      string `xml:"reason"`
}

type AdGroupServiceError struct {
	FieldPath   string `xml:"fieldPath"`
	Trigger     string `xml:"trigger"`
	ErrorString string `xml:"errorString"`
	Reason      string `xml:"reason"`
}

type NotEmptyError struct {
	FieldPath   string `xml:"fieldPath"`
	Trigger     string `xml:"trigger"`
	ErrorString string `xml:"errorString"`
	Reason      string `xml:"reason"`
}

type AdError struct {
	FieldPath   string `xml:"fieldPath"`
	Trigger     string `xml:"trigger"`
	ErrorString string `xml:"errorString"`
	Reason      string `xml:"reason"`
}

type ApiExceptionFault struct {
	Message string        `xml:"message"`
	Type    string        `xml:"ApplicationException.Type"`
	Errors  []interface{} `xml:"errors"`
}

func (aes *ApiExceptionFault) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		switch start := token.(type) {
		case xml.StartElement:
			switch start.Name.Local {
			case "message":
				if err := dec.DecodeElement(&aes.Message, &start); err != nil {
					return err
				}
			case "ApplicationException.Type":
				if err := dec.DecodeElement(&aes.Type, &start); err != nil {
					return err
				}
			case "errors":
				errorType, _ := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
				switch errorType {
				case "CriterionError":
					e := CriterionError{}
					dec.DecodeElement(&e, &start)
					aes.Errors = append(aes.Errors, e)
				case "TargetError":
					e := TargetError{}
					dec.DecodeElement(&e, &start)
					aes.Errors = append(aes.Errors, e)
				case "BudgetError":
					e := BudgetError{}
					dec.DecodeElement(&e, &start)
					aes.Errors = append(aes.Errors, e)
				case "AdGroupServiceError":
					e := AdGroupServiceError{}
					dec.DecodeElement(&e, &start)
					aes.Errors = append(aes.Errors, e)
				case "NotEmptyError":
					e := NotEmptyError{}
					dec.DecodeElement(&e, &start)
					aes.Errors = append(aes.Errors, e)
				case "AdError":
					e := AdError{}
					dec.DecodeElement(&e, &start)
					aes.Errors = append(aes.Errors, e)
				default:
					return fmt.Errorf("Unknown error type -> %s", start)
				}
			case "reason":
				break
			default:
				return fmt.Errorf("Unknown error field -> %s", start)
			}
		}
	}
	return err
}

type ErrorsType struct {
	ApiExceptionFaults []ApiExceptionFault `xml:"ApiExceptionFault"`
}

type Fault struct {
	XMLName     xml.Name   `xml:"Fault"`
	FaultCode   string     `xml:"faultcode"`
	FaultString string     `xml:"faultstring"`
	Errors      ErrorsType `xml:"detail"`
}

func (f *ErrorsType) Error() string {
	errors := []string{}
	for _, e := range f.ApiExceptionFaults {
		errors = append(errors, fmt.Sprintf("%s", e.Message))
	}
	return strings.Join(errors, "\n")
}

type ReportDownloadError struct {
	XMLName   xml.Name `xml:"reportDownloadError"`
	Type      string   `xml:"ApiError>type"`
	Trigger   string   `xml:"ApiError>trigger"`
	FieldPath string   `xml:"ApiError>fieldPath"`
}

func (f ReportDownloadError) Error() string {
	return fmt.Sprintf("Type = '%s', Trigger = '%s', FieldPath = '%s'", f.Type, f.Trigger, f.FieldPath)
}
