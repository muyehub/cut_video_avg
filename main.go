package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// 视频名称
	videoName := os.Args[1]
	// 按多少分钟分割
	splitMinute, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("请输入整数分钟：", err)
	}
	splitSecond := splitMinute * 60
	// 裁剪时长
	dur := fmt.Sprintf("%d", splitSecond)

	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=duration", "-of", "default=noprint_wrappers=1:nokey=1 ", videoName)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("获取视频时长错误: ", err)
	}

	// 视频时长，单位为秒
	stdoutStr := strings.Replace(string(stdout), "\n", "", -1)
	videoDuration, err := strconv.ParseFloat(stdoutStr, 64)
	if err != nil {
		log.Fatal("获取视频时长后转换描述时出错：", err)
	}

	// 向下取整
	vd := int(videoDuration)

	fmt.Println(vd, vd/splitSecond, vd%splitSecond)

	// 按照时间分为多少组
	group := vd / splitSecond
	groupMore := vd % splitSecond
	for i := 0; i < group; i++ {
		start := fmt.Sprintf("%d", i*splitSecond)
		go func(i int) {
			cutCmd := exec.Command("ffmpeg", "-i", videoName, "-ss", start, "-t", dur, fmt.Sprintf("%d.mp4", i))
			stdoutCut, err := cutCmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(stdoutCut))
				log.Fatal("视频截取错误: ", err)
			}
			fmt.Println(string(stdoutCut))
		}(i)
	}
	if groupMore > 0 {
		startMore := fmt.Sprintf("%d", group*splitSecond)
		cutCmdMore := exec.Command("ffmpeg", "-i", videoName, "-ss", startMore, "-t", dur, fmt.Sprintf("%d.mp4", group))
		stdoutCutMore, err := cutCmdMore.CombinedOutput()
		if err != nil {
			log.Fatal("视频截取错误: ", err)
		}
		fmt.Println(string(stdoutCutMore))
	}
}
