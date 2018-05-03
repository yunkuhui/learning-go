package services

import (
	"github.com/signintech/gopdf"
)

type PdfTable struct {
	pdf         *gopdf.GoPdf
	x           float64
	y           float64
	lineWidth   float64
	noWrite     []int
	defaultCell *PdfCell
	page        *page
	headRow     *PdfRow
	Rows        []*PdfRow
}

type PdfRow struct {
	h     float64
	Cells []*PdfCell
	table *PdfTable
}

type PdfCell struct {
	W          float64
	text       string
	lineText   []string
	lineTextW  float64
	lineTextH  float64
	topMargin  float64
	leftMargin float64
	align      string
	valign     string
	font       *font
	fontColor  *color
	backColor  *color
	row        *PdfRow
}

type font struct {
	size int
	style string
	family string
}

type color struct {
	r uint8
	g uint8
	b uint8
}

type page struct {
	pageWidth float64
	pageHeight float64
	pageBottomMargin float64
	pageTopMargin float64
}

func CreateTable(x float64, y float64, pdf *gopdf.GoPdf) (*PdfTable) {
	table := PdfTable{}
	defaultCell := PdfCell{}
	page := page{}
	fontColor := color{}
	backColor := color{}
	table.x = x
	table.y = y
	table.pdf = pdf
	// 初始化默认线宽
	table.lineWidth = 1.0
	// 初始化单元格边距
	defaultCell.topMargin = 2.5
	defaultCell.leftMargin = 3.0
	defaultCell.align = "center"
	defaultCell.valign = "middle"
	// 初始化字体属性默认值黑色
	fontColor.r = 0
	fontColor.g = 0
	fontColor.b = 0
	defaultCell.fontColor = &fontColor
	// 初始化表格填充颜色默认值白色
	backColor.r = 255
	backColor.g = 255
	backColor.b = 255
	defaultCell.backColor = &backColor
	table.defaultCell = &defaultCell
	// 默认页面属性
	page.pageWidth = 595.28
	page.pageHeight = 841.89
	page.pageBottomMargin = 50.0
	page.pageTopMargin = 50.0
	table.page = &page
	return &table
}

func (table *PdfTable) SetLineWidth(lineWidth float64) {
	table.lineWidth = lineWidth
}

func (table *PdfTable) SetLeftMargin(leftMargin float64) {
	table.defaultCell.leftMargin = leftMargin
}

func (table *PdfTable) SetTopMargin(topMargin float64) {
	table.defaultCell.topMargin = topMargin
}

func (table *PdfTable) SetFont(family string, style string, size int) {
	f := table.defaultCell.font
	if f == nil {
		f = &font{}
	}
	f.family = family
	f.style = style
	f.size = size
	table.defaultCell.font = f
}

func (table *PdfTable) SetFontColor(r uint8, g uint8 , b uint8) {
	table.defaultCell.fontColor.r = r
	table.defaultCell.fontColor.g = g
	table.defaultCell.fontColor.b = b
}

func (table *PdfTable) SetBackColor(r uint8, g uint8 , b uint8) {
	table.defaultCell.backColor.r = r
	table.defaultCell.backColor.g = g
	table.defaultCell.backColor.b = b
}

func (table *PdfTable) SetPage(width float64, height float64, bottomMargin float64, topMargin float64) {
	table.page.pageWidth = width
	table.page.pageHeight = height
	table.page.pageBottomMargin = bottomMargin
	table.page.pageTopMargin = topMargin
}

func (table *PdfTable) CreateRow(colHeight float64) (*PdfRow) {
	row := PdfRow{}
	row.h = colHeight
	row.table = table
	table.Rows = append(table.Rows, &row)
	return &row
}

func (table *PdfTable) CreateHeadRow(colHeight float64) (*PdfRow) {
	headRow := PdfRow{}
	headRow.h = colHeight
	headRow.table = table
	table.headRow = &headRow
	return table.headRow
}

func (table *PdfTable) GetHeadRow() (*PdfRow) {
	return table.headRow
}

