import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
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
  isWidthDown,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntProfessorship } from '../../api/models/EntProfessorship';
import { EntProfessor } from '../../api/models/EntProfessor';
import { EntPrefix } from '../../api/models/EntPrefix';
import { EntFaculty } from '../../api/models/EntFaculty';

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

interface Professors {
    tel: string;
    email: string;
    name: string;
    prefix: number;
    faculty: number;
    professorship: number;
    
}

const Professors: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();


  const [professors, setProfessors] = React.useState<Partial<Professors>>({});
  const [prefixs, setPrefixs] = React.useState<EntPrefix[]>([]);
  const [facultys, setFacultys] = React.useState<EntFaculty[]>([]);
  const [professorships, setProfessorships] = React.useState<EntProfessorship[]>([]);
  const [showInputError, setShowInputError] = React.useState(false); // for error input show
  const [emailError, setEmailError] = React.useState('');
  const [telError, setTelError] = React.useState('');
  



  const getPrefix = async () => {
    const res = await http.listPrefix({ limit: 4, offset: 0 });
    setPrefixs(res);
  };

  const getFaculty = async () => {
    const res = await http.listFaculty({ limit: 11, offset: 0 });
    setFacultys(res);
  };

  const getProfessorship = async () => {
    const res = await http.listProfessorship({ limit: 5, offset: 0 });
     setProfessorships(res);
  };

  
  useEffect(() => {
    getPrefix();
    getFaculty();
    getProfessorship();
  
    }, []);

 
   
    
    // set data to object Professor
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof Professors;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setProfessors({ ...professors, [name]: value });
    console.log(professors);
  };
  // ฟังก์ชั่นสำหรับ validate อีเมล
  const validateEmail = (email: string) => {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
  }

  // ฟังก์ชั่นสำหรับ validate เบอร์มือถือ
  const validateTel = (val: string) => {
    return val.length == 10 ? true : false;
  }

  
  

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'tel':
        validateTel(value) ? setTelError('') : setTelError('เบอร์มือถือ 10 หลัก');
        return;
      case 'email':
        validateEmail(value) ? setEmailError('') : setEmailError('รูปแบบอีเมลไม่ถูกต้อง')
        return;
      default:
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'tel':
        alertMessage("error","เบอร์มือถือ 10 หลัก");
        return;
      case 'phone':
          alertMessage("error","เบอร์โทรศัพท์ที่ทำงาน 9 หลัก");
          return;
      case 'email':
        alertMessage("error","รูปแบบอีเมลไม่ถูกต้อง");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }
    function clear() {
      setProfessors({});
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

    function save() {
      setShowInputError(true);
      let {prefix, faculty, professorship, name, tel, email} = professors;
      if (!prefix || !faculty || !professorship || !professors || !name || !tel || !email) {
        Toast.fire({
          icon: 'error',
          title: 'กรุณากรอกข้อมูลให้ครบถ้วน',
        });
        return;
      }
      const apiUrl = 'http://localhost:8080/api/v1/professors';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(professors),
      };
      console.log(professors);
  
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  };
  
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
      title={'Professor '}
      subtitle='Professor register system'
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
              <div className={classes.paper}>คำนำหน้า</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกคำนำหน้า</InputLabel>
                <Select
                  required
                  error={!professors.prefix && showInputError}
                  name="prefix"
                  //id="prefix"
                  value={professors.prefix || ''}
                  onChange={handleChange}
                >
                {
           
                  prefixs.map(item => {
                  return (
                <MenuItem key={item.id} value={item.id}>{item.prefix}</MenuItem>
                          );
                })
                }
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ - นามสุกล</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  required
                  error={!professors.name && showInputError}
                  label="กรอกชื่อ - นามสกุล"
                  name="name"
                  id="name"
                  value={professors.name || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>ตำแหน่ง</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกตำแหน่ง</InputLabel>
                <Select
                  required
                  error={!professors.professorship && showInputError}
                  name="professorship"
                  value={professors.professorship || ''}
                  onChange={handleChange}
         
                >
                  {professorships.map(item => {
                    return (
                  <MenuItem key={item.id} value={item.id}>{item.professorship}</MenuItem>
                          );
                  })
                  }    
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>สำนักวิชา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสำนักวิชา</InputLabel>
                <Select
                  required
                  error={!professors.faculty && showInputError}
                  name="faculty"
                  //id="faculty"
                  value={professors.faculty || ''}
                  onChange={handleChange}
         
                >
                {facultys.map(item => {
                    return (
                <MenuItem key={item.id} value={item.id}>{item.faculty}</MenuItem>
                            );
                })
                }
        </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์มือถือ</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  error = {telError ? true : false}
                  label="กรอกเบอร์มือถือ"
                  inputProps={{ maxLength: 10 }}
                  helperText= {telError}
                  name="tel"
                  id="tel"
                  value={professors.tel || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>E-mail</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  error = {emailError ? true : false}
                  label="E - mail"
                  name="email"
                  id="email"
                  helperText= {emailError}
                  value={professors.email || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
                
              >
                บันทึกข้อมูล
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};


export default Professors;