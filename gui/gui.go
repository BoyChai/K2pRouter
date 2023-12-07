package gui

import (
	"K2pRouter/control"
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/chromedp/chromedp"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"net"
	"os"
	"sync"
)

func RunGui() {
	// app
	gui := app.New()
	window := gui.NewWindow("K2p一键设置")
	// 整体布局
	grid := layout.NewGridLayout(3)

	// 获取输入框
	info, w, w2, w3, w4 := GetInfo()

	// 获取信息
	window.SetContent(container.New(grid, GetText(), container.NewVBox(info, w, w2, w3, w4), Btn(info, w, w2, w3, w4)))

	window.ShowAndRun()
}

func init() {
	fontPath, err := findfont.Find("simfang.ttf")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Found 'arial.ttf' in '%s'\n", fontPath)

	// load the font with the freetype library
	// 原作者使用的ioutil.ReadFile已经弃用
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}

func GetText() *fyne.Container {
	adminPas := widget.NewLabel("登录管理员密码")
	Wifi2GSsid := widget.NewLabel("2GWIFI名字")
	Wifi2GPas := widget.NewLabel("2GWIFI密码")
	Wifi5GSsid := widget.NewLabel("5GWIFI名字")
	Wifi5GPas := widget.NewLabel("5GWIFI密码")
	return container.NewVBox(adminPas, Wifi2GSsid, Wifi2GPas, Wifi5GSsid, Wifi5GPas)
}

func GetInfo() (*widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry) {
	adminPas := widget.NewEntry()
	Wifi2GSsid := widget.NewEntry()
	Wifi2GPas := widget.NewEntry()
	Wifi5GSsid := widget.NewEntry()
	Wifi5GPas := widget.NewEntry()
	adminPas.Text = "12345678"
	Wifi2GSsid.Text = "PHICN"
	Wifi2GPas.Text = "12345678"
	Wifi5GSsid.Text = "PHICN_5G"
	Wifi5GPas.Text = "12345678"
	return adminPas, Wifi2GSsid, Wifi2GPas, Wifi5GSsid, Wifi5GPas
}

