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
func (msgBuilder *MsgBuilder) Val(val Msg_MyEnum) *MsgBuilder {
	msgBuilder.msg.Val = val
	return msgBuilder
}

// Build builds the Msg
func (msgBuilder *MsgBuilder) Build() *Msg {
	return msgBuilder.msg
}
