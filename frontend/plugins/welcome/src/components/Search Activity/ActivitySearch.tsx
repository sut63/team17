import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
    InfoCard,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { Alert } from '@material-ui/lab';

import { makeStyles, Theme, createStyles, ThemeProvider } from '@material-ui/core/styles';
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

import { EntActivity } from '../../api/models/EntActivity';


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

export default function SearchRoom() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [activity, setActivity] = useState<EntActivity[]>([]);
    

    const [activitysearch, setActivitysearch] = useState(String);
    

  useEffect(() => {
    const getActivitys = async () => {
    const res = await api.listActivity();
      setLoading(false);
      setActivity(res);
    };
    getActivitys();
    }, [loading]);

    const SearchActivity = async () => {
        const res = await api.listActivity();
        const activitysearch = ActivitySearch(res);
        
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setActivity([]);
        if(activitysearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setActivity(activitysearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.activitysearchcheck = true;
        
    }
    const ActivitySearch = (res: any) => {
        const data = res.filter((filter: EntActivity) => filter.edges?.actiStud?.fname?.includes(activitysearch))
       // console.log(data.length)
        if (data.length != 0 && activitysearch != "") {
            return data;
        }
        else {
            searchcheck.activitysearchcheck = false;
            if(activitysearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const ActivitySearchhandleChange = (event: any) => {
        setActivitysearch(event.target.value as string);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="ค้นหาประวัติกิจกรรมนักศึกษา">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>ชื่อนักศึกษา</strong></div>
                    <TextField
                        id="name"
                       // label="ค้นหารหัสพนักงาน"
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
                                {activity.map((item:any ) => (
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
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => { setStatus(false) }}>
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