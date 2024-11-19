


type Report struct {
	previousEnd time.Time
	// ...
}

func (r *Report) sendReport() {
	newStart := r.nextDay(r.previousEnd)
	// ...
}

func nextDay(arg time.Time) time.Time {
	return arg.AddDate(0, 0, 1)
}

// class Report:
//     # ...
//     def sendReport(self):
//         newStart = self._nextDay(self.previousEnd)
//         # ...

//     def _nextDay(self, arg):
//         return Date(arg.getYear(), arg.getMonth(), arg.getDate() + 1)