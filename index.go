package dowser
import(
"github.com/shinpei/go-mecab"
)

type Index struct {
    inverseIndex map[string] int;
}

func (this *Index) IssueQuery(query string) []int {
  tagger := mecab.Create();
  headerNode := tagger.ParseToNode(query);
  var currentNode *mecab.Node = headerNode;

  // initialize map
  if (this.inverseIndex == nil) {
    this.inverseIndex = make(map[string]int);
  }


  for i := 0; currentNode.HasNext(); i++ {
    currentNode.Next();
    this.inverseIndex[currentNode.GetSurface()] = 1;
  }
  return make([]int, 10);
}

func (this *Index) GetDocument (position int) *DowserResponse {
  return nil;
}
