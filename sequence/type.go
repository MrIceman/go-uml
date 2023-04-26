package sequence

type participant struct {
	Name string
}

type edge struct {
	from        participant
	to          participant
	directional bool
	Label       string
}

func (e *edge) From() participant {
	return e.from
}

func (e *edge) To() participant {
	return e.to
}

type participantCoord struct {
	X float64
	Y float64
}
