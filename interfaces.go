package dowser

type Tokenizer interface {
	Tokenize(target string) ([]Token, error)
}

type TokenizerConfig struct {
}

type SpaceTokenizer struct {
	config TokenizerConfig
}

func NewSpaceTokenizer() *SpaceTokenizer {
	return &SpaceTokenizer{}
}

func (tok *SpaceTokenizer) Tokenize(target string) (tokens []Token, err error) {
	return nil, nil
}
