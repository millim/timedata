package pool

//Pool 池子
type Pool struct {
	Worker chan func()
	Max    chan struct{}
}

var pool *Pool


//同时处理的数量
var defaultSize = 100

//GetPool 获取池子,将任务丢给队列处理
func GetPool() *Pool {
	if pool == nil {
		New(defaultSize)
	}
	return pool
}

//New 返回一个线程池子
func New(size int) {
	defaultSize = size
	pool = &Pool{
		Worker: make(chan func()),
		Max:    make(chan struct{}, size),
	}
}

func (p *Pool) AddWork(work func()) {
	select {
	case p.Worker <- work:
	case p.Max <- struct{}{}:
		go p.worker(work)
	}
}

func (p *Pool) worker(work func()) {
	defer func() { <-p.Max }()
	for {
		work()
		work = <-p.Worker
	}
}
