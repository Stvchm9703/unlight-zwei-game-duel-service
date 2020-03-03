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

func (this *ECShortHand) Validate() error {
	return nil
}
func (this *GDCreateReq) Validate() error {
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
	for _, item := range this.HostExtraEc {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HostExtraEc", err)
			}
		}
	}
	for _, item := range this.DuelExtraEc {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DuelExtraEc", err)
			}
		}
	}
	return nil
}
func (this *GDGetInfoReq) Validate() error {
	return nil
}
func (this *GDBroadcastResp) Validate() error {
	for _, item := range this.InstanceSet {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("InstanceSet", err)
			}
		}
	}
	return nil
}
func (this *GDInstanceDT) Validate() error {
	for _, item := range this.UpdateCard {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateCard", err)
			}
		}
	}
	return nil
}
func (this *GDMoveConfirmReq) Validate() error {
	for _, item := range this.UpdateCard {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateCard", err)
			}
		}
	}
	for _, item := range this.TriggerSkl {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TriggerSkl", err)
			}
		}
	}
	return nil
}
func (this *GDMoveConfirmResp) Validate() error {
	return nil
}
func (this *GDADConfirmReq) Validate() error {
	for _, item := range this.TriggerSkl {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TriggerSkl", err)
			}
		}
	}
	for _, item := range this.UpdateCard {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateCard", err)
			}
		}
	}
	return nil
}
func (this *GDADResultResp) Validate() error {
	return nil
}
func (this *GDADDiceResult) Validate() error {
	return nil
}
func (this *GDPhaseConfirmReq) Validate() error {
	return nil
}
func (this *GDPhaseConfirmResp) Validate() error {
	return nil
}
func (this *GDChangeConfirmReq) Validate() error {
	return nil
}
