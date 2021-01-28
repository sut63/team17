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
import {EntRegion, EntProvince, EntCountry, EntContinent} from '../../api/models/';
import { Cookies } from '../../Cookie';

interface Province {
  province:string,
  postal:number,
  district:string,
  subdistrict:string,
  continent:string,
  country:string,
  name:string,
}

  const ProvinceUI: FC<{}> = () => {

    const handleClose = (event: React.SyntheticEvent | React.MouseEvent, reason?: string) => {
      if (reason === 'clickaway') {
        return;
      }
      setSuccess(false);
      setFail(false);
    };

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(false);
  const [loading, setLoading] = useState(true);
  const api = new DefaultApi();

  const [success, setSuccess] =useState(false);
  const [fail, setFail] =useState(false);

  const [coun,Setcoun] = useState<EntCountry[]>([]);
  const [cont,Setcont] = useState<EntContinent[]>([]);
  const [regi,Setregi] = useState<EntRegion[]>([]);

  const getCoun = async () => {
    const res = await api.listCountry({ limit: 10, offset: 0 });
    Setcoun(res);
  };
  const getCont = async () => {
    const res = await api.listContinent({ limit: 10, offset: 0 });
    Setcont(res);
  };
  const getRegi = async () => {
    const res = await api.listRegion({ limit: 10, offset: 0 });
    Setregi(res);
  };

  useEffect(() => {
    getCoun();
    getCont();
    getRegi();
    }, []);
  
  const [Province, setProvince] = React.useState<Partial<Province>>({});

  const handleChange = (
      event: React.ChangeEvent<{ province?: string; value: any }>,
    ) => {
      const province = event.target.province as keyof typeof Province;
      const { value } = event.target;
      setProvince({ ...Province, [province]: value });
      console.log(Province);
    };

    function clear() {
      setProvince({});
      getCoun();
      getCont();
      getRegi();
    }
  
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/provinces';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Province),
    };

    console.log(Province);

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.status);
        if (data.status === true) {
          clear();
          setSuccess(true);
        } else {
          setFail(true);
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
      title={'Province Management'}
      subtitle='Province Registration Department'
    >
        <Avatar alt="Remy Sharp"/>
        <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
        <Button variant="text" color="secondary" size="large"
          onClick={Clears} > Logout </Button>
    </Header>
      <Content>
        <ContentHeader title="Province Background"></ContentHeader>
        <TableContainer component={Paper}>
         <Table>
           <TableBody>
             <TableCell>


                  <Grid container spacing={1}>
                    <Grid item xs={2}>
                      <b>Continent</b>
                      <div>
                      <FormControl variant="outlined" fullWidth>
                      <Select name="cotinent" id='title'  
                      value={Province.continent}
                      onChange={handleChange}>
                        {cont.map((item) => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                          {item.continent}
                          </MenuItem>
                        );
                        })}
                      </Select>
                      </FormControl>
                      </div>
                    </Grid>
                  </Grid>

                  <Grid container spacing={1}>
                    <Grid item xs={2}>
                      <b>Country</b>
                      <div>
                      <FormControl variant="outlined" fullWidth>
                      <Select name="country" id='title'  
                      value={Province.country}
                      onChange={handleChange}>
                        {coun.map((item) => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                          {item.country}
                          </MenuItem>
                        );
                        })}
                      </Select>
                      </FormControl>
                      </div>
                    </Grid>
                  </Grid>

                  <Grid container spacing={1}>
                    <Grid item xs={2}>
                      <b>Region</b>
                      <div>
                      <FormControl variant="outlined" fullWidth>
                      <Select name="name" id='title'  
                      value={Province.name}
                      onChange={handleChange}>
                        {regi.map((item) => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                          {item.name}
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
                        <b>Province</b>
                        <div>
                        <TextField variant='outlined' type='string' value={Province.province}
                            onChange={handleChange}/>
                        </div>
                      </Grid>
                  </Grid>

                  <Grid container spacing={1}>
                      <Grid item xs={4}>
                        <b>District</b>
                        <div>
                        <TextField variant='outlined' type='string' value={Province.district}
                            onChange={handleChange}/>
                        </div>
                      </Grid>
                  </Grid>

                  <Grid container spacing={1}>
                      <Grid item xs={4}>
                        <b>Subdistrict</b>
                        <div>
                        <TextField variant='outlined' type='string' value={Province.subdistrict}
                            onChange={handleChange}/>
                        </div>
                      </Grid>
                  </Grid>

                  <Grid container spacing={1}>
                      <Grid item xs={4}>
                        <b>Postal</b>
                        <div>
                        <TextField variant='outlined' type='int' value={Province.postal}
                            onChange={handleChange}/>
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
              <Button variant='contained' color='primary' onClick={save}>Save</Button>
              <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: 'top', horizontal: 'center' }}>
                <Alert onClose={handleClose} severity="error">
                  This is a error message!
                </Alert>
              </Snackbar>
              <Snackbar open={success} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: 'top', horizontal: 'center' }}>
                <Alert onClose={handleClose} severity="success">
                  This is a success message!
                </Alert>
              </Snackbar>
             </TableCell>
           </TableBody>
           </Table>
       </TableContainer>
       </Content>
    </Page>
  );
};export default ProvinceUI;
  