func (row *PdfRow) CreateCell(colW float64) (*PdfCell) {
	cell := PdfCell{}
	cell.W = colW
	// 设置单元格默认值
	table := row.table
	cell.font = table.defaultCell.font
	cell.fontColor = table.defaultCell.fontColor
	cell.backColor = table.defaultCell.backColor
	cell.topMargin = table.defaultCell.topMargin
	cell.leftMargin = table.defaultCell.leftMargin
	cell.align = table.defaultCell.align
	cell.valign = table.defaultCell.valign
	cell.row = row
	row.Cells = append(row.Cells, &cell)
	return &cell
}

func (row *PdfRow) GetRowHeight() (float64) {
	return row.h
}

func (row *PdfRow) SetRowHeight(h float64) {
	row.h = h
}

func (cell *PdfCell) SetText(text string, leftMargin float64, topMargin float64) (error) {
	cell.text = text
	cell.leftMargin = leftMargin
	cell.topMargin = topMargin
	// 设置内容时就确认文字占用的高度
	table := cell.row.table
	h, err := cell.getTextHeight(table.pdf)
	if err != nil {
		return err
	}
	if h > cell.row.h {
		cell.row.h = h
	}
	return nil
}

func (cell *PdfCell) SetFont(family string, style string, size int) (error) {
	f := font{}
	f.family = family
	f.style = style
	f.size = size
	cell.font = &f
	// 设置内容时就确认文字占用的高度
	table := cell.row.table
	h, err := cell.getTextHeight(table.pdf)
	if err != nil {
		return err
	}
	if h > cell.row.h {
		cell.row.h = h
	}
	return nil
}

func (cell *PdfCell) SetFontColor(r uint8, g uint8 , b uint8) {
	cell.fontColor.r = r
	cell.fontColor.g = g
	cell.fontColor.b = b
}

func (cell *PdfCell) SetBackColor(r uint8, g uint8 , b uint8) {
	cell.backColor.r = r
	cell.backColor.g = g
	cell.backColor.b = b
}

/*
 * align 只能是left,center,right
 * valign 只能是top, bottom, middle
 */
func (cell *PdfCell) SetAlign(align string, valign string) {
	if align == "left" || align == "right" || align == "center" {
		cell.align = align
	}
	if valign == "top" || valign == "bottom" || valign == "middle" {
		cell.valign = valign
	}
}

func (table *PdfTable) Draw() (gopdf.GoPdf, error) {
	pdf := table.pdf
	rows := table.Rows
	cellX := table.x
	cellY := table.y
	// 表头行不能单独显示，如果有表格主体不是0行
	firstRowH := 0.0
	if len(rows) > 0 {
		firstRowH = rows[0].h
	}
	if table.headRow != nil {
		// 表头行
		if cellY + table.headRow.h + firstRowH > table.page.pageHeight - table.page.pageBottomMargin {
			pdf.AddPage()
			cellY = table.page.pageTopMargin
		}
		err := table.headRow.rangeRow(pdf, table, cellX, cellY)
		if err != nil {
			return *pdf, err
		}
		cellY += table.headRow.h
	} else if cellY + firstRowH > table.page.pageHeight - table.page.pageBottomMargin {
		pdf.AddPage()
		cellY = table.page.pageTopMargin
	}
	// 表格主体
	for i, row := range rows {
		err := row.rangeRow(pdf, table, cellX, cellY)
		if err != nil {
			return *pdf, err
		}
		// 行y坐标递增
		cellY += row.h
		if i + 1 < len(rows) && cellY + rows[i + 1].h > table.page.pageHeight - table.page.pageBottomMargin {
			pdf.AddPage()
			cellY = table.page.pageTopMargin
			// 表头行
			if table.headRow != nil {
				err := table.headRow.rangeRow(pdf, table, cellX, cellY)
				if err != nil {
					return *pdf, err
				}
				cellY += table.headRow.h
			}
		}
	}
	return *pdf, nil
}

