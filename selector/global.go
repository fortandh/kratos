package selector

import "fmt"

var globalSelector = &wrapSelector{}

var _ Builder = (*wrapSelector)(nil)

// wrapSelector wrapped Selector, help override global Selector implementation.
type wrapSelector struct{ Builder }

// GlobalSelector returns global selector builder.
func GlobalSelector() Builder {
	fmt.Println("selector/global.go:GlobalSelector(): start")
	if globalSelector.Builder != nil {
		fmt.Println("selector/global.go:GlobalSelector(): builder is not nil")
		fmt.Printf("selector/global.go:GlobalSelector(): builder.type: %v\n", globalSelector.Builder)
		return globalSelector
	}
	fmt.Println("selector/global.go:GlobalSelector(): builder is nil")
	return nil
}

// SetGlobalSelector set global selector builder.
func SetGlobalSelector(builder Builder) {
	fmt.Println("selector/global.go:SetGlobalSelector(): start")
	globalSelector.Builder = builder
}
