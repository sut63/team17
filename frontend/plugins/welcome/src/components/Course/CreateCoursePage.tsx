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
};

const initialCourseValidate = {
  coursename: false,
  degree: false,
  faculty: false,
  institution: false,
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
    let regExp = /ปริญญา+/i;
    let newCourseName = e.target.value;
    if (newCourseName && !newCourseName.match(regExp)) {
      setCourseNameError(true);
    } else {
      setCourseNameError(false);
    }
    setCourse({ ...other, coursename: e.target.value });
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
    setCourseValidate({
      coursename: course.coursename == '',
      degree: course.degree == '',
      faculty: course.faculty == '',
      institution: course.institution == '',
    });
  };

  const validateCourseData = () => {
    return (
      course.coursename != '' &&
      course.degree != '' &&
      course.faculty != '' &&
      course.institution != ''
    );
  };
  return (
    <>
      <Page theme={pageTheme.home}>
        <Header
          title="ระบบบันทึกข้อมูลหลักสูตร"
          subtitle="เพื่อเพิ่มข้อมูลหลักสูตรต่างๆภายในมหาลัย"
        ></Header>
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
