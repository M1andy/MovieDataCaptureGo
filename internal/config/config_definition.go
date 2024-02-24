package config

type Config struct {
	Main          Main          `toml:"main"`
	Actor         Actor         `toml:"actor"`
	DownloadFiles DownloadFiles `toml:"download_files"`
	Translation   Translation   `toml:"translation"`
	Proxy         Proxy         `toml:"proxy"`
	RenameRule    RenameRule    `toml:"rename_rule"`
	Update        Update        `toml:"update"`
	LoggerOptions LoggerOptions `toml:"logger"`
}
type Main struct {
	Mode                   string `toml:"mode"`
	SourceDirectory        string `toml:"source_directory"`
	SuccessOutputDirectory string `toml:"success_output_directory"`
	FailedOutputDirectory  string `toml:"failed_output_directory"`
	FailedMove             bool   `toml:"failed_move"`
}
type Actor struct {
	ActorGender string `toml:"actor_gender"`
}
type ThumbWatermark struct {
	Switch   bool   `toml:"switch"`
	Position string `toml:"position"`
}
type DownloadFiles struct {
	Poster         bool           `toml:"poster"`
	Nfo            bool           `toml:"nfo"`
	Thumb          bool           `toml:"thumb"`
	Fanart         bool           `toml:"fanart"`
	ExtraFanart    bool           `toml:"extra_fanart"`
	Trailer        bool           `toml:"trailer"`
	ThumbWatermark ThumbWatermark `toml:"thumb_watermark"`
}
type Translation struct {
	Switch         bool     `toml:"switch"`
	Engine         string   `toml:"engine"`
	TargetLanguage string   `toml:"target_language"`
	Text           []string `toml:"text"`
}
type Proxy struct {
	Switch  bool     `toml:"switch"`
	URL     []string `toml:"url"`
	Timeout int      `toml:"timeout"`
	Retry   int      `toml:"retry"`
}
type RenameRule struct {
	LocationRule        string `toml:"location_rule"`
	FileRule            string `toml:"file_rule"`
	RenameImgWithNumber bool   `toml:"rename_img_with_number"`
}
type Update struct {
	UpdateCheck bool `toml:"update_check"`
}
type LoggerOptions struct {
	Level   string `toml:"level"`
	LogPath string `toml:"log_path"`
}
