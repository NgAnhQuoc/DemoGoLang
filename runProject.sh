#!/bin/bash

# Kiểm tra xem Go đã được cài chưa
if ! command -v go &> /dev/null
then
    echo "Go chưa được cài đặt. Vui lòng cài Go trước khi chạy script này."
    exit 1
fi

# Biên dịch file Go
go build -o hello hello.go

# Kiểm tra nếu quá trình build thành công
if [ $? -eq 0 ]; then
    echo "Build thành công. Chạy chương trình..."
    ./hello
else
    echo "Build thất bại."
    exit 1
fi
