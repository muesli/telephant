package main

import (
	"github.com/therecipe/qt/core"
)

// Model Roles
const (
	AttachmentID = int(core.Qt__UserRole) + 1<<iota
	AttachmentPreview
)

// AttachmentModel holds a collection of attachments
type AttachmentModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Attachment            `property:"attachments"`

	_ func(*Attachment) `slot:"addAttachment"`
	_ func(row int)     `slot:"removeAttachment"`
	_ func()            `slot:"clear"`
}

// Attachment represents a single attachment
type Attachment struct {
	core.QObject

	ID      string
	Preview string
	URL     string
}

func (m *AttachmentModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		AttachmentID:      core.NewQByteArray2("attachmentID", -1),
		AttachmentPreview: core.NewQByteArray2("attachmentPreview", -1),
	})

	m.ConnectData(m.data)
	m.ConnectSetData(m.setData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddAttachment(m.addAttachment)
	m.ConnectRemoveAttachment(m.removeAttachment)
	m.ConnectClear(m.clear)
}

func (m *AttachmentModel) setData(index *core.QModelIndex, value *core.QVariant, role int) bool {
	if !index.IsValid() {
		return false
	}

	m.DataChanged(index, index, []int{Editing})

	return true
}

func (m *AttachmentModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}
	if index.Row() >= len(m.Attachments()) {
		return core.NewQVariant()
	}

	var p = m.Attachments()[index.Row()]
	switch role {
	case AttachmentID:
		{
			return core.NewQVariant14(p.ID)
		}
	case AttachmentPreview:
		{
			return core.NewQVariant14(p.Preview)
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *AttachmentModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Attachments())
}

func (m *AttachmentModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *AttachmentModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *AttachmentModel) clear() {
	m.BeginResetModel()
	m.SetAttachments([]*Attachment{})
	m.EndResetModel()
}

func (m *AttachmentModel) addAttachment(p *Attachment) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Attachments()), len(m.Attachments()))
	m.SetAttachments(append(m.Attachments(), p))
	m.EndInsertRows()
}

func (m *AttachmentModel) removeAttachment(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetAttachments(append(m.Attachments()[:row], m.Attachments()[row+1:]...))
	m.EndRemoveRows()
}

func init() {
	AttachmentModel_QRegisterMetaType()
	Attachment_QRegisterMetaType()
}