// Btn 设置按钮
// 老版本
func Btn(adminPas, Wifi2GSsid, Wifi2GPas, Wifi5GSsid, Wifi5GPas *widget.Entry) *fyne.Container {
	text := widget.NewLabel("K2p路由器一键设置DEBUG\n")
	out := container.New(layout.NewGridLayoutWithRows(1), text)
	//ip := "192.168.2.1"
	var wg sync.WaitGroup

	bt0 := widget.NewButton("扫描网段", func() {
		control.IpPool = []string{}
		_, ipNet, err := net.ParseCIDR(control.TargetCIDR)
		if err != nil {
			text.SetText("解析目标网段失败\n")
			return
		}
		ipChan := make(chan string, 100)
		var wg sync.WaitGroup

		for i := 0; i < cap(ipChan); i++ {
			go control.Detection(ipChan, &wg)
		}

		// 遍历ip
		for addr := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(addr); control.IncIP(addr) {
			wg.Add(1)
			ipChan <- addr.String()
		}
		wg.Wait()
		close(ipChan)
		text.SetText(fmt.Sprintf("总共解析目标数量:%v\n", len(control.IpPool)))
	})
	// 一键设置老
	bt1 := widget.NewButton("一键设置(老)", func() {
		text.SetText("检测版本...")
		pass := adminPas.Text
		name2g := Wifi2GSsid.Text
		pass2g := Wifi2GPas.Text
		name5g := Wifi5GSsid.Text
		pass5g := Wifi5GPas.Text
		//fmt.Println(1)
		if len(pass) <= 5 {

			//fmt.Println(2)
			text.SetText("管理员密码长度请设置5位以上\n")

		} else if len(pass2g) < 8 {

			//fmt.Println(3)
			//text.Text += "WIFI2G密码长度请设置6位以上\n"
			text.SetText("WIFI2G密码长度请设置7位以上\n")

		} else if len(pass5g) < 8 {
			//fmt.Println(4)
			//text.Text += "WIFI5G密码长度请设置6位以上\n"
			text.SetText("WIFI5G密码长度请设置7位以上\n")
		}

		//fmt.Println(5)
		for _, v := range control.IpPool {
			wg.Add(1)
			go control.SetRouter1(v, pass, name2g, pass2g, name5g, pass5g, text, &wg)
		}
		wg.Wait()
	})
	bt2 := widget.NewButton("一键设置(新)", func() {
		text.SetText("检测版本...")
		pass := adminPas.Text
		name2g := Wifi2GSsid.Text
		pass2g := Wifi2GPas.Text
		name5g := Wifi5GSsid.Text
		pass5g := Wifi5GPas.Text
		//fmt.Println(1)
		if len(pass) <= 5 {

			//fmt.Println(2)
			text.SetText("管理员密码长度请设置5位以上\n")

		} else if len(pass2g) < 8 {

			//fmt.Println(3)
			//text.Text += "WIFI2G密码长度请设置6位以上\n"
			text.SetText("WIFI2G密码长度请设置7位以上\n")

		} else if len(pass5g) < 8 {
			//fmt.Println(4)
			//text.Text += "WIFI5G密码长度请设置6位以上\n"
			text.SetText("WIFI5G密码长度请设置7位以上\n")
		}

		//fmt.Println(5)
		//control.SetRouter2(ip, pass, name2g, pass2g, name5g, pass5g, text)
		for _, v := range control.IpPool {
			wg.Add(1)
			go control.SetRouter1(v, pass, name2g, pass2g, name5g, pass5g, text, &wg)
		}
		wg.Wait()
	})
	bt3 := widget.NewButton("一键设置(自动)", func() {
		text.SetText("检测版本...")
		pass := adminPas.Text
		name2g := Wifi2GSsid.Text
		pass2g := Wifi2GPas.Text
		name5g := Wifi5GSsid.Text
		pass5g := Wifi5GPas.Text
		//fmt.Println(1)
		if len(pass) <= 5 {

			//fmt.Println(2)
			text.SetText("管理员密码长度请设置5位以上\n")

		} else if len(pass2g) < 8 {

			//fmt.Println(3)
			//text.Text += "WIFI2G密码长度请设置6位以上\n"
			text.SetText("WIFI2G密码长度请设置7位以上\n")

		} else if len(pass5g) < 8 {
			//fmt.Println(4)
			//text.Text += "WIFI5G密码长度请设置6位以上\n"
			text.SetText("WIFI5G密码长度请设置7位以上\n")
		}

		//fmt.Println(5)
		//if currentURL == "http://192.168.2.1/cgi-bin#/setLgPwd" {
		//	text.SetText("确定为新版本\n")
		//	control.SetRouter2(ip, pass, name2g, pass2g, name5g, pass5g, text)
		//} else {
		//	text.SetText("确定为老版本\n")
		//	control.SetRouter1(ip, pass, name2g, pass2g, name5g, pass5g, text)
		//}
		for _, v := range control.IpPool {
			ctx, _ := chromedp.NewContext(context.Background())
			var currentURL string

			err := chromedp.Run(ctx, chromedp.Navigate("http://"+v+"/"), chromedp.WaitReady("body"), chromedp.Location(&currentURL))
			if err != nil {
				text.SetText("请确定路由器已经连接")
				return
			}
			if currentURL == "http://"+v+"/cgi-bin#/setLgPwd" {
				//text.SetText("确定为新版本\n")
				wg.Add(1)
				go control.SetRouter2(v, pass, name2g, pass2g, name5g, pass5g, text, &wg)
			} else {
				//text.SetText("确定为老版本\n")
				wg.Add(1)
				control.SetRouter1(v, pass, name2g, pass2g, name5g, pass5g, text, &wg)
			}

		}
		wg.Wait()
	})

	return container.NewVBox(out, bt0, bt1, bt2, bt3)
}
