package main

import (
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
)

// maxMessageCount defines the max amount of messages stored in a model
const (
	maxMessageCount = 500
)

// Model Roles
const (
	Name = int(core.Qt__UserRole) + 1<<iota
	MessageID
	PostURL
	Author
	AuthorURL
	Avatar
	Body
	CreatedAt
	Actor
	ActorName
	Reply
	ReplyToID
	ReplyToAuthor
	Forward
	Mention
	Like
	Media
	Editing
	Liked
	Shared
)

// MessageModel holds a collection of messages
type MessageModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Message               `property:"messages"`

	_ func(*Message) `slot:"addMessage"`
	_ func(*Message) `slot:"appendMessage"`
	_ func(row int)  `slot:"removeMessage"`
}

// Message represents a single message
type Message struct {
	core.QObject

	Name          string
	MessageID     string
	PostURL       string
	Author        string
	AuthorURL     string
	Avatar        string
	Body          string
	CreatedAt     time.Time
	Actor         string
	ActorName     string
	Reply         bool
	ReplyToID     string
	ReplyToAuthor string
	Forward       bool
	Mention       bool
	Like          bool
	Media         string
	Editing       bool
	Liked         bool
	Shared        bool
}

func (m *MessageModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:          core.NewQByteArray2("name", -1),
		MessageID:     core.NewQByteArray2("messageid", -1),
		PostURL:       core.NewQByteArray2("posturl", -1),
		Author:        core.NewQByteArray2("author", -1),
		AuthorURL:     core.NewQByteArray2("authorurl", -1),
		Avatar:        core.NewQByteArray2("avatar", -1),
		Body:          core.NewQByteArray2("body", -1),
		CreatedAt:     core.NewQByteArray2("createdat", -1),
		Actor:         core.NewQByteArray2("actor", -1),
		ActorName:     core.NewQByteArray2("actorname", -1),
		Reply:         core.NewQByteArray2("reply", -1),
		ReplyToID:     core.NewQByteArray2("replytoid", -1),
		ReplyToAuthor: core.NewQByteArray2("replytoauthor", -1),
		Forward:       core.NewQByteArray2("forward", -1),
		Mention:       core.NewQByteArray2("mention", -1),
		Like:          core.NewQByteArray2("like", -1),
		Media:         core.NewQByteArray2("media", -1),
		Editing:       core.NewQByteArray2("editing", -1),
		Liked:         core.NewQByteArray2("liked", -1),
		Shared:        core.NewQByteArray2("shared", -1),
	})

	m.ConnectData(m.data)
	m.ConnectSetData(m.setData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddMessage(m.addMessage)
	m.ConnectAppendMessage(m.appendMessage)
	m.ConnectRemoveMessage(m.removeMessage)

	// keep time stamps ("1 minute ago") updated
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			m.updateMessageTime()
		}
	}()
}

func (m *MessageModel) setData(index *core.QModelIndex, value *core.QVariant, role int) bool {
	if !index.IsValid() {
		return false
	}

	var p = m.Messages()[len(m.Messages())-1-index.Row()]
	p.Editing = true

	m.DataChanged(index, index, []int{Editing})

	return true
}

func (m *MessageModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}
	if index.Row() >= len(m.Messages()) {
		return core.NewQVariant()
	}

	var p = m.Messages()[len(m.Messages())-1-index.Row()]
	switch role {
	case Name:
		{
			return core.NewQVariant14(p.Name)
		}
	case MessageID:
		{
			return core.NewQVariant14(p.MessageID)
		}
	case PostURL:
		{
			return core.NewQVariant14(p.PostURL)
		}
	case Author:
		{
			return core.NewQVariant14(p.Author)
		}
	case AuthorURL:
		{
			return core.NewQVariant14(p.AuthorURL)
		}
	case Avatar:
		{
			return core.NewQVariant14(p.Avatar)
		}
	case Body:
		{
			return core.NewQVariant14(p.Body)
		}
	case CreatedAt:
		{
			return core.NewQVariant14(humanize.Time(p.CreatedAt))
		}
	case Actor:
		{
			return core.NewQVariant14(p.Actor)
		}
	case ActorName:
		{
			return core.NewQVariant14(p.ActorName)
		}
	case Reply:
		{
			return core.NewQVariant11(p.Reply)
		}
	case ReplyToID:
		{
			return core.NewQVariant14(p.ReplyToID)
		}
	case ReplyToAuthor:
		{
			return core.NewQVariant14(p.ReplyToAuthor)
		}
	case Forward:
		{
			return core.NewQVariant11(p.Forward)
		}
	case Mention:
		{
			return core.NewQVariant11(p.Mention)
		}
	case Like:
		{
			return core.NewQVariant11(p.Like)
		}
	case Media:
		{
			return core.NewQVariant14(p.Media)
		}
	case Editing:
		{
			return core.NewQVariant11(p.Editing)
		}
	case Liked:
		{
			return core.NewQVariant11(p.Liked)
		}
	case Shared:
		{
			return core.NewQVariant11(p.Shared)
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *MessageModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Messages())
}

func (m *MessageModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *MessageModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *MessageModel) addMessage(p *Message) {
	m.BeginInsertRows(core.NewQModelIndex(), 0, 0)
	m.SetMessages(append(m.Messages(), p))
	m.EndInsertRows()

	if len(m.Messages()) > maxMessageCount {
		m.removeMessage(len(m.Messages()) - 1)
	}
}

func (m *MessageModel) appendMessage(p *Message) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Messages()), len(m.Messages()))
	m.SetMessages(append([]*Message{p}, m.Messages()...))
	m.EndInsertRows()
}

func (m *MessageModel) removeMessage(row int) {
	trow := len(m.Messages()) - 1 - row
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetMessages(append(m.Messages()[:trow], m.Messages()[trow+1:]...))
	m.EndRemoveRows()
}

func (m *MessageModel) updateMessageTime() {
	if len(m.Messages()) > 0 {
		var fIndex = m.Index(0, 0, core.NewQModelIndex())
		var lIndex = m.Index(len(m.Messages())-1, 0, core.NewQModelIndex())
		m.DataChanged(fIndex, lIndex, []int{CreatedAt})
	}
}

/*
func (m *MessageModel) editMessage(row int, param string) {
		var p = m.Messages()[row]

		var pIndex = m.Index(row, 0, core.NewQModelIndex())
		m.DataChanged(pIndex, pIndex, []int{roles})
}
*/

func init() {
	Message_QRegisterMetaType()
}
