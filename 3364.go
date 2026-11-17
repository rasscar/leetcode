func minimumSumSubarray(nums []int, l int, r int) int {
    	var n, pjg, i, k, jum, terbaik int
        	n = len(nums)
            	terbaik = 1 << 30

                	for pjg = l; pjg <= r; pjg++ {
                    		for i = 0; i+pjg <= n; i++ {
                            			jum = 0
                                        			for k = i; k < i+pjg; k++ {
                                                    				jum += nums[k]
                                                                    			}
                                                                                			if jum > 0 && jum < terbaik {
                                                                                            				terbaik = jum
                                                                                                            			}
                                                                                                                        		}
                                                                                                                                	}

                                                                                                                                    	if terbaik == 1<<30 {
                                                                                                                                        		return -1
                                                                                                                                                	}
                                                                                                                                                    	return terbaik
                                                                                                                                                        }
