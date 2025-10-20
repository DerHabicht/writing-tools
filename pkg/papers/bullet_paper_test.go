package papers_test

import (
	"testing"

	ppapers "github.com/derhabicht/writing-tools/pkg/papers"
)

func TestBulletPaperImplementsInterfaces(t *testing.T) {
	var _ ppapers.Paper = (*ppapers.BulletPaper)(nil)
}
