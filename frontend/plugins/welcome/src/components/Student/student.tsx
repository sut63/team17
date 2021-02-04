import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Alert from '@material-ui/lab/Alert';
import Swal from 'sweetalert2'; // alert
import {
    Container,
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    TextField,
    Button,
    CardMedia,
    Snackbar,
    Avatar,
  } from '@material-ui/core';
import{
  Content,
    InfoCard,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
import  { DefaultApi }  from '../../api/apis';
import { EntGender, EntDegree, EntPrefix, EntProvince} from '../../api/models/';
import { Cookies } from '../../Cookie';

interface Students {
  /**
   * 
   * @type {string}
   * @memberof ControllersStudent
   */
  addr?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  degree?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  district?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersStudent
   */
  email?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersStudent
   */
  fname?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersStudent
   */
  lname?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  postal?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  province?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersStudent
   */
  school?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  sex?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  tel?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  title?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersStudent
   */
  zone?: number;
}

const StudentUI: FC<{}> = () => {

  const api = new DefaultApi();
  const [pref,Setpref] = useState<EntPrefix[]>([]);
  const [prov,Setprov] = useState<EntProvince[]>([]);
  const [gend,Setgend] = useState<EntGender[]>([]);
  const [degr,Setdegr] = useState<EntDegree[]>([]);

  const getPref = async () => {
    const res = await api.listPrefix({ limit: 10, offset: 0 });
    Setpref(res);
  };
  const getProv= async () => {
    const res = await api.listProvince({ limit: 10, offset: 0 });
    Setprov(res);
  };
  const getGend = async () => {
    const res = await api.listGender({ limit: 10, offset: 0 });
    Setgend(res);
  };
  const getDegr = async () => {
    const res = await api.listDegree({ limit: 10, offset: 0 });
    Setdegr(res);
  };

  useEffect(() => {
    getDegr();
    getGend();
    getPref();
    getProv();
    }, []);

    const [Student, setStudent] = React.useState< Partial<Students>>({});
       
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

    const alertMessage = (icon: any, title: any) => {
      Toast.fire({
        icon: icon,
        title: title,
      });
    } 

    const [fnameError, setfnameError] = React.useState('');
    const [lnameError, setlnameError] = React.useState('');
    const [schoolError, setschoolError] = React.useState('');
    const [emailError, setemailError] = React.useState('');
    const [addrError, setaddrError] = React.useState('');
    const [telError, settelError] = React.useState('');

    const validateEmail = (email: string) => {
      const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    }

    const validateTextfield = (val: string) => {
      return val.match("[A-Za-z]");
    }

    const validateAddress = (val: string) => {
      return val.match("[A-Za-z0-9\\/]");
    }

    const validateTel = (val: string) => {
      return val.match("\\d{10}");
    }

    const checkPattern  = (id: string, value: string) => {
      switch(id) {
        case 'fname':
          validateTextfield(value) ? setfnameError('') : setfnameError('ชื่อใช้ภาษาอังกฤษเท่านั้น');
          return;
        case 'lname':
          validateTextfield(value) ? setlnameError('') : setlnameError('นามสกุลใช้ภาษาอังกฤษเท่านั้น');
          return;
        case 'school':
          validateTextfield(value) ? setschoolError('') : setschoolError('ชื่อโรงเรียนใช้ภาษาอังกฤษเท่านั้น');
          return;
        case 'tel':
          validateTel(value) ? settelError('') : settelError('กรุณาใส่เบอร์โทรศัพท์ให้ถูกต้อง');
          return;
        case 'addr':
          validateAddress(value) ? setaddrError('') : setaddrError('รูปแบบที่อยู่ไม่ถูกต้อง');
          return;
        case 'email':
          validateEmail(value) ?  setemailError('') :  setemailError('รูปแบบอีเมลไม่ถูกต้อง');
          return;
        default:
          return;
      }
    }

    const checkCaseSaveError = (field: string) => {
      switch(field) {
        case 'fname':
          alertMessage("error","ชื่อใช้ภาษาอังกฤษเท่านั้น");
          return;
        case 'lname':
          alertMessage("error","นามสกุลใช้ภาษาอังกฤษเท่านั้น");
          return;
        case 'schoolname':
          alertMessage("error","ชื่อโรงเรียนใช้ภาษาอังกฤษเท่านั้น");
          return;
        case 'telephone':
          alertMessage("error","กรุณาใส่เบอร์โทรศัพท์ให้ถูกต้อง");
          return;
        case 'recent_address':
          alertMessage("error","รูปแบบที่อยู่ไม่ถูกต้อง");
          return;
        case 'email':
          alertMessage("error","รูปแบบอีเมลไม่ถูกต้อง");
          return;
        default:
          alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
          return;
      }
    }

    const h = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name as keyof typeof Student;
      const { value } = event.target;
      setStudent({ ...Student, [name]: value });
      console.log(Student);
    };

    const i = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
      const id = event.target.id as keyof typeof Student;
      const { value } = event.target;
      const validateValue = value.toString()
      checkPattern(id, validateValue)
      setStudent({ ...Student, [id]: value });
    };
    
    function clear() {
      setStudent({});
    }

    console.log(Student);
    function save() {
      const apiUrl = 'http://localhost:8080/api/v1/students';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(Student),
      };
      console.log(Student);
  
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data.status);
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
    }

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
      title={'Student Management'}
      subtitle='Student Registration Department'>
        <Avatar alt="Remy Sharp"/>
        <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
        <Button variant="text" color="secondary" size="large"
          onClick={Clears} > Logout </Button>
    </Header>
    <Content>
      <ContentHeader title="Student Background"></ContentHeader>
      <TableContainer component={Paper}>
       <Table>
         <TableBody>
           <TableCell>

                <Grid container spacing={1}>
                  <Grid item xs={2}>
                    <b>Gender</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="sex" value={Student.sex||''}
                    onChange={h}>
                {gend.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.gender}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                  <b>Title</b>
                <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="title" value={Student.title||''}
                    onChange={h}>
                {pref.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.prefix}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                </div>
                  </Grid>
                </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>School</b>
                  <div>
                  <TextField variant='outlined' type='string' id="school" value={Student.school||''} helperText= {schoolError} error= {schoolError? true:false}
                    onChange={i}/>
                  </div>  
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>Province</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="province" value={Student.province||''}
                    onChange={h}>
                  {prov.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.province}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                  </div>
                </Grid>
              </Grid>
              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>Postal Code</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="postal" value={Student.postal||''}
                    onChange={h}>
                  {prov.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.postal}
                    </MenuItem>
                    );
                  })}
                  </Select>
                  </FormControl>
                  </div>
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={12}>

                </Grid>
              </Grid>

           </TableCell>
           
           <TableCell>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>First Name</b>
                  <div>
                  <TextField variant='outlined' type='string' id="fname" value={Student.fname||''} helperText= {fnameError} error= {fnameError ? true:false}
                    onChange={i}/>
                  </div>  
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>Degree</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="degree" value={Student.degree||''}
                    onChange={h}>
                  {degr.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.degree}
                    </MenuItem>
                    );
                  })}
                  </Select>
                  </FormControl>
                  </div>
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>District</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="district" value={Student.district||''}
                    onChange={h}>
                  {prov.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.district}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                  </div>
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>Telephone Number</b>
                  <div>
                  <TextField variant='outlined' type='tel' id="tel" value={Student.tel||''} helperText= {telError} error= {telError ? true:false}
                    onChange={i}/>
                  </div>
                </Grid>
              </Grid>

           </TableCell>
           
           <TableCell>
            <Grid container spacing={1}>
              <Grid item xs={4}>
                <b>Last Name</b>
                <div>
                <TextField variant='outlined' type='string' id="lname" value={Student.lname||''} helperText= {lnameError} error= {lnameError ? true:false}
                    onChange={i}/>
                </div>
              </Grid>
            </Grid>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                  <b>Recent Address</b>

                  <div>
                  <TextField variant='outlined' type='string' id="addr" value={Student.addr||''} helperText= {addrError} error= {addrError ? true:false}
                    onChange={i}/>
                  </div>
              </Grid>
            </Grid>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                  <b>Subdistrict</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="zone" value={Student.zone||''}
                    onChange={h}>
                  {prov.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.subdistrict}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                  </div>
              </Grid>
            </Grid>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                <b>Email</b>
                <div>
                <TextField variant='outlined' type='email' id="email" value={Student.email||''} helperText= {emailError} error= {emailError? true:false}
                    onChange={i}/>
                </div>
              </Grid>
            </Grid>

           </TableCell>
         </TableBody>
            <Grid container spacing={1}>
              <Grid item xs={4}>

                </Grid>
              </Grid>
         <TableBody>
           <TableCell></TableCell>
           <TableCell>
            <Button variant='contained' color='primary'  onClick={() => {
                  save(); 
              }}>Save</Button>
           </TableCell>
           <TableCell></TableCell>
         </TableBody>
         </Table>
     </TableContainer>
     </Content>
  </Page>
);
};export default StudentUI;