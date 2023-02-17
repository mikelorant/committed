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

func (m *Model) setSave() bool {
	save := m.snapshotToSave()

	switch {
	case m.currentSave.amend && save.amend:
		m.loadSave(save)
		return true
	case m.previousSave.amend && save.amend:
		m.swapSave()
		m.loadSave(save)
		return true
	case (save.body != "" || save.emoji.Name != "" || save.summary != "") && !m.currentSave.amend:
		m.loadSave(save)
		return true
	case (save.body != "" || save.emoji.Name != "" || save.summary != "") && !m.previousSave.amend:
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
