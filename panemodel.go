package main

import (
	"github.com/therecipe/qt/core"
)

// Model Roles
const (
	PaneName = int(core.Qt__UserRole) + 1<<iota
	MsgModel
)

// PaneModel holds a collection of panes
type PaneModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Pane                  `property:"panes"`

	_ func(*Pane)   `slot:"addPane"`
	_ func(row int) `slot:"removePane"`
}

// Pane represents a single pane
type Pane struct {
	core.QObject

	Name  string
	Model *MessageModel
}

func (m *PaneModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		PaneName: core.NewQByteArray2("panename", -1),
		MsgModel: core.NewQByteArray2("msgmodel", -1),
	})

	m.ConnectData(m.data)
	m.ConnectSetData(m.setData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddPane(m.addPane)
	m.ConnectRemovePane(m.removePane)
}

func (m *PaneModel) setData(index *core.QModelIndex, value *core.QVariant, role int) bool {
	if !index.IsValid() {
		return false
	}

	// var p = m.Panes()[len(m.Panes())-1-index.Row()]
	m.DataChanged(index, index, []int{Editing})

	return true
}

func (m *PaneModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}
	if index.Row() >= len(m.Panes()) {
		return core.NewQVariant()
	}

	var p = m.Panes()[index.Row()]
	switch role {
	case PaneName:
		{
			return core.NewQVariant14(p.Name)
		}
	case MsgModel:
		{
			return p.Model.ToVariant()
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *PaneModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Panes())
}

func (m *PaneModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *PaneModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *PaneModel) addPane(p *Pane) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Panes()), len(m.Panes()))
	m.SetPanes(append(m.Panes(), p))
	m.EndInsertRows()
}

func (m *PaneModel) removePane(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetPanes(append(m.Panes()[:row], m.Panes()[row+1:]...))
	m.EndRemoveRows()
}

func init() {
	PaneModel_QRegisterMetaType()
	Pane_QRegisterMetaType()
}
