package services

type ExportFormat string

const(
 FormatTXT ExportFormat="txt"
 FormatPDF ExportFormat="pdf"
)

type Exporter interface{ Export(content string,f ExportFormat)([]byte,error) }