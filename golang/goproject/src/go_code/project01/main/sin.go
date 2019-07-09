package main

import (
    "image"
    "image/color"
    "image/png"
    "log"
    "math"
    "os"
)

func main() {
	
	//  设置图片背景色
	
	
    // 图片大小
    const size = 300
    // 根据给定大小创建灰度图
    pic := image.NewGray(image.Rect(0, 0, size, size))

    // 遍历每个像素
    for x := 0; x < size; x++ {
        for y := 0; y < size; y++ {
            // 填充为白色
            pic.SetGray(x, y, color.Gray{255})
        }
    }
    /*
    代码说明如下：
    第 2 行，声明一个 size 常量，值为 300。
    第 5 行，使用 image 包的 NewGray() 函数创建一个图片对象，使用区域由 image.Rect 结构提供。image.Rect 描述一个方形的两个定位点 (x1,y1) 和 (x2,y2)。image.Rect(0,0,size,size) 表示使用完整灰度图像素，尺寸为宽 300，长 300。
    第 8 行和第 9 行，遍历灰度图的所有像素。
    第 11 行，将每一个像素的灰度设为 255，也就是白色。
    */
    
    //   绘制正弦函数轨迹
    
    
    // 从0到最大像素生成x坐标
    for x := 0; x < size; x++ {

        // 让sin的值的范围在0~2Pi之间
        s := float64(x) * 2 * math.Pi / size

        // sin的幅度为一半的像素。向下偏移一半像素并翻转
        y := size/2 - math.Sin(s)*size/2

        // 用黑色绘制sin轨迹
        pic.SetGray(x, int(y), color.Gray{0})
    }

	//  写入图片文件
	
	
    // 创建文件
    file, err := os.Create("sin.png")

    if err != nil {
        log.Fatal(err)
    }
    // 使用png格式将数据写入文件
    png.Encode(file, pic) //将image信息写入文件中

    // 关闭文件
    file.Close()
    /*
    第 2 行，创建 sin.png 的文件。
    第 4 行，如果创建文件失败，返回错误，打印错误并终止。
    第 8 行，使用 PNG 包，将图形对象写入文件中。
    第 11 行，关闭文件。
    */
}