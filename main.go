package main

import "ons/cmd"

func main() {
	cmd.Execute()
}

// 1. 配置文件：yaml、json、环境变量；
// 2. http服务的基础数据的CRUD；
// 3. 优先级：http、tcp、mqtt；
// 4. 边缘网关使用内存存储映射关系，手动维护过期

// 默认端口：
// HTTP: 8080
// TCP: 8888

// TCP和MQTT的payload格式一样，json格式，内容为：请求类型（0-5）、设备ID
