import time
from ctypes import cdll, c_char_p
 
 
start = time.time()
 
 
# 加载动态链接库
lib = cdll.LoadLibrary('./main.so')
 
 
# 配置输出参数变量类型
lib.count_time.restype = c_char_p
 
 
# 调用方法
lib.qrcodeText()
#print(ip)
rest = lib.count_time()
 
 
end = time.time()
 
 
print(f"Go 内部执行时间：{rest}")
print(f"Python 整体执行时间: {end - start}s")
