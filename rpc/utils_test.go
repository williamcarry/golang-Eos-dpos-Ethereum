
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有2016 Go Ethereum作者
//此文件是Go以太坊库的一部分。
//
//Go-Ethereum库是免费软件：您可以重新分发它和/或修改
//根据GNU发布的较低通用公共许可证的条款
//自由软件基金会，或者许可证的第3版，或者
//（由您选择）任何更高版本。
//
//Go以太坊图书馆的发行目的是希望它会有用，
//但没有任何保证；甚至没有
//适销性或特定用途的适用性。见
//GNU较低的通用公共许可证，了解更多详细信息。
//
//你应该收到一份GNU较低级别的公共许可证副本
//以及Go以太坊图书馆。如果没有，请参见<http://www.gnu.org/licenses/>。

package rpc

import (
	"strings"
	"testing"
)

func TestNewID(t *testing.T) {
	hexchars := "0123456789ABCDEFabcdef"
	for i := 0; i < 100; i++ {
		id := string(NewID())
		if !strings.HasPrefix(id, "0x") {
			t.Fatalf("invalid ID prefix, want '0x...', got %s", id)
		}

		id = id[2:]
		if len(id) == 0 || len(id) > 32 {
			t.Fatalf("invalid ID length, want len(id) > 0 && len(id) <= 32), got %d", len(id))
		}

		for i := 0; i < len(id); i++ {
			if strings.IndexByte(hexchars, id[i]) == -1 {
				t.Fatalf("unexpected byte, want any valid hex char, got %c", id[i])
			}
		}
	}
}