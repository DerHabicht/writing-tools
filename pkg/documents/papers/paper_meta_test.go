package papers

import (
	"testing"

	"github.com/derhabicht/writing-tools/pkg/documents"
)

func TestPaperMetaImplementsInterfaces(t *testing.T) {
	var _ documents.DocMeta = (PaperMeta)(PaperMeta{})
}
