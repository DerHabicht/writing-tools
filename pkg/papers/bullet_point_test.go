package papers_test

import (
	"testing"

	ppapers "github.com/derhabicht/writing-tools/pkg/papers"
)

func TestBPointsImplementInterfaces(t *testing.T) {
	var _ ppapers.PointList = (*ppapers.BulletPointList)(nil)
	var _ ppapers.Point = (*ppapers.BulletPoint)(nil)
}