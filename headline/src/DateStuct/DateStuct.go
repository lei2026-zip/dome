package DateStuct

type News struct{
   Reason  string
   Error_code int8
   Result Res
}

type  Res struct{
	 Data []List
	 Class string
	 Stat string
}

type List struct{
   Author_name string
   Category string
   Date string
   Title string
   Thumbnail_pic_s string
   Uniquekey string
   Url string
}
