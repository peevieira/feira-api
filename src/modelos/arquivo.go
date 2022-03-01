package modelos

type CampoArquivo struct {
	ID         int64   `json:"ID"`
	Long       float64 `json:"LONG"`
	Lat        float64 `json:"LAT"`
	SetCens    int64   `json:"SETCENS"`
	AreaP      float64 `json:"AREAP"`
	CodDist    int64   `json:"CODDIST"`
	Distrito   string  `json:"DISTRITO"`
	CodSubPref int64   `json:"CODSUBPREF"`
	SubPrefe   string  `json:"SUBPREFE"`
	Regiao5    string  `json:"REGIAO5"`
	Regiao8    string  `json:"REGIAO8"`
	Nome_Feira string  `json:"NOME_FEIRA"`
	Registro   string  `json:"REGISTRO"`
	Logradouro string  `json:"LOGRADOURO"`
	Numero     string  `json:"Numero"`
	Bairro     string  `json:"Bairro"`
	Referencia string  `json:"REFERENCIA"`
}
