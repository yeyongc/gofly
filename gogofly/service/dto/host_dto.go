package dto

type ShutdownHostDTO struct {
	HostIP string `json:"host_ip" form:"host_ip" xml:"host_ip" binding:"required" message:"Host ip is valid"`
}
