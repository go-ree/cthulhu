# Cthulhu AI MCP服务

这是一个基于 Model Context Protocol (MCP) 架构的 AI 服务示例。

## MCP架构说明

- **Model**: AI模型层
  - 定义多种AI模型接口
  - 处理输入输出
  - 模型参数配置

- **Context**: 上下文环境层
  - 对话历史管理
  - 用户信息管理
  - 会话状态追踪
  - 日志和元数据

- **Protocol**: 通信协议层
  - 请求/响应编解码
  - 路由管理
  - API接口定义
  - 错误处理

## 项目结构

```
.
├── cmd/
│   └── server/           # 服务器启动入口
├── internal/
│   ├── model/
│   │   └── ai/           # AI模型层 - 不同的模型实现
│   ├── context/
│   │   └── ai/           # AI上下文层 - 对话和会话管理
│   └── protocol/
│       └── ai/           # AI协议层 - 通信接口
├── pkg/
│   └── logger/           # 日志工具
└── main.go               # 主入口
```

## 支持的AI模型

- GPT-3
- GPT-4
- BERT

## API接口

- `GET /health` - 健康检查和当前模型信息
- `POST /api/ai/chat` - AI聊天接口
- `GET /api/ai/models` - 获取支持的模型列表

## API请求示例

### 聊天请求

```json
POST /api/ai/chat
{
  "message": "你好，请介绍一下你自己",
  "model_type": "gpt-3",
  "parameters": {
    "temperature": 0.7,
    "max_tokens": 100
  }
}
```

### 聊天响应

```json
{
  "code": 200,
  "message": "success",
  "response": "GPT3处理结果: 你好，请介绍一下你自己",
  "model": "gpt-3",
  "trace_id": "20231215123456-abcd1234",
  "parameters": {
    "temperature": 0.7,
    "max_tokens": 1000,
    "parameters": {
      "top_p": 0.9
    }
  }
}
```

## 如何运行

确保安装了 Go 1.21 或更高版本，然后运行：

```bash
go run main.go
```

或者直接运行服务器：

```bash
go run cmd/server/main.go
```

服务器默认在8080端口启动。