package trend

type SearchKeyword struct {
	StartDate              string `json:"startDate"`             //: "20220512",
	EndDate                string `json:"endDate"`               //": "20220611",
	TopN                   int    `json:"topN"`                  //": 500,
	Period                 string `json:"period"`                //": "1",
	AnalysisMonths         int    `json:"analysisMonths"`        //": 0,
	CategorySetName        string `json:"categorySetName"`       //": "SMT",
	Sources                string `json:"sources"`               //": "blog,news,twitter",
	Keyword                string `json:"keyword"`               //": "이정현",
	Synonym                string `json:"synonym"`               //": null,
	KeywordFilterIncludes  string `json:"keywordFilterIncludes"` //": null,
	KeyworkdFilterExcludes string `json:"keywordFilterExclude"`  //s": null,
	IncludeWordOperatros   string `json:"includeWordOperators"`  //": "||",
	ExcludeWordOperators   string `json:"excludeWordOperators"`  //": "||",
	ScoringKeyWord         string `json:"scoringKeyword"`        //": "",
	ExForHash              string `json:"exForHash"`             //": "",
	CategoryList           string `json:"categoryList"`          //": "politician,celebrity,sportsman,characterEtc,government,business,agency,groupEtc,tourism,restaurant,shopping,scene,placeEtc,brandFood,cafe,brandBeverage,brandElectronics,brandFurniture,brandBeauty,brandFashion,brandEtc,productFood,productBeverage,productElectronics,productFurniture,productBeauty,productFashion,productEtc,economy,social,medicine,education,culture,sports,cultureEtc,animal,plant,naturalPhenomenon,naturalEtc"
}

type Association struct {
	Item   Item   `json:"item"`
	Code   string `json:"code"`
	Errors string `json:"errors"`
	Error  string `json:"error"`
}

type Item struct {
	DataList []DataList `json:"dataList"`
	Keyword  string     `json:"keyworkd"`
}
type DataList struct {
	Source string `json:"source"`
	Data   Data   `json:"data"`
}
type Data struct {
	Rows        []Row  `json:"rows"`
	CategoryMap string `json:"categoryMap"`
}
type Row struct {
	date            string            `json:"date"`
	AssociationData []AssociationData `json:"associationData"`
}
type AssociationData struct {
	Label     string `json:"label"`
	Frequency int    `json:"frequency"`
}
