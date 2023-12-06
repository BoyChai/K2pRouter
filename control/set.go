package control

import (
	"context"
	"fyne.io/fyne/v2/widget"
	"github.com/chromedp/chromedp"
	"log"
)

func SetRouter(ip, adminPass, name2g, pass2g, name5g, pass5g string, text *widget.Label) {
	//var adminPass = "123456"
	//var wifiName = "testWifiName"
	//var wifiPass = "Qwer12345"

	// 创建一个带有10秒超时的上下文
	//ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	// 使用 chromedp.NewContext 创建一个带有超时的上下文
	// 1
	//ctx, _ := chromedp.NewExecAllocator(
	//	context.Background(),
	//	//chromedp.Flag("headless", false),
	//)
	// 2
	ctx, _ := chromedp.NewContext(context.Background())

	//var htmlContent string
	text.SetText("正在设置路由器管理员账户...\n")
	// 运行任务
	chromedp.Run(ctx,
		chromedp.Navigate("http://"+ip+"/#/guide"), // 修改为目标网站的登录页面URL
		chromedp.WaitVisible("body"),
		chromedp.Click("#Start"), // 修改为登录按钮的选择器
		chromedp.Navigate("http://"+ip+"/#/setLgPwd"),
		chromedp.WaitVisible("body"),
		chromedp.SendKeys("#PwdNew", adminPass),
		chromedp.SendKeys("#PwdCfm", adminPass),
		chromedp.Click("#Save"),
		//chromedp.Sleep(1*time.Second),
	)
	text.SetText("正在设置路由器WIFI账号密码...\n")
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://"+ip+"/#/guideWifiSet"),
		chromedp.WaitVisible("body"),
		//chromedp.Sleep(1*time.Second),
		//chromedp.Navigate("http://"+ip+"/#/guideWifiSet"),
		//chromedp.Sleep(1*time.Second),
		chromedp.Clear("#Ssid2G"),
		chromedp.SetValue("#Ssid2G", name2g),
		chromedp.SetValue("#Pwd2G", pass2g),
		chromedp.Clear("#Ssid5G"),
		chromedp.SetValue("#Ssid5G", name5g),
		chromedp.SetValue("#Pwd5G", pass5g),
		chromedp.Click("#SaveReboot"),
		//chromedp.OuterHTML("#WifiWait > div.wifi-tip > p:nth-child(1)", &htmlContent),
	)

	if err != nil {
		log.Fatal(err)
		return
		//text.Text += "路由器设置失败\n"
		text.SetText("路由器设置失败\n")
	}
	//text.Text += "路由器设置成功\n"

	text.SetText("路由器设置成功\n")
}
