import React, { FC,useState,useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import {
    Container,
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    TextField,
    Avatar,
    Button,
    FormHelperText,
  } from '@material-ui/core';
  import { EntYear } from '../../api/models/EntYear';
  import { EntResults } from '../../api/models/EntResults';
  import { EntTerm } from '../../api/models/EntTerm';
  import { EntSubject } from '../../api/models/EntSubject';
  import { EntStudent } from '../../api/models/EntStudent';
  import { DefaultApi } from '../../api/apis/DefaultApi'; // Api Gennerate From Command
  import Swal from 'sweetalert2'; // alert
  import { Cookies } from '../../Cookie';
// alert setting
const Toast = Swal.mixin({
  toast: true,
  position: 'top',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});
// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

 




const Results: FC<{}> = () => {

  const classes = useStyles();
  //const************************************************
  const api = new DefaultApi();
  const [loading, setLoading] = React.useState(false);
  const [years, setYears] = React.useState<EntYear[]>([]);
  const [resultss, setResultss] = React.useState<EntResults[]>([]);
  const [terms, setTerms] = React.useState<EntTerm[]>([]);
  const [subjects, setSubjects] = React.useState<EntSubject[]>([]);
  const [students, setStudents] = React.useState<EntStudent[]>([]);
 

  

  //Hook********************************************
  useEffect(() => {
    const getYears = async () => {
        const res = await api.listYear({ limit: 100, offset: 0 });
        //setLoading(true);
        setYears(res);
      };
    getYears();
  }, []);

  useEffect(() => {
    const getResultss = async () => {
        const res = await api.listResults({ limit: 100, offset: 0 });
        //setLoading(true);
        setResultss(res);
      };
    getResultss();
  }, []);

  useEffect(() => {
    const getTerms = async () => {
        const res = await api.listTerm({ limit: 100, offset: 0 });
        //setLoading(true);
        setTerms(res);
      };
    getTerms();
  }, []);

  useEffect(() => {
    const getSubjects = async () => {
        const res = await api.listSubject({ limit: 100, offset: 0 });
        //setLoading(true);
        setSubjects(res);
      };
    getSubjects();
  }, []);

  useEffect(() => {
    const getStudents = async () => {
        const res = await api.listStudent({ limit: 100, offset: 0 });
        //setLoading(true);
        setStudents(res);
      };
    getStudents();
  }, []);
   
  //Get Data By Textfile and ComboBox ************************************************
  const [yearx, setYearx] = React.useState('');
  const [gradex, setGradex] = React.useState('');
  const [studentx, setStudentx] = React.useState('');
  const [subx, setSubx] = React.useState('');
  const [termx, setTermx] = React.useState('');
  const [timex, setTimex] = React.useState('');
  const [groupx, setGroupx] = React.useState('');
  

  let yearID = Number(yearx)
  let grade = Number(gradex)
  let studentID = Number(studentx)
  let subjectID = Number(subx)
  let termID = Number(termx)
  let group = Number(groupx)
  let timed = timex + ":00+07:00"
  

  let results = {
	grade,    
	studentID,
	yearID,
  subjectID,
  termID,
  group,
  timed,
    };

  //Handle chang********************************************************************
  const handleInputYear = (event: any) => {
    setYearx(event.target.value);
    
  };
  const handleInputGrade = (event: any) => {
    setGradex(event.target.value);
    
  };
  const handleInputStudent = (event: any) => {
    setStudentx(event.target.value);
    
  };
  const handleInputSubject = (event: any) => {
    setSubx(event.target.value);
    
  };
  const handleInputTerm = (event: any) => {
    setTermx(event.target.value);
    
  };
  const handleInputTime = (event: any) => {
    setTimex(event.target.value);
    
  };
  
  
  
  ////--------------------------------------------
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  //Check Save Error Function text field
  const CheckSaveError = (field:string) =>{
    switch(field) {
      case 'grade':
      alertMessage("error","เกรดต้องมากกว่า0 และมากสุด=4")
      return
      case 'group':
      alertMessage("error","กลุ่มมีแค่1-4")
      return
    }
  }
  //Check Save Error Function combobox
  const CheckComboboxError = (field:string) =>{
    switch(field){
      case "Year not found":
      alertMessage("error","กรุณาใส่ข้อมูลปี")
      return
      case "Term not found":
      alertMessage("error","กรุณาใส่ข้อมูลเทอม")
      return
      case "Subject not found":
      alertMessage("error","กรุณาใส่ข้อมูลวิชา")
      return
      case "Student not found":
      alertMessage("error","กรุณาใส่ข้อมูลนักศึกษา")
      return
      case "time null":
      alertMessage("error","กรุณาเลือกวันเวลาที่ออกเกรด")
      return
    }
  }



  //Validate Select 
  const ResultFieldValidate = {
    grade1111: false,
    term1111: false,
    year1111: false,
    subject1111: false,
    student1111: false,
    group1111: false,
    time1111: false,
  };
  const [resultValidate, setResultValidate] = useState(ResultFieldValidate); //Select Validate
  const checkValidateData = () => {
    setResultValidate({
      grade1111: gradex == '',
      term1111: termx == '',
      year1111: yearx == '',
      subject1111: subx == '',
      student1111: studentx == '',
      group1111: groupx == '',
      time1111: timex == '',
    });
  };

  //validate TaxtField
  const [GradeError, setGradeError] = React.useState('');
  const [GroupError, setGroupError] = React.useState('');

  const validateTextfield = (val: number) => {
    if(val > 4)
    return false;
    else if (val < 0)
    return false;
    else
    return true;
  }
  const validateTextfieldgruop = (val: number) => {
    if(val > 4)
    return false;
    else if (val < 1)
    return false;
    else
    return true;
  }

  const checkPattern  = (id: any, value: string) => {
    switch(id) {
      case 'grade':
        validateTextfield(Number(value)) ? setGradeError('') : setGradeError('0-4 เท่านั้น');
        return;
        case 'group':
          validateTextfieldgruop(Number(value)) ? setGroupError('') : setGroupError('1-4 เท่านั้น');
          return;
      default:
        return;
    }
  }

  const g = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof gradex;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setGradex( value );
  };

  const ggruop = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof groupx;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setGroupx( value );
  };




  
  //seve**********************************************************
   // function save data
   function save() {
     setLoading(true);
     checkValidateData();
    const apiUrl = 'http://localhost:8080/api/v1/resultss';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(results),
    };

    console.log(results); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true ) {
          
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          CheckSaveError(data.error.Name)
          CheckComboboxError(data.error)
        }
      });
  }
  

  //console log****************************************************************
  console.log(years);
  console.log(resultss);
  console.log(terms);
  console.log(subjects);
  console.log(students);
  console.log(results)

  //cookie logout
  var cook = new Cookies()
  var cookieName = cook.GetCookie()

  function Clears() {
    cook.ClearCookie()
    window.location.reload(false)
  }

  return (
    <Page theme={pageTheme.home}>
      <Header
      title={'Results'}
      subtitle=''
    >
        <Avatar alt="Remy Sharp"/>
        <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
        <Button variant="text" color="secondary" size="large"
          onClick={Clears} > Logout </Button>
    </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>




            <Grid item xs={3}>
              <div className={classes.paper}>รหัสนักศึกษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl  error={resultValidate.student1111} variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรหัสนักศึกษา</InputLabel>
                <Select
                  error={resultValidate.student1111}
                  name="StudentID"
                  value={results.studentID || ''} // (undefined || '') = ''
                  onChange={handleInputStudent}
                >
                  {students.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
                {resultValidate.student1111 ? (
                  <FormHelperText>กรุณาเลือกรหัสนักศึกษา</FormHelperText>
                ) : null}
              </FormControl>
            </Grid>



            

            <Grid item xs={3}>
              <div className={classes.paper}>ปีการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl error={resultValidate.year1111} variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกปีการศึกษา</InputLabel>
                <Select
                  error={resultValidate.year1111}
                  name="YearID"
                  value={results.yearID || ''} // (undefined || '') = ''
                  onChange={handleInputYear}
                >
                  {years.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.years}
                      </MenuItem>
                    );
                  })}
                </Select>
                {resultValidate.year1111 ? (
                  <FormHelperText>กรุณาเลือกปีการศึกษา</FormHelperText>
                ) : null}
              </FormControl>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ภาคการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl error={resultValidate.term1111} variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกภาคการศึกษา</InputLabel>
                <Select
                error={resultValidate.term1111}
                  name="TermID"
                  value={results.termID || ''} // (undefined || '') = ''
                  onChange={handleInputTerm}
                >
                  {terms.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.semester}
                      </MenuItem>
                    );
                  })}
                </Select>
                {resultValidate.term1111 ? (
                  <FormHelperText>กรุณาเลือกภาคการศึกษา</FormHelperText>
                ) : null}
              </FormControl>
            </Grid>




            <Grid item xs={3}>
              <div className={classes.paper}>วิชา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl error={resultValidate.subject1111} variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกวิชา</InputLabel>
                <Select
                error={resultValidate.subject1111}
                  name="SubjectID"
                  value={results.subjectID || ''} // (undefined || '') = ''
                  onChange={handleInputSubject}
                >
                  {subjects.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.subjects}
                      </MenuItem>
                    );
                  })}
                </Select>
                {resultValidate.subject1111 ? (
                  <FormHelperText>กรุณาเลือกวิชา</FormHelperText>
                ) : null}
              </FormControl>
            </Grid>



            <Grid item xs={3}>
                  <div className={classes.paper}>เกรด</div>
                </Grid>
                <Grid item xs={9}>
                  <TextField variant="outlined" className={classes.textField}
                      helperText= {loading? GradeError:""} 
                      error= {GradeError ? true:false }
                  
                      id="grade" 
                      type="Number" 
                      name="grade"
                      value={results.grade || ''} // (undefined || '') = ''
                      onChange={g}>
                                                                            
                  </TextField>
                </Grid>

            

                <Grid item xs={3}>
                  <div className={classes.paper}>กลุ่ม</div>
                </Grid>
                <Grid item xs={9}>
                  <TextField variant="outlined" className={classes.textField}
                      helperText= {loading? GroupError:""} 
                      error= {GroupError ? true:false }
                  
                      id="group" 
                      type="Number" 
                      name="group"
                      value={results.group || ''} // (undefined || '') = ''
                      onChange={ggruop}>
                                                                            
                  </TextField>
                </Grid>



                <Grid item xs={3}>
              <div className={classes.paper}>วันที่ออกเกรด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl error={resultValidate.time1111} variant="outlined" className={classes.formControl}>
                <TextField
                error={resultValidate.time1111}
                  name="timex"
                  type="datetime-local"
                  defaultValue="2020-12-31"
                  onChange={handleInputTime}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
                 {resultValidate.time1111 ? (
                  <FormHelperText>กรุณาเลือกวันที่ออกเกรด</FormHelperText>
                ) : null}
              </FormControl>
            </Grid>

          

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                style={{width: 300 , marginTop: 20}}
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึก
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );

};

export default Results;
