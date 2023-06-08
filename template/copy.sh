#!/bin/bash

# 源文件路径
source_file="template"

# 目标文件路径
target_file="../v2/protogen.go"

# 参数替换
model_name="Protogen"
func_name_up="Protogen"
func_name="protogen"
model_name_application_name="protogenV22Anime_protogenV22"

# 复制文件并替换参数
cp "$source_file" "$target_file"
sed -i "" "s/{model_name_application_name}/$model_name_application_name/g" "$target_file"
sed -i "" "s/{func_name}/$func_name/g" "$target_file"
sed -i "" "s/{func_name_up}/$func_name_up/g" "$target_file"
sed -i "" "s/{model_name}/$model_name/g" "$target_file"

echo "文件复制完成，并成功替换参数。"
