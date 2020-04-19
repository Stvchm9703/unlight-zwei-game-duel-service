// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SESkillCalReq) Validate() error {
	for _, item := range this.IncomeCard {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("IncomeCard", err)
			}
		}
	}
	for _, item := range this.Feat {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Feat", err)
			}
		}
	}
	return nil
}
func (this *SESkillCalResp) Validate() error {
	for _, item := range this.EffectResult {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EffectResult", err)
			}
		}
	}
	return nil
}
func (this *SEDiceCalReq) Validate() error {
	return nil
}
func (this *SEDiceCalResp) Validate() error {
	for _, item := range this.DiceResult {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DiceResult", err)
			}
		}
	}
	return nil
}
func (this *DiceResultSet) Validate() error {
	return nil
}
func (this *SEEffectCalReq) Validate() error {
	if this.FromTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FromTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FromTime", err)
		}
	}
	if this.ToTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToTime", err)
		}
	}
	if this.GamesetInstant != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GamesetInstant); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GamesetInstant", err)
		}
	}
	return nil
}
func (this *SEEffectCalResp) Validate() error {
	if this.GamesetResult != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GamesetResult); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GamesetResult", err)
		}
	}
	return nil
}