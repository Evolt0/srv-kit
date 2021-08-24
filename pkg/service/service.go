package service

type BookInter interface {
	GetBookInfoByID(int32) string
}
type Book struct {
}

func (b *Book) GetBookInfoByID(id int32) string {
	switch id {
	case 1:
		return "西游记"
	case 2:
		return "三国演义"
	case 3:
		return "水浒传"
	case 4:
		return "红楼梦"
	case 5:
		return "朝花夕拾"
	}
	return "未知id"
}

type Hello interface {
	Hello() string
}

type HelloImpl struct {
}

func (h *HelloImpl) Hello() string {
	return "hello world ettetetete"
}
