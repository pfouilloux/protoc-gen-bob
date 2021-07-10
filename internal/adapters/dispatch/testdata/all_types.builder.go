package __

// MsgBuilder provides a fluent api for building an instance of Msg
type MsgBuilder struct {
	msg *Msg
}

// NewMsg starts building the Msg
func NewMsg() *MsgBuilder {
	return &MsgBuilder{&Msg{}}
}

// Field0 sets Msg.Field0
func (msgBuilder *MsgBuilder) Field0(field0 float32) *MsgBuilder {
	msgBuilder.msg.Field0 = field0
	return msgBuilder
}

// Field1 sets Msg.Field1
func (msgBuilder *MsgBuilder) Field1(field1 float64) *MsgBuilder {
	msgBuilder.msg.Field1 = field1
	return msgBuilder
}

// Field2 sets Msg.Field2
func (msgBuilder *MsgBuilder) Field2(field2 int64) *MsgBuilder {
	msgBuilder.msg.Field2 = field2
	return msgBuilder
}

// Field3 sets Msg.Field3
func (msgBuilder *MsgBuilder) Field3(field3 uint64) *MsgBuilder {
	msgBuilder.msg.Field3 = field3
	return msgBuilder
}

// Field4 sets Msg.Field4
func (msgBuilder *MsgBuilder) Field4(field4 int32) *MsgBuilder {
	msgBuilder.msg.Field4 = field4
	return msgBuilder
}

// Field5 sets Msg.Field5
func (msgBuilder *MsgBuilder) Field5(field5 uint64) *MsgBuilder {
	msgBuilder.msg.Field5 = field5
	return msgBuilder
}

// Field6 sets Msg.Field6
func (msgBuilder *MsgBuilder) Field6(field6 uint32) *MsgBuilder {
	msgBuilder.msg.Field6 = field6
	return msgBuilder
}

// Field7 sets Msg.Field7
func (msgBuilder *MsgBuilder) Field7() *MsgBuilder {
	msgBuilder.msg.Field7 = true
	return msgBuilder
}

// Field8 sets Msg.Field8
func (msgBuilder *MsgBuilder) Field8(field8 string) *MsgBuilder {
	msgBuilder.msg.Field8 = field8
	return msgBuilder
}

// Field10 sets Msg.Field10
func (msgBuilder *MsgBuilder) Field10(field10 Msg) *MsgBuilder {
	msgBuilder.msg.Field10 = field10
	return msgBuilder
}

// Field11 sets Msg.Field11
func (msgBuilder *MsgBuilder) Field11(field11 []byte) *MsgBuilder {
	msgBuilder.msg.Field11 = field11
	return msgBuilder
}

// Field12 sets Msg.Field12
func (msgBuilder *MsgBuilder) Field12(field12 uint32) *MsgBuilder {
	msgBuilder.msg.Field12 = field12
	return msgBuilder
}

// Field13 sets Msg.Field13
func (msgBuilder *MsgBuilder) Field13(field13 MyEnum) *MsgBuilder {
	msgBuilder.msg.Field13 = field13
	return msgBuilder
}

// Field14 sets Msg.Field14
func (msgBuilder *MsgBuilder) Field14(field14 int32) *MsgBuilder {
	msgBuilder.msg.Field14 = field14
	return msgBuilder
}

// Field15 sets Msg.Field15
func (msgBuilder *MsgBuilder) Field15(field15 int64) *MsgBuilder {
	msgBuilder.msg.Field15 = field15
	return msgBuilder
}

// Field16 sets Msg.Field16
func (msgBuilder *MsgBuilder) Field16(field16 int32) *MsgBuilder {
	msgBuilder.msg.Field16 = field16
	return msgBuilder
}

// Field17 sets Msg.Field17
func (msgBuilder *MsgBuilder) Field17(field17 int64) *MsgBuilder {
	msgBuilder.msg.Field17 = field17
	return msgBuilder
}

// Build builds the Msg
func (msgBuilder *MsgBuilder) Build() *Msg {
	return msgBuilder.msg
}
