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
  const [loading, setLoading] = useState(true);
  const [years, setYears] = React.useState<EntYear[]>([]);
  const [resultss, setResultss] = React.useState<EntResults[]>([]);
  const [terms, setTerms] = React.useState<EntTerm[]>([]);
  const [subjects, setSubjects] = React.useState<EntSubject[]>([]);
  const [students, setStudents] = React.useState<EntStudent[]>([]);
 


  //Hook********************************************
  useEffect(() => {
    const getYears = async () => {
        const res = await api.listYear({ limit: 100, offset: 0 });
        setLoading(true);
        setYears(res);
      };
    getYears();
  }, [loading]);

  useEffect(() => {
    const getResultss = async () => {
        const res = await api.listResults({ limit: 100, offset: 0 });
        setLoading(true);
        setResultss(res);
      };
    getResultss();
  }, [loading]);

  useEffect(() => {
    const getTerms = async () => {
        const res = await api.listTerm({ limit: 100, offset: 0 });
        setLoading(true);
        setTerms(res);
      };
    getTerms();
  }, [loading]);

  useEffect(() => {
    const getSubjects = async () => {
        const res = await api.listSubject({ limit: 100, offset: 0 });
        setLoading(true);
        setSubjects(res);
      };
    getSubjects();
  }, [loading]);

  useEffect(() => {
    const getStudents = async () => {
        const res = await api.listStudent({ limit: 100, offset: 0 });
        setLoading(true);
        setStudents(res);
      };
    getStudents();
  }, [loading]);
   
  //Get Data By Textfile and ComboBox ************************************************
  const [yearx, setYearx] = React.useState();
  const [gradex, setGradex] = React.useState(String);
  const [studentx, setStudentx] = React.useState();
  const [subx, setSubx] = React.useState();
  const [termx, setTermx] = React.useState();
  

  let yearID = Number(yearx)
  let grade = Number(gradex)
  let studentID = Number(studentx)
  let subjectID = Number(subx)
  let termID = Number(termx)

  let results = {
	grade,    
	studentID,
	yearID,
  subjectID,
  termID,
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
  console.log(results)
  
  //seve**********************************************************
   // function save data
   function save() {
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
        if (data.id != null ) {
          
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
  

  //console log****************************************************************
  console.log(years);
  console.log(resultss);
  console.log(terms);
  console.log(subjects);
  console.log(students);

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
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรหัสนักศึกษา</InputLabel>
                <Select
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
              </FormControl>
            </Grid>



            

            <Grid item xs={3}>
              <div className={classes.paper}>ปีการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกปีการศึกษา</InputLabel>
                <Select
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
              </FormControl>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ภาคการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกภาคการศึกษา</InputLabel>
                <Select
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
              </FormControl>
            </Grid>




            <Grid item xs={3}>
              <div className={classes.paper}>วิชา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกวิชา</InputLabel>
                <Select
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
              </FormControl>
            </Grid>



            <Grid item xs={3}>
                  <div className={classes.paper}>เกรด</div>
                </Grid>
                <Grid item xs={9}>
                  <TextField variant="outlined" className={classes.textField}  
                      name="grade"
                      value={results.grade || ''} // (undefined || '') = ''
                      onChange={handleInputGrade}>
                                                                            
                  </TextField>
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
