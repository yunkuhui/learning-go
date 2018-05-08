# 使用 go 语言生成 pdf

当前 go 语言生成 pdf 的第 3 方依赖库，还不是很成熟。由于生成的 pdf 中需要中文字符，本人选择使用支持中文的 [signintech/gopdf](https://github.com/signintech/gopdf) 类库。该类库目前有几个缺陷，但是已经可以制作简单的 pdf 文件。以下是该类库的中的部分缺陷：
  1. 不能读取 pdf 文件
  2. 设置字体大小只能是整数，不能是浮点数值
  3. 未封装生成表格
  4. 无法设置线条的颜色
  
********************************************************************************

## 使用 [signintech/gopdf](https://github.com/signintech/gopdf) 类库封装制作 pdf 表格

### pdf 页面

* 页面大小以 A4 大小为标准，使用 72 分辨率。
  * A4宽 595.28 px
  * A4高 841.89 px
  
* 厘米（cm）与像素（px）的计算公式为：

      像素 = 分辨率 * 厘米 / 2.54
        
* 计算字符（字体大小：size）在 pdf 页面中占用的高度（单位：px）：

      pdf := gopdf.GoPdf{}
      var h float64
      h = pdf.Br(size)
        
* 计算字符串在 pdf 页面中所占宽度（单位：px）：

      pdf := gopdf.GoPdf{}
      var w float64
      w,_ = pdf.MeasureTextWidth("字符串")
        
  **注意：** 在 pdf 页面中，一个中文字符占宽是一个英文字符的双倍。
  
* X，Y 的起始位置：从页面左上角开始（X：0，Y：0），横向为 X 轴，竖向为 Y 轴，坐标往右和往下为正向

### 使用自己个人封装的代码生成 pdf 表格

封装的代码尚未经过实践，不管是构思或是细节，还需在实践中不断修改和补充。
(源代码文件：leiyunran_learning_note/code/pdfTableService.go)

#### 表格对象结构

表格对象中，包含表格左上角的位置和行数组，每个行元素中，包含单元格数组。每个单元格元素中，包含各自的字体配置、内容以及背景颜色等。绘制表格时，需要双 for 循环，分别遍历每一行以及行中的每一个单元格。

#### 功能使用

* 创建表格对象

      // x，y 为表格左上角坐标
      var table *PdfTable
      table = CreateTable(x float64, y float64, pdf *gopdf.GoPdf)
    
* 设置表格线宽

      // 默认线宽为 1.0 xp
      table.SetLineWidth(lineWidth float64)
    
* 设置单元格中，内容与左右边框的间隔距离
      
      // 默认 3.0 xp
      table.SetLeftMargin(leftMargin float64)
      
* 设置单元格中，内容与上下边框的间隔距离

      // 默认 2.5 xp
      table.SetTopMargin
    
* 设置表格默认字体

      // 嵌入字体，family 为自定义字体名称，ttfpath 为 ttf 文件相对路径
      pdf.AddTTFFont(family string, ttfpath string)
      // style 为 "B" 或 "" 或 "I" 或 "U"；size 为字体像素大小
      table.SetFont(family string, style string, size int)
    
* 设置表格字体颜色
    
      // 默认黑色(0, 0, 0)
      table.SetFontColor(r uint8, g uint8 , b uint8) 
    
* 设置表格填充颜色(背景颜色)

      // 默认白色(255, 255, 255)
      table.SetBackColor(r uint8, g uint8 , b uint8)
    
* 设置页面宽高以及上下边距(暂不支持左右边距)

      // 默认宽595.28，高841.89，上下页边距 50.0
      table.SetPage(width float64, height float64, bottomMargin float64, topMargin float64)
      
* 执行绘制pdf表格

      table.Draw() (gopdf.GoPdf, error)    
    
* 创建表头行

      // 暂无默认高度
      // 多次创建表头行，则最后取最后一次
      var row *PdfRow
      row = table.CreateHeadRow(colHeight float64) (*PdfRow)
      
  **注意：**行高小于内容需要的高度时，会因为单元格自动换行，动态匹配内容的高度
  
* 创建行

      // 暂无默认高度
      // 一个表格有多行，需要多次创建行
      var row *PdfRow
      row = table.CreateRow(colHeight float64) (*PdfRow)
      
* 创建单元格

      // colW 为单元格宽度，单位 xp
      // 一行有多个单元格，需要多次创建单元格
      var cell *PdfCell
      row.CreateCell(colW float64) (*PdfCell)
      
* 设置单元格内容

      cell.SetText(text string, leftMargin float64, topMargin float64)
      
  **注意：**设置单元格内容时，会自动计算单元格需要的高度，如果高度大于行高，则修改行高
      
* 设置单元格字体
      
      // 默认使用 table 中设置的字体样式
      cell.SetFont(family string, style string, size int)
 
  **注意：**设置单元格字体时，会自动计算单元格需要的高度，如果高度大于行高，则修改行高
      
* 设置单元格字体颜色

      // 默认使用 table 中设置的字体颜色
      cell.SetFontColor(r uint8, g uint8 , b uint8)
      
* 设置单元格背景颜色

      // 默认使用 table 中设置的背景颜色
      cell.SetBackColor(r uint8, g uint8 , b uint8) 
      
* 设置单元格内容的位置

      // align表示水平位置，只能是 left、center、right，默认是 center
      // valign表示垂直位置，只能是 top、 bottom、middle，默认是 middle
      cell.SetAlign(align string, valign string)
      

