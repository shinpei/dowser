package dowser

type Tokenizer interface {
	Tokenize() ([]Token, error)
}

type TokenizerConfig struct {
}

type SpaceTokenizer struct {
	config TokenizerConfig
}

func NewSpaceTokenizer() *SpaceTokenizer {
	return &SpaceTokenizer{}
}

func (tok *SpaceTokenizer) Tokenize() (tokens []Token, err error) {
	return nil, nil
}
