package types

// AccountResourceTag returns the path tag to the Account resource, which is 0x01+hash(0x0.LibraAccount.T)
func AccountResourceTag() AccessPathTag {
	return &StructTag{
		Module: "LibraAccount",
		Name:   "T",
	}
}

// AccountResourcePath returns the raw path to the Account resource, which is 0x01+hash(0x0.LibraAccount.T)
func AccountResourcePath() []byte {
	p := &DecodedPath{
		Tag: AccountResourceTag(),
	}
	b, _ := p.MarshalBinary()
	return b
}

// AccountSentEventPath returns the raw path to the coin sent events, which is
// "0x01+hash(0x0.LibraAccount.T)"/sent_events_count/
func AccountSentEventPath() []byte {
	p := &DecodedPath{
		Tag:      AccountResourceTag(),
		Accesses: []string{"sent_events_count", ""},
	}
	b, _ := p.MarshalBinary()
	return b
}

// AccountReceivedEventPath returns the raw path to the coin received events, which is
// "0x01+hash(0x0.LibraAccount.T)"/received_events_count/
func AccountReceivedEventPath() []byte {
	p := &DecodedPath{
		Tag:      AccountResourceTag(),
		Accesses: []string{"received_events_count", ""},
	}
	b, _ := p.MarshalBinary()
	return b
}

var (
	pathTagNameMap = map[string]string{
		string(AccountResourcePath()): "0x0.LibraAccount.T",
	}
)

// InferPathTagName returns the name of known path root tags, by tag hash and type.
// Known tags:
//  - 0x0.LibraAccount.T
func InferPathTagName(tag AccessPathTag) string {
	key := string(append([]byte{tag.TypePrefix()}, tag.Hash()...))
	if name, ok := pathTagNameMap[key]; ok {
		return name
	}
	return "unknown"
}
