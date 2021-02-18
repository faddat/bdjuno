package types

// ConsensusEvent represents a consensus event
type ConsensusEvent struct {
	Height     int64
	Round      int32
	Step       string
	Percentage float64
}

// NewConsensusEvent allows to easily build a new ConsensusEvent object
func NewConsensusEvent(height int64, round int32, step string, percentage float64) *ConsensusEvent {
	return &ConsensusEvent{
		Height:     height,
		Round:      round,
		Step:       step,
		Percentage: percentage,
	}
}

// Equal tells whether c and other contain the same data
func (c ConsensusEvent) Equal(other ConsensusEvent) bool {
	return c.Height == other.Height &&
		c.Round == other.Round &&
		c.Step == other.Step &&
		c.Percentage == other.Percentage
}
