package dowser

type Document struct {
    bodyField string;
}

func (this Document) index() Index {
    return Index{};
}
