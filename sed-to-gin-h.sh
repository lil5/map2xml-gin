#!/bin/bash

echo map2xml.go | sed 's/map[string]interface{}/gin.H/' > map2xml-gin.go
