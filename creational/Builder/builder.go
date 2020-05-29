package Builder

type Builder interface {
	Build()
}

type Directer struct {
	builder Builder
}

func NewBuilder(b Builder) Directer {
	return Directer{builder:b}
}

func (d *Directer) Construct()  {
	d.builder.Build()
}
func NewDirector(b Builder) Directer {
	return Directer{builder:b}
}

//ConcreateBuilder实现了Build方法，是Builder的子类，
//NewDirector(b Builder)也即ConcreateBuilder可以传入给Builder
type ConcreateBuilder struct {
	built bool
}

func NewConcreateBuilder() ConcreateBuilder {
	return ConcreateBuilder{false}
}

func (b *ConcreateBuilder) Build()  {
	b.built = true
}

type Product struct {
	Built bool
}

func (b *ConcreateBuilder) GetResult() Product {
	return Product{b.built}
}