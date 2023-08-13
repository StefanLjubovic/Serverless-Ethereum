package dto

type AddVideoDTO struct {
	SectionName string `json:"section_name"`
	VideoName   string `json:"video_name"`
	CourseID    int    `json:"course_id"`
	VideoPath   string `json:"video_path"`
	Length      int    `json:"length"`
}
