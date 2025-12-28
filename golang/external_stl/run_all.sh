#!/bin/bash

# Script to run all example files

echo "Installing dependencies..."
go get github.com/emirpasic/gods@latest

echo ""
echo "========================================="
echo "Running set.go examples"
echo "========================================="
go run set.go

echo ""
echo "========================================="
echo "Running map.go examples"
echo "========================================="
go run map.go

echo ""
echo "========================================="
echo "Running priority_queue.go examples"
echo "========================================="
go run priority_queue.go

echo ""
echo "========================================="
echo "Running stack.go examples"
echo "========================================="
go run stack.go

echo ""
echo "========================================="
echo "Running queue.go examples"
echo "========================================="
go run queue.go

echo ""
echo "========================================="
echo "Running deque.go examples"
echo "========================================="
go run deque.go

echo ""
echo "All examples completed!"

