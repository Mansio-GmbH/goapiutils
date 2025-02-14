package types

import (
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/mansio-gmbh/goapiutils/ct"
)

type DepotDoc struct {
	BaseDoc
	Name                string          `json:"name"`
	TenantID            string          `json:"tenantId"`
	Location            *ct.Location    `json:"location"`
	CoreOperationalArea OperationalArea `json:"coreOperationalArea"`
}

type OperationalArea struct {
	IncludedPostalCodePrefixes *[]string `json:"includedPostalCodePrefixes" dynamodbav:"includedPostalCodePrefixes"`
	ExcludedPostalCodePrefixes *[]string `json:"excludedPostalCodePrefixes" dynamodbav:"excludedPostalCodePrefixes"`
}

func (oa OperationalArea) Includes(postalCodes ...string) bool {
	if oa.IncludedPostalCodePrefixes == nil {
		return false
	}
	included := pie.SortUsing(*oa.IncludedPostalCodePrefixes, func(i, j string) bool {
		return len(i) <= len(j)
	})
	excluded := make([]string, 0)
	if oa.ExcludedPostalCodePrefixes != nil {
		excluded = pie.SortUsing(*oa.ExcludedPostalCodePrefixes, func(i, j string) bool {
			return len(i) <= len(j)
		})
	}

	lenIncluded := len(included)
	lenExcluded := len(excluded)

	maxLength := len(included[lenIncluded-1])
	if lenExcluded > 0 && maxLength < len(excluded[lenExcluded-1]) {
		maxLength = len(excluded[lenExcluded-1])
	}

	for _, code := range postalCodes {
		currentIncludedIndex := 0
		currentExcludedIndex := 0
		pcIsIncluded := false
		for currentLength := 1; currentLength <= maxLength; currentLength++ {
			for {
				if lenIncluded <= currentIncludedIndex || len(included[currentIncludedIndex]) != currentLength {
					break
				}
				if strings.HasPrefix(code, included[currentIncludedIndex]) {
					pcIsIncluded = true
				}
				currentIncludedIndex++
			}
			for {
				if lenExcluded <= currentExcludedIndex || len(excluded[currentExcludedIndex]) != currentLength {
					break
				}
				if strings.HasPrefix(code, excluded[currentExcludedIndex]) {
					pcIsIncluded = false
				}
				currentExcludedIndex++
			}
		}
		if !pcIsIncluded {
			return false
		}
	}

	return true
}
