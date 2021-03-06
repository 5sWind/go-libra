package stdscript

var (
	// PeerToPeerTransfer byte code
	PeerToPeerTransfer = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 6, 0, 0, 0, 13, 86, 0, 0, 0, 6, 0, 0, 0, 5, 92, 0, 0, 0, 41, 0, 0, 0, 4, 133, 0, 0, 0, 32, 0, 0, 0, 7, 165, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 2, 5, 3, 0, 3, 2, 5, 3, 3, 0, 6, 60, 83, 69, 76, 70, 62, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 4, 109, 97, 105, 110, 15, 112, 97, 121, 95, 102, 114, 111, 109, 95, 115, 101, 110, 100, 101, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 4, 0, 11, 0, 11, 1, 18, 1, 1, 2}

	// CreateAccount byte code
	CreateAccount = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 6, 0, 0, 0, 13, 86, 0, 0, 0, 6, 0, 0, 0, 5, 92, 0, 0, 0, 44, 0, 0, 0, 4, 136, 0, 0, 0, 32, 0, 0, 0, 7, 168, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 2, 5, 3, 0, 3, 2, 5, 3, 3, 0, 6, 60, 83, 69, 76, 70, 62, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 4, 109, 97, 105, 110, 18, 99, 114, 101, 97, 116, 101, 95, 110, 101, 119, 95, 97, 99, 99, 111, 117, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 4, 0, 11, 0, 11, 1, 18, 1, 1, 2}

	// Mint byte code
	Mint = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 6, 0, 0, 0, 3, 76, 0, 0, 0, 6, 0, 0, 0, 12, 82, 0, 0, 0, 6, 0, 0, 0, 13, 88, 0, 0, 0, 6, 0, 0, 0, 5, 94, 0, 0, 0, 51, 0, 0, 0, 4, 145, 0, 0, 0, 32, 0, 0, 0, 7, 177, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 3, 0, 1, 4, 0, 2, 0, 2, 5, 3, 0, 3, 2, 5, 3, 3, 0, 6, 60, 83, 69, 76, 70, 62, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 9, 76, 105, 98, 114, 97, 67, 111, 105, 110, 4, 109, 97, 105, 110, 15, 109, 105, 110, 116, 95, 116, 111, 95, 97, 100, 100, 114, 101, 115, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 4, 0, 11, 0, 11, 1, 18, 1, 1, 2}

	// RotateAuthKey byte code
	RotateAuthKey = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 5, 0, 0, 0, 13, 85, 0, 0, 0, 5, 0, 0, 0, 5, 90, 0, 0, 0, 51, 0, 0, 0, 4, 141, 0, 0, 0, 32, 0, 0, 0, 7, 173, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 1, 9, 0, 3, 1, 9, 3, 0, 6, 60, 83, 69, 76, 70, 62, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 4, 109, 97, 105, 110, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 3, 0, 11, 0, 18, 1, 1, 2}

	// RotateConsensusKey byte code
	RotateConsensusKey = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 5, 0, 0, 0, 13, 85, 0, 0, 0, 5, 0, 0, 0, 5, 90, 0, 0, 0, 52, 0, 0, 0, 4, 142, 0, 0, 0, 32, 0, 0, 0, 7, 174, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 1, 9, 0, 3, 1, 9, 3, 0, 6, 60, 83, 69, 76, 70, 62, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 4, 109, 97, 105, 110, 23, 114, 111, 116, 97, 116, 101, 95, 99, 111, 110, 115, 101, 110, 115, 117, 115, 95, 112, 117, 98, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 3, 0, 11, 0, 18, 1, 1, 2}

	// AddValidator byte code
	AddValidator = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 5, 0, 0, 0, 13, 85, 0, 0, 0, 5, 0, 0, 0, 5, 90, 0, 0, 0, 38, 0, 0, 0, 4, 128, 0, 0, 0, 32, 0, 0, 0, 7, 160, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 1, 5, 0, 3, 1, 5, 3, 0, 6, 60, 83, 69, 76, 70, 62, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 4, 109, 97, 105, 110, 13, 97, 100, 100, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 3, 0, 11, 0, 18, 1, 1, 2}

	// RemoveValidator byte code
	RemoveValidator = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 5, 0, 0, 0, 13, 85, 0, 0, 0, 5, 0, 0, 0, 5, 90, 0, 0, 0, 41, 0, 0, 0, 4, 131, 0, 0, 0, 32, 0, 0, 0, 7, 163, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 1, 5, 0, 3, 1, 5, 3, 0, 6, 60, 83, 69, 76, 70, 62, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 4, 109, 97, 105, 110, 16, 114, 101, 109, 111, 118, 101, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 3, 0, 11, 0, 18, 1, 1, 2}

	// RegisterValidator byte code
	RegisterValidator = []byte{161, 28, 235, 11, 1, 0, 7, 1, 70, 0, 0, 0, 4, 0, 0, 0, 3, 74, 0, 0, 0, 6, 0, 0, 0, 12, 80, 0, 0, 0, 10, 0, 0, 0, 13, 90, 0, 0, 0, 10, 0, 0, 0, 5, 100, 0, 0, 0, 57, 0, 0, 0, 4, 157, 0, 0, 0, 32, 0, 0, 0, 7, 189, 0, 0, 0, 23, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 3, 0, 2, 0, 6, 9, 9, 9, 9, 9, 9, 0, 3, 6, 9, 9, 9, 9, 9, 9, 3, 0, 6, 60, 83, 69, 76, 70, 62, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 4, 109, 97, 105, 110, 28, 114, 101, 103, 105, 115, 116, 101, 114, 95, 99, 97, 110, 100, 105, 100, 97, 116, 101, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 6, 0, 8, 0, 11, 0, 11, 1, 11, 2, 11, 3, 11, 4, 11, 5, 18, 1, 1, 2}
)

var (
	programNameMap = map[string]string{
		string(PeerToPeerTransfer): "peer_to_peer_transfer",
		string(CreateAccount):      "create_account",
		string(Mint):               "mint",
		string(RotateAuthKey):      "rotate_authentication_key",
		string(RotateConsensusKey): "rotate_consensus_key",
		string(AddValidator):       "add_validator",
		string(RemoveValidator):    "remove_validator",
		string(RegisterValidator):  "register_validator",
	}
)

// InferProgramName infers a human-readable name from a transaction script byte code
func InferProgramName(program []byte) string {
	if name, ok := programNameMap[string(program)]; ok {
		return name
	}
	return "unknown"
}
