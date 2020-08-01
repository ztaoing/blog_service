/**
* @Author:zhoutao
* @Date:2020/8/1 上午8:49
* @Desc:
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
)

//对上传后的文件名进行格式化
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	//Sum appends the current hash to b and returns the resulting slice
	return hex.EncodeToString(m.Sum(nil))
}
