package entproto

import "google.golang.org/protobuf/types/descriptorpb"

const (
	TypeBool   = descriptorpb.FieldDescriptorProto_TYPE_BOOL
	TypeString = descriptorpb.FieldDescriptorProto_TYPE_STRING
)

func NamedMessages(messages ...*namedMessage) MessageOption {
	return func(msg *message) {
		msg.NamedMessages = append(msg.NamedMessages, messages...)
	}
}

func NamedMessage(name string) *namedMessage {
	return &namedMessage{
		ProtoMessageOptions: protoMessageOptions{},
		Name:                name,
	}
}

type namedMessage struct {
	ProtoMessageOptions protoMessageOptions
	Name                string
	Groups              FieldGroups
	ExtraFields         []pbfield
}

func (nm *namedMessage) WithGroups(groups *FieldGroups) *namedMessage {
	nm.Groups = *groups
	return nm
}

func (nm *namedMessage) WithSkipID(skip bool) *namedMessage {
	nm.ProtoMessageOptions.SkipID = skip
	return nm
}

func (nm *namedMessage) WithExtraFields(fields ...*extraField) *namedMessage {
	nm.ProtoMessageOptions.ExtraFields = append(nm.ProtoMessageOptions.ExtraFields, fields...)
	return nm
}

func ExtraField(name string, number int) *extraField {
	return &extraField{
		Name: name,
		Descriptor: pbfield{
			Number: number,
		},
	}
}

func (ef *extraField) WithType(t descriptorpb.FieldDescriptorProto_Type) *extraField {
	ef.Descriptor.Type = t
	return ef
}

func (ef *extraField) WithTypeName(name string) *extraField {
	ef.Descriptor.TypeName = name
	return ef
}
