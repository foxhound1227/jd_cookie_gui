package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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

func (j *JDCookieExtractor) getEdgeVersion() (string, error) {
	cmd := exec.Command("reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Edge\\BLBeacon", "/v", "version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("无法获取Edge版本: %v", err)
	}

	re := regexp.MustCompile(`version\s+REG_SZ\s+([0-9.]+)`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return "", fmt.Errorf("无法解析Edge版本信息")
	}

	return matches[1], nil
}

func (j *JDCookieExtractor) downloadEdgeDriver(version, driverPath string) error {
	// 更新UI状态
	j.statusLabel.SetText("正在下载EdgeDriver...")
	j.progressBar.SetValue(0.4)
	
	// 提取主版本号
	versionParts := strings.Split(version, ".")
	if len(versionParts) < 3 {
		return fmt.Errorf("版本格式无效: %s", version)
	}
	majorVersion := strings.Join(versionParts[:3], ".")

	// 构建下载URL
	downloadURL := fmt.Sprintf("https://msedgedriver.azureedge.net/%s/edgedriver_win64.zip", majorVersion)
	log.Printf("正在下载EdgeDriver: %s", downloadURL)

	// 下载文件
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("下载失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("下载失败，状态码: %d", resp.StatusCode)
	}

	// 创建临时文件
	tempFile := driverPath + ".zip"
	out, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("创建临时文件失败: %v", err)
	}
	defer out.Close()
	defer os.Remove(tempFile)

	// 保存文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}
	out.Close()

	// 解压文件（简单实现，假设zip中只有一个exe文件）
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Expand-Archive -Path '%s' -DestinationPath '%s' -Force", tempFile, filepath.Dir(driverPath)))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("解压失败: %v", err)
	}

	// 重命名解压出的文件
	extractedPath := filepath.Join(filepath.Dir(driverPath), "msedgedriver.exe")
	if _, err := os.Stat(extractedPath); err == nil {
		return nil // 文件已存在
	}

	// 查找解压出的driver文件
	files, err := filepath.Glob(filepath.Join(filepath.Dir(driverPath), "*driver*.exe"))
	if err != nil || len(files) == 0 {
		return fmt.Errorf("未找到解压的driver文件")
	}

	// 重命名为msedgedriver.exe
	err = os.Rename(files[0], driverPath)
	if err != nil {
		return fmt.Errorf("重命名driver文件失败: %v", err)
	}

	return nil
}

func (j *JDCookieExtractor) getEdgeDriverPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("获取可执行文件路径失败: %v", err)
	}

	baseDir := filepath.Dir(execPath)
	driverPath := filepath.Join(baseDir, "msedgedriver.exe")

	// 检查EdgeDriver是否存在
	if _, err := os.Stat(driverPath); os.IsNotExist(err) {
		log.Println("EdgeDriver不存在，尝试自动下载...")
		
		// 获取Edge版本
		version, err := j.getEdgeVersion()
		if err != nil {
			return "", fmt.Errorf("EdgeDriver文件不存在且无法获取Edge版本\n\n文件路径: %s\n\n解决方案:\n1. 请访问 https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/\n2. 下载与您的Edge浏览器版本匹配的EdgeDriver\n3. 将下载的msedgedriver.exe文件放置在程序目录中\n4. 重新启动程序\n\n错误详情: %v", driverPath, err)
		}

		log.Printf("检测到Edge版本: %s", version)
		
		// 尝试下载EdgeDriver
		err = j.downloadEdgeDriver(version, driverPath)
		if err != nil {
			return "", fmt.Errorf("EdgeDriver自动下载失败\n\n文件路径: %s\nEdge版本: %s\n\n请手动下载:\n1. 访问 https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/\n2. 下载版本 %s 的EdgeDriver\n3. 将msedgedriver.exe放置在程序目录中\n\n错误详情: %v", driverPath, version, version, err)
		}
		
		log.Println("EdgeDriver下载成功")
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

		j.statusLabel.SetText("检查EdgeDriver...")
		j.progressBar.SetValue(0.3)

		// 获取EdgeDriver路径（可能会自动下载）
		driverPath, err := j.getEdgeDriverPath()
		if err != nil {
			dialog.ShowError(err, j.myWindow)
			return
		}

		j.statusLabel.SetText("初始化浏览器...")
		j.progressBar.SetValue(0.6)

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