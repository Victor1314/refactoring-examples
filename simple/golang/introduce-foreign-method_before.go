


type Report struct {
	previousEnd time.Time
	// ...
}

func (r *Report) sendReport() {
	nextDay := r.previousEnd.AddDate(0, 0, 1)
	// ...
}

// class Report:
//     # ...
//     def sendReport(self):
//         nextDay = Date(self.previousEnd.getYear(),
//             self.previousEnd.getMonth(), self.previousEnd.getDate() + 1)
//         # ...