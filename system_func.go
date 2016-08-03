package main

import(

)

func ParseBool(str string) (value bool) {
    switch str {
    	case "1","true", "TRUE", "True":
    			return true
    	case "0","false", "FALSE", "False":
    			return false
    }
    return false
}

