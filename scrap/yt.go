package scrape

import (
	"encoding/json"
	"os/exec"
	"fmt"
	"net/url"
)

func Ytdownload(id, k string) (string, error) {
    encodeK := url.QueryEscape(k)
    out, _ := exec.Command("bash", "-c", "curl 'https://www.yt1s.com/api/ajaxConvert/convert' --data-raw 'vid="+id+"&k="+encodeK+"'").Output()
	var result map[string]interface{}
	json.Unmarshal(out, &result)
	dlink := fmt.Sprintf("%v", result["dlink"]) 
	return dlink, nil
}

func Ytinfo(url string) (map[string]interface{}, error) {
    out, _ := exec.Command("bash", "-c", "curl 'https://www.yt1s.com/api/ajaxSearch/index' --data-raw 'q="+url+"&vt=home'").Output()
    var result map[string]interface{}
    json.Unmarshal(out, &result)
    return result, nil
}
 
func Ytmetadata(id string) (map[string]interface{}, error) {
    out, _ := exec.Command("bash", "-c", "curl 'https://www.googleapis.com/youtube/v3/videos?key=AIzaSyASTMQck-jttF8qy9rtEnt1HyEYw5AmhE8&quotaUser=ALyRzwjclEBKVTHlQ5nLjE8unFfIELQvrD33vWma&part=snippet%2Cstatistics%2CrecordingDetails%2Cstatus%2CliveStreamingDetails%2Clocalizations%2CcontentDetails%2CtopicDetails&id="+id+"' -H 'Referer: https://mattw.io/youtube-metadata/' --compressed").Output()
	var result map[string]interface{}
	json.Unmarshal(out, &result)
	return result, nil
}

func Ytdl(url string) (map[string]interface{}, error) {
	info, _ := Ytinfo(url)
	download_mp4 := make(map[string]map[string]interface{})
	download_mp3 := make(map[string]map[string]interface{})
	for _, v := range info["links"].(map[string]interface{})["mp4"].(map[string]interface{}) {
		v := v.(map[string]interface{})
		id := fmt.Sprintf("%v", v["k"])
		vid := fmt.Sprintf("%v", info["vid"])
        quality := fmt.Sprintf("%v", v["q"])
        size := fmt.Sprintf("%v", v["size"])
		dlink, _ := Ytdownload(vid, id)
		download_mp4[quality] = map[string]interface{}{"url": dlink, "size": size}
	}
	for _, v := range info["links"].(map[string]interface{})["mp3"].(map[string]interface{}) {
		v := v.(map[string]interface{})
		id := fmt.Sprintf("%v", v["k"])
		vid := fmt.Sprintf("%v", info["vid"])
        quality := fmt.Sprintf("%v", v["q"])
        size := fmt.Sprintf("%v", v["size"])
		dlink, _ := Ytdownload(vid, id)
		download_mp3[quality] = map[string]interface{}{"url": dlink, "size": size}
	}
	for _, v := range info["links"].(map[string]interface{})["m4a"].(map[string]interface{}) {
		v := v.(map[string]interface{})
		id := fmt.Sprintf("%v", v["k"])
		vid := fmt.Sprintf("%v", info["vid"])
        quality := fmt.Sprintf("%v", v["q"])
        size := fmt.Sprintf("%v", v["size"])
		dlink, _ := Ytdownload(vid, id)
		download_mp3[quality] = map[string]interface{}{"url": dlink, "size": size}
	}		
	hasil := make(map[string]interface{})
	vid := fmt.Sprintf("%v", info["vid"])
	hsl, _ := Ytmetadata(vid)
	hasil["information"] = hsl["items"].([]interface{})[0].(map[string]interface{})["snippet"]
	hasil["mp4"] = download_mp4
	hasil["mp3"] = download_mp3 
	return hasil, nil
	}
