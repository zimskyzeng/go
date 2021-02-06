package main

import (
	"fmt"
	"time"
)

/*
	time包 主要用来处理时间相关
*/

func main() {
	t1 := time.Now()                                      // 获取当前时间的时间对象
	t2 := time.Date(2020, 3, 2, 11, 20, 0, 0, time.Local) // 获取指定时间

	// 1.格式化输出，这些数字都是固定的
	fmt.Println(t1.Format("2006-1-2 15:04:05"))
	fmt.Println(t2.Format("2006年1月2日 15时04分05秒"))

	// 2.格式化解析字符串时间=> time
	t3, err := time.Parse("2006-1-2 15:04:05", "2020-12-12 19:00:00")
	fmt.Println(t3, err)

	// 3.通过Unix时间戳获取time对象, 后面那个参数是纳秒
	t4 := time.Unix(1611995048, 0)
	fmt.Println(t4)

	// 4.获取纳秒时间戳
	fmt.Println(time.Now().Unix())

	// 5.获取纳秒时间戳
	fmt.Println(time.Now().UnixNano())

	// 6.根据时间对象获取指定内容
	//t1.Date()-->year,month,day
	//t1.Clock()-->hour,minute,second
	//t1.Year()-->year
	//t1.Month()-->Month
	//t1.Day()-->day
	//t1.Week()
	//t1.Hour()
	//t1.Minute()
	//t1.Second()
	//t1.NanoSecond()

	// 7.时间间隔 Duration 本质是一个int64数，最大表示范围290年
	// 常用的Duration time.Second time.Minute time.Hour
	t5 := t1.Add(10 * time.Second)
	fmt.Println(t5)

	// 8.休眠
	time.Sleep(1 * time.Second)

	// 9.计时器1
	// 原理分析：本质上就是过一段时间向返回的管道中入定时后的时间Time
	/*
		func NewTimer(d Duration) *Timer {
			c := make(chan Time, 1)
			t := &Timer{
				C: c,                 程序中获取C就是获取到Time管道
				r: runtimeTimer{
					when: when(d),
					f:    sendTime,
					arg:  c,          传入Time管道c，
				},
			}
			startTimer(&t.r)          这里实现了等待Duration时间后，向管道中写入等待后的当前时间对象
			return t                  返回Timer的地址
		}
	*/

	timer1 := time.NewTimer(3 * time.Second)
	fmt.Printf("类型：%T\n", timer1)
	t6 := <-timer1.C
	fmt.Println(t6)

	// 9.计时器2
	ch := time.After(3* time.Second)  // 返回的就是 NewTimer(d).C
	fmt.Println("time2", time.Now())
	timer2 := <- ch
	fmt.Println(timer2)
}
