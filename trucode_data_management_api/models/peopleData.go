package models

import (
	"gorm.io/gorm"
)

type Workclass string
type Education string
type MaritalStatus string
type Occupation string
type Relationship string
type Race string
type Sex string
type Income string

const (
	// Workclass options
	WorkclassUnknown Workclass = "?"
	FederalGov       Workclass = "Federal-gov"
	LocalGov         Workclass = "Local-gov"
	NeverWorked      Workclass = "Never-worked"
	Private          Workclass = "Private"
	SelfEmpInc       Workclass = "Self-emp-inc"
	SelfEmpNotInc    Workclass = "Self-emp-not-inc"
	StateGov         Workclass = "State-gov"
	WithoutPay       Workclass = "Without-pay"

	// Education options
	EducationUnknown Education = "?"
	Education10th    Education = "10th"
	Education11th    Education = "11th"
	Education12th    Education = "12th"
	Education1st4th  Education = "1st-4th"
	Education5th6th  Education = "5th-6th"
	Education7th8th  Education = "7th-8th"
	Education9th     Education = "9th"
	AssocAcdm        Education = "Assoc-acdm"
	AssocVoc         Education = "Assoc-voc"
	Bachelors        Education = "Bachelors"
	Doctorate        Education = "Doctorate"
	HSGrad           Education = "HS-grad"
	Masters          Education = "Masters"
	Preschool        Education = "Preschool"
	ProfSchool       Education = "Prof-school"
	SomeCollege      Education = "Some-college"

	// MaritalStatus options
	MaritalStatusUnknown MaritalStatus = "?"
	Divorced             MaritalStatus = "Divorced"
	MarriedAFSpouse      MaritalStatus = "Married-AF-spouse"
	MarriedCivSpouse     MaritalStatus = "Married-civ-spouse"
	MarriedSpouseAbsent  MaritalStatus = "Married-spouse-absent"
	NeverMarried         MaritalStatus = "Never-married"
	Separated            MaritalStatus = "Separated"
	Widowed              MaritalStatus = "Widowed"

	// Occupation options
	OccupationUnknown Occupation = "?"
	AdmClerical       Occupation = "Adm-clerical"
	ArmedForces       Occupation = "Armed-Forces"
	CraftRepair       Occupation = "Craft-repair"
	ExecManagerial    Occupation = "Exec-managerial"
	FarmingFishing    Occupation = "Farming-fishing"
	HandlersCleaners  Occupation = "Handlers-cleaners"
	MachineOpInspct   Occupation = "Machine-op-inspct"
	OtherService      Occupation = "Other-service"
	PrivHouseServ     Occupation = "Priv-house-serv"
	ProfSpecialty     Occupation = "Prof-specialty"
	ProtectiveServ    Occupation = "Protective-serv"
	Sales             Occupation = "Sales"
	TechSupport       Occupation = "Tech-support"
	TransportMoving   Occupation = "Transport-moving"

	// Relationship options
	RelationshipUnknown Relationship = "?"
	Husband             Relationship = "Husband"
	NotInFamily         Relationship = "Not-in-family"
	OtherRelative       Relationship = "Other-relative"
	OwnChild            Relationship = "Own-child"
	Unmarried           Relationship = "Unmarried"
	Wife                Relationship = "Wife"

	// Race options
	RaceUnknown      Race = "?"
	AmerIndianEskimo Race = "Amer-Indian-Eskimo"
	AsianPacIslander Race = "Asian-Pac-Islander"
	Black            Race = "Black"
	Other            Race = "Other"
	White            Race = "White"

	// Sex options
	SexUnknown Sex = "?"
	Female     Sex = "Female"
	Male       Sex = "Male"

	// Income options
	IncomeUnknown    Income = "?"
	LessThan50k      Income = "<=50K"
	MoreThan50k      Income = ">50K"
	MoreThanEqual50k Income = ">=50K"
)

type PeopleData struct {
	gorm.Model    `json:"-"`
	Age           int           `json:"age"`
	Workclass     Workclass     `json:"workclass"`
	Fnlwgt        int           `json:"fnlwgt"`
	Education     Education     `json:"education"`
	EducationNum  int           `json:"education_num"`
	MaritalStatus MaritalStatus `json:"marital_status"`
	Occupation    Occupation    `json:"occupation"`
	Relationship  Relationship  `json:"relationship"`
	Race          Race          `json:"race"`
	Sex           Sex           `json:"sex"`
	CapitalGain   int           `json:"capital_gain"`
	CapitalLoss   int           `json:"capital_loss"`
	HoursPerWeek  int           `json:"hours_per_week"`
	NativeCountry string        `json:"native_country"`
	Income        Income        `json:"income"`
}

type PeopleDataFrontend struct {
	ID            uint   `json:"id"`
	Age           int    `json:"age"`
	Education     string `json:"education"`
	MaritalStatus string `json:"marital_status"`
	Occupation    string `json:"occupation"`
	Income        string `json:"income"`
}

// Funciones helper para obtener todas las opciones
func GetAllWorkclassOptions() []string {
	return []string{
		string(WorkclassUnknown),
		string(FederalGov),
		string(LocalGov),
		string(NeverWorked),
		string(Private),
		string(SelfEmpInc),
		string(SelfEmpNotInc),
		string(StateGov),
		string(WithoutPay),
	}
}

func GetAllEducationOptions() []string {
	return []string{
		string(EducationUnknown),
		string(Education10th),
		string(Education11th),
		string(Education12th),
		string(Education1st4th),
		string(Education5th6th),
		string(Education7th8th),
		string(Education9th),
		string(AssocAcdm),
		string(AssocVoc),
		string(Bachelors),
		string(Doctorate),
		string(HSGrad),
		string(Masters),
		string(Preschool),
		string(ProfSchool),
		string(SomeCollege),
	}
}

func GetAllMaritalStatusOptions() []string {
	return []string{
		string(MaritalStatusUnknown),
		string(Divorced),
		string(MarriedAFSpouse),
		string(MarriedCivSpouse),
		string(MarriedSpouseAbsent),
		string(NeverMarried),
		string(Separated),
		string(Widowed),
	}
}

func GetAllOccupationOptions() []string {
	return []string{
		string(OccupationUnknown),
		string(AdmClerical),
		string(ArmedForces),
		string(CraftRepair),
		string(ExecManagerial),
		string(FarmingFishing),
		string(HandlersCleaners),
		string(MachineOpInspct),
		string(OtherService),
		string(PrivHouseServ),
		string(ProfSpecialty),
		string(ProtectiveServ),
		string(Sales),
		string(TechSupport),
		string(TransportMoving),
	}
}

func GetAllRelationshipOptions() []string {
	return []string{
		string(RelationshipUnknown),
		string(Husband),
		string(NotInFamily),
		string(OtherRelative),
		string(OwnChild),
		string(Unmarried),
		string(Wife),
	}
}

func GetAllRaceOptions() []string {
	return []string{
		string(RaceUnknown),
		string(AmerIndianEskimo),
		string(AsianPacIslander),
		string(Black),
		string(Other),
		string(White),
	}
}

func GetAllSexOptions() []string {
	return []string{
		string(SexUnknown),
		string(Female),
		string(Male),
	}
}

func GetAllIncomeOptions() []string {
	return []string{
		string(IncomeUnknown),
		string(LessThan50k),
		string(MoreThan50k),
		string(MoreThanEqual50k),
	}
}
