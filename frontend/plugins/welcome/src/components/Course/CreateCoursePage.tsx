import React, { useState, useEffect } from 'react';

import { Header, Page, pageTheme, Content } from '@backstage/core';
import { Alert } from '@material-ui/lab';
import {
  makeStyles,
  createStyles,
  Theme,
  FormControl,
  TextField,
  Button,
  Select,
  MenuItem,
  Avatar,
  InputLabel,
  FormHelperText,
} from '@material-ui/core';
import {
  DefaultApi,
  EntDegree,
  EntInstitution,
  EntFaculty,
  CreateCourseRequest,
} from '../../api';

import { Cookies } from '../../Cookie';
//
// header css
const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    select: {
      width: 500,
    },
    margin: {
      margin: theme.spacing(2),
    },

    flexcenter: {
      display: 'flex',
      justifyContent: 'center',
    },

    column: {
      flexDirection: 'column',
      display: 'flex',
      justifyContent: 'center',
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },

    textField: {
      margin: theme.spacing(2),
      width: 500,
    },
  }),
);

const initialCourseState = {
  coursename: '',
  degree: '',
  faculty: '',
  institution: '',
  courseId: '',
  annotation: '',
  credit: '',
};

const initialCourseValidate = {
  coursename: false,
  degree: false,
  faculty: false,
  institution: false,
  courseId: false,
  annotation: false,
  credit: false,
};
const CreateCoursePage = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [course, setCourse] = useState(initialCourseState);
  const [degree, setDegree] = useState<EntDegree[]>([]);
  const [faculties, setFaculties] = useState<EntFaculty[]>([]);
  const [institutions, setInstitutions] = useState<EntInstitution[]>([]);
  const [courseValidate, setCourseValidate] = useState(initialCourseValidate);
  const [courseNameError, setCourseNameError] = useState<boolean>(false);
  const [courseIdError, setCourseIdError] = useState<boolean>(false);
  const [annotationError, setAnnotationError] = useState<boolean>(false);
  const [creditError, setCreditError] = useState<boolean>(false);
  const [alert, setAlert] = useState(false);
  const [saveResult, setSaveResult] = useState(false);

  useEffect(() => {
    getDegree();
    getFaculty();
    getInstitution();
  }, []);

  const getDegree = async () => {
    let degree: EntDegree[] = await api.listDegree({});
    setDegree(degree);
  };

  const getFaculty = async () => {
    let faculties: EntFaculty[] = await api.listFaculty({});
    setFaculties(faculties);
  };

  const getInstitution = async () => {
    let institutions: EntInstitution[] = await api.listInstitution({});
    setInstitutions(institutions);
  };

  const handleInputChange = (e: any) => {
    const { coursename, ...other } = course;
    setCourse({ ...other, coursename: e.target.value });
  };

  const handleCourseIdChange = (e: any) => {
    const { courseId, ...other } = course;
    setCourse({ ...other, courseId: e.target.value });
  };

  const handleCreditChange = (e: any) => {
    const { credit, ...other } = course;
    setCourse({ ...other, credit: e.target.value });
  };

  const handleAnnotationChange = (e: any) => {
    const { annotation, ...other } = course;
    setCourse({ ...other, annotation: e.target.value });
  };

  const handleDegreeChange = (e: any) => {
    const { degree, ...other } = course;
    setCourse({ ...other, degree: e.target.value });
  };

  const handleFacultyChange = (e: any) => {
    const { faculty, ...other } = course;
    setCourse({ ...other, faculty: e.target.value });
  };

  const handleInstitutionChange = (e: any) => {
    const { institution, ...other } = course;
    setCourse({ ...other, institution: e.target.value });
  };

  const onSaveCourse = async () => {
    checkValidateData();
    if (validateCourseData()) {
      let createCourseRequest: CreateCourseRequest = {
        course: {
          coursename: course.coursename,
          degree: Number(course.degree),
          faculty: Number(course.faculty),
          institution: Number(course.institution),
          credit: Number(course.credit),
          annotation: course.annotation,
          courseID: Number(course.courseId),
        },
      };
      try {
        await api.createCourseRaw(createCourseRequest);
        clearData();
        showAlert(true);
      } catch (error) {
        showAlert(false);
        console.log(error);
      }
    }
  };

  const showAlert = (isSuccess: boolean) => {
    setAlert(true);
    setSaveResult(isSuccess);
    setTimeout(() => {
      setAlert(false);
      setSaveResult(false);
    }, 2000);
  };

  const clearData = () => {
    setCourse(initialCourseState);
    setCourseValidate(initialCourseValidate);
  };

  const checkValidateData = () => {
    validateCourseId();
    validateAnnotation();
    validateCredit();
    setCourseValidate({
      coursename: course.coursename == '',
      degree: course.degree == '',
      faculty: course.faculty == '',
      institution: course.institution == '',
      courseId: course.courseId == '',
      annotation: course.annotation == '',
      credit: course.credit == '',
    });
  };

  const validateCourseId = () => {
    let regExp = /\d{6}/i;
    if (
      course.courseId &&
      !(course.courseId.match(regExp) && course.courseId.length == 6)
    ) {
      setCourseIdError(true);
    } else {
      setCourseIdError(false);
    }
  };

  const validateAnnotation = () => {
    let regExp = /[^a-zA-Z]+$/i;
    if (course.annotation && !course.annotation.match(regExp)) {
      setAnnotationError(true);
    } else {
      setAnnotationError(false);
    }
  };

  const validateCredit = () => {
    let regExp = /\d+/i;
    if (course.credit && !course.credit.match(regExp)) {
      setCreditError(true);
    } else {
      setCreditError(false);
    }
  };

  const validateCourseData = () => {
    let regExp = /ปริญญา+/i;
    if (course.coursename && !course.coursename.match(regExp)) {
      setCourseNameError(true);
    } else {
      setCourseNameError(false);
    }
    return (
      !courseIdError &&
      !annotationError &&
      !creditError &&
      course.coursename != '' &&
      course.coursename.match(regExp) &&
      course.degree != '' &&
      course.faculty != '' &&
      course.institution != ''
    );
  };

  //cookie logout
  var cook = new Cookies();
  var cookieName = cook.GetCookie();

  function Clears() {
    cook.ClearCookie();
    window.location.reload(false);
  }

  return (
    <>
      <Page theme={pageTheme.home}>
        <Header
          title="ระบบบันทึกข้อมูลหลักสูตร"
          subtitle="เพื่อเพิ่มข้อมูลหลักสูตรต่างๆภายในมหาลัย"
        >
          <Avatar alt="Remy Sharp" />
          <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
          <Button
            variant="text"
            color="secondary"
            size="large"
            onClick={Clears}
          >
            {' '}
            Logout{' '}
          </Button>
        </Header>
        <Content>
          {alert ? (
            <div>
              {saveResult ? (
                <Alert severity="success">บันทึกข้อมูลหลักสูตรสำเร็จ</Alert>
              ) : (
                <Alert severity="error">
                  ไม่สามารถบันทึกข้อมูลหลักสูตรได้ กรุณาลองใหม่ภายหลัง
                </Alert>
              )}
            </div>
          ) : null}

          <div className={classes.root}>
            <form autoComplete="off" className={classes.column}>
              <FormControl
                fullWidth
                className={classes.textField}
                variant="outlined"
              >
                <TextField
                  id="courseId"
                  label="Course Id (รหัสหลักสูตร)"
                  variant="outlined"
                  type="string"
                  size="medium"
                  helperText={
                    courseIdError
                      ? 'รหัสหลักสูตรต้องเป็นตัวเลขจำนวน 6 หลักเท่านั้น'
                      : courseValidate.courseId
                      ? 'กรุณาระบุรหัสหลักสูตร'
                      : null
                  }
                  error={courseIdError || courseValidate.courseId}
                  value={course.courseId}
                  onChange={handleCourseIdChange}
                />
              </FormControl>
              <FormControl
                fullWidth
                className={classes.textField}
                variant="outlined"
              >
                <TextField
                  id="coursename"
                  label="Course name"
                  variant="outlined"
                  type="string"
                  size="medium"
                  helperText={
                    courseNameError
                      ? 'ชื่อหลักสูตรไม่ถูกต้อง ต้องขึ้นตัวด้วยคำว่า ปริญญา'
                      : courseValidate.coursename
                      ? 'กรุณาระบุชื่อหลักสูตร'
                      : null
                  }
                  error={courseNameError || courseValidate.coursename}
                  value={course.coursename}
                  onChange={handleInputChange}
                />
              </FormControl>
              <FormControl
                fullWidth
                className={classes.textField}
                variant="outlined"
              >
                <TextField
                  id="annotation"
                  label="Annotation (หมายเหตุ)"
                  variant="outlined"
                  type="string"
                  size="medium"
                  helperText={
                    annotationError
                      ? 'หมายเหตุสามารถเป็นตัวเลข หรือตัวอักษรภาษาไทย เท่านั้น'
                      : courseValidate.annotation
                      ? 'กรุณาระบุหมายเหตุ'
                      : null
                  }
                  error={annotationError || courseValidate.annotation}
                  value={course.annotation}
                  onChange={handleAnnotationChange}
                />
              </FormControl>
              <FormControl
                error={courseValidate.degree}
                className={`${classes.margin} ${classes.select}`}
                variant="outlined"
              >
                <InputLabel>Degree</InputLabel>
                <Select
                  label="Degree"
                  error={courseValidate.degree}
                  variant="outlined"
                  value={course.degree}
                  onChange={handleDegreeChange}
                >
                  {degree.map((item, key) => (
                    <MenuItem key={key} value={item.id}>
                      {item.degree}
                    </MenuItem>
                  ))}
                </Select>
                {courseValidate.degree ? (
                  <FormHelperText>กรุณาเลือกระดับปริญญา</FormHelperText>
                ) : null}
              </FormControl>

              <FormControl
                className={`${classes.margin} ${classes.select}`}
                error={courseValidate.faculty}
                variant="outlined"
              >
                <InputLabel>Faculty</InputLabel>
                <Select
                  label="Faculty"
                  error={courseValidate.faculty}
                  variant="outlined"
                  value={course.faculty}
                  onChange={handleFacultyChange}
                >
                  {faculties.map((item, key) => (
                    <MenuItem key={key} value={item.id}>
                      {item.faculty}
                    </MenuItem>
                  ))}
                </Select>
                {courseValidate.faculty ? (
                  <FormHelperText>กรุณาเลือกสำนักวิชา</FormHelperText>
                ) : null}
              </FormControl>

              <FormControl
                className={`${classes.margin} ${classes.select}`}
                error={courseValidate.institution}
                variant="outlined"
              >
                <InputLabel>Institution</InputLabel>
                <Select
                  label="Institution"
                  error={courseValidate.institution}
                  variant="outlined"
                  value={course.institution}
                  onChange={handleInstitutionChange}
                >
                  {institutions
                    .filter(
                      inst =>
                        inst.edges?.instFacu?.id == Number(course.faculty),
                    )
                    .map((item, key) => (
                      <MenuItem key={key} value={item.id}>
                        {item.institution}
                      </MenuItem>
                    ))}
                </Select>
                {courseValidate.institution ? (
                  <FormHelperText>กรุณาเลือกสาขาวิชา</FormHelperText>
                ) : null}
              </FormControl>

              <FormControl
                fullWidth
                className={classes.textField}
                variant="outlined"
              >
                <TextField
                  id="credit"
                  label="Credit (หน่วยกิต)"
                  variant="outlined"
                  type="string"
                  size="medium"
                  helperText={
                    creditError
                      ? 'หน่วยกิตไม่ถูกต้อง ต้องเป็นตัวเลขเท่านั้น'
                      : courseValidate.credit
                      ? 'กรุณาระบุหน่วยกิต'
                      : null
                  }
                  error={creditError || courseValidate.credit}
                  value={course.credit}
                  onChange={handleCreditChange}
                />
              </FormControl>
              <div className={`${classes.margin} ${classes.flexcenter}`}>
                <Button onClick={() => clearData()} variant="contained">
                  ล้างข้อมูล
                </Button>
                <Button
                  style={{ marginLeft: 20 }}
                  onClick={onSaveCourse}
                  variant="contained"
                  color="primary"
                >
                  บันทึก
                </Button>
              </div>
            </form>
          </div>
        </Content>
      </Page>
    </>
  );
};

export default CreateCoursePage;
