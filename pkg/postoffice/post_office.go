package postoffice

var Singleton = &PostOffice{}

type PostOffice struct {
	isWorker    bool
	isServer    bool
	isScheduler bool
}

func (p *PostOffice) IsScheduler() bool {
	return p.isScheduler
}

func (p *PostOffice) IsWorker() bool {
	return p.isWorker
}
