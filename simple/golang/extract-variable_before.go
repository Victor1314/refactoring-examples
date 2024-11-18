



func renderBanner(platform string, browser string, wasInitialized bool, resize int) {
	if strings.Contains(strings.ToUpper(platform), "MAC") &&
		strings.Contains(strings.ToUpper(browser), "IE") &&
		wasInitialized && resize > 0 {
		// do something
	}
}

// def renderBanner(self):
//     if (self.platform.toUpperCase().indexOf("MAC") > -1) and \
//        (self.browser.toUpperCase().indexOf("IE") > -1) and \
//        self.wasInitialized() and (self.resize > 0):
//         # do something