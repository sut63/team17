import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2';


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
import { DefaultApi } from '../../api/apis/DefaultApi'; // Api Gennerate From Command

import { EntPlace } from '../../api/models/EntPlace'; // import interface Place
import { EntAgency } from '../../api/models/EntAgency'; // import interface Agency
import { EntYear } from '../../api/models/EntYear'; // import interface Year
import { EntStudent } from '../../api/models/EntStudent'; // import interface Student
import { EntTerm } from '../../api/models/EntTerm'; // import interface Term
import { Cookies } from '../../Cookie';

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

interface Activity{
  activity: Date;
  activityname: String;
    place: number;
    agency: number;
    added:  String;
    hours: String;
    student: number;
    year: number;
    term: number;
    
}

const Activity: FC<{}> = () => {
    

    const classes = useStyles();
    const http = new DefaultApi();

    

    const [activity, setActivitys] = React.useState<Partial<Activity>>({});
    const [place, setPlaces] = React.useState<EntPlace[]>([]);
    const [agency, setAgencys] = React.useState<EntAgency[]>([]);
    const [year, setYears] = React.useState<EntYear[]>([]);
    const [student, setStudents] = React.useState<EntStudent[]>([]);
    const [term, setTerms] = React.useState<EntTerm[]>([]);
    const [showInputError, setShowInputError] = React.useState(false); // for error input show

    
 
    const getPlace = async () => {
        const res = await http.listPlace({ limit: 10, offset: 0 });
        setPlaces(res);
      };

    const getAgency = async () => {
        const res = await http.listAgency({ limit: 10, offset: 0 });
        setAgencys(res);
      };

    const getYear = async () => {
        const res = await http.listYear({ limit: 10, offset: 0 });
        setYears(res);
      };

    const getStudent = async () => {
        const res = await http.listStudent({ limit: 10, offset: 0 });
        setStudents(res);
      };
    const getTerm = async () => {
        const res = await http.listTerm({ limit: 10, offset: 0 });
        setTerms(res);
    };
    

    // Lifecycle Hooks
    useEffect(() => {
        getPlace();
        getAgency();
        getYear();
        getStudent();
        getTerm();
        
    }, []);

    // set data to object activitys
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
        const name = event.target.name as keyof typeof Activity;
        const { value } = event.target;
        setActivitys({ ...activity, [name]: value });
        console.log(activity);
        };

    // clear input form
    function clear() {
        setActivitys({});
        setShowInputError(false);
    }

    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });

    // function save data
    function save() {
      setShowInputError(true);
      let {hours, activityname, added} = activity;
      if (!hours || !activityname || !added) {
          Toast.fire({
            icon: 'error',
            title: 'กรุณากรอกข้อมูลให้ครบถ้วน',
          });
      return;
      }

        const apiUrl = 'http://localhost:8080/api/v1/activitys';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(activity),
        };

        console.log(activity); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
        .then(response => {
          console.log(response)
          response.json()

        if (response.ok === true) {
          clear();
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
        })
    }


    //cookie logout
  var cook = new Cookies()
  var cookieName = cook.GetCookie()

  function Clears() {
    cook.ClearCookie()
    window.location.reload(false)
  }

  

    function redirecLogOut() {
      // redirect Page ... http://localhost:3000/
      window.location.href = "http://localhost:3000/";
    }

    return (
        <Page theme={pageTheme.home}>
          <Header  
          
          title={'Student Management'}
          subtitle='Student Registration Department'
          >
            <Avatar alt="Remy Sharp"/>
            <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
            <Button variant="text" color="secondary" size="large"
              onClick={Clears} > Logout </Button>
          </Header>
          <Content>
            <Container maxWidth="sm">

            <Grid container spacing={3}>
                <Grid item xs={3}>
                   <div className={classes.paper}>ชื่อนักศึกษา</div>
                </Grid>
                <Grid item xs={9}>
            
                  <FormControl variant="outlined" className={classes.formControl}>
                    <InputLabel>เลือกชื่อนักศึกษา</InputLabel>
                    <Select
                      name="student"
                      value={activity.student || ''} // (undefined || '') = ''
                      onChange={handleChange}
                    >
                      {student.map(item => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                            {item.fname}
                          </MenuItem>
                        );
                      })}
                    </Select>
                  </FormControl>
                </Grid>
                </Grid>
                <Grid container spacing={3}>
                <Grid item xs={12}></Grid>
                <Grid item xs={3}>
                  <div className={classes.paper}>กิจกรรม</div>
                </Grid>
                <Grid item xs={9}>
                <form className={classes.container} noValidate>
                <TextField
                  label="ชื่อกิจกรรม"
                  name="activityname"
                  value={activity.activityname || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                ></TextField>
              </form>

                </Grid>

                <Grid item xs={3}>
                  <div className={classes.paper}>สถานที่จัดกิจกรรม</div>
                </Grid>
                <Grid item xs={9}>
                  <FormControl variant="outlined" className={classes.formControl}>
                    <InputLabel>เลือกสถานที่</InputLabel>
                    <Select
                      name="place"
                      value={activity.place || ''} // (undefined || '') = ''
                      onChange={handleChange}
                    >
                      {place.map(item => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                            {item.pLACE}
                          </MenuItem>
                        );
                      })}
                    </Select>
                  </FormControl>
                </Grid>
    
                <Grid item xs={3}>
                  <div className={classes.paper}>หน่วยงานที่จัดกิจกรรม</div>
                </Grid>
                <Grid item xs={9}>
                  <FormControl variant="outlined" className={classes.formControl}>
                    <InputLabel>เลือกหน่วยงาน</InputLabel>
                    <Select
                      name="agency"
                      value={activity.agency || ''} // (undefined || '') = ''
                      onChange={handleChange}
                    >
                      {agency.map(item => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                            {item.aGENCY}
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
                        name="year"
                        value={activity.year || ''} // (undefined || '') = ''
                        onChange={handleChange}
                      
                    >
                      {year.map(item => {
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
                        name="term"
                        value={activity.term || ''} // (undefined || '') = ''
                        onChange={handleChange}
                      
                    >
                      {term.map(item => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                            {item.semester}
                          </MenuItem>
                        );
                      })}
                    </Select>
                  </FormControl>
                </Grid>

                <Grid item xs={12}></Grid>
                <Grid item xs={3}>
                  <div className={classes.paper}>ชั่วโมงจิตอาสา</div>
                </Grid>
                <Grid item xs={9}>
                <form className={classes.container} noValidate>
                <TextField
                  label="ชั่วโมงจิตอาสา"
                  name="hours"
                  value={activity.hours || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                  
                />
              </form>

                </Grid>      

    
                <Grid item xs={3}>
                  <div className={classes.paper}>วันที่ / เวลา</div>
                </Grid>
                <Grid item xs={9}>
                  <form className={classes.container} noValidate>
                    <TextField 
                      label="เลือกเวลา"
                      name="added"
                      type="datetime-local"
                      value={activity.added || ''}
                      className={classes.textField}
                      InputLabelProps={{
                          shrink: true,
                      }}
                      onChange={handleChange}
                    ></TextField>
                    
                  </form>
                </Grid>
    
                <Grid item xs={4}></Grid>
                <Grid item xs={7}>
                  <Button
                    variant="contained"
                    color="primary"
                    size="large"
                    startIcon={<SaveIcon />}
                    onClick={save}
                  >
                    บันทึกรายการ
                  </Button>
                </Grid>
              </Grid>
            </Container>
          </Content>
        </Page>
      );
  
};

export default Activity;


