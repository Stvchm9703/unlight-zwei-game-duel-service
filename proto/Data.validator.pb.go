// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Data.proto

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

func (this *GameDataSet) Validate() error {
	for _, item := range this.HostCardDeck {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HostCardDeck", err)
			}
		}
	}
	for _, item := range this.DuelCardDeck {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DuelCardDeck", err)
			}
		}
	}
	for _, item := range this.HostEventCardDeck {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HostEventCardDeck", err)
			}
		}
	}
	for _, item := range this.DuelEventCardDeck {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DuelEventCardDeck", err)
			}
		}
	}
	for _, item := range this.EffectCounter {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EffectCounter", err)
			}
		}
	}
	return nil
}
func (this *CharCardSet) Validate() error {
	for _, item := range this.StatusInst {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StatusInst", err)
			}
		}
	}
	if this.EquSet != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EquSet); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EquSet", err)
		}
	}
	return nil
}
func (this *CharCardEquSet) Validate() error {
	return nil
}
func (this *EventCard) Validate() error {
	return nil
}
func (this *SkillSet) Validate() error {
	for _, item := range this.CondCard {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CondCard", err)
			}
		}
	}
	return nil
}
func (this *SkillCardCond) Validate() error {
	return nil
}
func (this *StatusSet) Validate() error {
	return nil
}
func (this *Status_Effect) Validate() error {
	return nil
}
func (this *EffectResult) Validate() error {
	return nil
}
