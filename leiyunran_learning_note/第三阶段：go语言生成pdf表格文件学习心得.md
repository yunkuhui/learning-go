#使用go语言生成pdf

当前go语言生成pdf的第3方依赖库，还不是很成熟。由于生成的pdf中需要中文字符，本人选择使用支持中文的[signintech/gopdf](https://github.com/signintech/gopdf)类库。该类库目前有几个缺陷，但是已经可以制作简单的pdf文件。以下是该类库的中的部分缺陷：
  1. 不能读取pdf文件
  2. 设置字体大小只能是整数，不能是浮点数值
  3. 未封装生成表格
  4. 无法设置线条的颜色
  
********************************************************************************

## 使用[signintech/gopdf](https://github.com/signintech/gopdf)类库封装制作pdf表格

### pdf页面

* 页面大小以A4大小为标准，使用72分辨率。
  * A4宽 595.28px
  * A4高 841.89px
  
* 厘米（cm）与像素（px）的计算公式为：

      像素 = 分辨率 * 厘米 / 2.54
        
* 计算字符（字体大小：size）在pdf页面中占用的高度（单位：px）：

      pdf := gopdf.GoPdf{}
      var h float64
      h = pdf.Br(size)
        
* 计算字符串在pdf页面中所占宽度（单位：px）：

      pdf := gopdf.GoPdf{}
      var w float64
      w,_ = pdf.MeasureTextWidth("字符串")
        
  **注意：**在pdf页面中，一个中文字符占宽是一个英文字符的双倍。
* X，Y的起始位置：从页面左上角开始（X：0，Y：0），横向为X轴，竖向为Y轴，坐标往右和往下为正向

### 使用自己个人封装的代码生成pdf表格(源代码文件：code/pdfTableService.go)

#### 表格对象结构



#### 功能调用

* 创建表格对象

      // x，y为表格左上角坐标
      var table *PdfTable
      table = CreateTable(x float64, y float64, pdf *gopdf.GoPdf)
    
* 设置表格线宽

      // 默认线宽为1.0xp
      table.SetLineWidth(lineWidth float64)
    
* 设置单元格中，内容与左右边框的间隔距离
      
      // 默认3.0xp
      table.SetLeftMargin(leftMargin float64)
      
* 设置单元格中，内容与上下边框的间隔距离

      // 默认2.5xp
      table.SetTopMargin
    
* 设置表格默认字体

      // 嵌入字体,family为自定义字体名称，ttfpath为ttf文件相对路径
      pdf.AddTTFFont(family string, ttfpath string)
      // style为"B"或""或"I"或"U";size为字体像素大小
      table.SetFont(family string, style string, size int)
    
* 设置表格字体颜色
    
      // 默认黑色(0, 0, 0)
      table.SetFontColor(r uint8, g uint8 , b uint8) 
    
* 设置表格填充颜色(背景颜色)

      // 默认白色(255, 255, 255)
      table.SetBackColor(r uint8, g uint8 , b uint8)
    
* 设置页面宽高以及上下边距(暂不支持左右边距)

      // 默认宽595.28，高841.89，上下页边距50.0
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

      // colW为单元格宽度，单位xp
      // 一行有多个单元格，需要多次创建单元格
      var cell *PdfCell
      row.CreateCell(colW float64) (*PdfCell)
      
* 设置单元格内容

      cell.SetText(text string, leftMargin float64, topMargin float64)
      
* 设置单元格字体
      
      // 默认使用table中设置的字体样式
      cell.SetFont(family string, style string, size int)
      
* 设置单元格字体颜色

      // 默认使用table中设置的字体颜色
      cell.SetFontColor(r uint8, g uint8 , b uint8)
      
* 设置单元格背景颜色

      // 默认使用table中设置的背景颜色
      cell.SetBackColor(r uint8, g uint8 , b uint8) 
      
* 设置单元格内容的位置

      // align表示水平位置，只能是left,center,right，默认是center
      // valign表示垂直位置，只能是top, bottom, middle，默认是middle
      cell.SetAlign(align string, valign string)
      

