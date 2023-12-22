package converter

import (
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
)

func getStringValidationRules(r *validate.StringRules, ignoreEmpty bool) []string {
	if r == nil {
		return []string{}
	}

	rules := []string{}

	if r.Const != nil {
		rules = append(rules, fmt.Sprintf("Must be `%s`", *r.Const))
	}

	if r.Contains != nil {
		rules = append(rules, fmt.Sprintf("Must contain `%s`", *r.Contains))
	}

	if r.In != nil {
		rules = append(rules, fmt.Sprintf("Must be one of: `%s`", strings.Join(r.In, "`, `")))
	}

	if r.Len != nil {
		msg := "Exactly %d character%s long"
		if ignoreEmpty {
			msg = "Must be empty, or exactly %d character%s long"
		}

		rules = append(rules, fmt.Sprintf(msg, *r.Len, utils.PluralSuffix(int(*r.Len), "s", "")))
	}

	if r.MinLen != nil {
		msg := "Must be at least %d character%s long"
		if ignoreEmpty {
			msg = "Must be empty, or at least %d character%s long"
		}

		rules = append(rules, fmt.Sprintf(msg, *r.MinLen, utils.PluralSuffix(int(*r.MinLen), "s", "")))
	}

	if r.MaxLen != nil {
		rules = append(rules, fmt.Sprintf("Must be %d or fewer character%s long", *r.MaxLen, utils.PluralSuffix(int(*r.MaxLen), "s", "")))
	}

	if r.LenBytes != nil {
		msg := "Exactly %d byte%s long"
		if ignoreEmpty {
			msg = "Must be empty, or exactly %d byte%s long"
		}

		rules = append(rules, fmt.Sprintf(msg, *r.LenBytes, utils.PluralSuffix(int(*r.LenBytes), "s", "")))
	}

	if r.MinBytes != nil {
		msg := "Must be at least %d byte%s long"
		if ignoreEmpty {
			msg = "Must be empty, or at least %d byte%s long"
		}

		rules = append(rules, fmt.Sprintf(msg, *r.MinBytes, utils.PluralSuffix(int(*r.MinBytes), "s", "")))
	}

	if r.NotContains != nil {
		rules = append(rules, fmt.Sprintf("Must not contain `%s`", *r.NotContains))
	}

	if r.NotIn != nil {
		rules = append(rules, fmt.Sprintf("Must not be one of: `%s`", strings.Join(r.NotIn, "`, `")))
	}

	if r.Pattern != nil {
		rules = append(rules, fmt.Sprintf("Must match the regex pattern `%s`", *r.Pattern))
	}

	if r.Prefix != nil {
		rules = append(rules, fmt.Sprintf("Must start with `%s`", *r.Prefix))
	}

	if r.Suffix != nil {
		rules = append(rules, fmt.Sprintf("Must end with `%s`", *r.Suffix))
	}

	if r.GetAddress() {
		rules = append(rules, "Must be a valid hostname or IP address")
	}

	if r.GetEmail() {
		rules = append(rules, "Must be a valid email address")
	}

	if r.GetHostname() {
		rules = append(rules, "Must be a valid hostname")
	}

	if r.GetIp() {
		rules = append(rules, "Must be a valid IP address")
	}

	if r.GetIpPrefix() {
		rules = append(rules, "Must be a valid IP prefix (eg, 20.0.0.0/16)")
	}

	if r.GetUri() {
		rules = append(rules, "Must be a valid URI")
	}

	if r.GetUriRef() {
		rules = append(rules, "Must be a valid URI reference")
	}

	if r.GetUuid() {
		rules = append(rules, "Must be a valid UUID")
	}

	return rules
}

