#!/bin/bash

# 配置参数
APP_NAME="myapp-back"           # 应用名称
BUILD_DIR="./bin"          # 构建输出目录
LOG_DIR="./logs"           # 日志目录
CONFIG_FILE="./cfg.json" # 配置文件路径

# 自动检测系统和架构
OS="unknown"
ARCH="unknown"
case "$(uname -s)" in
    Linux*)     OS="linux";;
    Darwin*)    OS="darwin";;
    CYGWIN*)    OS="windows";;
    MINGW*)     OS="windows";;
    *)          OS="unknown"
esac

case "$(uname -m)" in
    x86_64)     ARCH="amd64";;
    aarch64)    ARCH="arm64";;
    armv7l)     ARCH="arm";;
    *)          ARCH="unknown"
esac

# 构建函数
build() {
    echo "Building binaries..."
    mkdir -p $BUILD_DIR

    # 构建Linux amd64
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "${BUILD_DIR}/${APP_NAME}" main.go
    echo "Linux amd64 binary: ${BUILD_DIR}/${APP_NAME}"

    # 构建Windows amd64
    # GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o "${BUILD_DIR}/${APP_NAME}.exe" main.go
    # echo "Windows amd64 binary: ${BUILD_DIR}/${APP_NAME}.exe"

    # 构建当前平台二进制
    # CGO_ENABLED=0 go build -o "${BUILD_DIR}/${APP_NAME}" main.go
    # echo "Current platform binary: ${BUILD_DIR}/${APP_NAME}"
}

# 启动函数
start() {
    echo "Starting service..."
    mkdir -p $LOG_DIR

    local BINARY_PATH="${BUILD_DIR}/${APP_NAME}"
    if [ "$OS" = "windows" ]; then
        BINARY_PATH="${BUILD_DIR}/${APP_NAME}.exe"
    fi

    if [ ! -f "$BINARY_PATH" ]; then
        echo "Error: Binary not found. Run build first."
        exit 1
    fi

    # 检查是否已运行
    if [ -f "${BUILD_DIR}/pid" ]; then
        echo "Service is already running (PID: $(cat ${BUILD_DIR}/pid))"
        exit 1
    fi

    # 启动服务并记录PID
    if [ "$OS" = "windows" ]; then
        start /B "$BINARY_PATH" -c "$CONFIG_FILE" > "${LOG_DIR}/output.log" 2>&1
        echo $! > "${BUILD_DIR}/pid"
    else
        "$BINARY_PATH" -c "$CONFIG_FILE" > "${LOG_DIR}/output.log" 2>&1 &
        echo $! > "${BUILD_DIR}/pid"
    fi

    echo "Service started (PID: $(cat ${BUILD_DIR}/pid))"
}

# 停止函数
stop() {
    echo "Stopping service..."
    if [ ! -f "${BUILD_DIR}/pid" ]; then
        echo "Service is not running"
        return
    fi

    local PID=$(cat "${BUILD_DIR}/pid")
    if [ "$OS" = "windows" ]; then
        taskkill //F //PID $PID
    else
        kill $PID
    fi

    rm "${BUILD_DIR}/pid"
    echo "Service stopped"
}

# 状态函数
status() {
    if [ -f "${BUILD_DIR}/pid" ]; then
        local PID=$(cat "${BUILD_DIR}/pid")
        if ps -p $PID > /dev/null 2>&1; then
            echo "Service is running (PID: $PID)"
            return
        else
            echo "Service pid file exists but process not found"
            rm "${BUILD_DIR}/pid"
        fi
    fi
    echo "Service is not running"
}

# 清理函数
clean() {
    echo "Cleaning build artifacts..."
    rm -rf $BUILD_DIR
    rm -rf $LOG_DIR
}

# 主逻辑
case "$1" in
    build)
        build
        ;;
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        sleep 2
        start
        ;;
    status)
        status
        ;;
    clean)
        clean
        ;;
    *)
        echo "Usage: $0 {build|start|stop|restart|status|clean}"
        exit 1
esac