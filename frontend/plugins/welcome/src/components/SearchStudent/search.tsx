import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Swal from 'sweetalert2'; // alert
import {
    TextField,
    Button,
    Typography,
    makeStyles,
    Avatar,
    Link,
  } from '@material-ui/core';
import{
    Content,
    Header,
    Page,
    pageTheme,
} from '@backstage/core';
import { Cookies } from '../../Cookie';
import { DefaultApi } from '../../api/apis'; 
import { EntStudent } from '../../api/models';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  table: {

    minWidth: 650,
 
  },
}));

const StudentSearchUI: FC<{}> = () => {
  
  const api = new DefaultApi();
  const classes = useStyles();
  const [fname,setfname] = useState("")
  const [lname,setlname] = useState("")
  const [tf,settf] = useState("")
  const [tl,settl] = useState("")
  const [found,setfound] = useState(Boolean);
  const [st,setst] = useState<EntStudent[]>([]);

  const fnh = (event: React.ChangeEvent<{ value: unknown }>) => {
    setfname(event.target.value as string);
  };
  const lnh = (event: React.ChangeEvent<{ value: unknown }>) => {
    setlname(event.target.value as string);
  };

  const getSt = async () => {
    const res = await api.listStudent({ limit: 10, offset: 0 });
    setst(res)
  };

  useEffect(() => {
    getSt();
  }, []);

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

  const CheckFound = async () => {
    let x =''
    st.forEach((item) =>{
      if(item.fname===fname||item.lname===lname){
        x += 'f'
      }
    })
    if(x!=''){
      setfound(true);
      settf(fname);
      settl(lname);
      alertMessage('success','ค้นหาสำเร็จ')
    }
    else{
      setfound(false);
      alertMessage('error','ค้นหาไม่พบ')
  }};
  
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
    <TableContainer component={Paper}>
       <Table>
         <TableBody>
           <TableRow>
             <TableCell>
                <TextField value={fname||''} type='string' name='f' label="Enter First Name" onChange={fnh}/>
              </TableCell>
              <TableCell>
                <TextField value={lname||''} type='string' name='l' label="Enter Last Name" onChange={lnh}/>
              </TableCell>
              <TableCell>
      <Button  
          onClick={() => {
            CheckFound();
          }}
          variant="contained" 
          color="primary" 
          > 
          Search
          </Button>
              </TableCell>
            </TableRow>
          </TableBody>
       </Table>
            <Table className={classes.table}>
             <TableRow>
             <TableCell align='center'><b>Title</b></TableCell>
             <TableCell align='center'><b>First Name</b></TableCell>
             <TableCell align='center'><b>Last Name</b></TableCell>
             <TableCell align='center'><b>Sex</b></TableCell>
             <TableCell align='center'><b>Degree</b></TableCell>
             <TableCell align='center'><b>Telephone</b></TableCell>
             <TableCell align='center'><b>Email</b></TableCell>
            </TableRow>
            {found
            ?
            st.filter((filter=>filter.fname===tf||filter.lname===tl)).map((item:EntStudent)=>(
          <TableRow key={item.id}>
          <TableCell align='center'><b>{item.edges?.studPref?.prefix}</b></TableCell>
          <TableCell align='center'><b>{item.fname}</b></TableCell>
          <TableCell align='center'><b>{item.lname}</b></TableCell>
          <TableCell align='center'><b>{item.edges?.studGend?.gender}</b></TableCell>
          <TableCell align='center'><b>{item.edges?.studDegr?.degree}</b></TableCell>
          <TableCell align='center'><b>{item.telephone}</b></TableCell>
          <TableCell align='center'><b>{item.email}</b></TableCell>
         </TableRow>))
            :
            null
        }
            </Table>
    </TableContainer>
    </Content>
    </Page>
  );
};

export default StudentSearchUI;