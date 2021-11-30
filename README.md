### 根据给定时间平均分割视频

在微信给别人发视频的时候会有 15 分钟的时间限制，如果要发大一点的视频要切割成不同的小视频，这个工具就是用来把一个大的视频平均分割成小视频用的，自己写着玩的，所以要严格按照以下步骤操作，以免报错。

首先要安装 ffmpeg 和 ffprobe 两个命令行工具到环境变量中

然后 go build main.go 创建可执行文件

执行 ./main video_name.mp4 2 

第一个参数为视频名称，第二个参数为裁剪时长

PS：可执行文件要和视频在同一目录下 

PPS：该程序只在 mac 系统下调试通过 

PPPS：最好不要裁太太大的文件，否则电脑可能会爆炸，比如一个小时的电影按照 2 分钟分割为 30 段，我没试过，不过我觉得够呛