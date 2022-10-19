package main

import "fmt"

// Browsers: Firefox
// Link: RegionBlockedSite, NormalSite
// Adapter: VPN

// Firefox browser
type Firefox struct{}

func (f *Firefox) OpenSite(s NormalSite) { fmt.Println(s.NormalData()) }

// Normal site
type NormalSite interface {
	NormalData() string
}

type DuckDuckGo struct{}

func (ddg *DuckDuckGo) NormalData() string {
	return "You can search stuff here more anonymously than via Google :)"
}

// Region blocked site
type RegionBlockedSite interface {
	RegionBlockedData() string
}

type RefactoringGuru struct{}

func (rg *RefactoringGuru) RegionBlockedData() string {
	return "We got blocked because of racoon image xD"
}

// Adapter, VPNAccess
type VPNAccess struct {
	site RegionBlockedSite
}

func (vpn *VPNAccess) Access(site RegionBlockedSite) (self *VPNAccess) {
	vpn.site = site
	return vpn
}

func (vpn *VPNAccess) NormalData() string {
	return vpn.site.RegionBlockedData() // `RegionBlockedSite` now can be accessed as a normal one
}

func main() {
	f := Firefox{}

	ddg := DuckDuckGo{}
	f.OpenSite(&ddg) // Firefox can work with DuckDuckGo's interface due to it being normal site

	rg := RefactoringGuru{}
	// f.OpenSite(&rg) // ...`*RefactoringGuru` does not implement `NormalSite`
	vpn := (&VPNAccess{}).Access(&rg)
	f.OpenSite(vpn)
}
