# 设置控制台输出编码为 UTF-8
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$OutputEncoding = [System.Text.Encoding]::UTF8

# 设置工作目录为当前脚本所在目录
Set-Location -Path $PSScriptRoot

# 获取上一级目录
$parentDir = Split-Path -Parent $PSScriptRoot

# 遍历上一级目录下的所有文件夹
Get-ChildItem -Path $parentDir -Directory | ForEach-Object {
    $folderName = $_.Name
    
    # 排除 tests 和 symbols 目录
    if ($folderName -ne "tests" -and $folderName -ne "symbols") {
        Write-Host "Processing: $folderName" -ForegroundColor Green
        
        # 执行 yaegi extract 命令
        yaegi extract "github.com/yanlingrpa/protocol/$folderName"
        
        if ($LASTEXITCODE -eq 0) {
            Write-Host "  Success: extracted $folderName" -ForegroundColor Cyan
        } else {
            Write-Host "  Failed: $folderName" -ForegroundColor Red
        }
    }
}

Write-Host "`nAll done!" -ForegroundColor Yellow