// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/sut63/team17/app/ent/activity"
	"github.com/sut63/team17/app/ent/agency"
	"github.com/sut63/team17/app/ent/continent"
	"github.com/sut63/team17/app/ent/country"
	"github.com/sut63/team17/app/ent/course"
	"github.com/sut63/team17/app/ent/degree"
	"github.com/sut63/team17/app/ent/emp"
	"github.com/sut63/team17/app/ent/faculty"
	"github.com/sut63/team17/app/ent/institution"
	"github.com/sut63/team17/app/ent/place"
	"github.com/sut63/team17/app/ent/prefix"
	"github.com/sut63/team17/app/ent/professor"
	"github.com/sut63/team17/app/ent/professorship"
	"github.com/sut63/team17/app/ent/province"
	"github.com/sut63/team17/app/ent/region"
	"github.com/sut63/team17/app/ent/results"
	"github.com/sut63/team17/app/ent/schema"
	"github.com/sut63/team17/app/ent/student"
	"github.com/sut63/team17/app/ent/subject"
	"github.com/sut63/team17/app/ent/term"
	"github.com/sut63/team17/app/ent/year"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	activityFields := schema.Activity{}.Fields()
	_ = activityFields
	// activityDescACTIVITYNAME is the schema descriptor for ACTIVITYNAME field.
	activityDescACTIVITYNAME := activityFields[0].Descriptor()
	// activity.ACTIVITYNAMEValidator is a validator for the "ACTIVITYNAME" field. It is called by the builders before save.
	activity.ACTIVITYNAMEValidator = activityDescACTIVITYNAME.Validators[0].(func(string) error)
	// activityDescHours is the schema descriptor for hours field.
	activityDescHours := activityFields[2].Descriptor()
	// activity.HoursValidator is a validator for the "hours" field. It is called by the builders before save.
	activity.HoursValidator = func() func(int) error {
		validators := activityDescHours.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(hours int) error {
			for _, fn := range fns {
				if err := fn(hours); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	agencyFields := schema.Agency{}.Fields()
	_ = agencyFields
	// agencyDescAGENCY is the schema descriptor for AGENCY field.
	agencyDescAGENCY := agencyFields[0].Descriptor()
	// agency.AGENCYValidator is a validator for the "AGENCY" field. It is called by the builders before save.
	agency.AGENCYValidator = agencyDescAGENCY.Validators[0].(func(string) error)
	continentFields := schema.Continent{}.Fields()
	_ = continentFields
	// continentDescContinent is the schema descriptor for continent field.
	continentDescContinent := continentFields[0].Descriptor()
	// continent.ContinentValidator is a validator for the "continent" field. It is called by the builders before save.
	continent.ContinentValidator = continentDescContinent.Validators[0].(func(string) error)
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescCountry is the schema descriptor for country field.
	countryDescCountry := countryFields[0].Descriptor()
	// country.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	country.CountryValidator = countryDescCountry.Validators[0].(func(string) error)
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescCourse is the schema descriptor for course field.
	courseDescCourse := courseFields[0].Descriptor()
	// course.CourseValidator is a validator for the "course" field. It is called by the builders before save.
	course.CourseValidator = courseDescCourse.Validators[0].(func(string) error)
	// courseDescAnnotation is the schema descriptor for annotation field.
	courseDescAnnotation := courseFields[1].Descriptor()
	// course.AnnotationValidator is a validator for the "annotation" field. It is called by the builders before save.
	course.AnnotationValidator = courseDescAnnotation.Validators[0].(func(string) error)
	// courseDescCredit is the schema descriptor for credit field.
	courseDescCredit := courseFields[2].Descriptor()
	// course.CreditValidator is a validator for the "credit" field. It is called by the builders before save.
	course.CreditValidator = courseDescCredit.Validators[0].(func(int) error)
	// courseDescCourseID is the schema descriptor for course_id field.
	courseDescCourseID := courseFields[3].Descriptor()
	// course.CourseIDValidator is a validator for the "course_id" field. It is called by the builders before save.
	course.CourseIDValidator = courseDescCourseID.Validators[0].(func(int) error)
	degreeFields := schema.Degree{}.Fields()
	_ = degreeFields
	// degreeDescDegree is the schema descriptor for degree field.
	degreeDescDegree := degreeFields[0].Descriptor()
	// degree.DegreeValidator is a validator for the "degree" field. It is called by the builders before save.
	degree.DegreeValidator = degreeDescDegree.Validators[0].(func(string) error)
	empFields := schema.Emp{}.Fields()
	_ = empFields
	// empDescUser is the schema descriptor for user field.
	empDescUser := empFields[0].Descriptor()
	// emp.UserValidator is a validator for the "user" field. It is called by the builders before save.
	emp.UserValidator = empDescUser.Validators[0].(func(string) error)
	// empDescPass is the schema descriptor for pass field.
	empDescPass := empFields[1].Descriptor()
	// emp.PassValidator is a validator for the "pass" field. It is called by the builders before save.
	emp.PassValidator = empDescPass.Validators[0].(func(string) error)
	facultyFields := schema.Faculty{}.Fields()
	_ = facultyFields
	// facultyDescFaculty is the schema descriptor for faculty field.
	facultyDescFaculty := facultyFields[0].Descriptor()
	// faculty.FacultyValidator is a validator for the "faculty" field. It is called by the builders before save.
	faculty.FacultyValidator = facultyDescFaculty.Validators[0].(func(string) error)
	institutionFields := schema.Institution{}.Fields()
	_ = institutionFields
	// institutionDescInstitution is the schema descriptor for institution field.
	institutionDescInstitution := institutionFields[0].Descriptor()
	// institution.InstitutionValidator is a validator for the "institution" field. It is called by the builders before save.
	institution.InstitutionValidator = institutionDescInstitution.Validators[0].(func(string) error)
	placeFields := schema.Place{}.Fields()
	_ = placeFields
	// placeDescPLACE is the schema descriptor for PLACE field.
	placeDescPLACE := placeFields[0].Descriptor()
	// place.PLACEValidator is a validator for the "PLACE" field. It is called by the builders before save.
	place.PLACEValidator = placeDescPLACE.Validators[0].(func(string) error)
	prefixFields := schema.Prefix{}.Fields()
	_ = prefixFields
	// prefixDescPrefix is the schema descriptor for prefix field.
	prefixDescPrefix := prefixFields[0].Descriptor()
	// prefix.PrefixValidator is a validator for the "prefix" field. It is called by the builders before save.
	prefix.PrefixValidator = prefixDescPrefix.Validators[0].(func(string) error)
	professorFields := schema.Professor{}.Fields()
	_ = professorFields
	// professorDescTel is the schema descriptor for tel field.
	professorDescTel := professorFields[1].Descriptor()
	// professor.TelValidator is a validator for the "tel" field. It is called by the builders before save.
	professor.TelValidator = func() func(string) error {
		validators := professorDescTel.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(tel string) error {
			for _, fn := range fns {
				if err := fn(tel); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// professorDescEmail is the schema descriptor for email field.
	professorDescEmail := professorFields[2].Descriptor()
	// professor.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	professor.EmailValidator = professorDescEmail.Validators[0].(func(string) error)
	professorshipFields := schema.Professorship{}.Fields()
	_ = professorshipFields
	// professorshipDescProfessorship is the schema descriptor for professorship field.
	professorshipDescProfessorship := professorshipFields[0].Descriptor()
	// professorship.ProfessorshipValidator is a validator for the "professorship" field. It is called by the builders before save.
	professorship.ProfessorshipValidator = professorshipDescProfessorship.Validators[0].(func(string) error)
	provinceFields := schema.Province{}.Fields()
	_ = provinceFields
	// provinceDescProvince is the schema descriptor for province field.
	provinceDescProvince := provinceFields[0].Descriptor()
	// province.ProvinceValidator is a validator for the "province" field. It is called by the builders before save.
	province.ProvinceValidator = provinceDescProvince.Validators[0].(func(string) error)
	// provinceDescDistrict is the schema descriptor for district field.
	provinceDescDistrict := provinceFields[1].Descriptor()
	// province.DistrictValidator is a validator for the "district" field. It is called by the builders before save.
	province.DistrictValidator = provinceDescDistrict.Validators[0].(func(string) error)
	// provinceDescSubdistrict is the schema descriptor for subdistrict field.
	provinceDescSubdistrict := provinceFields[2].Descriptor()
	// province.SubdistrictValidator is a validator for the "subdistrict" field. It is called by the builders before save.
	province.SubdistrictValidator = provinceDescSubdistrict.Validators[0].(func(string) error)
	// provinceDescPostal is the schema descriptor for postal field.
	provinceDescPostal := provinceFields[3].Descriptor()
	// province.PostalValidator is a validator for the "postal" field. It is called by the builders before save.
	province.PostalValidator = func() func(string) error {
		validators := provinceDescPostal.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(postal string) error {
			for _, fn := range fns {
				if err := fn(postal); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	regionFields := schema.Region{}.Fields()
	_ = regionFields
	// regionDescName is the schema descriptor for name field.
	regionDescName := regionFields[0].Descriptor()
	// region.NameValidator is a validator for the "name" field. It is called by the builders before save.
	region.NameValidator = regionDescName.Validators[0].(func(string) error)
	resultsFields := schema.Results{}.Fields()
	_ = resultsFields
	// resultsDescGrade is the schema descriptor for grade field.
	resultsDescGrade := resultsFields[0].Descriptor()
	// results.GradeValidator is a validator for the "grade" field. It is called by the builders before save.
	results.GradeValidator = func() func(float64) error {
		validators := resultsDescGrade.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
			validators[2].(func(float64) error),
		}
		return func(grade float64) error {
			for _, fn := range fns {
				if err := fn(grade); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// resultsDescGroup is the schema descriptor for group field.
	resultsDescGroup := resultsFields[1].Descriptor()
	// results.GroupValidator is a validator for the "group" field. It is called by the builders before save.
	results.GroupValidator = func() func(int) error {
		validators := resultsDescGroup.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
			validators[2].(func(int) error),
		}
		return func(group int) error {
			for _, fn := range fns {
				if err := fn(group); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescFname is the schema descriptor for fname field.
	studentDescFname := studentFields[0].Descriptor()
	// student.FnameValidator is a validator for the "fname" field. It is called by the builders before save.
	student.FnameValidator = func() func(string) error {
		validators := studentDescFname.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(fname string) error {
			for _, fn := range fns {
				if err := fn(fname); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// studentDescLname is the schema descriptor for lname field.
	studentDescLname := studentFields[1].Descriptor()
	// student.LnameValidator is a validator for the "lname" field. It is called by the builders before save.
	student.LnameValidator = func() func(string) error {
		validators := studentDescLname.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(lname string) error {
			for _, fn := range fns {
				if err := fn(lname); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// studentDescSchoolname is the schema descriptor for schoolname field.
	studentDescSchoolname := studentFields[2].Descriptor()
	// student.SchoolnameValidator is a validator for the "schoolname" field. It is called by the builders before save.
	student.SchoolnameValidator = func() func(string) error {
		validators := studentDescSchoolname.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(schoolname string) error {
			for _, fn := range fns {
				if err := fn(schoolname); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// studentDescRecentAddress is the schema descriptor for recent_address field.
	studentDescRecentAddress := studentFields[3].Descriptor()
	// student.RecentAddressValidator is a validator for the "recent_address" field. It is called by the builders before save.
	student.RecentAddressValidator = func() func(string) error {
		validators := studentDescRecentAddress.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(recent_address string) error {
			for _, fn := range fns {
				if err := fn(recent_address); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// studentDescTelephone is the schema descriptor for telephone field.
	studentDescTelephone := studentFields[4].Descriptor()
	// student.TelephoneValidator is a validator for the "telephone" field. It is called by the builders before save.
	student.TelephoneValidator = studentDescTelephone.Validators[0].(func(string) error)
	// studentDescEmail is the schema descriptor for email field.
	studentDescEmail := studentFields[5].Descriptor()
	// student.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	student.EmailValidator = studentDescEmail.Validators[0].(func(string) error)
	subjectFields := schema.Subject{}.Fields()
	_ = subjectFields
	// subjectDescCode is the schema descriptor for code field.
	subjectDescCode := subjectFields[0].Descriptor()
	// subject.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	subject.CodeValidator = subjectDescCode.Validators[0].(func(int) error)
	// subjectDescSubjects is the schema descriptor for subjects field.
	subjectDescSubjects := subjectFields[1].Descriptor()
	// subject.SubjectsValidator is a validator for the "subjects" field. It is called by the builders before save.
	subject.SubjectsValidator = subjectDescSubjects.Validators[0].(func(string) error)
	// subjectDescCreditpiont is the schema descriptor for creditpiont field.
	subjectDescCreditpiont := subjectFields[2].Descriptor()
	// subject.CreditpiontValidator is a validator for the "creditpiont" field. It is called by the builders before save.
	subject.CreditpiontValidator = subjectDescCreditpiont.Validators[0].(func(int) error)
	termFields := schema.Term{}.Fields()
	_ = termFields
	// termDescSemester is the schema descriptor for semester field.
	termDescSemester := termFields[0].Descriptor()
	// term.SemesterValidator is a validator for the "semester" field. It is called by the builders before save.
	term.SemesterValidator = termDescSemester.Validators[0].(func(int) error)
	yearFields := schema.Year{}.Fields()
	_ = yearFields
	// yearDescYears is the schema descriptor for years field.
	yearDescYears := yearFields[0].Descriptor()
	// year.YearsValidator is a validator for the "years" field. It is called by the builders before save.
	year.YearsValidator = yearDescYears.Validators[0].(func(int) error)
}
