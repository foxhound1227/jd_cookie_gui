import tkinter as tk
from tkinter import ttk, scrolledtext, messagebox
from selenium import webdriver
from selenium.webdriver.edge.options import Options
from selenium.webdriver.edge.service import Service
from selenium.common.exceptions import InvalidSessionIdException
import os
import sys
import traceback
import requests
import zipfile
import subprocess
import re

def get_edge_version():
    try:
        # 尝试从注册表获取Edge版本
        import winreg
        try:
            key = winreg.OpenKey(winreg.HKEY_CURRENT_USER, r"Software\Microsoft\Edge\BLBeacon")
            version, _ = winreg.QueryValueEx(key, "version")
            winreg.CloseKey(key)
            return version
        except:
            pass
        
        # 尝试从程序文件获取版本
        edge_paths = [
            os.path.join(os.environ.get('PROGRAMFILES(X86)', ''), 'Microsoft', 'Edge', 'Application', 'msedge.exe'),
            os.path.join(os.environ.get('PROGRAMFILES', ''), 'Microsoft', 'Edge', 'Application', 'msedge.exe')
        ]
        
        for edge_path in edge_paths:
            if os.path.exists(edge_path):
                try:
                    result = subprocess.run([edge_path, '--version'], capture_output=True, text=True, timeout=10)
                    if result.returncode == 0:
                        version_match = re.search(r'(\d+\.\d+\.\d+\.\d+)', result.stdout)
                        if version_match:
                            return version_match.group(1)
                except:
                    continue
        
        return None
    except Exception as e:
        print(f"获取Edge版本时出现异常: {str(e)}")
        return None

def download_edge_driver(version, status_callback=None):
    try:
        if status_callback:
            status_callback("正在下载EdgeDriver...")
        
        # 构建下载URL (使用32位版本以提高兼容性)
        download_url = f"https://msedgedriver.microsoft.com/{version}/edgedriver_win32.zip"
        print(f"下载EdgeDriver: {download_url}")
        
        # 下载文件
        response = requests.get(download_url, timeout=30)
        response.raise_for_status()
        
        # 保存到临时文件
        zip_path = "edgedriver.zip"
        with open(zip_path, 'wb') as f:
            f.write(response.content)
        
        # 解压文件
        with zipfile.ZipFile(zip_path, 'r') as zip_ref:
            zip_ref.extractall(".")
        
        # 删除临时文件
        os.remove(zip_path)
        
        if status_callback:
            status_callback("EdgeDriver下载完成")
        
        return True
    except Exception as e:
        print(f"下载EdgeDriver失败: {str(e)}")
        return False

def check_edge_installed():
    try:
        edge_paths = [
            os.path.join(os.environ.get('PROGRAMFILES(X86)', ''), 'Microsoft', 'Edge', 'Application', 'msedge.exe'),
            os.path.join(os.environ.get('PROGRAMFILES', ''), 'Microsoft', 'Edge', 'Application', 'msedge.exe')
        ]
        for path in edge_paths:
            print(f"检查Edge浏览器路径: {path}")
            if os.path.exists(path):
                print(f"找到Edge浏览器: {path}")
                return True
        print("未找到Edge浏览器")
        return False
    except Exception as e:
        print(f"检查Edge浏览器安装时出现异常: {str(e)}")
        traceback.print_exc()
        return False





def get_edge_driver_path(status_callback=None):
    try:
        # 获取EdgeDriver路径
        if getattr(sys, 'frozen', False):
            # 如果是打包后的exe
            base_path = sys._MEIPASS
            driver_path = os.path.join(base_path, 'msedgedriver.exe')
        else:
            # 如果是开发环境
            base_path = os.path.dirname(os.path.abspath(__file__))
            driver_path = os.path.join(base_path, 'msedgedriver.exe')
        
        # 如果EdgeDriver不存在，尝试自动下载
        if not os.path.exists(driver_path):
            print("EdgeDriver不存在，尝试自动下载...")
            if status_callback:
                status_callback("检测Edge版本...")
            
            # 获取Edge版本
            edge_version = get_edge_version()
            if not edge_version:
                raise Exception("无法获取Edge浏览器版本，请确保已安装Microsoft Edge")
            
            print(f"检测到Edge版本: {edge_version}")
            
            # 下载EdgeDriver
            if download_edge_driver(edge_version, status_callback):
                print("EdgeDriver下载成功")
            else:
                raise Exception("EdgeDriver下载失败，请手动下载EdgeDriver并放置在程序目录中")
        
        if not os.path.exists(driver_path):
            raise Exception(f"EdgeDriver文件不存在: {driver_path}\n请手动下载EdgeDriver并放置在程序目录中")
        
        print(f"使用EdgeDriver: {driver_path}")
        return driver_path
    except Exception as e:
        print(f"获取EdgeDriver时出现异常: {str(e)}")
        traceback.print_exc()
        raise

