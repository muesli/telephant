package main

import (
	"github.com/therecipe/qt/core"
)

// Model Roles
const (
	PaneID = int(core.Qt__UserRole) + iota
	PaneName
	PaneSticky
	PaneDefault
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
	_ func()        `slot:"clear"`
}

// Pane represents a single pane
type Pane struct {
	core.QObject

	ID      string
	Name    string
	Sticky  bool
	Default bool
	Model   *MessageModel
}

func (m *PaneModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		PaneName:   core.NewQByteArray2("panename", -1),
		PaneSticky: core.NewQByteArray2("panesticky", -1),
		MsgModel:   core.NewQByteArray2("msgmodel", -1),
	})

	m.ConnectData(m.data)
	m.ConnectSetData(m.setData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddPane(m.addPane)
	m.ConnectRemovePane(m.removePane)
	m.ConnectClear(m.clear)
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
			return core.NewQVariant1(p.Name)
		}
	case PaneSticky:
		{
			return core.NewQVariant1(p.Sticky)
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

func (m *PaneModel) clear() {
	m.BeginResetModel()
	m.SetPanes([]*Pane{})
	m.EndResetModel()
}

func (m *PaneModel) addPane(p *Pane) {
	// add pane before the last pane, which is always the notifications pane
	if len(m.Panes()) == 0 {
		m.BeginInsertRows(core.NewQModelIndex(), 0, 0)
		m.SetPanes(append(m.Panes(), p))
	} else {
		m.BeginInsertRows(core.NewQModelIndex(), len(m.Panes())-1, len(m.Panes())-1)
		m.SetPanes(append(m.Panes()[:len(m.Panes())-1], p, m.Panes()[len(m.Panes())-1]))
	}
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