func getTimestampValidationRules(r *validate.TimestampRules, ignoreEmpty bool) []string {
	if r == nil {
		return []string{}
	}

	rules := []string{}

	if r.Const != nil {
		rules = append(rules, fmt.Sprintf("Must be `%s`", r.Const.String()))
	}

	if r.Within != nil {
		rules = append(rules, fmt.Sprintf("Must be within %s of now", r.Within.String()))
	}

	if r.GetGt() != nil {
		rules = append(rules, fmt.Sprintf("Must be after %s", r.GetGt().String()))
	}

	if r.GetGtNow() {
		rules = append(rules, "Must be after now")
	}

	if r.GetGte() != nil {
		rules = append(rules, fmt.Sprintf("Must be equal to or after %s", r.GetGte().String()))
	}

	if r.GetLt() != nil {
		rules = append(rules, fmt.Sprintf("Must be before %s", r.GetLt().String()))
	}

	if r.GetLtNow() {
		rules = append(rules, "Must be before now")
	}

	if r.GetLte() != nil {
		rules = append(rules, fmt.Sprintf("Must be equal to or before %s", r.GetLte().String()))
	}

	return rules
}

func getInt32ValidationRules(r *validate.Int32Rules, ignoreEmpty bool) []string {
	if r == nil {
		return []string{}
	}

	rules := []string{}

	if r.Const != nil {
		rules = append(rules, fmt.Sprintf("Must be `%d`", *r.Const))
	}

	if r.In != nil {
		rules = append(rules, fmt.Sprintf("Must be one of: %v", r.In))
	}

	if r.NotIn != nil {
		rules = append(rules, fmt.Sprintf("Must not be one of: %v", r.NotIn))
	}

	if r.GreaterThan != nil {
		if gt, ok := r.GreaterThan.(*validate.Int32Rules_Gt); ok {
			rules = append(rules, fmt.Sprintf("Must be greater than `%d`", gt.Gt))
		} else if gte, ok := r.GreaterThan.(*validate.Int32Rules_Gte); ok {
			rules = append(rules, fmt.Sprintf("Must be greater than or equal to `%d`", gte.Gte))
		} else {
			rules = append(rules, "Unknown greater-than rule, fix the generator!")
		}
	}

	if r.LessThan != nil {
		if lt, ok := r.LessThan.(*validate.Int32Rules_Lt); ok {
			rules = append(rules, fmt.Sprintf("Must be less than `%d`", lt.Lt))
		} else if lte, ok := r.LessThan.(*validate.Int32Rules_Lte); ok {
			rules = append(rules, fmt.Sprintf("Must be less than or equal to `%d`", lte.Lte))
		} else {
			rules = append(rules, "Unknown less-than rule, fix the generator!")
		}
	}

	return rules
}

func getUInt32ValidationRules(r *validate.UInt32Rules, ignoreEmpty bool) []string {
	if r == nil {
		return []string{}
	}

	rules := []string{}

	if r.Const != nil {
		rules = append(rules, fmt.Sprintf("Must be `%d`", *r.Const))
	}

	if r.In != nil {
		rules = append(rules, fmt.Sprintf("Must be one of: %v", r.In))
	}

	if r.NotIn != nil {
		rules = append(rules, fmt.Sprintf("Must not be one of: %v", r.NotIn))
	}

	if r.GreaterThan != nil {
		if gt, ok := r.GreaterThan.(*validate.UInt32Rules_Gt); ok {
			rules = append(rules, fmt.Sprintf("Must be greater than `%d`", gt.Gt))
		} else if gte, ok := r.GreaterThan.(*validate.UInt32Rules_Gte); ok {
			rules = append(rules, fmt.Sprintf("Must be greater than or equal to `%d`", gte.Gte))
		} else {
			rules = append(rules, "Unknown greater-than rule, fix the generator!")
		}
	}

	if r.LessThan != nil {
		if lt, ok := r.LessThan.(*validate.UInt32Rules_Lt); ok {
			rules = append(rules, fmt.Sprintf("Must be less than `%d`", lt.Lt))
		} else if lte, ok := r.LessThan.(*validate.UInt32Rules_Lte); ok {
			rules = append(rules, fmt.Sprintf("Must be less than or equal to `%d`", lte.Lte))
		} else {
			rules = append(rules, "Unknown less-than rule, fix the generator!")
		}
	}

	return rules
}
