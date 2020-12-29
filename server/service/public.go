package service

import "github.com/millim/timedata/lib/pool"

var p *pool.Pool

func getPool() *pool.Pool{
	if p == nil {
		p = pool.GetPool()
	}
	return p
}

