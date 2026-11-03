package main

import "testing"

func BenchmarkLoadBalancer(b *testing.B) {
	const connectionsCount = 100
	const triesCount = 1000

	conns := make([]*Connection, connectionsCount)
	for i := 0; i < connectionsCount; i++ {
		conns = append(conns, &Connection{})
	}

	loadBalancerChannel := NewLoadBalancerChan(conns)
	loadBalancerChannel.Init()
	defer loadBalancerChannel.Close()

	loadBalancerAtomic := NewLoadBalancerAtomic(conns)

	loadBalancerMutex := NewLoadBalancerMutex(conns)

	b.ResetTimer()

	b.Run("channel", func(b *testing.B) {
		for i := 0; i < triesCount; i++ {
			loadBalancerChannel.NextConn()
		}
	})
	b.Run("atomic", func(b *testing.B) {
		for i := 0; i < triesCount; i++ {
			loadBalancerAtomic.NextConn()
		}
	})
	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < triesCount; i++ {
			loadBalancerMutex.NextConn()
		}
	})
}
