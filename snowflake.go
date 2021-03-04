/* coding: utf-8
@Time :   2021/3/4 下午1:55
@Author : legend
@File :   snowflake.go
*/
package utils_go

import (
	"errors"
	"sync"
	"time"
)

/*
 * 算法解释
 * SnowFlake的结构如下(每部分用-分开):<br>
 * 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000 <br>
 * 1位标识，由于long基本类型在Java中是带符号的，最高位是符号位，正数是0，负数是1，所以id一般是正数，最高位是0<br>
 * 41位时间截(毫秒级)，注意，41位时间截不是存储当前时间的时间截，而是存储时间截的差值（当前时间截 - 开始时间截)
 * 得到的值），这里的的开始时间截，一般是我们的id生成器开始使用的时间，由我们程序来指定的（如下的epoch属性）。
 * 41位的时间截，可以使用69年，年T = (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69<br>
 * 10位的数据机器位，可以部署在1024个节点，包括5位datacenterId和5位workerId<br>
 * 12位序列，毫秒内的计数，12位的计数顺序号支持每个节点每毫秒(同一机器，同一时间截)产生4096个ID序号<br>
 * 加起来刚好64位，为一个Long型。<br>
 * SnowFlake的优点是，整体上按照时间自增排序，并且整个分布式系统内不会产生ID碰撞(由数据中心ID和机器ID作区分)，并且效率较高，经测试，SnowFlake每秒能够产生26万ID左右。
 */
const (
	nodeBits  uint8 = 10                    // 节点 ID 的位数
	stepBits  uint8 = 12                    // 序列号的位数
	nodeMax   int64 = -1 ^ (-1 << nodeBits) // 节点 ID 的最大值，用于检测溢出
	stepMax   int64 = -1 ^ (-1 << stepBits) // 序列号的最大值，用于检测溢出
	timeShift uint8 = nodeBits + stepBits   // 时间戳向左的偏移量
	nodeShift uint8 = stepBits              // 节点 ID 向左的偏移量
)

var Epoch int64 = 1288834974657

type ID int64
type Node struct {
	mu        sync.Mutex // 添加互斥锁，保证并发安全
	timestamp int64      // 时间戳部分
	node      int64      // 节点 ID 部分
	step      int64      // 序列号 ID 部分
}

func NewNode(node int64) (*Node, error) {
	// 如果超出节点的最大范围，产生一个 error
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	// 生成并返回节点实例的指针
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

func Generate(n *Node) ID {

	n.mu.Lock()         // 保证并发安全, 加锁
	defer n.mu.Unlock() // 方法运行完毕后解锁

	// 获取当前时间的时间戳 (毫秒数显示)
	now := time.Now().UnixNano() / 1e6

	if n.timestamp == now {
		// step 步进 1
		n.step++

		// 当前 step 用完
		if n.step > stepMax {
			// 等待本毫秒结束
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 本毫秒内 step 用完
		n.step = 0
	}

	n.timestamp = now
	// 移位运算，生产最终 ID
	result := ID((now-Epoch)<<timeShift | (n.node << nodeShift) | (n.step))

	return result
}
