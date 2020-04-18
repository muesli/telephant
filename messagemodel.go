package main

import (
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
)

// maxMessageCount defines the max amount of messages stored in a model
const (
	maxMessageCount = 80
)

// Model Roles
const (
	Name = int(core.Qt__UserRole) + iota
	MessageID
	PostID
	PostURL
	Author
	AuthorURL
	AuthorID
	Avatar
	Body
	Sensitive
	Warning
	CreatedAt
	Actor
	ActorName
	ActorID
	Reply
	ReplyToID
	ReplyToAuthor
	MentionIDs
	MentionNames
	Mentions
	Forward
	Mention
	Like
	Followed
	Following
	FollowedBy
	MediaPreview
	MediaURL
	Liked
	Shared
	RepliesCount
	SharesCount
	LikesCount
	Editing
	Visibility
)

// MessageModel holds a collection of messages
type MessageModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Message               `property:"messages"`

	_ func(*Message) `slot:"addMessage"`
	_ func(*Message) `slot:"appendMessage"`
	_ func(int)      `slot:"removeMessage"`
	_ func(string)   `slot:"removeMessageID"`
	_ func()         `slot:"clear"`
}

func (m *MessageModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:          core.NewQByteArray2("name", -1),
		MessageID:     core.NewQByteArray2("messageid", -1),
		PostID:        core.NewQByteArray2("postid", -1),
		PostURL:       core.NewQByteArray2("posturl", -1),
		Author:        core.NewQByteArray2("author", -1),
		AuthorURL:     core.NewQByteArray2("authorurl", -1),
		AuthorID:      core.NewQByteArray2("authorid", -1),
		Avatar:        core.NewQByteArray2("avatar", -1),
		Body:          core.NewQByteArray2("body", -1),
		Sensitive:     core.NewQByteArray2("sensitive", -1),
		Warning:       core.NewQByteArray2("warning", -1),
		CreatedAt:     core.NewQByteArray2("createdat", -1),
		Actor:         core.NewQByteArray2("actor", -1),
		ActorName:     core.NewQByteArray2("actorname", -1),
		ActorID:       core.NewQByteArray2("actorid", -1),
		Reply:         core.NewQByteArray2("reply", -1),
		ReplyToID:     core.NewQByteArray2("replytoid", -1),
		ReplyToAuthor: core.NewQByteArray2("replytoauthor", -1),
		MentionIDs:    core.NewQByteArray2("mentionids", -1),
		MentionNames:  core.NewQByteArray2("mentionnames", -1),
		Mentions:      core.NewQByteArray2("mentions", -1),
		Forward:       core.NewQByteArray2("forward", -1),
		Mention:       core.NewQByteArray2("mention", -1),
		Like:          core.NewQByteArray2("like", -1),
		Followed:      core.NewQByteArray2("followed", -1),
		Following:     core.NewQByteArray2("following", -1),
		FollowedBy:    core.NewQByteArray2("followedby", -1),
		MediaPreview:  core.NewQByteArray2("mediapreview", -1),
		MediaURL:      core.NewQByteArray2("mediaurl", -1),
		Liked:         core.NewQByteArray2("liked", -1),
		Shared:        core.NewQByteArray2("shared", -1),
		RepliesCount:  core.NewQByteArray2("repliescount", -1),
		SharesCount:   core.NewQByteArray2("sharescount", -1),
		LikesCount:    core.NewQByteArray2("likescount", -1),
		Visibility:    core.NewQByteArray2("visibility", -1),
	})

	m.ConnectData(m.data)
	m.ConnectSetData(m.setData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddMessage(m.addMessage)
	m.ConnectAppendMessage(m.appendMessage)
	m.ConnectRemoveMessage(m.removeMessage)
	m.ConnectClear(m.clear)

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
	if p == nil {
		return core.NewQVariant()
	}

	switch role {
	case Name:
		{
			return core.NewQVariant1(p.Name)
		}
	case MessageID:
		{
			return core.NewQVariant1(p.MessageID)
		}
	case PostID:
		{
			return core.NewQVariant1(p.PostID)
		}
	case PostURL:
		{
			return core.NewQVariant1(p.PostURL)
		}
	case Author:
		{
			return core.NewQVariant1(p.Author)
		}
	case AuthorURL:
		{
			return core.NewQVariant1(p.AuthorURL)
		}
	case AuthorID:
		{
			return core.NewQVariant1(p.AuthorID)
		}
	case Avatar:
		{
			return core.NewQVariant1(p.Avatar)
		}
	case Body:
		{
			return core.NewQVariant1(p.Body)
		}
	case Sensitive:
		{
			return core.NewQVariant1(p.Sensitive)
		}
	case Warning:
		{
			return core.NewQVariant1(p.Warning)
		}
	case CreatedAt:
		{
			if time.Since(p.CreatedAt) < time.Minute {
				return core.NewQVariant1("just now")
			}
			return core.NewQVariant1(humanize.Time(p.CreatedAt))
		}
	case Actor:
		{
			return core.NewQVariant1(p.Actor)
		}
	case ActorName:
		{
			return core.NewQVariant1(p.ActorName)
		}
	case ActorID:
		{
			return core.NewQVariant1(p.ActorID)
		}
	case Reply:
		{
			return core.NewQVariant1(p.Reply)
		}
	case ReplyToID:
		{
			return core.NewQVariant1(p.ReplyToID)
		}
	case ReplyToAuthor:
		{
			return core.NewQVariant1(p.ReplyToAuthor)
		}
	case MentionIDs:
		{
			return core.NewQVariant1(p.MentionIDs)
		}
	case MentionNames:
		{
			return core.NewQVariant1(p.MentionNames)
		}
	case Mentions:
		{
			var s string
			for _, v := range p.MentionNames {
				s += "@" + v + " "
			}
			return core.NewQVariant1(s)
		}
	case Forward:
		{
			return core.NewQVariant1(p.Forward)
		}
	case Mention:
		{
			return core.NewQVariant1(p.Mention)
		}
	case Like:
		{
			return core.NewQVariant1(p.Like)
		}
	case Followed:
		{
			return core.NewQVariant1(p.Followed)
		}
	case Following:
		{
			return core.NewQVariant1(p.Following)
		}
	case FollowedBy:
		{
			return core.NewQVariant1(p.FollowedBy)
		}
	case MediaPreview:
		{
			return core.NewQVariant1(p.MediaPreview)
		}
	case MediaURL:
		{
			return core.NewQVariant1(p.MediaURL)
		}
	case Liked:
		{
			return core.NewQVariant1(p.Liked)
		}
	case Shared:
		{
			return core.NewQVariant1(p.Shared)
		}
	case RepliesCount:
		{
			return core.NewQVariant1(p.RepliesCount)
		}
	case SharesCount:
		{
			return core.NewQVariant1(p.SharesCount)
		}
	case LikesCount:
		{
			return core.NewQVariant1(p.LikesCount)
		}
	case Visibility:
		{
			return core.NewQVariant1(p.Visibility)
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

func (m *MessageModel) clear() {
	m.BeginResetModel()
	m.SetMessages([]*Message{})
	m.EndResetModel()
}

func (m *MessageModel) addMessage(p *Message) {
	m.BeginInsertRows(core.NewQModelIndex(), 0, 0)
	addMessage(m, p)
	m.SetMessages(append(m.Messages(), p))
	m.EndInsertRows()

	if len(m.Messages()) > maxMessageCount {
		m.removeMessage(len(m.Messages()) - 1)
	}
}

func (m *MessageModel) appendMessage(p *Message) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Messages()), len(m.Messages()))
	addMessage(m, p)
	m.SetMessages(append([]*Message{p}, m.Messages()...))
	m.EndInsertRows()
}

func (m *MessageModel) removeMessage(row int) {
	trow := len(m.Messages()) - 1 - row
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	removeMessage(m, m.Messages()[trow])
	m.SetMessages(append(m.Messages()[:trow], m.Messages()[trow+1:]...))
	m.EndRemoveRows()
}

func (m *MessageModel) removeMessageID(id string) {
	for idx, v := range m.Messages() {
		if v.MessageID == id {
			trow := len(m.Messages()) - 1 - idx
			m.removeMessage(trow)
			break
		}
	}
}

func (m *MessageModel) updateMessage(id string) {
	for _, v := range m.Messages() {
		if v.MessageID == id {
			// FIXME: only update the affected rows
			var fIndex = m.Index(0, 0, core.NewQModelIndex())
			var lIndex = m.Index(len(m.Messages())-1, 0, core.NewQModelIndex())
			m.DataChanged(fIndex, lIndex, []int{})
			break
		}
	}
}

func (m *MessageModel) updateMessageTime() {
	// debugln("Updating timestamps...")
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
	MessageModel_QRegisterMetaType()
	Message_QRegisterMetaType()
}
