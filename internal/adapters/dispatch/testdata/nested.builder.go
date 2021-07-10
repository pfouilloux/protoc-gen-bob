package __

// MsgBuilder provides a fluent api for building an instance of Msg
type MsgBuilder struct {
	msg *Msg
}

// NewMsg starts building the Msg
func NewMsg() *MsgBuilder {
	return &MsgBuilder{&Msg{}}
}

// Val sets Msg.Val
func (msgBuilder *MsgBuilder) Val(val *Msg_NestedBuilder) *MsgBuilder {
	msgBuilder.msg.Val = val.Build()
	return msgBuilder
}

// Build builds the Msg
func (msgBuilder *MsgBuilder) Build() *Msg {
	return msgBuilder.msg
}

// Msg_NestedBuilder provides a fluent api for building an instance of Msg_Nested
type Msg_NestedBuilder struct {
	msg_nested *Msg_Nested
}

// NewMsg_Nested starts building the Msg_Nested
func NewMsg_Nested() *Msg_NestedBuilder {
	return &Msg_NestedBuilder{&Msg_Nested{}}
}

// Build builds the Msg_Nested
func (msg_nestedBuilder *Msg_NestedBuilder) Build() *Msg_Nested {
	return msg_nestedBuilder.msg_nested
}
