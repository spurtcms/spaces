package spaces

import (
	"fmt"

	"gorm.io/gorm"
)

// get pages by spaceid
func GetPageBySpaceIdORPageIds(pagereq GetPageReq, DB *gorm.DB) ([]Pages, error) {

	var pages []Pages

	pagess, _, err := Spacemodel.SelectPage(pagereq, DB)

	if err != nil {

		return []Pages{}, err
	}

	var ids []int

	for _, page := range pagess {

		ids = append(ids, page.Id)

	}

	pagelog, err := Spacemodel.GetPageLogDetails(pagereq.Spaceid, 0, ids, DB)

	if err != nil {

		fmt.Println(err)
	}

	var finallog []PageLog

	for _, val := range pagelog {

		var log PageLog

		log.Username = val.Username

		if val.ModifiedOn.IsZero() {

			log.Status = "draft"

		} else {

			log.Status = "Updated"
		}

		if val.Status == "publish" {

			log.Status = val.Status

		}

		log.Date = val.CreatedOn

		finallog = append(finallog, log)

	}

	var one_page Pages

	page_aliases, _, _ := Spacemodel.PageAliases(pagereq, DB)

	for _, val := range page_aliases {

		one_page.PageId = val.PageId

		one_page.Title = val.PageTitle

		one_page.Content = val.PageDescription

		one_page.OrderIndex = val.OrderIndex

		one_page.Pgroupid = val.PageGroupId

		one_page.ParentId = val.ParentId

		one_page.CreatedDate = val.CreatedOn

		one_page.LastUpdate = val.ModifiedOn

		one_page.Username = val.Username

		one_page.ReadTime = val.ReadTime

		one_page.Log = finallog

		pages = append(pages, one_page)
	}

	return pages, nil

}

// get page by pageid return individual page
func GetPageByPageId(pagereq GetPageReq, DB *gorm.DB) (Pages, error) {

	_, singlepage, err := Spacemodel.SelectPage(pagereq, DB)

	if err != nil {

		return Pages{}, err
	}

	pagelog, err := Spacemodel.GetPageLogDetails(pagereq.Spaceid, pagereq.PageId, []int{}, DB)

	if err != nil {

		fmt.Println(err)
	}

	var finallog []PageLog

	for _, val := range pagelog {

		var log PageLog

		log.Username = val.Username

		if val.ModifiedOn.IsZero() {

			log.Status = "draft"

		} else {

			log.Status = "Updated"
		}

		if val.Status == "publish" {

			log.Status = val.Status

		}

		log.Date = val.CreatedOn

		finallog = append(finallog, log)

	}

	pagereq.PageId = singlepage.Id

	var one_page Pages

	_, page_aliases, _ := Spacemodel.PageAliases(pagereq, DB)

	one_page.PageId = page_aliases.PageId

	one_page.Title = page_aliases.PageTitle

	one_page.Content = page_aliases.PageDescription

	one_page.OrderIndex = page_aliases.OrderIndex

	one_page.Pgroupid = page_aliases.PageGroupId

	one_page.ParentId = page_aliases.ParentId

	one_page.CreatedDate = page_aliases.CreatedOn

	one_page.LastUpdate = page_aliases.ModifiedOn

	one_page.Username = page_aliases.Username

	one_page.ReadTime = page_aliases.ReadTime

	one_page.Log = finallog

	return one_page, nil

}

// get pages by spaceid
func GetSubPageBySpaceIdORPageIds(pagereq GetPageReq, DB *gorm.DB) ([]SubPages, error) {

	var subpages []SubPages

	pagess, _, err := Spacemodel.SelectPage(pagereq, DB)

	if err != nil {

		return []SubPages{}, err
	}

	var ids []int

	for _, page := range pagess {

		ids = append(ids, page.Id)

	}

	pagelog, err := Spacemodel.GetPageLogDetails(pagereq.Spaceid, 0, ids, DB)

	if err != nil {

		fmt.Println(err)
	}

	var finallog []PageLog

	for _, val := range pagelog {

		var log PageLog

		log.Username = val.Username

		if val.ModifiedOn.IsZero() {

			log.Status = "draft"

		} else {

			log.Status = "Updated"
		}

		if val.Status == "publish" {

			log.Status = val.Status

		}

		log.Date = val.CreatedOn

		finallog = append(finallog, log)

	}

	var subpage SubPages

	page_aliases, _, _ := Spacemodel.PageAliases(pagereq, DB)

	for _, val := range page_aliases {

		subpage.SubPageId = val.PageId

		subpage.Title = val.PageTitle

		subpage.Content = val.PageDescription

		subpage.ParentId = val.ParentId

		subpage.OrderIndex = val.PageSuborder

		subpage.CreatedDate = val.CreatedOn

		subpage.LastUpdate = val.ModifiedOn

		subpage.Username = val.Username

		subpage.ReadTime = val.ReadTime

		subpage.Log = finallog

		subpages = append(subpages, subpage)

	}

	return subpages, nil

}

// get page by pageid return individual page
func GetSubPageByPageId(pagereq GetPageReq, DB *gorm.DB) (SubPages, error) {

	_, singlepage, err := Spacemodel.SelectPage(pagereq, DB)

	if err != nil {

		return SubPages{}, err
	}

	pagelog, err := Spacemodel.GetPageLogDetails(pagereq.Spaceid, pagereq.PageId, []int{}, DB)

	if err != nil {

		fmt.Println(err)
	}

	var finallog []PageLog

	for _, val := range pagelog {

		var log PageLog

		log.Username = val.Username

		if val.ModifiedOn.IsZero() {

			log.Status = "draft"

		} else {

			log.Status = "Updated"
		}

		if val.Status == "publish" {

			log.Status = val.Status

		}

		log.Date = val.CreatedOn

		finallog = append(finallog, log)

	}

	pagereq.PageId = singlepage.Id

	var one_page SubPages

	_, page_aliases, _ := Spacemodel.PageAliases(pagereq, DB)

	one_page.Title = page_aliases.PageTitle

	one_page.Content = page_aliases.PageDescription

	one_page.OrderIndex = page_aliases.OrderIndex

	one_page.ParentId = page_aliases.ParentId

	one_page.CreatedDate = page_aliases.CreatedOn

	one_page.LastUpdate = page_aliases.ModifiedOn

	one_page.Username = page_aliases.Username

	one_page.ReadTime = page_aliases.ReadTime

	one_page.Log = finallog

	return one_page, nil

}
