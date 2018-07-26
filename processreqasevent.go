package core

// ProcessReqAsEvent reschedules the request as an event at next clock cycle
// determined by frequency.
func ProcessReqAsEvent(request Req, engine Engine, freq Freq) {
	recvTime := request.RecvTime()
	eventTime := freq.ThisTick(recvTime)
	request.SetEventTime(eventTime)
	engine.Schedule(request)
}