class JDCookieExtractor:
    def __init__(self, root):
        try:
            self.root = root
            self.root.title("京东Cookie提取工具")
            self.root.geometry("400x300")
            
            # 创建主Frame
            self.main_frame = ttk.Frame(root)
            self.main_frame.pack(fill=tk.BOTH, expand=True)
            
            # 添加状态标签
            self.status_label = ttk.Label(self.main_frame, text="正在初始化...")
            self.status_label.pack(pady=5)
            
            # 添加进度条
            self.progress = ttk.Progressbar(self.main_frame, mode='determinate')
            self.progress.pack(fill=tk.X, padx=10, pady=5)
            
            # 添加获取Cookie按钮（初始状态为禁用）
            self.get_cookie_btn = ttk.Button(self.main_frame, text="获取Cookie", command=self.get_cookies, state='disabled')
            self.get_cookie_btn.pack(pady=5)
            
            # 添加Cookie显示区域
            self.cookie_text = scrolledtext.ScrolledText(self.main_frame, width=50, height=20)
            self.cookie_text.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)
            
            # 使用after方法延迟执行初始化
            self.root.after(100, self.delayed_init)
        except Exception as e:
            print(f"初始化程序时出现异常: {str(e)}")
            traceback.print_exc()
            messagebox.showerror("错误", f"程序初始化失败: {str(e)}")
            self.root.destroy()
            sys.exit(1)

    def delayed_init(self):
        try:
            # 更新进度条
            self.progress['value'] = 0
            self.status_label.config(text="检查Edge浏览器...")
            self.root.update()
            
            if not check_edge_installed():
                messagebox.showerror("错误", "未检测到Edge浏览器，请先安装Edge浏览器后再运行程序。\n\n下载地址：https://www.microsoft.com/edge")
                self.root.destroy()
                sys.exit(1)
            
            # 更新进度条
            self.progress['value'] = 50
            self.status_label.config(text="初始化浏览器...")
            self.root.update()
            
            # 初始化浏览器
            self.init_browser()
            
            # 完成初始化
            self.progress['value'] = 100
            self.status_label.config(text="准备就绪")
            self.get_cookie_btn.config(state='normal')
            
        except Exception as e:
            print(f"初始化过程出现异常: {str(e)}")
            traceback.print_exc()
            self.status_label.config(text="初始化失败")
            self.cookie_text.delete(1.0, tk.END)
            self.cookie_text.insert(tk.END, f"初始化失败: {str(e)}")
            messagebox.showerror("错误", f"程序初始化失败: {str(e)}")
    
    def init_browser(self):
        try:
            edge_options = Options()
            edge_options.add_argument("--disable-infobars")
            edge_options.add_argument("--disable-extensions")
            edge_options.add_argument("--disable-gpu")
            edge_options.add_argument("--no-sandbox")
            edge_options.add_argument("--disable-dev-shm-usage")
            edge_options.add_argument("--start-maximized")
            edge_options.add_argument("--disable-blink-features=AutomationControlled")
            edge_options.add_argument("--disable-features=msEdgeTranslate")
            edge_options.add_experimental_option('excludeSwitches', ['enable-automation'])
            edge_options.add_experimental_option('useAutomationExtension', False)
            
            # 获取EdgeDriver路径（传递状态回调）
            def update_status(message):
                self.status_label.config(text=message)
                self.root.update()
            
            driver_path = get_edge_driver_path(update_status)
            service = Service(driver_path)
            service.creation_flags = 0x08000000  # CREATE_NO_WINDOW标志
            
            # 初始化Edge浏览器
            self.driver = webdriver.Edge(service=service, options=edge_options)
            
            # 设置页面加载超时时间
            self.driver.set_page_load_timeout(10)
            self.driver.implicitly_wait(5)
            
            # 获取屏幕尺寸
            screen_width = self.root.winfo_screenwidth()
            screen_height = self.root.winfo_screenheight()
            
            # 设置浏览器窗口位置和大小
            self.driver.set_window_size(screen_width // 2, screen_height)
            self.driver.set_window_position(screen_width // 2, 0)
            
            # 尝试加载京东页面
            try:
                self.driver.execute_cdp_cmd('Page.addScriptToEvaluateOnNewDocument', {
                    'source': '''
                        Object.defineProperty(navigator, 'webdriver', {
                            get: () => undefined
                        })
                    '''
                })
                self.driver.get("https://plogin.m.jd.com/login/login")
            except Exception as e:
                print(f"加载京东页面失败: {str(e)}")
                # 即使加载失败也继续运行，让用户可以手动刷新页面
                pass
            
            print("浏览器初始化完成")
            
        except Exception as e:
            print(f"初始化浏览器时出现异常: {str(e)}")
            traceback.print_exc()
            raise
    
    def get_cookies(self):
        try:
            self.status_label.config(text="正在获取Cookie...")
            cookies = self.driver.get_cookies()
            self.cookie_text.delete(1.0, tk.END)
            pt_pin = None
            pt_key = None
            for cookie in cookies:
                if cookie['name'] == 'pt_pin':
                    pt_pin = cookie['value']
                elif cookie['name'] == 'pt_key':
                    pt_key = cookie['value']
            
            if pt_pin and pt_key:
                self.cookie_text.insert(tk.END, f'pt_pin={pt_pin};pt_key={pt_key};')
                self.status_label.config(text="Cookie获取成功")
            else:
                self.cookie_text.insert(tk.END, "未找到所需的Cookie，请确保已登录京东")
                self.status_label.config(text="Cookie获取失败")
                
        except InvalidSessionIdException:
            print("浏览器会话已失效，正在重新初始化...")
            self.init_browser()
            self.get_cookies()
        except Exception as e:
            print(f"获取Cookie时出现异常: {str(e)}")
            traceback.print_exc()
            self.status_label.config(text="Cookie获取失败")
            self.cookie_text.delete(1.0, tk.END)
            self.cookie_text.insert(tk.END, f"获取Cookie失败: {str(e)}")
    
    def on_closing(self):
        if hasattr(self, 'driver'):
            try:
                self.driver.quit()
            except:
                pass
        self.root.destroy()

if __name__ == "__main__":
    try:
        print("程序启动...")
        root = tk.Tk()
        app = JDCookieExtractor(root)
        root.protocol("WM_DELETE_WINDOW", app.on_closing)
        root.mainloop()
    except Exception as e:
        print(f"程序运行时出现异常: {str(e)}")
        traceback.print_exc()