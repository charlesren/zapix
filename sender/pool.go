package sender

import (
	"net"
	"sync"
	"time"

	"io"
)

// ConnectionPool 管理可复用的 TCP 连接
type ConnectionPool struct {
	pool     []net.Conn
	mutex    sync.Mutex
	poolSize int
}

// NewConnectionPool 初始化连接池
func NewConnectionPool(poolSize int) *ConnectionPool {
	return &ConnectionPool{
		pool:     make([]net.Conn, 0, poolSize),
		poolSize: poolSize,
	}
}

// Get 从池中获取一个连接，如果没有则新建
func (cp *ConnectionPool) Get(addr string, timeout time.Duration) (net.Conn, error) {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	for i, conn := range cp.pool {
		if conn != nil {
			// 检查连接是否有效（如发送空包或设置短暂超时）
			if err := checkConnAlive(conn); err == nil {
				cp.pool = append(cp.pool[:i], cp.pool[i+1:]...)
				return conn, nil
			}
			// 无效连接直接关闭
			conn.Close()
			cp.pool[i] = nil // 标记为 nil，避免重复关闭
		}
	}

	return net.DialTimeout("tcp", addr, timeout)
}

// checkConnAlive 检查连接是否有效
func checkConnAlive(conn net.Conn) error {
	conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	_, err := conn.Read(make([]byte, 1))
	if err == nil || err == io.EOF {
		return nil
	}
	return err
}

// Put 将连接归还到池中（如果池未满）或直接关闭
func (cp *ConnectionPool) Put(conn net.Conn) {
	if conn == nil {
		return // 忽略 nil 连接
	}

	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	if len(cp.pool) < cp.poolSize {
		cp.pool = append(cp.pool, conn)
	} else {
		conn.Close()
	}
}

// GetStats 获取连接池统计信息
func (cp *ConnectionPool) GetStats() map[string]interface{} {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	activeConnections := 0
	for _, conn := range cp.pool {
		if conn != nil {
			activeConnections++
		}
	}

	return map[string]interface{}{
		"pool_size":           cp.poolSize,
		"active_connections":  activeConnections,
		"available_capacity":  cp.poolSize - len(cp.pool),
		"current_pool_length": len(cp.pool),
	}
}

// GetActiveConnectionCount 获取当前活跃连接数量
func (cp *ConnectionPool) GetActiveConnectionCount() int {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	count := 0
	for _, conn := range cp.pool {
		if conn != nil {
			count++
		}
	}
	return count
}

// GetPoolSize 获取连接池大小
func (cp *ConnectionPool) GetPoolSize() int {
	return cp.poolSize
}
