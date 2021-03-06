import React, { useState, useEffect } from 'react';
import {
    Content,
    Page,
    pageTheme,
    InfoCard,
    Header
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import { Alert } from '@material-ui/lab';
import { Avatar } from '@material-ui/core';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import SearchIcon from '@material-ui/icons/Search';
import { Cookies } from '../../Cookie';
import { EntActivity } from '../../api/models/EntActivity';
import Swal from 'sweetalert2'; // alert


const Toast = Swal.mixin({
  toast: true,
  position: 'center',
  width: '400px',
  padding: '100px',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 400 ,
      marginLeft:7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    center: {
      marginTop: theme.spacing(2),
      marginLeft: theme.spacing(23),
    },
    cardtable: {
      marginTop: theme.spacing(2),
    },
    fieldText: {
      width: 200,
      marginLeft:7,
    },
    fieldLabel: {
      marginLeft:8,
      marginRight: 20,
    },
    table: {
        minWidth: 650,
    }
  }),
);

const searchcheck = {
   activitysearchcheck: true,
   
}

export default function SearchActivity() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [activity, setActivity] = useState<EntActivity[]>([]);
    

    const [activitysearch, setActivitysearch] = useState(String);
    

    const SearchActivity = async () => {
        const res = await api.getActivity({fname: activitysearch});
        setActivity(res)
        var  lenActivity: number
        lenActivity = res.length
        if (lenActivity > 0) {
          //setOpen(true)
          Toast.fire({
            icon: 'success',
            title: 'ค้นหาข้อมูลสำเร็จ',
          })
        } else {
          //setFail(true)
          Toast.fire({
            icon: 'error',
            title: 'ค้นหาข้อมูลไม่พบ',
          })
        }
      }
    


  
    const ResetSearchCheck = () => {
        searchcheck.activitysearchcheck = true;
        
    }

    const ActivitySearchhandleChange = (event: any) => {
        setActivitysearch(event.target.value as string);
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
          
          title={'Search Student Activity '}
          subtitle='Student Activities Department'
          >
            <Avatar alt="Remy Sharp"/>
            <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
            <Button variant="text" color="secondary" size="large"
              onClick={Clears} > Logout </Button>
          </Header>
            <Content>
            <InfoCard title="ค้นหาประวัติกิจกรรมนักศึกษา">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>ชื่อนักศึกษา</strong></div>
                    <TextField
                        id="name"
                        type="string"
                        size="medium"
                        value={activitysearch}
                        onChange={ActivitySearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

           

                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchActivity();
                }}
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                ค้นหา
             </Button>
             
        
             <div><br></br></div>
             <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                <TableCell align="center">ชื่อนักศึกษา</TableCell>
                                <TableCell align="center">ชื่อกิจกรรม</TableCell>
                                <TableCell align="center">สถานที่จัดกิจกรรม</TableCell>
                                <TableCell align="center">หน่วยงานที่จัดกิจกรรม</TableCell>
                                <TableCell align="center">จำนวนชั่วโมงจิตอาสา</TableCell>
                                <TableCell align="center">ภาคการศึกษา</TableCell>
                                <TableCell align="center">ปีการศึกษา</TableCell>
                                <TableCell align="center">วันที่ - เวลา</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {activity.map((item, index ) => (
                                <TableRow key={item.id}>
                                <TableCell align="center">{item.edges?.actiStud?.fname}</TableCell>
                                <TableCell align="center">{item.aCTIVITYNAME}</TableCell>
                                <TableCell align="center">{item.edges?.actiPlace?.pLACE}</TableCell>
                                <TableCell align="center">{item.edges?.actiAgen?.aGENCY}</TableCell>
                                <TableCell align="center">{item.hours}</TableCell>
                                <TableCell align="center">{item.edges?.actiTerm?.semester}</TableCell>
                                <TableCell align="center">{item.edges?.actiYear?.years}</TableCell>
                                <TableCell align="center">{item.added}</TableCell>
                                </TableRow>
                                ))}
                            </TableBody>
                            </Table>
                        </TableContainer>

                        <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity="error" style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>

                </InfoCard>

                        

            </Content>
     </Page>
    );
}