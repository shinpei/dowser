package dowser

import (
	"log"
	"os/user"
)

type Dowser struct {
	docRoot string;

}

func Create () *Dowser{
		return &Dowser { docRoot: ""};
}

func (this *Dowser) Initiate(docRoot string) {
	if docRoot == "" {
		user, err := user.Current();
		if err != nil {
			log.Fatal(err);
		}
		docRoot = user.HomeDir + "/.dowsing";
		log.Printf("set default docRoot: %s\n", docRoot);
	}

}

func (this *Dowser) Encode2Document (data string) *Document {
	return &Document{ bodyField: data};
}

func (this *Dowser) UpdateIndex(doc *Document) {


}

func (this *Dowser) Search (query string) *DowserResponse {
	return &DowserResponse{"hoge"};
}
