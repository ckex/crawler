package conf

import (
	"os"
	"path/filepath"
)

// 软件信息。
const (
	VERSION string = "v01"                                                   // 版本号
	AUTHOR string = "Ckex"                                             // 作者
	PROJECT_DIR = "crawler"                                           // 项目目录名z
	TAG string = PROJECT_DIR                                           // 软件标识符
)

// 默认配置。
const (
	WORK_TMP string = TAG + "_tmp"                                   // 运行时的目录名称
	WORK_BIN string = PROJECT_DIR + "_bin"                                   // 运行时bin目录
	PHANTOMJS_DIR string = WORK_BIN                                   // Surfer-Phantom下载器：js文件临时目录
)

// 来自配置文件的配置项。
var (
	workDir, _ = os.Getwd()
	PHANTOMJS string = filepath.Join(workDir, PHANTOMJS_DIR, "phantomjs") // Surfer-Phantom下载器：phantomjs程序路径
	CACHE_DIR string = filepath.Join(workDir, WORK_BIN, "cache")                         // 缓存文件目录
)