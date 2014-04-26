package dowser
import(
"github.com/shinpei/go-mecab"
)

type Index struct {
    inverseIndex map[string] int;
}

func (this *Index) IssueQuery(query string) []int {


  return make ([]int, 10);
}

func (this *Index) GetDocument (position int) *DowserResponse {
  return nil;
}
