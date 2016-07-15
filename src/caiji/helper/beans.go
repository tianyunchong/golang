package helper
/**
	队列的操作
 */
import("github.com/kr/beanstalk"
	"time"
	"strings"
	"fmt"
)

type BeansCon struct {
	conn *beanstalk.Conn
	connWrite *beanstalk.Conn
	beansName string
}

/**
	连接队列
 */
func (beansObj *BeansCon)Init(config [2]string, beansName string) {
	beansObj.beansName = beansName
	c, err := beanstalk.Dial(config[0], config[1])
	if err != nil {
		panic(err)
	}
	c.Tube.Name = beansName
	c.TubeSet.Name[beansName] = true
	beansObj.conn = c
}

/**
	读取队列,获取一条信息并返回
 */
func (beansObj *BeansCon)Read() string{
	//从队列中取出
	id, body, err := beansObj.conn.Reserve(1 * time.Second)
	if err != nil {
		if !strings.Contains(err.Error(), "timeout") {
			fmt.Println("timeout", " [Consumer] [", beansObj.conn.Tube.Name, "] err:", err, " id:", id)
		}
		return ""
	}
	//从队列中清理掉
	err = beansObj.conn.Delete(id)
	return string(body)

}

/**
	写入队列
 */
func (beansObj *BeansCon)Write(content string) int64 {
	if content == "" {
		return 0
	}
	c := beansObj.conn
	msg := fmt.Sprintf("%s", content)
	rs, err := c.Put([]byte(msg), 30, 0, 120*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)
	return 1
}

/**
	关闭队列
 */
func (beansObj *BeansCon)Close() {
	beansObj.conn.Close()
}