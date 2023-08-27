package internal

import (
	"fmt"
	"time"
)

func NotGoRoutineSample() {
	printTitle("NotGoRoutineSample()")

	fmt.Println("not goroutine")
	process(2, "A")
	process(2, "B")
	fmt.Println("finish")
}

func GoRoutineSample() {
	printTitle("GoRoutineSample()")

	fmt.Println("goroutine")

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		process(2, "A")
		ch1 <- true
	}()

	go func() {
		process(2, "B")
		ch2 <- true
	}()

	<-ch1
	<-ch2

	fmt.Println("finish")
}

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}

func ChannelSample() {
	printTitle("ChannelSample()")

	//channelの作成
	messages := make(chan string)

	//作成されたchannelに値(str)を送信
	go func() { messages <- "str" }()

	//channelから値を受信
	msg := <-messages
	fmt.Println(msg) //=> str
}

// https://qiita.com/gold-kou/items/8e5342d8a30ae8f34dff
func BufferSample() {
	printTitle("BufferSample()")

	ch := make(chan int, 5) // バッファ付きchannel宣言

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	// 念のためchにデータが書き込まれるのを待つ
	time.Sleep(time.Second)

	// 1〜5を読み込んで出力
	for i := 1; i <= 5; i++ {
		tmp := <-ch
		fmt.Println(tmp)
	}

	// 6〜10がchに書き込まれるのを待つ
	fmt.Println("waiting")
	time.Sleep(time.Second)

	// 6〜10を読み込んで出力
	for i := 1; i <= 5; i++ {
		tmp := <-ch
		fmt.Println(tmp)
	}
}
