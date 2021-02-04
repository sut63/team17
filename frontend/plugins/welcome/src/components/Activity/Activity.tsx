import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2';
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';

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
  
  activityname: String;
    place: number;
    agency: number;
    added:  Date;
    hours: number;
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
  activity.hours = Number(activity.hours);
  

  
    //cookie logout
    var cook = new Cookies()
    var cookieName = cook.GetCookie()
  
    function Clears() {
      cook.ClearCookie()
      window.location.reload(false)
    }
  

  
   
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
    const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
        const name = event.target.name as keyof typeof Activity;
        const { value } = event.target;
        setActivitys({ ...activity, [name]: value });
        const validateValue = value.toString() 
        checkPattern(name, validateValue)
        console.log(activity);
    };

    // clear input form
    function clear() {
        setActivitys({});
        setShowInputError(false);
    }
    
  // alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);
    

    

   // สำหรับตรวจสอบความถูกต้อง
   const [activitynameError, setActivityNameError] = React.useState('');
   const [hoursError, setHoursError] = React.useState('');
   const [addedError, setAddedError] = React.useState('');
   const [errors, setError] = React.useState(String);


    //validate ชื่อกิจกรรม
    const validateActivityName = (val: string) => {
      return val.length < 5 ? false : true;
    }

    //validate จำนวนชั่วโมงจิตอาสา
    const validateHours = (val: string) => {
      return val;
    }
  //validate วันที่ เวลา
   const validateAdded = (val: string) => {
      return val;
    }

    
  

    // checkPattern
    const checkPattern  = (id: string, value: string) => {
      switch(id) { 
        case 'activityname':
          validateActivityName(value) ? setActivityNameError('') : setActivityNameError('ใส่ตัวอักษรไม่น้อยกว่า 5 ตัวอักษร');
          return;
        case 'hours':
          validateHours(value) ? setHoursError('') : setHoursError('จำนวนชั่วโมงจิตอาสาต้องมากกว่า 0');
          return;
        case 'added':
          validateAdded(value) ? setAddedError('') : setAddedError('ใส่วันที่จัดกิจกรรม');
          return;     
        default:
          return;
  }
}

    const checkerror = (s: string) => {
      switch(s) {
        case 'activityname':
          setError("ใส่ชื่อกิจกรรมใหม่");
          return;
          
        case 'hours':
          setError("จำนวนชั่วโมงจิตอาสาต้องเป็นจำนวนเต็มบวก");

        case 'added':
          setError("กรุณาใส่วันที่จัดกิจกรรม");
          return;
      default:
          setError("กรุณากรอกข้อมูลให้ครบถ้วน");
          return;
      }
    }

  //close alert 
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setFail(false);
    setOpen(false);
  };
    

    // function save data
    function save() {
     
      const apiUrl = 'http://localhost:8080/api/v1/activitys';
      const requestOptions = {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(activity),
      };

      console.log(activity); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console


     // setShowInputError(true);
     // let {hours, activityname, added} = activity;
     // if (!hours || !activityname || !added) {
     //     Toast.fire({
     //       icon: 'error',
     //       title: 'กรุณากรอกข้อมูลให้ครบถ้วน',
     //     });
     //       return;
    //  }

        

        fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (data.status === true) {
            clear();
            setOpen(true);
          } else {
            checkerror(data.error.Name)
            setFail(true);
          }
        });
    }


  

    return (
        <Page theme={pageTheme.home}>
          <Header  
          
          title={'Student Activity '}
          subtitle='Student Activities Department'
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
            
                  <FormControl 
                    variant="outlined" 
                    className={classes.formControl}>
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
                  error = {activitynameError ? true : false}
                  helperText={activitynameError}
                  variant="outlined"
                  name="activityname"
                  type="string"
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
                  error = {hoursError ? true : false}
                  helperText={hoursError}
                  variant="outlined"
                  label="ชั่วโมงจิตอาสา"
                  name="hours"
                  type="number"
                  value={activity.hours || ''} // (undefined || '') = ''
                  className={classes.textField}
                  
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
                     error = {addedError ? true : false}
                     helperText={addedError}
                      variant="outlined"
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
                    บันทึก
                  </Button>
                </Grid>
              </Grid>


              <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              บันทึกข้อมูลสำเร็จ!
        </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              {errors}
        </Alert>
          </Snackbar>


            </Container>
          </Content>
        </Page>
      );
  
};

export default Activity;