func (row *PdfRow) rangeRow(pdf *gopdf.GoPdf, table *PdfTable, cellX float64, cellY float64) (error) {
	cells := row.Cells
	for _, cell := range cells {
		textX, textY := cell.getCoordinate(cellX, cellY)
		pdf.SetY(textY)
		pdf.SetFont(cell.font.family, cell.font.style, cell.font.size)
		pdf.SetFillColor(cell.fontColor.r, cell.fontColor.g, cell.fontColor.b)
		lineTexts := cell.lineText
		for _, lineText := range lineTexts {
			pdf.SetX(textX)
			pdf.Cell(nil, lineText)
			pdf.Br(float64(cell.font.size))
		}
		// 遍历完一行文字，再次遍历，边框画入pdf
		if table.lineWidth > 0.0 {
			pdf.SetFillColor(cell.backColor.r, cell.backColor.g, cell.backColor.b)
			pdf.RectFromUpperLeft(cellX, cellY, cell.W, row.h)
		}
		cellX += cell.W
	}
	return nil
}

func (cell *PdfCell) getCoordinate(cellX float64, cellY float64) (float64, float64) {
	// 水平对齐坐标
	if cell.align == "left" {
		cellX += cell.leftMargin
	} else if cell.align == "right" {
		cellX += cell.W - cell.leftMargin - cell.lineTextW
	} else {
		space := (cell.W - cell.lineTextW) / 2
		if space > cell.leftMargin {
			cellX += space
		} else {
			cellX += cell.leftMargin
		}
	}
	// 垂直对齐坐标
	if cell.valign == "top" {
		cellY += cell.topMargin
	} else if cell.valign == "bottom" {
		space := cell.row.h - cell.topMargin - cell.lineTextH
		if space > cell.topMargin {
			cellY += space
		} else {
			cellY += cell.topMargin
		}
	} else {
		space := (cell.row.h - cell.lineTextH) / 2
		if space > cell.topMargin {
			cellY += space
		} else {
			cellY += cell.topMargin
		}
	}
	return cellX, cellY
}

func (cell *PdfCell) getTextHeight(pdf *gopdf.GoPdf) (float64, error) {
	err := pdf.SetFont(cell.font.family, cell.font.style, cell.font.size)
	if err != nil {
		return 0.0, err
	}
	textRune := []rune(cell.text)
	textLen := len(textRune)
	var lineText []string
	// 截取长度初始化
	subLen := 1
	strLen := 0
	subText := ""
	startIndex := 0
	h := 0.0
	// 字符串占宽
	lineTextW := 0.0
	subTextW := 0.0
	for subBool := true; subBool; {
		// 截取结束下标
		endIndex := subLen + startIndex
		if endIndex > textLen {
			endIndex = textLen
		}
		subStr := string(textRune[startIndex : endIndex])
		// 截取文字内容像素长
		textw, err := pdf.MeasureTextWidth(subStr)
		if err != nil {
			return h, err
		}
		if textw >= (cell.W - cell.leftMargin) {
			// 单元格越界，长度截取减一
			subLen--
			// 字符长度截取适合长度
			if subLen == strLen {
				lineText = append(lineText, subText)
				if lineTextW < subTextW {
					lineTextW = subTextW
				}
				// 文字占高累加
				h += float64(cell.font.size)
				// 下一行的文字截取位置
				startIndex += subLen
				strLen = 0
			}
		} else if textw < (cell.W - cell.leftMargin) {
			// 确认最后一行内容
			if subText == subStr {
				lineText = append(lineText, subText)
				if lineTextW < subTextW {
					lineTextW = subTextW
				}
				h += float64(cell.font.size)
				// 已写入最后一段文字，设置循环条件为false不再继续遍历
				subBool = false
			}
			// 临时保存当前截取内容的长度和内容
			strLen = subLen
			subText = subStr
			subTextW = textw
			// 单元格未填满，长度截取+1
			subLen++
		}
	}
	cell.lineText = lineText
	cell.lineTextW = lineTextW
	cell.lineTextH = h
	h += 2 * cell.topMargin
	return h, nil
}
