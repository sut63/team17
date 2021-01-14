import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Alert from '@material-ui/lab/Alert';
import { Link as RouterLink } from 'react-router-dom';
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
    Typography,
    makeStyles,
    CssBaseline,
    Avatar,
    Link,
  } from '@material-ui/core';
import{
  Content,
    InfoCard,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { Cookies } from '../../Cookie';
import { DefaultApi } from '../../api/apis'; 
import { EntEmp } from '../../api';

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
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {

  const classes = useStyles();
  const api = new DefaultApi();
  var ck = new Cookies();
  var check : boolean
  const [path, setPath] = React.useState("");

 const [emp,setEmp] = React.useState<EntEmp[]>([])
  const listEmp = async() => {
        const res = await api.listEmp({})
        setEmp(res)
  }

  useEffect(() => {
    listEmp();
},[])

  
 // setUser
  const [user, setUser] = React.useState("")
  const handleUser = (event : any) => {
      setUser(event.target.value)
  }

  // setPass
  const [pass, setPass] = React.useState("")
  const handlePass = (event : any) => {
      setPass(event.target.value)
  }

  // handleCookies
  function handleCookies() {
    check = ck.CheckLogin(emp,user,pass)
    console.log("check => "+check)
    if(check === true){
      setPath("/WelcomePage")
      ck.SetCookie("user_email",user,30)
      ck.SetCookie("user_id",ck.SetID(emp,user,pass),30)
      ck.SetCookie("user_role","ทะเบียน",30)
      window.location.reload(false)
    }else if(check === false){
      alert("The wrong password or email was entered.!!!")
      setPath("/")
    }
    console.log(user);
    console.log(pass);
  }

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={handleUser}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={handlePass}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleCookies}
            component={RouterLink}
            to={path}
          >
            Sign In
          </Button>
        </form>
      </div>
    </Container>
  );
};

export default SignIn;