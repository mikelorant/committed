package ui

func (m *Model) restoreModel(save savedState) {
	m.models.header.Amend = save.amend
	m.models.header.Emoji = save.emoji
	m.models.header.SetSummary(save.summary)
	m.models.body.SetValue(save.body)
}

func (m *Model) backupModel() savedState {
	var save savedState

	save.amend = m.models.header.Amend
	save.emoji = m.models.header.Emoji
	save.summary = m.models.header.Summary()
	save.body = m.models.body.RawValue()

	return save
}

func (m *Model) setSaves() {
	m.file = m.state.Options.File.MessageFile != ""
	m.amend = m.state.Options.Amend || m.state.File.Amend

	switch m.amend {
	case true:
		switch {
		case m.state.File.Amend:
			m.currentSave = defaultHookEditorSave(m.state)
		default:
			m.currentSave = defaultAmendSave(m.state)
		}

	case false:
		if m.file {
			m.currentSave = defaultHookEditorSave(m.state)
		}

		switch {
		case m.state.File.Amend:
			m.previousSave = defaultHookEditorSave(m.state)
		default:
			m.previousSave = defaultAmendSave(m.state)
		}
	}
}

func (m *Model) setSave() bool {
	save := m.snapshotToSave()

	hasSave := (save.body != "" || save.emoji.Name != "" || save.summary != "")

	switch {
	case m.currentSave.amend && save.amend:
		m.loadSave(save)
		return true
	case m.previousSave.amend && save.amend:
		m.swapSave()
		m.loadSave(save)
		return true
	case hasSave && !m.currentSave.amend:
		m.loadSave(save)
		return true
	case hasSave && !m.previousSave.amend:
		m.swapSave()
		m.loadSave(save)
		return true
	}

	return false
}

func (m *Model) swapSave() {
	m.currentSave = m.backupModel()

	m.models.header.ResetSummary()
	m.models.body.Reset()

	m.currentSave, m.previousSave = m.previousSave, m.currentSave

	m.restoreModel(m.currentSave)
}

func (m *Model) loadSave(st savedState) {
	m.models.header.ResetSummary()
	m.models.body.Reset()

	m.restoreModel(st)
}

func (m Model) snapshotToSave() savedState {
	s := savedState{
		amend:   m.state.Snapshot.Amend,
		summary: m.state.Snapshot.Summary,
		body:    m.state.Snapshot.Body,
	}

	if e := m.state.Emojis.Find(m.state.Snapshot.Emoji); e.Valid {
		s.emoji = e.Emoji
	}

	return s
}
