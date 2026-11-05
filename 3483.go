func totalNumbers(digits []int) int{

var x[10]int
var a,b,c int
var hasil int
	for i:=0; i<len(digits); i++{
	 x[digits[i]] ++
	}
	for i:= 100; i<999; i += 2{
	 a = i % 10 //satuan
	 c = i / 100 //ratusan
	 b = (i / 10) % 10 // puluhan
	 if (a == b) && (b == c) {
		if x[a] >= 3 {
		hasil ++
		}
	 }else if a == b { 
		  if (x[a] >= 2) && (x[c] >= 1){
		  hasil++
		  } 
	 }else if b == c {
		  if (x[b] >= 2) && (x[a] >=1){
	    		hasil ++
		  	}
	 }else if c == a {
		   if (x[a] >= 2) && (x[b] >=1){
		    hasil ++
		   }
	 }else{
	 	   if (( x[a] >=1) && (x[b] >=1)) && (x[c] >= 1) {
			hasil++
	           } 
	 }
	}
return hasil

}
