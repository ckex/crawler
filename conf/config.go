package conf

// 软件信息。
const (
	VERSION string = "v01"                                                   // 版本号
	AUTHOR string = "Ckex"                                             // 作者
	PROJECT_DIR = "robot"                                           // 项目目录名
	TAG string = PROJECT_DIR                                           // 软件标识符
)

// 默认配置。
const (
	WORK_ROOT string = TAG + "_tmp"                                   // 运行时的目录名称
	WORK_BIN string = PROJECT_DIR + "_bin"                                   // 运行时bin目录
	CONFIG string = WORK_ROOT + "/conf/config.ini"                   // 配置文件路径
	CACHE_DIR string = WORK_ROOT + "/cache"                           // 缓存文件目录
	PHANTOMJS_DIR string = WORK_BIN                                   // Surfer-Phantom下载器：js文件临时目录
)

// 来自配置文件的配置项。
var (
	PHANTOMJS string = PHANTOMJS_DIR + "/phantomjs"                    // Surfer-Phantom下载器：phantomjs程序路径
)