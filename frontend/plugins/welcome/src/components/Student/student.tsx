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
  tel?: number;
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

const handleClose = (event: React.SyntheticEvent | React.MouseEvent, reason?: string) => {
  if (reason === 'clickaway') {
    return;
  }
  setSuccess(false);
  setFail(false);
};

  const [success, setSuccess] =useState(false);
  const [fail, setFail] =useState(false);
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

    const h = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name as keyof typeof Student;
      const { value } = event.target;
      setStudent({ ...Student, [name]: value });
      console.log(Student);
    };

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
            Toast.fire({
              icon: 'error',
              title: 'บันทึกข้อมูลไม่สำเร็จ',
            });
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
      subtitle='Student Registration Department'
    >
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
                <Select name="sex" id='sex' value={Student.sex||''}
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
                <Select name="title" id='title' value={Student.title||''}
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
                  <TextField variant='outlined' type='string' name="school" value={Student.school||''}
                    onChange={h}/>
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
                  <TextField variant='outlined' type='string' name="fname" value={Student.fname||''}
                    onChange={h}/>
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
                  <TextField variant='outlined' type='number' name="tel" value={Student.tel||''}
                    onChange={h}/>
                  </div>
                </Grid>
              </Grid>

           </TableCell>
           
           <TableCell>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                <b>Last Name</b>
                <div>
                <TextField variant='outlined' type='string' name="lname" value={Student.lname||''}
                    onChange={h}/>
                </div>
              </Grid>
            </Grid>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                  <b>Recent Address</b>
                  <div>
                  <TextField variant='outlined' type='string' name="addr" value={Student.addr||''}
                    onChange={h}/>
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
                <TextField variant='outlined' type='email' name="email" value={Student.email||''}
                    onChange={h}/>
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
            <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: 'top', horizontal: 'center' }}>
              <Alert onClose={handleClose} severity="error">
                Error
              </Alert>
            </Snackbar>
            <Snackbar open={success} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: 'top', horizontal: 'center' }}>
              <Alert onClose={handleClose} severity="success">
                Successful
              </Alert>
            </Snackbar>
           </TableCell>
           <TableCell></TableCell>
         </TableBody>
         </Table>
     </TableContainer>
     </Content>
  </Page>
);
};export default StudentUI;