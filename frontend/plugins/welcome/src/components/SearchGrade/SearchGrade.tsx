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
import { EntResults } from '../../api';
import { EntStudent } from '../../api';
import { EntYear } from '../../api';
import { EntTerm } from '../../api';




// alert setting
const Toast = Swal.mixin({
    toast: true,
    position: 'top',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
});
////--------------------------------------------
const alertMessage = (icon: any, title: any) => {
    Toast.fire({
        icon: icon,
        title: title,
    });
}




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
        width: 200,
        marginLeft: -70,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        width: 200,
    },
    bottom: {
        marginLeft: theme.spacing(0),
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(2),
    },
    divider: {
        margin: theme.spacing(2, 0),
    },
    table: {
        minWidth: 650,
    },
}));




const SearchGrade: FC<{}> = () => {
    const classes = useStyles();
    //const************************************************
    const api = new DefaultApi();
    //const [results, setResults] = React.useState(Array);
    const [years, setYears] = React.useState<EntYear[]>([]);
    const [terms, setTerms] = React.useState<EntTerm[]>([]);
    const [students, setStudents] = React.useState<EntStudent[]>([]);
    const [result111, setResult111] = React.useState<EntResults[]>([]);



    //Hook********************************************
    useEffect(() => {
        const getYears = async () => {
            const res = await api.listYear({ limit: 100, offset: 0 });
            //setLoading(true);
            setYears(res);
        };
        getYears();
    }, []);
    useEffect(() => {
        const getTerms = async () => {
            const res = await api.listTerm({ limit: 100, offset: 0 });
            //setLoading(true);
            setTerms(res);
        };
        getTerms();
    }, []);
    useEffect(() => {
        const getStudents = async () => {
            const res = await api.listStudent({ limit: 100, offset: 0 });
            //setLoading(true);
            setStudents(res);
        };
        getStudents();
    }, []);



    //Get Data By Textfile and ComboBox ************************************************
  const [yearx, setYearx] = React.useState('');
  const [studentx, setStudentx] = React.useState('');
  const [termx, setTermx] = React.useState('');
  
  

  let yearID = Number(yearx)
  let studentID = Number(studentx)
  let termID = Number(termx)
  

  let results = {   
	studentID,
	yearID,
    termID,
    };





   //Handle chang********************************************************************
  const handleInputYear = (event: any) => {
    setYearx(event.target.value);
    
  };
  const handleInputStudent = (event: any) => {
    setStudentx(event.target.value);
    
  };
  const handleInputTerm = (event: any) => {
    setTermx(event.target.value);
    
  };





    //clear
    //function clear() {
     //   setResults([])

    //}
    const getResult = async () => {
        const res = await api.listResultssomting({ id:results.studentID,year:results.yearID,term:results.termID });
        //setLoading(true);
        let temp = res.length
        if (temp > 0 ) {
          
            Toast.fire({
              icon: 'success',
              title: 'ค้นหาสำเร็จ',
            });
          } else {
            Toast.fire({
                icon: 'error',
                title: 'ค้นหาไม่สำเร็จ',
              });
          }
          console.log(res)
        setResult111(res);
    };

  //ค้นหา**********************************************************
   // function save data
   function search() {
    getResult()
   }
 
   



    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาประวัติผลการศึกษา`}>
                <div style={{ marginLeft: 10, marginRight: 20 }}></div>

            </Header>
            <Content>
                <Grid container spacing={1}>


                    <Grid item xs={1}>
                        <div className={classes.paper}>รหัสนักศึกษา</div>
                    </Grid>
                    <Grid item xs={1}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>เลือกรหัสนักศึกษา</InputLabel>
                            <Select
                                name="StudentID"
                                value={results.studentID || ''} // (undefined || '') = ''
                                onChange={handleInputStudent}
                            >
                                {students.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.id}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>


                    <Grid item xs={1}>
                        <div className={classes.paper}>ปีการศึกษา</div>
                    </Grid>
                    <Grid item xs={1}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>เลือกปีภาคการศึกษา</InputLabel>
                            <Select
                                name="YearID"
                                value={results.yearID || ''} // (undefined || '') = ''
                                onChange={handleInputYear}
                            >
                                {years.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.years}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>


                    <Grid item xs={1}>
                        <div className={classes.paper}>ภาคการศึกษา</div>
                    </Grid>
                    <Grid item xs={1}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>เลือกภาคการศึกษา</InputLabel>
                            <Select
                                name="TermID"
                                value={results.termID || ''} // (undefined || '') = ''
                                onChange={handleInputTerm}
                            >
                                {terms.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.semester}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={1}>
                        <Button
                            className={classes.bottom}
                            variant="contained"
                            color="primary"
                            size="large"
                            onClick={search}
                        >
                            ค้นหา
                        </Button>
                    </Grid>



                </Grid>


                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ผลการค้นหา
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">วิชา</TableCell>
                                <TableCell align="center">เกรด</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {result111.map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.resuSubj.subjects}</TableCell>
                                    <TableCell align="center">{item.grade}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
};

export default SearchGrade;