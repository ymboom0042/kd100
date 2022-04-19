/**
 * @Author: YMBoom
 * @Description:
 * @File:  err
 * @Version: 1.0.0
 * @Date: 2022/04/19 13:24
 */
package mian

import "errors"

func newErr(text string) error {
	return errors.New(text)
}
