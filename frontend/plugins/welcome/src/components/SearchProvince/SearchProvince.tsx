import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
    MuiPickersUtilsProvider,
    KeyboardTimePicker,
    KeyboardDatePicker,
} from '@material-ui/pickers';
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
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntProvince } from '../../api';
import { Cookies } from '../../Cookie';



// alert setting
const Toast = Swal.mixin({
    toast: true,
    position: 'top' ,
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
});

const alertMessage = ( icon: any, title: any ) => {
    Toast.fire({
        icon: icon,
        title: title,
    });
}

// header Css
const HeaderCustom = {
    minHeight: '50px',
};


// Css Style
const useStyles = makeStyles(theme => ({
    root: {
        flexGrown: 1,
    },
    paper: {
        marginTop: theme.spacing(2),
        marginBottom: theme.spacing(2),
    },
    formControl: {
        width: 300,
        marginLeft: -70,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    bottom: {
        marginLeft: theme.spacing(0),
        marginTop: theme.spacing(2),
        marginBottom: theme.spacing(2),
    },
    divider: {
        margin: theme.spacing(2,0),
    },
    table: {
        minWidth: 650,
    },
}));


const SearchProvince: FC<{}> = () => {

    const classes = useStyles();
    const api = new DefaultApi();
    const [province, setProvince] = useState<EntProvince[]>([]);
    const [search,setsearch] = useState("");

    const sh = (event: React.ChangeEvent<{ value: unknown }>) => {
        setsearch(event.target.value as string);
        Clear();
      };

      const getSt = async () => {
        const res = await api.listProvince({limit: 10, offset: 0})
        setProvince(res)
        var boo = false
        var arr = ""
        for(let i = 0; i < res.length ; i++){
          if(res[i].province === search){ 
            boo = true
            arr += "f"
            console.log(res[i].id)
            console.log(res[i].edges)
          }
          else{
            boo = false
          }
        }
        if(boo!=false||arr!=""){
          alertMessage("success","ค้นหาสำเร็จ")
        }else{
          alertMessage("error","ค้นหาไม่พบ")
        }
      }


    // function clear
    function Clear (){
        setProvince([])
    }

    //cookie logout
    var cook = new Cookies()
    var cookieName = cook.GetCookie()

    function Clears() {
        cook.ClearCookie()
        window.location.reload(false)
    }

    return(
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
                <TextField value={search||''} type='string' name='search' label="Enter First Name" onChange={sh}/>
              </TableCell>
              <TableCell>
      <Button  
          onClick={() => {
            getSt();
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
                    <TableCell align='center'><b>Continent</b></TableCell>
                    <TableCell align='center'><b>Country</b></TableCell>
                    <TableCell align='center'><b>Region</b></TableCell>
                    <TableCell align='center'><b>Province</b></TableCell>
                    <TableCell align='center'><b>District</b></TableCell>
                    <TableCell align='center'><b>Subdistrict</b></TableCell>
                    <TableCell align='center'><b>Postal</b></TableCell>
                </TableRow>
                    {province.filter((filter)=> (filter.province === search)).map((item) => (
                <TableRow key={item.id}>
                    <TableCell align='center'><b>{item.edges?.provCont?.continent}</b></TableCell>
                    <TableCell align='center'><b>{item.edges?.provCoun?.country}</b></TableCell>
                    <TableCell align='center'><b>{item.edges?.provRegi?.name}</b></TableCell>
                    <TableCell align='center'><b>{item.province}</b></TableCell>
                    <TableCell align='center'><b>{item.district}</b></TableCell>
                    <TableCell align='center'><b>{item.subdistrict}</b></TableCell>
                    <TableCell align='center'><b>{item.postal}</b></TableCell>
                </TableRow>))}
            </Table>
            </TableContainer>
        </Content>
    </Page>
    
  );

};

export default SearchProvince;