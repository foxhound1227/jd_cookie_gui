# 京东Cookie提取工具

一个京东Cookie自动提取工具，提供Python和Go两个版本，支持图形界面操作和GitHub自动编译。

## 版本说明

- **Python版本**: 基于tkinter和Selenium的传统版本
- **Go版本**: 基于Fyne GUI框架的现代化版本，支持跨平台编译

## 功能特点

- 🖥️ 简洁直观的图形用户界面
- 🔄 自动检测Microsoft Edge浏览器
- 🍪 一键获取京东登录Cookie
- 📋 自动复制Cookie到剪贴板
- 🚀 支持打包为独立可执行文件
- ⚡ 精简优化的代码结构

## 系统要求

- Windows 10/11 64位系统
- Microsoft Edge浏览器
- EdgeDriver（已包含在项目中）

## 使用方法

### Go版本（推荐）

#### 方式一：下载预编译版本
1. 从GitHub Releases页面下载对应平台的可执行文件
2. 确保已安装Microsoft Edge浏览器
3. 将`msedgedriver.exe`放置在程序同目录下（Windows版本）
4. 直接运行可执行文件
5. 点击"获取Cookie"按钮，在弹出的浏览器中登录京东账号
6. 登录成功后，Cookie会自动显示在程序界面中

#### 方式二：从源码编译
1. 安装Go 1.21+
2. 克隆项目：`git clone <repository-url>`
3. 进入项目目录：`cd jd-cookie-gui`
4. 安装依赖：`go mod tidy`
5. 编译运行：`go run main.go`
6. 或者编译可执行文件：`go build -o jd-cookie-gui main.go`

#### 使用Makefile（可选）
```bash
# 安装依赖
make deps

# 编译当前平台版本
make build

# 编译所有平台版本
make build-all

# 运行程序
make run
```

### Python版本

#### 方式一：直接运行可执行文件

1. 下载 `dist/jd_cookie_gui.exe`
2. 双击运行程序
3. 点击"获取Cookie"按钮
4. 在弹出的浏览器中登录京东账号
5. 登录成功后，Cookie会自动显示并复制到剪贴板

#### 方式二：运行Python源码

1. 安装Python 3.7+
2. 安装依赖包：
   ```bash
   pip install selenium tkinter
   ```
3. 运行程序：
   ```bash
   python jd_cookie_gui.py
   ```

## 项目结构

```
jd_cookie_gui/
├── jd_cookie_gui.py      # 主程序源码
├── jd_cookie_gui.spec    # PyInstaller打包配置
├── msedgedriver.exe      # Edge WebDriver
├── dist/                 # 打包后的可执行文件
│   └── jd_cookie_gui.exe
├── build/                # 构建临时文件
├── 使用说明.txt          # 详细使用说明
└── README.md             # 项目说明文档
```

## 技术栈

### Go版本
- **GUI框架**: Fyne v2.4.3
- **浏览器自动化**: Selenium WebDriver (github.com/tebeka/selenium)
- **浏览器**: Microsoft Edge
- **编译**: Go 1.21+
- **CI/CD**: GitHub Actions

### Python版本
- **GUI框架**: tkinter
- **浏览器自动化**: Selenium WebDriver
- **浏览器**: Microsoft Edge
- **打包工具**: PyInstaller
- **Microsoft Edge WebDriver** - 浏览器驱动

## GitHub Actions自动编译

Go版本支持GitHub Actions自动编译，每次推送代码到main分支或创建tag时会自动构建多平台版本：

### 支持平台
- **Windows** (amd64)
- **Linux** (amd64) 
- **macOS** (amd64)

### 自动化流程
1. 代码推送触发构建
2. 多平台并行编译
3. 自动上传构建产物
4. 创建tag时自动发布Release

### 获取编译版本
- 访问项目的GitHub Actions页面查看构建状态
- 从Artifacts下载对应平台的可执行文件
- 或从Releases页面下载正式版本

## 开发特色

### Go版本优势
- 跨平台编译，一次编写多处运行
- 现代化GUI框架，界面美观
- 静态编译，无需运行时依赖
- GitHub Actions自动化构建和发布
- 更好的性能和资源占用

### Python版本特点
- 移除了复杂的驱动下载逻辑，简化部署
- 精简错误处理机制，提高程序稳定性
- 优化代码结构，从366行精简至228行
- 清理不必要的依赖和导入
- 自动检测浏览器和驱动状态
- 友好的错误提示信息
- 一键复制功能，方便使用
- 独立可执行文件，无需安装Python环境

## 注意事项

1. 首次使用需要手动登录京东账号
2. 确保Microsoft Edge浏览器已安装
3. 如果遇到驱动问题，请检查Edge版本兼容性
4. Cookie有时效性，建议定期重新获取

## 故障排除

### 常见问题

**Q: 提示"未找到msedgedriver.exe"**
A: 确保msedgedriver.exe文件与程序在同一目录下

**Q: 浏览器无法启动**
A: 检查Microsoft Edge是否正确安装，版本是否兼容

**Q: Cookie获取失败**
A: 确保已成功登录京东账号，刷新页面后重试

## 版本信息

- **当前版本**: v1.0
- **最后更新**: 2025年1月
- **兼容性**: Windows 10/11, Microsoft Edge

## 许可证

本项目仅供学习和个人使用，请遵守相关法律法规和网站服务条款。

## 贡献

欢迎提交Issue和Pull Request来改进这个项目。

---

**免责声明**: 本工具仅用于技术学习和个人使用，使用者需自行承担使用风险，开发者不承担任何法律责任。