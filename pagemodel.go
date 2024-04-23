package spaces

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Tblpagesgroup struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	SpacesId   int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
}

type Tblpagesgroupaliases struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	PageGroupId      int       `gorm:"type:integer"`
	LanguageId       int       `gorm:"type:integer"`
	GroupName        string    `gorm:"type:character varying"`
	GroupSlug        string    `gorm:"type:character varying"`
	GroupDescription string    `gorm:"type:character varying"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	OrderIndex       int       `gorm:"type:integer"`
}

type Tblpagealiases struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	PageId           int       `gorm:"type:integer"`
	LanguageId       int       `gorm:"type:integer"`
	PageTitle        string    `gorm:"type:character varying"`
	PageSlug         string    `gorm:"type:character varying"`
	PageDescription  string    `gorm:"type:character varying"`
	PublishedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Author           string    `gorm:"type:character varying"`
	Excerpt          string    `gorm:"type:character varying"`
	FeaturedImages   string    `gorm:"type:character varying"`
	Access           string    `gorm:"type:character varying"`
	MetaDetails      datatypes.JSONType[MetaDetails]
	Status           string `gorm:"type:character varying"`
	AllowComments    bool
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	OrderIndex       int       `gorm:"type:integer"`
	PageSuborder     int       `gorm:"type:integer"`
	CreatedDate      string    `gorm:"-"`
	ModifiedDate     string    `gorm:"-"`
	Username         string    `gorm:"<-:false"`
	PageGroupId      int       `gorm:"-:migration;<-:false"`
	ParentId         int       `gorm:"-:migration;<-:false"`
	LastRevisionDate time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	LastRevisionNo   int       `gorm:"type:integer"`
	ReadTime         int       `gorm:"type:integer"`
}

type Tblpagealiaseslog struct {
	Id              int       `gorm:"primaryKey;auto_increment"`
	PageId          int       `gorm:"type:integer"`
	LanguageId      int       `gorm:"type:integer"`
	PageTitle       string    `gorm:"type:character varying"`
	PageSlug        string    `gorm:"type:character varying"`
	PageDescription string    `gorm:"type:character varying"`
	PublishedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Author          string    `gorm:"type:character varying"`
	Excerpt         string    `gorm:"type:character varying"`
	FeaturedImages  string    `gorm:"type:character varying"`
	Access          string    `gorm:"type:character varying"`
	MetaDetails     datatypes.JSONType[MetaDetails]
	Status          string `gorm:"type:character varying"`
	AllowComments   bool
	CreatedOn       time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy       int       `gorm:"type:integer"`
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	CreatedDate     string    `gorm:"-"`
	ModifiedDate    string    `gorm:"-"`
	Username        string    `gorm:"-:migration;<-:false"`
	PageGroupId     int       `gorm:"-:migration;<-:false"`
	ParentId        int       `gorm:"-:migration;<-:false"`
	ReadTime        int       `gorm:"type:integer"`
}

type PageCreate struct {
	SpaceId       int          //spaceid
	NewPages      []Pages      //pages only
	NewGroup      []PageGroups //groups only
	NewSubPage    []SubPages   //subpages only
	UpdatePages   []Pages      //pages only
	UpdateGroup   []PageGroups //groups only
	UpdateSubPage []SubPages   //subpages only
	DeletePages   []Pages      //delete pages only
	DeleteGroup   []PageGroups //delete groups only
	DeleteSubPage []SubPages   //delete subpages only
	Status        string       //publish,draft
}

type PageLog struct {
	Username string
	Status   string
	Date     time.Time
}

type PageGroups struct {
	GroupId    int
	NewGroupId int
	Name       string
	OrderIndex int `json:"OrderIndex"`
}

type Pages struct {
	PgId        int
	NewPgId     int
	Name        string
	Content     string `json:"Content"`
	Pgroupid    int
	NewGrpId    int
	OrderIndex  int `json:"OrderIndex"`
	ParentId    int
	ReadTime    int
	CreatedDate time.Time
	LastUpdate  time.Time
	Status      string
	Date        string
	Username    string
	Log         []PageLog
}

type SubPages struct {
	SpgId       int
	NewSpId     int
	Name        string
	Content     string
	ParentId    int
	NewParentId int
	PgroupId    int
	NewPgroupId int
	ReadTime    int
	OrderIndex  int `json:"OrderIndex"`
	CreatedDate time.Time
	LastUpdate  time.Time
	Status      string
	Date        string
	Username    string
	Log         []PageLog
}

// pass any one only-- (ids,id,groupids,groupid,spaceid,spaceids)
type DeletePagereq struct {
	Ids       []int //bulk delete using id
	Id        int   //individual id
	GroupIds  []int //bulk group child pages delete
	GroupId   int   //individual group pages delete
	SpaceId   int   //delete page using spaceid
	SpaceIds  []int //buik spaces pages delete
	DeletedBy int
}

// pass any one only-- (groupids,groupid,spaceid,spaceids)
type DeletePageGroupreq struct {
	GroupIds  []int //bulk group child group delete
	GroupId   int   //individual group delete
	SpaceId   int   //delete pagegroup using spaceid
	SpaceIds  []int //bulk pagesgroup delete using spaceid
	DeletedBy int
}

// delete page group
func (SpaceModel) DeletePageGroup(tblpage *TblPagesGroup, id int, DB *gorm.DB) error {

	if err := DB.Model(TblPagesGroup{}).Where("tbl_pages_groups.id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": tblpage.IsDeleted, "deleted_on": tblpage.DeletedOn, "deleted_by": tblpage.DeletedBy}).Error; err != nil {

		return err
	}

	return nil
}

func (SpaceModel) GetPageDetailsBySpaceId(getpg *[]TblPage, id int, DB *gorm.DB) (*[]TblPage, error) {

	if err := DB.Table("tbl_pages").Where("tbl_pages.is_deleted = ? and tbl_pages.spaces_id = ?", 0, id).Find(&getpg).Error; err != nil {

		return &[]TblPage{}, err
	}

	return getpg, nil
}
