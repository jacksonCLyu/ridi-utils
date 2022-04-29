package convutil

import (
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func TestUtf82GBK(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		t.Errorf("TestUtf82GBK err: %v \n", err)
	})
	data := []byte("中文")
	t.Logf("isUtf8: %v", IsUtf8(data))
	gbk := assignutil.Assign(Utf82GBK(data))
	t.Logf("isGBK: %v", IsGBK(gbk))
	t.Log(gbk)
}
