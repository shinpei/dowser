package dowser
import(
"github.com/shinpei/go-mecab"
"fmt"
)

type Index struct {
    inverseIndex map[string] int;
}

func (this *Index) IssueQuery(query string) []int {
  tagger := mecab.Create();
  result := tagger.Parse(query)
	fmt.Println(result)
  return make([]int, 10);
}

func (this *Index) GetDocument (position int) *DowserResponse {
  return nil;
}
