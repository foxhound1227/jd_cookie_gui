package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type JDCookieExtractor struct {
	myApp       fyne.App
	myWindow    fyne.Window
	driver      selenium.WebDriver
	statusLabel *widget.Label
	progressBar *widget.ProgressBar
	cookieText  *widget.Entry
	getButton   *widget.Button
}

func main() {
	myApp := app.New()
	myApp.SetIcon(nil)

	myWindow := myApp.NewWindow("京东Cookie提取工具")
	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.CenterOnScreen()

	extractor := &JDCookieExtractor{
		myApp:    myApp,
		myWindow: myWindow,
	}

	extractor.setupUI()
	extractor.initBrowser()

	myWindow.ShowAndRun()
}

func (j *JDCookieExtractor) setupUI() {
	// 状态标签
	j.statusLabel = widget.NewLabel("正在初始化...")

	// 进度条
	j.progressBar = widget.NewProgressBar()
	j.progressBar.SetValue(0)

	// 获取Cookie按钮
	j.getButton = widget.NewButton("获取Cookie", j.getCookies)
	j.getButton.Disable()

	// Cookie显示区域
	j.cookieText = widget.NewMultiLineEntry()
	j.cookieText.SetPlaceHolder("Cookie将显示在这里...")
	j.cookieText.Resize(fyne.NewSize(480, 200))

	// 布局
	content := container.NewVBox(
		j.statusLabel,
		j.progressBar,
		j.getButton,
		j.cookieText,
	)

	j.myWindow.SetContent(content)

	// 设置窗口关闭事件
	j.myWindow.SetCloseIntercept(func() {
		j.cleanup()
		j.myWindow.Close()
	})
}

func (j *JDCookieExtractor) checkEdgeInstalled() bool {
	var edgePaths []string

	if runtime.GOOS == "windows" {
		programFiles := os.Getenv("PROGRAMFILES")
		programFilesX86 := os.Getenv("PROGRAMFILES(X86)")

		if programFiles != "" {
			edgePaths = append(edgePaths, filepath.Join(programFiles, "Microsoft", "Edge", "Application", "msedge.exe"))
		}
		if programFilesX86 != "" {
			edgePaths = append(edgePaths, filepath.Join(programFilesX86, "Microsoft", "Edge", "Application", "msedge.exe"))
		}
	}

	for _, path := range edgePaths {
		log.Printf("检查Edge浏览器路径: %s", path)
		if _, err := os.Stat(path); err == nil {
			log.Printf("找到Edge浏览器: %s", path)
			return true
		}
	}

	log.Println("未找到Edge浏览器")
	return false
}

func (j *JDCookieExtractor) getEdgeDriverPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("获取可执行文件路径失败: %v", err)
	}

	baseDir := filepath.Dir(execPath)
	driverPath := filepath.Join(baseDir, "msedgedriver.exe")

	if _, err := os.Stat(driverPath); os.IsNotExist(err) {
		return "", fmt.Errorf("EdgeDriver文件不存在: %s\n请手动下载EdgeDriver并放置在程序目录中", driverPath)
	}

	log.Printf("使用EdgeDriver: %s", driverPath)
	return driverPath, nil
}

func (j *JDCookieExtractor) initBrowser() {
	go func() {
		// 更新UI状态
		j.statusLabel.SetText("检查Edge浏览器...")
		j.progressBar.SetValue(0.1)

		// 检查Edge浏览器
		if !j.checkEdgeInstalled() {
			dialog.ShowError(
				fmt.Errorf("未检测到Edge浏览器，请先安装Edge浏览器后再运行程序\n\n下载地址：https://www.microsoft.com/edge"),
				j.myWindow,
			)
			return
		}

		j.statusLabel.SetText("初始化浏览器...")
		j.progressBar.SetValue(0.5)

		// 获取EdgeDriver路径
		driverPath, err := j.getEdgeDriverPath()
		if err != nil {
			dialog.ShowError(err, j.myWindow)
			return
		}

		// 启动Selenium服务
		selenium.SetDebug(false)
		service, err := selenium.NewChromeDriverService(driverPath, 9515)
		if err != nil {
			dialog.ShowError(fmt.Errorf("启动EdgeDriver服务失败: %v", err), j.myWindow)
			return
		}
		defer service.Stop()

		// 配置浏览器选项
		caps := selenium.Capabilities{"browserName": "chrome"}
		chromeCaps := chrome.Capabilities{
			Args: []string{
				"--disable-infobars",
				"--disable-extensions",
				"--disable-gpu",
				"--no-sandbox",
				"--disable-dev-shm-usage",
				"--start-maximized",
				"--disable-blink-features=AutomationControlled",
				"--disable-features=msEdgeTranslate",
			},
			ExcludeSwitches: []string{"enable-automation"},
		}
		caps.AddChrome(chromeCaps)

		// 创建WebDriver
		j.driver, err = selenium.NewRemote(caps, "http://localhost:9515")
		if err != nil {
			dialog.ShowError(fmt.Errorf("创建WebDriver失败: %v", err), j.myWindow)
			return
		}

		// 设置超时
		j.driver.SetPageLoadTimeout(10 * time.Second)
		j.driver.SetImplicitWaitTimeout(5 * time.Second)

		// 加载京东登录页面
		err = j.driver.Get("https://plogin.m.jd.com/login/login?appid=300&returnurl=https%3A%2F%2Fm.jd.com%2F&source=wq_passport")
		if err != nil {
			log.Printf("加载京东页面失败: %v", err)
			// 即使加载失败也继续运行
		}

		// 完成初始化
		j.statusLabel.SetText("准备就绪")
		j.progressBar.SetValue(1.0)
		j.getButton.Enable()

		log.Println("浏览器初始化完成")
	}()
}

func (j *JDCookieExtractor) getCookies() {
	if j.driver == nil {
		dialog.ShowError(fmt.Errorf("浏览器未初始化"), j.myWindow)
		return
	}

	j.statusLabel.SetText("正在获取Cookie...")

	cookies, err := j.driver.GetCookies()
	if err != nil {
		dialog.ShowError(fmt.Errorf("获取Cookie失败: %v", err), j.myWindow)
		j.statusLabel.SetText("Cookie获取失败")
		return
	}

	var ptPin, ptKey string
	for _, cookie := range cookies {
		if cookie.Name == "pt_pin" {
			ptPin = cookie.Value
		} else if cookie.Name == "pt_key" {
			ptKey = cookie.Value
		}
	}

	if ptPin != "" && ptKey != "" {
		cookieStr := fmt.Sprintf("pt_pin=%s;pt_key=%s;", ptPin, ptKey)
		j.cookieText.SetText(cookieStr)
		j.statusLabel.SetText("Cookie获取成功")
	} else {
		j.cookieText.SetText("未找到所需的Cookie，请确保已登录京东")
		j.statusLabel.SetText("Cookie获取失败")
	}
}

func (j *JDCookieExtractor) cleanup() {
	if j.driver != nil {
		j.driver.Quit()
	}
}