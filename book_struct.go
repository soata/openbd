package openbd

import "encoding/json"

const tagName = "explain"

const (
	textTypeBrief       = "02"
	textTypeDescription = "03"
	textTypeToc         = "04"
)

// Book :struct to map whole response
type Book struct {
	Onix    `json:"onix"`    // JPRO-onix準拠項目
	Hanmoto `json:"hanmoto"` // 版元ドットコム独自項目
	Summary `json:"summary"` // 基本事項
}

// Onix :JPRO-onix準拠項目
type Onix struct {
	RecordReference string // ISBNコード
	ProductIdentifier
	DescriptiveDetail
	Measure
	CollateralDetail
	PublishingDetail
	NotificationType string
	ProductSupply
}

type ProductIdentifier struct {
	IDValue       string // ISBNコード
	ProductIDType string // ISBNコード
}

type Measure []struct { // 判型(実寸)
	MeasureType     string
	MeasureUnitCode string
	Measurement     string
}

type ProductPart []struct { // 付録情報
	NumberOfItemsOfThisForm json.Number
	ProductForm             string
	ProductFormDescription  string
}

type TitleText struct {
	Content      string `json:"content"`
	CollationKey string `json:"collationkey"`
}
type TitleElement []struct {
	TitleElementLevel string
	TitleText
}
type TitleDetail struct { // シリーズ名/レーベル名
	TitleType string
	TitleElement
}
type Collection struct {
	TitleDetail
	CollectionType string
}

type PersonName struct { // 著者名
	Content      string // 氏名
	Collationkey string // 照合キー
}
type Contributor []struct { // 著者情報
	ContributorRole []string // 著者区分
	PersonName
	BiographicalNote string // 著者略歴
	SequenceNumber   json.Number
}
type Language []struct { // 言語
	LanguageCode string
	LanguageRole string
	CountryCode  string
}
type Extent []struct { // ページ数
	ExtentValue json.Number
	ExtentUnit  string
	ExtentType  string
}
type Subject []struct { // Cコード/ジャンルコード/キーワード
	SubjectSchemeIdentifier string
	SubjectCode             string
	SubjectHeadingText      string
}
type Audience []struct { // 読者対象/成人指定
	AudienceCodeType  string
	AudienceCodeValue string
}

type TextContent []struct { // 内容紹介/目次
	Text            string
	TextType        string
	ContentAudience string
}

type ResourceVersionFeature []struct {
	ResourceVersionFeatureType string
	FeatureValue               string
}
type ResourceVersion []struct {
	ResourceLink string // 画像パス
	ResourceForm string
	ResourceVersionFeature
}
type SupportingResource []struct { // 画像
	ResourceContentType string // ISBNコード
	ResourceMode        string
	ContentAudience     string
	ResourceVersion
}
type CollateralDetail struct { // 商品付帯項目
	TextContent

	SupportingResource
}

type DescriptiveDetail struct { // シリーズ物等の情報
	ProductComposition string // セット商品分売可否
	ProductFormDetail  string // 判型
	Measure
	ProductPart
	Collection
	TitleDetail
	Contributor

	Language

	Extent

	ProductForm string
	Subject

	Audience
}

type Imprint struct { // 発行元出版社
	ImprintName string
	ImprintIdentifier
}
type ImprintIdentifier []struct {
	IDValue       string
	ImprintIDType string
}
type PublisherIdentifier []struct {
	PublisherIDType string
	IDValue         string
}
type Publisher struct { // 発売元出版社
	PublisherIdentifier

	PublisherName  string
	PublishingRole string
}
type PublishingDate []struct { // 発売予定日等
	Date               string
	PublishingDateRole string
}
type PublishingDetail struct { // 出版社情報
	Imprint

	Publisher

	PublishingDate
}

type Price []struct { // 価格情報
	PriceAmount  json.Number
	CurrencyCode string
	PriceType    string
}
type ReturnsConditions struct { // 販売条件
	ReturnsCode     string
	ReturnsCodeType string
}

type SupplyDetail struct {
	ReturnsConditions
	ProductAvailability string
	Price
}
type ProductSupply struct { // 出版状況等
	SupplyDetail
}

type Reviews []struct { // 書評情報
	Han        string // 版数
	Appearance string // 掲載日
	PostUser   string `json:"post_user"` // 入力者
	KubunID    int    `json:"kubun_id"`  // 掲載元区分ID
	Source     string // 掲載元
	Choyukan   string // 朝刊/夕刊
	SourceID   int    `json:"source_id"` // 掲載元ID
	Reviewer   string // 書評者
	Link       string // リンク
}

type Author []struct { // 独自著者情報
	Dokujikubun string // 独自著者区分
	Listseq     int    // リスト順
}

type Jyuhan []struct { // 重版情報
	Date    string // 重版日
	Ctime   string // レコード作成日時
	Suri    int    // 刷り数
	Comment string // 重版コメント
}

// Hanmoto :版元ドットコム独自項目
type Hanmoto struct {
	Dateshuppan    string // 出版年月日
	Datemodified   string // 情報更新日時
	Datecreated    string // 情報作成日時
	Lanove         bool   // ライトノベルフラグ
	Hasshohyo      string // 書評データありフラグ
	Hastameshiyomi bool   // ためしよみありフラグ
	Reviews

	// 版元ドットコム会員社独自項目 ->
	Genshomei                       string // 原書名
	Han                             string // 版"` // 新版、改訂などの版次(エディショ
	Datejuuhanyotei                 string // 重版予定日
	Datezeppan                      string // 絶版日
	Toji                            string // 製本"`  // 上製="上製",並製="並
	Zaiko                           int    // 在庫"` // 在庫あり=11／在庫僅少=21／重版中=／品切れ・重版未定=33／絶版=34
	Maegakinado                     string // まえがきなど
	Hanmotokarahitokoto             string // 版元からひとこと
	Kaisetsu105W                    string // 内容紹介TRC(105字)
	Kanrensho                       string // 関連書
	KanrenLink                      string // 関連リンク	htmlで記入
	Tsuiki                          string // サイト追記
	Genrecodetrc                    int    // ジャンルコード
	Genrecodetrcjidou               string // ジャンルコード(TRC児童書)
	Rubynoumu                       string // ルビの有無
	Ndccode                         string // NDCコード
	Kankoukeitai                    string // 刊行形態
	Sonotatokkijikou                string // その他特記事項
	Jushoujouhou                    string // 受賞情報
	Furoku                          string // 付録
	Furokusonota                    string // 付録その他
	Dokushakakikomi                 string // 読者書き込み
	Dokushakakikomipagesuu          int    // 読者書き込みページ数
	Fuzokushiryounokangaikashidashi int    // 付属資料館外貸出 可=1,不可=2,館内のみ=3
	Obinaiyou                       string // 帯内容
	Ruishokyougousho                string // 類書・競合書
	Bessoushiryou                   string // 別送資料
	Zasshicode                      string // 雑誌コード
	Bikoujpo                        string // 備考(JPO)
	Bikoutrc                        string // 備考(TRC)
	Kanrenshoisbn                   string // 関連書ISBN
	Author

	Jyuhan
	// 版元ドットコム会員社独自項目 <-
}

// Summary :使いやすい基本事項
type Summary struct {
	Publisher string
	ISBN      string
	Pubdate   string
	Title     string
	Series    string
	Author    string
	Cover     string
	Volume    string
}
