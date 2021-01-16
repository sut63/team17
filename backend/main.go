package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sut63/team17/app/controllers"
	_ "github.com/sut63/team17/app/docs"
	"github.com/sut63/team17/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Terms struct {
	Term []Term
}
type Term struct {
	Term int
}
type Years struct {
	Year []Year
}
type Year struct {
	Years int
}
type Subjects struct {
	Subject []Subject
}
type Subject struct {
	Code        int
	Subjects    string
	Creditpiont int
}
type Agencys struct {
	Agency []Agency
}
type Agency struct {
	Agency string
}
type Places struct {
	Place []Place
}
type Place struct {
	Place string
}
type Prefixs struct {
	Prefix []Prefix
}

type Prefix struct {
	Prefix string
}

type Facultys struct {
	Faculty []Faculty
}

type Faculty struct {
	Faculty string
}

type Professorships struct {
	Professorship []Professorship
}

type Professorship struct {
	Professorship string
}
type Degrees struct {
	Degree []Degree
}
type Degree struct {
	Degree string
}
type Institutions struct {
	Institution []Institution
}
type Institution struct {
	Institution string
}
type Genders struct {
	Gender []Gender
}
type Gender struct {
	Gender string
}
type Students struct {
	Student []Student
}
type Student struct {
	Fname  string
	Lname  string
	Addr   string
	Email  string
	School string
	Tel    int
}
type Countrys struct {
	Country []Country
}
type Country struct {
	Country string
}
type Continents struct {
	Continent []Continent
}
type Continent struct {
	Continent string
}
type Regions struct {
	Region []Region
}
type Region struct {
	Region string
}
type Provinces struct {
	Province []Province
}
type Province struct {
	Province    string
	District    string
	Subdistrict string
	Postal      int
}

