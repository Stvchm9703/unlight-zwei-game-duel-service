package proto

import "fmt"

// RdsKeyName : the common modal naming for PhaseSnapMod
func (*PhaseSnapMod) RdsKeyName() string {
	return ":PhaseState"
}

// RdsKeyName : the common model naming for MovePhaseSnapMod
func (*MovePhaseSnapMod) RdsKeyName() string {
	return ":MvPhMod"
}

// RdsKeyName : the common model naming for MovePhaseResult
func (*GDMoveConfirmResp) RdsKeyName() string {
	return ":MvPhModResult"
}

// RdsKeyName : the common model naming for ADPhaseSnapMod
func (*ADPhaseSnapMod) RdsKeyName() string {
	return ":ADPhMod"
}

// RdsKeyName : the common model naming for EffectNodeSnapMod
func (*EffectNodeSnapMod) RdsKeyName() string {
	return ":EfMod"
}

func CleanAfterExec(orig []*EffectResult, changed []*EffectResult) {
	fmt.Println("in orig, before exec")
	for k := range orig {
		fmt.Printf("%v :\n %#v \n\n", orig[k], *orig[k])
	}
	fmt.Println("in changed, before exec")
	for k := range changed {
		fmt.Printf("%v :\n %#v \n\n", orig[k], *orig[k])
		changed[k].RemainCd--
	}
	fmt.Println("in orig, after exec")
	for k := 0; k < len(orig); k++ {
		fmt.Printf("%v :\n %#v \n\n", orig[k], *orig[k])
		if orig[k].RemainCd == 0 {
			RemoveEffREsult(orig, k)
			k--
		}
	}
	fmt.Println("in orig, final exec")
	for k := range orig {
		fmt.Printf("%v :\n %#v \n\n", orig[k], *orig[k])
	}
}

func RemoveEffREsult(slice []*EffectResult, s int) []*EffectResult {
	return append(slice[:s], slice[s+1:]...)
}

// EffectResult sorting
func NodeFilter(vs []*EffectResult, f func(*EffectResult) bool) []*EffectResult {
	vsf := make([]*EffectResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// EventCard filter
func EventCardFilter(vs []*EventCard, f func(*EventCard) bool) []*EventCard {
	vsf := make([]*EventCard, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

type EventCardListSet struct {
	Set []*EventCard `json:"set,omitempty"`
}

func (set *EventCardListSet) RdsKeyName(side PlayerSide) string {
	if side == PlayerSide_HOST {
		return ":HtEvtCrdDk"
	} else if side == PlayerSide_DUELER {
		return ":DlEvtCrdDk"
	}
	return ":EvtCrdDk"
}

func (card *EventCard) ToECShostHand() *ECShortHand {
	return &ECShortHand{
		CardId:   card.Id,
		Position: card.Position,
		IsInvert: card.IsInvert,
	}
}
