   function CaculateWeekDay(year1, month1, day1)        //通过年月日计算该日的星期
				{
				    var week = 1;
					if(month1 == 1 || month1 == 2)
					{
						month1 += 12;
						year1--;
						
					}		
         //    基姆拉尔森计算公式
					///1582年9月3日后：w = (day + 2*month+3*(month+1)/5+y+y/4-y/100+y/400)%7;
	                //1582年9月3日前：w = (d+2*m+3*(m+1)/5+y+y/4+5） % 7;
					week = (day1+2*month1+parseInt(3*((month1+1)/5))+year1+parseInt(year1/4)-parseInt(year1/100)+parseInt(year1/400))%7;		
					week++;
					if(week==7)
					{
						week=0;
					}
					return (week);
				}
 function judge_month_length(year1,month1){   //判断这某年某月的总日数===========
		    	if(month1>12)
		    	{
		    		year1++;
		    		month1 = month1-12;
		    	}
		        if(month1<1)
		        {
		        	year1--;
		        	month1 = month1+12;
		        }
		    	if((((month1%2==1)&&(month1<8))||((month1%2==0)&&(month1>7)))&&(month1!=2))
		    	{
		    		return 31;
		    	}
		    	else if(month1==2){
		    		if(judge_year(year1))
		    		{
		    			return 29;
		    		}
		    		else{
		    			return 28;
		    		}
		    	}
		    	else{
		    		return 30;
		    	}
		    }
    function judge_month(month1){
		    	if(month1>12)
		    	{
		    		return month1-12;
		    	}
		    	else if(month1<1)
		    	{
		    		return month1+12;
		    	}
		    	else{
		    		return month1
		    	}
		    }
       //============================================================
    //判断是否是闰年=================================
    function judge_year(year1){
    	if(((year1%4==0)&&(year1%100!=0))||(year1%400==0))
    	{
    		return 1;
    	}
    	else{
    		return 0;
    	}
    }
    
    
    	function judge_festival(month1,day1){return "";
			    switch(month1){
			    	case 1: switch(day1){
			    		case 1: return "元旦";
			    		case 2: return "腊八";
			    		case 23: return "小年";
			    		case 30: return "除夕";
			    		case 31: return "春节";
			    	}break;
			    	case 2: switch(day1){
			    		case 4: return "立春";
			    		case 14: return "元宵节";
			    	}break;
			    	case 3: switch(day1){
			    		case 7: return "女生节";
			    		case 8: return "妇女节";
			    		case 12: return "植树节";
			    	}break;
			    	case 4: switch(day1){
			    		case 1:return "愚人节";
			    		case 4:return "清明节";
			    		case 22:return "地球日";
			    	}break;
			    	case 5: switch(day1){
			    		case 1:return "劳动节";
			    		case 4:return "青年节";
			    		case 5:return "立夏";
			    		case 8:return "母亲节";
			    		case 20:return "小满";
			    	}break;
			    	case 6: switch(day1){
			    		case 1:return "儿童节";
			    		case 5:return "芒种";
			    		case 9:return "端午节";
			    		case 19:return "父亲节";
			    		case 21:return "夏至";
			    	}break;
			    	case 7: switch(day1){
			    		case 1:return "建党节";
			    		case 7:return "小暑";
			    		case 22:return "大暑";
			    	}break;
			    	case 8: switch(day1){
			    		case 1:return "建军节";
			    		case 7:return "立秋";
			    		case 9:return "七夕";
			    	}break;
			    	case 9: switch(day1){
			    		 case 10:return "教师节";
			    		// case 15:return "中秋节";
			    	}break;
			    	case 10: switch(day1){
			    		case 1: return "国庆";
			    		case 9: return "重阳节";
			    	};
			    	case 11: switch(day1){
			    		case 7:return "立冬";
			    		case 11:return "光棍节";
			    		case 22:return "小雪";
			    		case 24:return "感恩节";
			    	}break;
			    	case 12: switch(day1){
			    		case 7: return "大雪";
			    		case 22: return "冬至";
			    		case 24: return "平安夜";
			    		case 25: return "圣诞节";
               case 26:return  "腊月" ;
			    	}break;	
			    	default:break;
			    }
			    return "";
			} //判断节假日