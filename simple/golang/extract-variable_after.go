

func renderBanner(platform string, browser string, wasInitialized bool, resize int) {

	isMacOs := strings.Contains(strings.ToUpper(platform), "MAC")
	isIE := strings.Contains(strings.ToUpper(browser), "IE")

	if isMacOs && isIE && wasInitialized && resize > 0 {
		// do something
	}

}
