package treasury

type BookAccount struct {
	BookName    string
	WinRate     float64
	IsRestricted bool
}

func MonitorAccountHealth(accs []BookAccount) []string {
	var restrictedBooks []string
	for _, acc := range accs {
		if acc.WinRate > 0.62 || acc.IsRestricted {
			restrictedBooks = append(restrictedBooks, acc.BookName)
		}
	}
	return restrictedBooks // Signals Vampire Agent to drain balance and shift routing
}
