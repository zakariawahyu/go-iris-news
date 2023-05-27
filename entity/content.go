package entity

import "time"

type Content struct {
	ID                 int64                 `db:"id,primary" json:"id"`
	Type               string                `db:"type" json:"type"`
	TypeID             *int64                `db:"type_id" json:"type_id"`
	TypeChildID        *int64                `db:"type_child_id" json:"type_child_id"`
	SuplemenID         *int64                `db:"suplemen_id" json:"suplemen_id"`
	UpperTitle         string                `db:"uppertitle" json:"upper_title"`
	Title              string                `db:"title" json:"title"`
	Slug               string                `db:"slug" json:"slug"`
	Excerpt            string                `db:"excerpt" json:"excerpt"`
	Content            string                `db:"content" json:"content"`
	Caption            *string               `db:"caption" json:"caption"`
	Image              *string               `db:"image" json:"image"`
	Thumbnail          *string               `db:"thumbnail" json:"thumbnail"`
	HeadlineType       int64                 `db:"headline_type" json:"headline_type"`
	IsActive           bool                  `db:"is_active" json:"is_active"`
	IsFeatured         bool                  `db:"is_featured" json:"is_featured"`
	IsEditorChoice     bool                  `db:"is_editor_choice" json:"is_editor_choice"`
	IsNational         bool                  `db:"is_national" json:"is_national"`
	Token              string                `db:"token" json:"token"`
	AdsPosition        int                   `db:"ads_position" json:"ads_position"`
	AdsExpiredDate     time.Time             `db:"ads_expired_date" json:"ads_expired_date"`
	UrlVideo           *string               `db:"url_video" json:"url_video"`
	IsPopular          bool                  `db:"is_popular" json:"is_popular"`
	IsAdult            bool                  `db:"is_adult" json:"is_adult"`
	IsComment          bool                  `db:"is_comment" json:"is_comment"`
	LocationCity       *string               `db:"location_city" json:"location_city"`
	LocationDistrict   *string               `db:"location_citydistrict" json:"location_district"`
	LocationSuburb     *string               `db:"location_suburb" json:"location_suburb"`
	CreatedBy          int64                 `db:"created_by" json:"created_by"`
	User               *UserResponse         `ref:"created_by" fk:"id" json:"user"`
	Region             *RegionResponse       `ref:"type_id" fk:"id" json:"region"`
	Channel            *ChannelResponse      `ref:"type_id" fk:"id" json:"channel"`
	SubChannel         *SubChannelResponse   `ref:"type_child_id" fk:"id" json:"sub_channel"`
	ContentHasTag      []*ContentHasTag      `ref:"id" fk:"content_id" json:"tags"`
	ContentHasTopic    []*ContentHasTopic    `ref:"id" fk:"content_id" json:"topics"`
	ContentHasReporter []*ContentHasReporter `ref:"id" fk:"content_id" json:"reporters"`
	SubPhoto           []*SubPhoto           `ref:"id" fk:"content_id" json:"sub_photo"`
	PublishedDate      time.Time             `db:"published_date" json:"published_date"`
	CreatedAt          time.Time             `db:"created" json:"created_at"`
	UpdatedAt          *time.Time            `db:"modified" json:"updated_at"`
}

type ContentResponse struct {
	ID                 int64                 `db:"id,primary" json:"id"`
	Type               string                `db:"type" json:"type"`
	TypeID             *int64                `db:"type_id" json:"type_id"`
	TypeChildID        *int64                `db:"type_child_id" json:"type_child_id"`
	SuplemenID         *int64                `db:"suplemen_id" json:"suplemen_id"`
	UpperTitle         string                `db:"uppertitle" json:"upper_title"`
	Title              string                `db:"title" json:"title"`
	Slug               string                `db:"slug" json:"slug"`
	Excerpt            string                `db:"excerpt" json:"excerpt"`
	Content            string                `db:"content" json:"content"`
	Caption            *string               `db:"caption" json:"caption"`
	Image              *string               `db:"image" json:"image"`
	Thumbnail          *string               `db:"thumbnail" json:"thumbnail"`
	UrlVideo           *string               `db:"url_video" json:"url_video"`
	IsPopular          bool                  `db:"is_popular" json:"is_popular"`
	IsAdult            bool                  `db:"is_adult" json:"is_adult"`
	IsComment          bool                  `db:"is_comment" json:"is_comment"`
	CreatedBy          int64                 `db:"created_by" json:"created_by"`
	User               *UserResponse         `ref:"created_by" fk:"id" json:"user"`
	Region             *RegionResponse       `ref:"type_id" fk:"id" json:"region"`
	Channel            *ChannelResponse      `ref:"type_id" fk:"id" json:"channel"`
	SubChannel         *SubChannelResponse   `ref:"type_child_id" fk:"id" json:"sub_channel"`
	ContentHasTag      []*ContentHasTag      `ref:"id" fk:"content_id" json:"tags"`
	ContentHasTopic    []*ContentHasTopic    `ref:"id" fk:"content_id" json:"topics"`
	ContentHasReporter []*ContentHasReporter `ref:"id" fk:"content_id" json:"reporters"`
	SubPhoto           []*SubPhoto           `ref:"id" fk:"content_id" json:"sub_photo"`
	PublishedDate      time.Time             `db:"published_date" json:"published_date"`
	UpdatedAt          *time.Time            `db:"modified" json:"updated_at"`
}

func (u Content) Table() string {
	return "contents_new"
}

func (u ContentResponse) Table() string {
	return "contents_new"
}

func NewContentResponse(content Content) ContentResponse {
	return ContentResponse{
		ID:                 content.ID,
		Type:               content.Type,
		TypeID:             content.TypeID,
		TypeChildID:        content.TypeChildID,
		SuplemenID:         content.SuplemenID,
		UpperTitle:         content.UpperTitle,
		Title:              content.Title,
		Slug:               content.Slug,
		Excerpt:            content.Excerpt,
		Content:            content.Content,
		Caption:            content.Caption,
		Image:              content.Image,
		Thumbnail:          content.Thumbnail,
		UrlVideo:           content.UrlVideo,
		IsPopular:          content.IsPopular,
		IsAdult:            content.IsAdult,
		IsComment:          content.IsComment,
		CreatedBy:          content.CreatedBy,
		User:               content.User,
		Region:             content.Region,
		Channel:            content.Channel,
		SubChannel:         content.SubChannel,
		ContentHasTag:      content.ContentHasTag,
		ContentHasTopic:    content.ContentHasTopic,
		ContentHasReporter: content.ContentHasReporter,
		SubPhoto:           content.SubPhoto,
		PublishedDate:      content.PublishedDate,
		UpdatedAt:          content.UpdatedAt,
	}
}
