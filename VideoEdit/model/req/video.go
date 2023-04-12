// @User CPR
package req

type EditVideo struct {
	Start    string `json:"start" validate:"required" label:"开始时间"`
	Duration string `json:"duration" validate:"required" label:"持续时间"`
}
