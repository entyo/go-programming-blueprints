#!/bin/zsh
echo domainfinder をビルドしますよ〜しましま
go build -o domainfinder

echo synonyms をビルドしますよ〜しましま
cd ../synonyms/
go build -o ../domainfinder/lib/synonyms

echo available をビルドしますよ〜しましま
cd ../available/
go build -o ../domainfinder/lib/available

echo sprinkle をビルドしますよ〜しましま
cd ../sprinkle/
go build -o ../domainfinder/lib/sprinkle

echo coolify をビルドしますよ〜しましま
cd ../coolify/
go build -o ../domainfinder/lib/coolify

echo domainify をビルドしますよ〜しましま
cd ../domainify/
go build -o ../domainfinder/lib/domainify

echo "よしっ(適当)"
