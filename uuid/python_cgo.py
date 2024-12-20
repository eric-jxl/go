import ctypes

# 加载 Go 生成的 C 共享库
go_tool = ctypes.CDLL('./libgo_tool.so')

# 定义 Go 函数的参数类型和返回值类型
go_tool.AddNumbers.argtypes = [ctypes.c_int, ctypes.c_int]
go_tool.AddNumbers.restype = ctypes.c_int

go_tool.Snowflake1.restype = ctypes.c_char_p

# 调用 Go 函数
result = go_tool.AddNumbers(10, 20)
print(f"Result from Go: {result}")

c_string = go_tool.Snowflake1()
python_string = ctypes.c_char_p(c_string).value.decode('utf-8')
print(f"Result from Go: {python_string}")
