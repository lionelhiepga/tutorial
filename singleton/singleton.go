package singleton


type Singleton interface{
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

//HÀM INIT: khi khơi tạo thì đối tượng intance được tạo ra luôn cho chúng ra không chờ cta gọi nữa

func init() {
	instance = &singleton{count : 100}
}

func GetInstance() Singleton {

	//chỉ muốn làm đứng 1 lần thôi
	// once.Do(func() {
	// 	time.Sleep(time.Second)
    //     instance = &singleton{count: 100}  
    // })
    return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}