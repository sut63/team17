import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Alert from '@material-ui/lab/Alert';
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

    const [addr,Setaddr] = useState(String);
    const [email,Setemail] = useState(String);
    const [fname,Setfname] = useState(String);
    const [lname,Setlname] = useState(String);
    const [school,Setschool] = useState(String);
    const [tel,Settel] = useState(String);

    const [sex,Setsex] = useState(Number);
    const [province,Setprovince] = useState(Number);
    const [title,Settitle] = useState(Number);
    const [degree,Setdegree] = useState(Number);

    const [district,Setdistrict] = useState(Number);
    const [postal,Setpostal] = useState(Number);
    const [zone,Setzone] = useState(Number);
  
    const AddrhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setaddr(event.target.value as string);
    };
    const EmailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setemail(event.target.value as string);
    };
    const FnamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setfname(event.target.value as string);
    };
    const LnamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setlname(event.target.value as string);
    };
    const SchoolhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setschool(event.target.value as string);
    };
    const TelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Settel(event.target.value as string);
    };
    const SexhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setsex(event.target.value as number);
    };
    const ProvincehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setprovince(event.target.value as number);
    };
    const TitlehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Settitle(event.target.value as number);
    };
    const DegreehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setdegree(event.target.value as number);
    };
    const DistricthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setdistrict(event.target.value as number);
    };
    const PostalhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setpostal(event.target.value as number);
    };
    const ZonehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      Setzone(event.target.value as number);
    };

  const Student = {
      addr:addr,
      degree:degree,
      email:email,
      fname:fname,
      lname:lname,
      province:province,
      school:school,
      sex:sex,
      tel:tel,
      title:title,
      zone:zone,
      district:district,
      postal:postal
    };
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
          setSuccess(true);
          setFail(false);
        } else {
          setFail(true);
          setSuccess(false);
        }
      })
  }

  return ( 
    <Page theme={pageTheme.home}>
    <Header
      title={'Student Management'}
      subtitle='Student Registration Department'
    >
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
                <Select name="sex" id='sex' value={sex}
                    onChange={SexhandleChange}>
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
                <Select name="title" id='title' value={title}
                    onChange={TitlehandleChange}>
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
                  <TextField variant='outlined' type='string' name="school" value={school}
                    onChange={SchoolhandleChange}/>
                  </div>
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>Province</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="province" value={province}
                    onChange={ProvincehandleChange}>
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
                  <Select name="postal" value={postal}
                    onChange={PostalhandleChange}>
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
                  <TextField variant='outlined' type='string' name="fname" value={fname}
                    onChange={FnamehandleChange}/>
                  </div>  
                </Grid>
              </Grid>

              <Grid container spacing={1}>
                <Grid item xs={4}>
                  <b>Degree</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="degree" value={degree}
                    onChange={DegreehandleChange}>
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
                  <Select name="district" value={district}
                    onChange={DistricthandleChange}>
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
                  <TextField variant='outlined' type='number' name="tel" value={tel}
                    onChange={TelhandleChange}/>
                  </div>
                </Grid>
              </Grid>

           </TableCell>
           
           <TableCell>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                <b>Last Name</b>
                <div>
                <TextField variant='outlined' type='string' name="lname" value={lname}
                    onChange={LnamehandleChange}/>
                </div>
              </Grid>
            </Grid>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                  <b>Recent Address</b>
                  <div>
                  <TextField variant='outlined' type='string' name="addr" value={addr}
                    onChange={AddrhandleChange}/>
                  </div>
              </Grid>
            </Grid>

            <Grid container spacing={1}>
              <Grid item xs={4}>
                  <b>Subdistrict</b>
                  <div>
                  <FormControl variant="outlined" fullWidth>
                  <Select name="zone" value={zone}
                    onChange={ZonehandleChange}>
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
                <TextField variant='outlined' type='email' name="email" value={email}
                    onChange={EmailhandleChange}/>
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