type Emps struct {
	Emp []Emp
}
type Emp struct {
	User string
	Pass string
	Role string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
//comment eiei
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewTermController(v1, client)
	controllers.NewYearController(v1, client)
	controllers.NewSubjectController(v1, client)
	controllers.NewResultsController(v1, client)
	controllers.NewActivityController(v1, client)
	controllers.NewAgencyController(v1, client)
	controllers.NewPlaceController(v1, client)
	controllers.NewCourseController(v1, client)
	controllers.NewDegreeController(v1, client)
	controllers.NewFacultyController(v1, client)
	controllers.NewContinentController(v1, client)
	controllers.NewProvinceController(v1, client)
	controllers.NewRegionController(v1, client)
	controllers.NewCountryController(v1, client)
	controllers.NewPrefixController(v1, client)
	controllers.NewProfessorController(v1, client)
	controllers.NewProfessorshipController(v1, client)
	controllers.NewStudentController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewEmpController(v1, client)
	controllers.NewInstitutionController(v1, client)

	// Set Genders Data
	Genders := Genders{
		Gender: []Gender{
			Gender{"Male"},
			Gender{"Female"},
		},
	}

	for _, u := range Genders.Gender {
		client.Gender.
			Create().
			SetGender(u.Gender).
			Save(context.Background())
	}

	// Set Genders Data
	Students := Students{
		Student: []Student{
			Student{"Max", "Alask", "west", "A", "De", 1},
			Student{"Tom", "Ronwe", "north", "B", "Ce", 2},
		},
	}

	for _, u := range Students.Student {
		client.Student.
			Create().
			SetFname(u.Fname).
			SetLname(u.Lname).
			SetEmail(u.Email).
			SetRecentAddress(u.Addr).
			SetSchoolname(u.School).
			SetTelephone(u.Tel).
			Save(context.Background())
	}

	// Set Term Data
	terms := Terms{
		Term: []Term{
			Term{1},
			Term{2},
			Term{3},
		},
	}

	for _, rt := range terms.Term {
		client.Term.
			Create().
			SetSemester(rt.Term).
			Save(context.Background())
	}

	// Set Subject Data
	subjects := Subjects{
		Subject: []Subject{
			Subject{111111, "วิชาท่องนภาขั้น1", 2},
			Subject{111112, "วิชาท่องนภาขั้น2", 2},
			Subject{111113, "วิชาท่องนภาขั้น3", 2},
			Subject{111114, "วิชาท่องนภาขั้น4", 2},
			Subject{211111, "เงานินจาขั้น1", 4},
			Subject{211112, "เงานินจาขั้น2", 4},
			Subject{211112, "เงานินจาขั้น3", 4},
			Subject{211112, "เงานินจาขั้น4", 4},
			Subject{214101, "Art App", 3},
			Subject{214102, "Man'Solcail", 3},
			Subject{214103, "Man'Economy", 3},
			Subject{214104, "Digital Literacy", 3},
			Subject{214105, "วิชาญี่ปุ่นสอนกินซูชิ", 3},
			Subject{523111, "ComputerFinfin1", 1},
			Subject{523112, "ComputerFinfin2", 1},
			Subject{523113, "ComputerFinfin3", 1},
		},
	}

	for _, rs := range subjects.Subject {
		client.Subject.
			Create().
			SetCode(rs.Code).
			SetSubjects(rs.Subjects).
			SetCreditpiont(rs.Creditpiont).
			Save(context.Background())
	}

	// Set Year Data
	years := Years{
		Year: []Year{
			Year{2551},
			Year{2552},
			Year{2553},
			Year{2554},
			Year{2555},
			Year{2556},
			Year{2557},
			Year{2558},
			Year{2559},
			Year{2560},
			Year{2561},
			Year{2562},
			Year{2563},
			Year{2564},
			Year{2565},
		},
	}

	for _, ry := range years.Year {
		client.Year.
			Create().
			SetYears(ry.Years).
			Save(context.Background())
	}

	// Set Agency Data
	agencys := Agencys{
		Agency: []Agency{
			Agency{"ส่วนกิจการนักศึกษา"},
			Agency{"สำนักวิชาแพทยศาสตร์"},
			Agency{"สำนักวิชาเทคโนโลยีการเกษตร"},
			Agency{"สำนักวิชาวิทยาศาสตร์"},
			Agency{"สำนักวิชาเทคโนโลยีสังคม"},
			Agency{"สำนักวิชาพยาบาลศาสตร์"},
			Agency{"สำนักวิชาวิศวกรรมศาสตร์"},
			Agency{"สถานกีฬาและสุขภาพ"},
			Agency{"เทคโนธานี"},
			Agency{"สุรสัมนาคาร"},
			Agency{"โรงพยาบาล"},
			Agency{"ฟาร์มมหาวิทยาลัย"},
			Agency{"ศูนย์บรรณสารและสื่อการศึกษา"},
		},
	}

	for _, ag := range agencys.Agency {
		client.Agency.
			Create().
			SetAGENCY(ag.Agency).
			Save(context.Background())
	}

	// Set Place Data
	places := Places{
		Place: []Place{
			Place{"เทคโนธานี"},
			Place{"ศูนย์บรรณสารและสื่อการศึกษา"},
			Place{"สุรสัมนาคาร"},
			Place{"โรงพยาบาล"},
			Place{"สถานกีฬาและสุขภาพ"},
			Place{"อาคารเรียนรวม 1"},
			Place{"อาคารเรียนรวม 2"},
			Place{"อาคารสิรินธรวิศวพัฒน์ F11"},
		},
	}

	for _, p := range places.Place {
		client.Place.
			Create().
			SetPLACE(p.Place).
			Save(context.Background())
	}

	//Set Prefix Data
	prefixs := Prefixs{
		Prefix: []Prefix{
			Prefix{"นาย"},
			Prefix{"นาง"},
			Prefix{"นางสาว"},
		},
	}
	for _, pf := range prefixs.Prefix {
		client.Prefix.
			Create().
			SetPrefix(pf.Prefix).
			Save(context.Background())
	}

	// Set Faculty Data
	facultys := Facultys{
		Faculty: []Faculty{
			Faculty{"สำนักวิชาวิทยาศาสตร์"},
			Faculty{"สำนักวิชาเทคโนโลยีสังคม"},
			Faculty{"สำนักวิชาเทคโนโลยีการเกษตร"},
			Faculty{"สำนักวิชาวิศวกรรมศาสตร์"},
			Faculty{"สำนักวิชาแพทยศาสตร์"},
			Faculty{"สำนักวิชาพยาบาลศาสตร์"},
			Faculty{"สำนักวิชาทันตแพทยศาสตร์"},
			Faculty{"สำนักวิชาสาธารณสุขศาสตร์"},
		},
	}
	for _, fa := range facultys.Faculty {
		client.Faculty.
			Create().
			SetFaculty(fa.Faculty).
			Save(context.Background())
	}

	// Set Professorship Data
	professorships := Professorships{
		Professorship: []Professorship{
			Professorship{"ศาสตราจารย์"},
			Professorship{"รองศาสตราจารย์"},
			Professorship{"ผู้ช่วยศาสตราจารย์"},
			Professorship{"อาจารย์"},
		},
	}
	for _, ps := range professorships.Professorship {
		client.Professorship.
			Create().
			SetProfessorship(ps.Professorship).
			Save(context.Background())
	}

	//Set Country Data
	countrys := Countrys{
		Country: []Country{
			Country{"ไทย"},
			Country{"จีน"},
			Country{"กัมพูชา"},
			Country{"อังกฤษ"},
			Country{"อเมริกา"},
			Country{"ออสเตรเลีย"},
			Country{"สิงคโปร์"},
		},
	}
	for _, cu := range countrys.Country {
		client.Country.
			Create().
			SetCountry(cu.Country).
			Save(context.Background())
	}

	//Set Continent Data
	continents := Continents{
		Continent: []Continent{
			Continent{"อเมริกา"},
			Continent{"เอเซีย"},
			Continent{"ยุโรป"},
			Continent{"แอฟริกา"},
			Continent{"ออสเตรเลีย"},
			Continent{"แอนตาร์กติกา"},
		},
	}
	for _, cn := range continents.Continent {
		client.Continent.
			Create().
			SetContinent(cn.Continent).
			Save(context.Background())
	}

	//Set Region Data
	regions := Regions{
		Region: []Region{
			Region{"ภาคกลาง"},
			Region{"ภาคตะวันออกเฉียงเหนือ"},
			Region{"ภาคเหนือ"},
			Region{"ภาคใต้"},
		},
	}
	for _, st := range regions.Region {
		client.Region.
			Create().
			SetName(st.Region).
			Save(context.Background())
	}

	// Set Province Data
	provinces := Provinces{
		Province: []Province{
			Province{"d", "เมืองนครราชสีมา", "ในเมือง", 30000},
			Province{"a", "เมืองนครราชสีมา", "จอหอ", 30310},
			Province{"b", "สีดา", "โพนทอง", 30430},
			Province{"c", "โนนสูง", "โนนสูง", 30280},
		},
	}
	for _, pv := range provinces.Province {
		client.Province.
			Create().
			SetProvince(pv.Province).
			SetDistrict(pv.District).
			SetSubdistrict(pv.Subdistrict).
			SetPostal(pv.Postal).
			Save(context.Background())
	}

	// Set Institution Data
	institutions := Institutions{
		Institution: []Institution{
			Institution{"สาขาวิชาเคมี"},
			Institution{"สาขาวิชาคณิตศาสตร์"},
			Institution{"สาขาวิชาฟิสิกส์"},
			Institution{"สาขาวิทยาศาสตร์การกีฬา"},
			Institution{"สาขาวิชาวิศวกรรมคอมพิวเตอร์"},
			Institution{"สาขาวิชาวิศวกรรมเครื่องกล"},
			Institution{"สาขาวิชาวิศวกรรมโยธา"},
			Institution{"สาขาวิชาวิศวกรรมไฟฟ้า"},
			Institution{"สาขาวิชาวิศวกรรมขนส่งและโลจิสติกส์"},
			Institution{"สาขาวิชาวิศวกรรมเซรามิก"},
			Institution{"สาขาวิชาวิศวกรรมโทรคมนาคม"},
			Institution{"สาขาวิชาวิศวกรรมปิโตรเลียมและเทคโนโลยีธรณี"},
			Institution{"สาขาวิชาวิศวกรรมพอลิเมอร์"},
			Institution{"สาขาวิชาวิศวกรรมอุตสาหการ"},
			Institution{"สาขาวิชาวิศวกรรมยานยนต์"},
			Institution{"สาขาวิชาวิศวกรรมอากาศยาน"},
			Institution{"สาขาวิชาวิศวกรรมอิเล็กทรอนิกส์"},
			Institution{"สาขาวิชาวิศวกรรมเกษตรและอาหาร"},
			Institution{"สาขาวิชาวิศวกรรมนวัตกรรมและการออกแบบวัสดุ"},
			Institution{"สาขาวิชาวิศวกรรมโยธาและโครงสร้างพื้นฐาน"},
			Institution{"สาขาวิชาวิศวกรรมเมคคาทรอนิกส์"},
			Institution{"สาขาวิชาเทคโนโลยีการจัดการ"},
			Institution{"สาขาวิชาเทคโนโลยีอาหาร"},
			Institution{"สาขาวิชาแพทยศาสตรบัณฑิต"},
			Institution{"สาขาวิชาพยาบาลศาสตรบัณฑิต"},
			Institution{"สาขาวิชาทันตแพทยศาสตรบัณฑิต"},
			Institution{"สาขาวิชาอาชีวอนามัยและความปลอดภัย"},
			Institution{"สาขาวิชาอนามัยสิ่งแวดล้อม"},
		},
	}
	for _, in := range institutions.Institution {
		client.Institution.
			Create().
			SetInstitution(in.Institution).
			Save(context.Background())
	}

	// Set degree Data
	degrees := Degrees{
		Degree: []Degree{
			Degree{"ปริญญาบัณฑิต"},
			Degree{"ปริญญามหาบัณฑิต"},
			Degree{"ปริญญาดุษฎีบัณฑิต"},
			Degree{"ปริญญากิตติมศักดิ์"},
			Degree{"มัธยมศึกษาตอนปลาย"},
			Degree{"ปวช.,ปวส."},
		},
	}
	for _, de := range degrees.Degree {
		client.Degree.
			Create().
			SetDegree(de.Degree).
			Save(context.Background())
	}
	// Set Emp Data
	Emps := Emps{
		Emp: []Emp{
			Emp{"Prayut", "1234", "ทะเบียน"},
			Emp{"Prawit", "1234", "ทะเบียน"},
			Emp{"Tummanas", "1234", "ทะเบียน"},
			Emp{"Pareena", "1234", "ทะเบียน"},
		},
	}
	for _, de := range Emps.Emp {
		client.Emp.
			Create().
			SetUser(de.User).
			SetPass(de.Pass).
			SetRole(de.Role).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
