# 适配器模式
适配器模式（Adapter Pattern）是作为两个不兼容的接口之间的桥梁。这种类型的设计模式属于结构型模式，它结合了两个独立接口的功能。

这种模式涉及到一个单一的类，该类负责加入独立的或不兼容的接口功能。举个真实的例子，读卡器是作为内存卡和笔记本之间的适配器。您将内存卡插入读卡器，再将读卡器插入笔记本，这样就可以通过笔记本来读取内存卡。
# 示例
## 场景描述
我们有一个 MediaPlayer 接口和一个实现了 MediaPlayer 接口的实体类 AudioPlayer。默认情况下，AudioPlayer 可以播放 mp3 格式的音频文件。

我们还有另一个接口 AdvancedMediaPlayer 和实现了 AdvancedMediaPlayer 接口的实体类。该类可以播放 vlc 和 mp4 格式的文件。

我们想要让 AudioPlayer 播放其他格式的音频文件。为了实现这个功能，我们需要创建一个实现了 MediaPlayer 接口的适配器类 MediaAdapter，并使用 AdvancedMediaPlayer 对象来播放所需的格式。

AudioPlayer 使用适配器类 MediaAdapter 传递所需的音频类型，不需要知道能播放所需格式音频的实际类。AdapterPatternDemo 类使用 AudioPlayer 类来播放各种格式。
## 代码展示
* MediaPlayer接口
``` go
package adapter

type MediaPlayer interface {
	play(string)
}
```
* MediaPlayer接口的实现AudioPlayer
``` go
package adapter

import "fmt"

type AudioPlayer struct {
}

func (a *AudioPlayer) play(fileName string) {
	fmt.Println("audio player：" + fileName)
}
```
* AdvancedMediaPlayer接口
``` go
package adapter

type AdvancedMediaPlayer interface {
	playVlc(string)
	playMp4(string)
}
```
* AdvancedMediaPlayer接口的实现Mp4Player
``` go
package adapter

import "fmt"

type Mp4Player struct {
}

func (*Mp4Player) playVlc(fileName string) {
	panic("unsupported file format: " + fileName)
}
func (*Mp4Player) playMp4(fileName string) {
	fmt.Println("Play mp4 file: " + fileName)
}
```

* AdvancedMediaPlayer接口的实现VlcPlayer
``` go
package adapter

import "fmt"

type VlcPlayer struct {
}

func (*VlcPlayer) playVlc(fileName string) {
	fmt.Println("Play vlc file: " + fileName)
}
func (*VlcPlayer) playMp4(fileName string) {
	panic("unsupported file format")
}
```

* MediaPlayer接口的Mp4适配器实现
``` go
package adapter

type Mp4MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m *Mp4MediaAdapter) play(fileName string) {
	m.advancedMediaPlayer.playMp4(fileName)
}
```
* MediaPlayer接口的Vlc适配器实现
``` go
package adapter

type VlcMediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m *VlcMediaAdapter) play(fileName string) {
	m.advancedMediaPlayer.playVlc(fileName)
}

```
* 测试类
``` go
package adapter

import "testing"

func TestAdapter(t *testing.T) {
	// user audio play
	var mediaPlayer MediaPlayer
	mediaPlayer = &AudioPlayer{}
	mediaPlayer.play("map3")

	// user VlcMediaAdapter play
	mediaPlayer = &VlcMediaAdapter{
		advancedMediaPlayer: &VlcPlayer{},
	}
	mediaPlayer.play("vlc")

	// user Mp4MediaAdapter play
	mediaPlayer = &Mp4MediaAdapter{
		advancedMediaPlayer: &Mp4Player{},
	}
	mediaPlayer.play("mp4")
}
```
* 执行命令
```shell
go test -v ./
```

* 输出结果
```
=== RUN   TestAdapter
audio player：map3
Play vlc file: vlc
Play mp4 file: mp4
--- PASS: TestAdapter (0.00s)
PASS
ok      go-design-patterns/structural/adapter   0.169s
```
## 类图
![类图](https://caixunshi.github.io/document/go-design-patterns/adapter.jpg)

## 优点
* 可以让任何两个没有关联的类一起运行。
* 提高了类的复用。 
* 增加了类的透明度。 
* 灵活性好。

## 缺点
* 过多地使用适配器，会让系统非常零乱，不易整体进行把握。比如，明明看到调用的是 A 接口，其实内部被适配成了 B 接口的实现，一个系统如果太多出现这种情况，无异于一场灾难。因此如果不是很有必要，可以不使用适配器，而是直接对系统进行重构。

## 使用场景
有动机地修改一个正常运行的系统的接口，这时应该考虑使用适配器模式。
## 注意事项
适配器不是在详细设计时添加的，而是解决正在服役的项目的问题。