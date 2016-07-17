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
}

/**
	连接队列
 */
func (beansObj *BeansCon)Init(config [2]string) {
	c, err := beanstalk.Dial(config[0], config[1])
	if err != nil {
		panic(err)
	}
	beansObj.conn = c
}

/**
	读取队列,获取一条信息并返回
 */
func (beansObj *BeansCon)Read(beansName string) string{
	beansObj.conn.Tube.Name = beansName
	beansObj.conn.TubeSet.Name[beansName] = true
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
func (beansObj *BeansCon)Write(content string, beansName string) int64 {
	beansObj.conn.Tube.Name = beansName
	beansObj.conn.TubeSet.Name[beansName] = true
	if content == "" {
		return 0
	}
	c := beansObj.conn
	msg := fmt.Sprintf("%s", content)
	_, err := c.Put([]byte(msg), 30, 0, 120*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return 1
}

/**
	关闭队列
 */
func (beansObj *BeansCon)Close() {
	beansObj.conn.Close()
}