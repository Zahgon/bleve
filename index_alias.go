package bleve

type IndexAlias interface {
	Index

	Add(i ...Index)
	Remove(i ...Index)
	Swap(in, out []Index)
}
