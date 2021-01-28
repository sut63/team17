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
    const [results, setResults] = React.useState<EntResults[]>([]);
    //const [results, setResults] = React.useState(Array);
 


  
    
    

  //Get Data By Textfile and ComboBox ************************************************
  const [yearx, setYearx] = React.useState();
  const [studentx, setStudentx] = React.useState();
  const [termx, setTermx] = React.useState();
  




   //Handle chang********************************************************************
  const handleInputYear = (event: any) => {
    setYearx(event.target.value);
    clear();
  };
  const handleInputStudent = (event: any) => {
    setStudentx(event.target.value);
    clear();
  };
  const handleInputTerm = (event: any) => {
    setTermx(event.target.value);
    clear();
  };

 

  //function ต่อเนิ่อง
  function ss2 (){
      clear();
  }

 //clear
  function clear (){
      setResults([])
      
  }
    

  
  //get
  const getResultss = async () => {
    const res = await api.listResults({ limit: 10, offset: 0 });
    setResults(res);
    var check = false
    var chpap = ""
    for(let i = 0; i < res.length ; i++){
        if(res[i].edges?.resuTerm?.semester == termx && res[i].edges?.resuYear?.years == yearx && res[i].edges?.resuStud?.id == studentx){
            check = true
            chpap += "s"
        }
    }
    if(check == true || chpap != ""){
        alertMessage("success","ค้นหาสำเร็จ")
    }else{
        alertMessage("error","ค้นหาไม่พบ")
    }
    
  };

  

  

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
                        <TextField variant="outlined" className={classes.textField} 
                        type="String" 
                        name="grade"
                        value={studentx} // (undefined || '') = ''
                        onChange={handleInputStudent}
                        >                                                   
                        </TextField>                          
                        </FormControl>
                    </Grid>
                    

                    <Grid item xs={1}>
                        <div className={classes.paper}>ปีการศึกษา</div>
                    </Grid>
                    <Grid item xs={1}>
                        <FormControl variant="outlined" className={classes.formControl}>
                        <TextField variant="outlined" className={classes.textField} 
                        type="Number" 
                        name="grade"
                        value={yearx} // (undefined || '') = ''
                        onChange={handleInputYear}
                        >                                                   
                        </TextField>   
                        </FormControl>
                    </Grid>


                    <Grid item xs={1}>
                        <div className={classes.paper}>เทอม</div>
                    </Grid>
                    <Grid item xs={2}>
                        <FormControl variant="outlined" className={classes.formControl}>
                        <TextField variant="outlined" className={classes.textField} 
                        type="Number" 
                        name="grade"
                        value={termx} // (undefined || '') = ''
                        onChange={handleInputTerm}
                        >                                                   
                        </TextField>   
                        </FormControl>
                    </Grid>

                    <Grid item xs={1}>
                        <Button
                            className={classes.bottom}
                            variant="contained"
                            color="primary"
                            size="large"
                            onClick={getResultss}
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
                        {results.filter((filter: any) => filter.edges?.resuTerm?.semester == termx && filter.edges?.resuYear?.years == yearx && filter.edges?.resuStud?.id == studentx).map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges.resuSubj.subjects}</TableCell>
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