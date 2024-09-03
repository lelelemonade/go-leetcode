package problem165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
    version1Split:=strings.Split(version1,".")
	version1Int:=make([]int,0)
    version2Split:=strings.Split(version2,".")
	version2Int:=make([]int,0)
	for _,v:=range version1Split{
		i,err:=strconv.Atoi(v)
		if err!=nil{
			panic(err)
		}
		version1Int=append(version1Int,i)
	}
	for _,v:=range version2Split{
		i,err:=strconv.Atoi(v)
		if err!=nil{
			panic(err)
		}
		version2Int=append(version2Int,i)
	}

	i:=0
	for ;i<min(len(version1Int),len(version2Int));i++{
		if version1Int[i]==version2Int[i]{
			continue
		}
		if version1Int[i]>version2Int[i]{
			return 1
		}else{
			return -1
		}
	}
	if i<len(version1Int){
		for ;i<len(version1Int);i++{
			if version1Int[i]!=0{
				return 1
			}
		}
	}else{
		for ;i<len(version2Int);i++{
			if version2Int[i]!=0{
				return -1
			}
		}
	}
	return 0
}