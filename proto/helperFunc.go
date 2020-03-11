package proto

import "fmt"

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
