package ingame

// Register event subscribers
func (p EtherProx) registerSubscribers() error {
	// Show community treasury at top of screens
	// event.Subscribe(p.Event(), 0, p.treasuryBar())

	return nil
}

func (p *EtherProx) Init() error {
	// Initialize our commands and event ubscribers
	p.EtherTrust()
	return p.registerSubscribers()
}

func (p *EtherProx) EtherTrust() {

	p.Command().Register(makeExchangeUser())
	p.Command().Register(sellGameCurrency())
	p.grabBalance()
}
