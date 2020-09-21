package factory

type Type uint8

// 事件类型定义
const (
	Start Type = iota
	End
)

// 事件抽象接口
type Event interface {
	EventType() Type
	Content() string
}

// 开始事件，实现了 Event接口
type StartEvent struct {
	content string
}

func (se StartEvent) EventType() Type {
	return 0
}
func (se StartEvent) Content() string {
	return se.content
}

// 结束事件，实现了 Event接口
type EndEvent struct {
	content string
}

func (ee EndEvent) EventType() Type {
	return 1
}
func (ee EndEvent) Content() string {
	return ee.content
}

type Factory struct{}

// 根据 事件类型 创建 具体事件
func (e *Factory) Create(etype Type) Event {
	switch etype {
	case Start:
		return &StartEvent{
			content: "this is start event",
		}
	case End:
		return &EndEvent{
			content: "this is end event",
		}
	default:
		return nil
	}
}

// 按照第二种实现方式，分别给Start和End类型的Event单独提供一个工厂方法
func OfStart() Event {
	return &StartEvent{
		content: "this is start event",
	}
}
func OfEnd() Event {
	return &EndEvent{
		content: "this is end event",
	}
}
