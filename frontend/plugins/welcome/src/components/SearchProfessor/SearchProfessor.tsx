import React, { FC, useEffect, useState } from 'react';
import clsx from 'clsx';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';

import SearchIcon from '@material-ui/icons/Search';
import DeleteIcon from '@material-ui/icons/Delete';
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
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import { DefaultApi } from '../../api/apis';
import { EntProfessor } from '../../api/models/EntProfessor';
import Swal from 'sweetalert2'; // alert

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


const Toast = Swal.mixin({
  //toast: true,
  position: 'center',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  showCloseButton: true,

});




export default function ComponentsTable() {

  //--------------------------
  ;

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  //---------------------------
  const [checkname, setNames] = useState(false);
  const [professor, setProfessor] = useState<EntProfessor[]>([])

  //--------------------------
  const [name, setName] = useState(String);
  const profile = { givenName: 'ระบบค้นหาข้อมูลอาจารย์' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getProfessors = async () => {
      const res = await api.listProfessor({ offset: 0 });
      setLoading(false);
      setProfessor(res);
    };
    getProfessors();
  }, [loading]);

  const PNamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setNames(false);
    setName(event.target.value as string);

  };

  const cleardata = () => {
    setName("");
    setSearch(false);
    setNames(false);
    setSearch(false);

  }

  const checkresearch = async () => {
    var check = false;
    professor.map(item => {
      if (name != "") {
        if (item.name?.includes(name)) {
          setNames(true);
          alertMessage("success", "ค้นหาสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูล");
    }
    console.log(checkname)
  };



  return (
    <Page theme={pageTheme.home}>
      <Header
      title={'Search Professor '}
      subtitle='ค้นหาข้อมูลอาจารย์'
    >
    </Header>
      <Content>
      
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>กรอกชื่อ - นามสุกล</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
              <TextField
                id="name"
                value={name}
                onChange = {PNamehandlehange}
                multiline
                variant="outlined"
                fullWidth
              />
              </form>
            </Grid>
            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
            
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SearchIcon/>}
                onClick={() => {
                  checkresearch();
                  setSearch(true);

                }}
              >
                ค้นหา
              </Button>
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            
              <Button
                variant="contained"
                color="secondary"
                size="large"
                startIcon={<DeleteIcon />}
                onClick={() => {
                  cleardata();

                }}
                
              >
                ล้างข้อมูล
              </Button>
              
            </Grid>
          </Grid>
        </Container>
      
        <br></br>
        <Paper>
              {search ? (
                <div>
                  {  checkname? (
                    
                    <TableContainer component={Paper}>
                      <Table>
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">คำนำหน้า</TableCell>
                                <TableCell align="center">ชื่อ - นามสกุล</TableCell>
                                <TableCell align="center">ตำแหน่ง</TableCell>
                                <TableCell align="center">สำนักวิชา</TableCell>
                                <TableCell align="center">เบอร์มือถือ</TableCell>
                                <TableCell align="center">E - mail</TableCell>
                              </TableRow>
                        </TableHead>
                        <TableBody>

                          {professor.filter((filter: any) => filter.name.includes(name)).map((item => (
                            <TableRow key={item.id}>
                                <TableCell align="center">{item.edges?.profPre?.prefix}</TableCell>
                                <TableCell align="center">{item.name}</TableCell>
                                <TableCell align="center">{item.edges?.profPros?.professorship}</TableCell>
                                <TableCell align="center">{item.edges?.profFacu?.faculty}</TableCell>
                                <TableCell align="center">{item.tel}</TableCell>
                                <TableCell align="center">{item.email}</TableCell>
                            </TableRow>
                          )))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                    
                  )
                     : null}
                </div>
              ) : null}
            </Paper>
      </Content>
    </Page>
  );
}
