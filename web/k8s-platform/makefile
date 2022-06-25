# 定义变量
APP_NAME := k8s-platform
PKG := github.com/iewiewiew/Learn-Go/web/k8s-platform
BUILD_DIR := build
SRC_FILES := $(shell find . -name "*.go")

# 默认目标，执行编译和构建
.PHONY: all
#all: clean build run
all: run

# 清理构建目录
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# 构建目标
.PHONY: build
build: $(BUILD_DIR)/$(APP_NAME)

# 编译可执行文件
$(BUILD_DIR)/$(APP_NAME): $(SRC_FILES)
	#go build -o $@ $(PKG)
	go build -o $@ ./main.go

# 安装依赖
.PHONY: deps
deps:
	go get -d -v ./...

# 运行项目
.PHONY: run
run:
	#go run $(PKG)
	go run main.go

# 测试项目
.PHONY: test
test:
	go test -v ./...

# 帮助文档
.PHONY: help
help:
	@echo "可用的目标："
	@echo "  all       编译和构建项目（默认目标）"
	@echo "  clean     清理构建目录"
	@echo "  build     编译可执行文件"
	@echo "  deps      安装依赖"
	@echo "  run       运行项目"
	@echo "  test      测试项目"
	@echo "  help      显示帮助信息"