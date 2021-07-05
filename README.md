#util-go 功能说明
##algorithm 为常见算法实现
    (1) snowflake为雪花算法实现
    (2) sort为排序相关
        a, QuickSort快速排序
        b, BubbleSort冒泡排序
        c, BinarySearch二分查找法
        d, FibonacciRecursion斐波拉契数列
        e, FibonacciFind迭代法
##data为常见数据类型转换
    (1) JsonToMap json转map
    (2) JsonToMaps json转map数组
    (3) MapToJson map转json
    (4) Struct2Map 结构体转map
    (5) DecConvertToBin 16进制字符串转二进制
    (6) SliceReverse 数组反转
    (7)ReverseString 字符串反转
    (8)ClearStringByRegex 清除正则表达式匹配到的字符
    (9)ClearStringUnusual 清楚字符串中给的"\x00"，空格和换行符
    (10)XmlToMap xml转map
            a,最外层legend-root添加根节点, 规避xml不规范,格式化问题
            b,删除注释,规避注释里面不规范的写法

##devops为devops相关组件方法
    (1)cpu为获取cpu名称和总数
    (2)disk为获取磁盘信息
    (3)memory为获取内存信息
    (4)net-interface为获取网络接口信息

##file为文件常见相关操作
    (1) CheckFileIsExist检查文件是否存在
    (2) ReadFile读文件
    (3)WriteFile写文件

##int-list为int类型的数组相关方法
## logger中主要实现日志相关借助zapcore
## net为网络相关 包含下载和上传文件
    (1)UploadFile
    (2)DownloadFile
## scrypt为加密解密和密码相关
    (1)ScryptPassword为密码hash
    (2)ComparePassword比较密码
    (3)ProductToken生产token,根据uuid生产
    (4)Md5Sum获取文件md5
    (5)Md5SumString获取字符串md5
    (6)GetRandomString从数字和字母中获取指定长度的随机字符串

##shell为shell操作相关
    (1)RunCommand为运行shell命令，返回结果

##time_utils为时间相关方法
    (1)GetDate获取YYYY-MM-DD hh:mm:ss格式的时间
    (2)GetUnix获取Unix时间戳
    (3)GetMilliUnix获取毫秒级时间戳
    (4)GetNanoUnix获取纳秒级时间戳
    (5)TimestampToDate时间戳转时间
    (6)DateToTimestamp时间转时间戳
