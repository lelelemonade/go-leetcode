package problem535

import "encoding/base64"

type Codec struct {
    
}


func Constructor() Codec {
    return Codec{}
}

// Encodes a URL to a shorten ed URL.
func (this *Codec) encode(longUrl string) string {
	i:=0
	slashCount:=0
	for ;i<len(longUrl);i++{
		if longUrl[i]=='/'{
			slashCount++
		}
		if slashCount>=3{
			break
		}
	}
	return longUrl[:i+1]+base64.StdEncoding.EncodeToString([]byte(longUrl[i+1:]))
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	i:=0
	slashCount:=0
	for ;i<len(shortUrl);i++{
		if shortUrl[i]=='/'{
			slashCount++
		}
		if slashCount>=3{
			break
		}
	}
	raw,_:=base64.StdEncoding.DecodeString(shortUrl[i+1:])
	return shortUrl[:i+1]+string(raw)
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
