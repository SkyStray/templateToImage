package entity

type Data struct {
	Title   string `json:"Title" comment:"标题"`
	Name    string `json:"Name" comment:"名称"`
	Content string `json:"Content" comment:"正文"`
	Text    string `json:"Text" comment:"正文"`
	Year    string `json:"Year" comment:"年"`
	Month   string `json:"Month" comment:"月"`
	Day     string `json:"Day" comment:"日"`
	No      string `json:"No" comment:"编号"`
}
