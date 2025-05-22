#!/bin/bash
set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <module_name>"
  exit 1
fi

module_name="$1"
module_dir="internal/$module_name"
stub_dir="stubs/[module]"

if grep -q "^$module_name$" ./modules.txt; then
  echo "Error: Module '$module_name' already exists"
  exit 1
fi

mkdir -p "$module_dir"


cp "$stub_dir/entity.go.stub" "$module_dir/entity.go"
cp "$stub_dir/handler.go.stub" "$module_dir/handler.go"
cp "$stub_dir/register.go.stub" "$module_dir/register.go"
cp "$stub_dir/repository.go.stub" "$module_dir/repository.go"
cp "$stub_dir/usecase.go.stub" "$module_dir/usecase.go"

find "$module_dir" -name "*.go" -print0 | while IFS= read -r -d $'\0' file; do
  sed -i "" "s/\[module]/$module_name/g" "$file"
done

echo "Module '$module_name' generated successfully in '$module_dir'"

echo "$module_name" >> modules.txt
