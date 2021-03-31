package main

func searchMatrix(matrix [][]int, target int) bool {
	for i:=0;i< len(matrix);i++{
		for j:=0;j< len(matrix[i]);j++{
			if(matrix[i][j]==target){
				return true
			}
		}
	}
	return false

}
/**
 看大佬思路优化的版本
 虽然是一个二维矩阵，但是由于它特殊的有序性，所以完全可以按照下标把它看成一个一维矩阵，只不过需要行列坐标转换。最后利用二分搜索直接搜索即可。
 */
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	//m为每一行的长度
	m,low,high:=len(matrix[0]),0,len(matrix[0])*len(matrix)-1
	for low<=high{
		//(low+high)/2在low和high值比较大的时候可能内存溢出，需要改成low+(high-low)/2
		//除以2可以改成位运算>>1,这样效率会更高
		mid:=low+(high-low)>>1
		if(matrix[mid/m][mid%m]<target){
			low=mid+1
		}else if(matrix[mid/m][mid%m]>target){
			high=mid-1
		}else {
			return true
		}

	}
	return false


}
func main() {
	
}
