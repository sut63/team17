import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Alert from '@material-ui/lab/Alert';
import Swal from 'sweetalert2'; // alert
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
import { EntRegion, EntContinent, EntCountry} from '../../api/models/';
import { Cookies } from '../../Cookie';


const ProvinceUI: FC<{}> = () => {
  const api = new DefaultApi();
  const [region,Setregion] = useState<EntRegion[]>([]);
  const [country,Setcountry] = useState<EntCountry[]>([]);
  const [continent,Setcontinent] = useState<EntContinent[]>([]);

  const getRegion = async () => {
      const res = await api.listRegion({ limit: 10, offset: 0 });
      Setregion(res);
    };
  const getCount = async () => {
      const res = await api.listCountry({ limit: 10, offset: 0 });
      Setcountry(res);
    };
  const getconti = async () => {
      const res = await api.listContinent({ limit: 10, offset: 0 });
      Setcontinent(res);
    };
  useEffect(() => {
      getCount();
      getRegion();
      getconti();
  }, []);

  const [p,setp]=useState(String)
  const ph = (event: React.ChangeEvent<{ value: unknown }>) => {
    setp(event.target.value as string);
  };
  const [d,setd]=useState(String)
  const dh = (event: React.ChangeEvent<{ value: unknown }>) => {
    setd(event.target.value as string);
  };
  const [s,sets]=useState(String)
  const sh = (event: React.ChangeEvent<{ value: unknown }>) => {
    sets(event.target.value as string);
  };
  const [n,setn]=useState(String)
  const nh = (event: React.ChangeEvent<{ value: unknown }>) => {
    setn(event.target.value as string);
  };
  const [r,setr]=useState(Number)
  const rh = (event: React.ChangeEvent<{ value: unknown }>) => {
    setr(event.target.value as number);
  };
  const [u,setu]=useState(Number)
  const uh = (event: React.ChangeEvent<{ value: unknown }>) => {
    setu(event.target.value as number);
  };
  const [t,sett]=useState(Number)
  const th = (event: React.ChangeEvent<{ value: unknown }>) => {
    sett(event.target.value as number);
  };

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

  const Create = async () => {
    const nnn = {
      cont:t,
      coun:u,
      dist:d,
      post:n,
      prov:p,
      regi:r,
      subd:s,
    }
    const res1: any = await api.createProvince({ province : nnn });
    if (res1.id != '') {
      Toast.fire({
        icon: 'success',
        title: 'successful',
      });
    } else {
      Toast.fire({
        icon: 'error',
        title: 'failure',
      });
    }
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
          title={'Student Management'}
          subtitle='Student Registration Department'>
            <Avatar alt="Remy Sharp"/>
            <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
            <Button variant="text" color="secondary" size="large"
              onClick={Clears} > Logout </Button>
        </Header>
        <Content>
          <ContentHeader title="Student Background"></ContentHeader>
          <TableContainer component={Paper}>
          <Grid container spacing={1}>
                  <Grid item xs={2}>
                    <b>Continent</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="ct" value={t||''}
                    onChange={th}>
                {continent.map((item) => {
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
                  <Grid item xs={2}>
                    <b>Country</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="cn" value={u||''}
                    onChange={uh}>
                {country.map((item) => {
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
                  <Grid item xs={2}>
                    <b>Region</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="re" value={r||''}
                    onChange={rh}>
                {region.map((item) => {
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
                  <Grid item xs={2}>
                    <b>Province</b>
                    <div>
               <TextField name='p' type='string' value={p||''} onChange={ph}/>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>District</b>
                    <div>
                <TextField name='d' type='string' value={d||''} onChange={dh}/>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>Subdistrict</b>
                    <div>
                <TextField name='s' type='string' value={s||''} onChange={sh}/>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>Post Code</b>
                    <div>
                <TextField name='p' type='string' value={n||''} onChange={nh}/>
                </div>
                  </Grid>
                </Grid>

                <Button variant="contained" color="primary" disableElevation  onClick={() => {
                    Create(); 
                }} > Save </Button>
          </TableContainer>
          </Content>
        </Page>
    );
};export default ProvinceUI;