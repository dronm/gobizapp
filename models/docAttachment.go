package models

import (
	"time"
	// fields "github.com/dronm/crudifier/fields"
)

// type DocAttachment struct {
// 	Id              int                    `json:"id" primaryKey:"true" srvCalc:"true"`
// 	Date_time       fields.FieldDateTimeTZ `json:"date_time"`
// 	Ref             fields.FieldText       `json:"ref"`
// 	Content_info    fields.FieldText       `json:"content_info"`
// 	Content_data    fields.FieldText       `json:"content_data"`
// 	Content_preview fields.FieldText       `json:"content_preview"`
// }

type DocAttachmentContentInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type DocAttachment struct {
	ID             int                      `json:"id" primaryKey:"true" srvCalc:"true"`
	DateTime       time.Time                `json:"date_time"`
	Ref            Ref                      `json:"ref"`
	ContentInfo    DocAttachmentContentInfo `json:"content_info"`
	ContentData    []byte                   `json:"content_data"`
	ContentPreview []byte                   `json:"content_preview"`
}

type DocAttachment_get_file struct {
	Ref       Ref    `form:"ref" json:"ref" binding:"required"`
	ContentId string `form:"content_id" json:"content_id" binding:"required"`
	Inline    bool   `form:"inline" json:"inline" binding:"required"`
}

type DocAttachment_delete_file struct {
	Ref       Ref    `form:"ref" json:"ref" required:"true"`
	ContentId string `form:"content_id" json:"content_id" binding:"required"`
}
